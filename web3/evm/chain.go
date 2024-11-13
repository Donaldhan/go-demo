package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

func initClient() *ethclient.Client {
	// 连接到以太坊节点
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		log.Fatal("init client error", err)
	}
	return client
}
func chainDemo() {
	client := initClient()
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("chainId:", chainId)
	// 获取最新区块号
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block number:", blockNumber)
	getBalance(client, "0x43186408725f64229A67e706f7523cB5e5A44279")

}
func getBalance(client *ethclient.Client, addr string) {
	// 替换为要查询余额的以太坊地址
	address := common.HexToAddress(addr)

	// 获取最新区块中的账户余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 打印余额（以 wei 为单位）
	fmt.Printf("账户余额（wei）: %s\n", balance.String())

	// 转换为以太（ETH）单位并打印
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	fmt.Printf("账户余额（ETH）: %f\n", ethValue)
}

func transfer() {
	client := initClient()

	// 发起账户的私钥
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 获取发送方的公钥地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("发送方地址:", fromAddress.Hex())

	// 目标地址
	toAddress := common.HexToAddress("0x20D6B2602757FC40e697d6f11b2f68AA0Fc0665c")

	// 获取发送方地址的 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	eth := big.NewFloat(0.0002) // 0.05 ETH
	wei := ethToWei(eth)
	fmt.Printf("value: %s wei\n", wei.String())

	// 获取当前 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易
	tx := types.NewTransaction(nonce, toAddress, wei, uint64(21000), gasPrice, nil)

	// 签名交易
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易已发送，交易哈希: %s\n", signedTx.Hash().Hex())

	// 设置超时时间为 1 分钟
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	receipt, err := waitForReceipt(ctx, client, signedTx.Hash())
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	//fmt.Printf("Transaction receipt received: %+v\n", receipt)
	fmt.Printf("Transaction receipt logs: %+v\n", receipt.Logs)
	fmt.Printf("Transaction receipt BlockNumber: %+v\n", receipt.BlockNumber)
	fmt.Printf("Transaction receipt GasUsed: %+v\n", receipt.GasUsed)
	fmt.Printf("Transaction receipt Status: %+v\n", receipt.Status)
	fmt.Printf("Transaction receipt ContractAddress: %+v\n", receipt.ContractAddress)
	fmt.Printf("Transaction receipt CumulativeGasUsed: %+v\n", receipt.CumulativeGasUsed)
	fmt.Printf("Transaction receipt EffectiveGasPrice: %+v\n", receipt.EffectiveGasPrice)

}

func transferBaseNewTx() {
	client := initClient()

	// 发起账户的私钥
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 获取发送方的公钥地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("发送方地址:", fromAddress.Hex())

	// 目标地址
	toAddress := common.HexToAddress("0x20D6B2602757FC40e697d6f11b2f68AA0Fc0665c")

	// 获取发送方地址的 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	eth := big.NewFloat(0.0002) // 0.05 ETH
	value := ethToWei(eth)
	fmt.Printf("value: %s wei\n", value.String())

	// 获取当前 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasLimit := uint64(21000)
	// 构建交易对象
	txData := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
	})

	// 签名交易
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 签名交易
	signedTx, err := types.SignTx(txData, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易已发送，交易哈希: %s\n", signedTx.Hash().Hex())

	// 设置超时时间为 1 分钟
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	receipt, err := waitForReceipt(ctx, client, signedTx.Hash())
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	//fmt.Printf("Transaction receipt received: %+v\n", receipt)
	fmt.Printf("Transaction receipt logs: %+v\n", receipt.Logs)
	fmt.Printf("Transaction receipt BlockNumber: %+v\n", receipt.BlockNumber)
	fmt.Printf("Transaction receipt GasUsed: %+v\n", receipt.GasUsed)
	fmt.Printf("Transaction receipt Status: %+v\n", receipt.Status)
	fmt.Printf("Transaction receipt ContractAddress: %+v\n", receipt.ContractAddress)
	fmt.Printf("Transaction receipt CumulativeGasUsed: %+v\n", receipt.CumulativeGasUsed)
	fmt.Printf("Transaction receipt EffectiveGasPrice: %+v\n", receipt.EffectiveGasPrice)
}

// waitForReceipt 等待并返回交易回执
func waitForReceipt(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		// 检查是否超过超时
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timed out waiting for transaction receipt")
		default:
			fmt.Println("Waiting for transaction to be mined...")
			time.Sleep(3 * time.Second)
		}
	}
}

// 获取交易hash
func transactionReceipt(txHashString string) {
	client := initClient()
	// 假设我们有一个交易哈希的字符串
	//txHashString := "0x5e9e2de37c5a907fe59e7e0bcb7c4d8c93c67f5f1a49b7a7a9e3edbba033d144"

	// 使用 HexToHash 将字符串转换为 common.Hash
	txHash := common.HexToHash(txHashString)
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}
	//fmt.Printf("Transaction receipt received: %+v\n", receipt)
	fmt.Printf("Transaction receipt logs: %+v\n", receipt.Logs)
	fmt.Printf("Transaction receipt BlockNumber: %+v\n", receipt.BlockNumber)
	fmt.Printf("Transaction receipt GasUsed: %+v\n", receipt.GasUsed)
	fmt.Printf("Transaction receipt Status: %+v\n", receipt.Status)
	fmt.Printf("Transaction receipt ContractAddress: %+v\n", receipt.ContractAddress)
	fmt.Printf("Transaction receipt CumulativeGasUsed: %+v\n", receipt.CumulativeGasUsed)
	fmt.Printf("Transaction receipt EffectiveGasPrice: %+v\n", receipt.EffectiveGasPrice)
}

func transactionInfo(txHashString string) {
	client := initClient()
	// 使用 HexToHash 将字符串转换为 common.Hash
	txHash := common.HexToHash(txHashString)

	// 获取交易信息
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Failed to retrieve transaction: %v", err)
	}

	// 输出交易基本信息
	fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
	//fmt.Printf("From: %s\n", tx.)
	fmt.Printf("To: %s\n", tx.To().Hex())
	fmt.Printf("Value: %s\n", tx.Value().String())
	fmt.Printf("Gas: %d\n", tx.Gas())
	fmt.Printf("Gas Price: %s\n", tx.GasPrice().String())
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("Data: %x\n", tx.Data())
	fmt.Printf("Transaction is pending: %v\n", isPending)
}
