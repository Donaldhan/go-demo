package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/net/context"
	"log"
	"math/big"
	"strings"
	"time"
)

func initAbi() abi.ABI {
	abiString := `[
		  {
			"inputs": [
			  {
				"internalType": "string",
				"name": "_greeting",
				"type": "string"
			  }
			],
			"stateMutability": "nonpayable",
			"type": "constructor"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": true,
				"internalType": "string",
				"name": "greeting",
				"type": "string"
			  }
			],
			"name": "GreeterChange",
			"type": "event"
		  },
		  {
			"inputs": [],
			"name": "greet",
			"outputs": [
			  {
				"internalType": "string",
				"name": "",
				"type": "string"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "string",
				"name": "_greeting",
				"type": "string"
			  }
			],
			"name": "setGreeting",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  }
		]
		`
	// 合约 ABI JSON
	contractABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		log.Fatal(err)
	}
	return contractABI

}

func abiTransaction(contractAdr string) {
	contractABI := initAbi()
	// 连接到以太坊节点
	client := initClient()
	// 合约地址
	contractAddress := common.HexToAddress(contractAdr)

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 获取发送方的公钥地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("发送方地址:", fromAddress.Hex())

	// 获取账户 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置 Gas 价格和限制
	suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	msg := "hello rain!"
	// 打包调用数据
	data, err := contractABI.Pack("setGreeting", msg)
	if err != nil {
		log.Fatal(err)
	}

	// 签名交易
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 创建消息调用的参数
	callMsg := ethereum.CallMsg{
		From:     fromAddress,
		To:       &contractAddress,
		Value:    big.NewInt(0),
		Data:     data, // 如果是合约调用，需要填写合约方法的编码数据
		GasPrice: suggestedGasPrice,
	}

	gasLimit := gasLimitBaseEstimateGas(client, callMsg, 1)

	// 提高 gasPrice，例如增加 20%
	increasedGasPrice := new(big.Int).Mul(suggestedGasPrice, big.NewInt(12))
	increasedGasPrice.Div(increasedGasPrice, big.NewInt(10)) // 相当于增加 20%

	fmt.Printf("建议的 gasPrice: %s wei\n", suggestedGasPrice.String())
	fmt.Printf("增加 20%% 后的 gasPrice: %s wei\n", increasedGasPrice.String())

	// 构建交易对象
	txData := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &contractAddress,
		Value:    big.NewInt(0),
		Gas:      gasLimit,
		GasPrice: increasedGasPrice,
		Data:     data,
	})
	// 签名交易
	signedTx, err := types.SignTx(txData, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

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
func getBaseFeeAndPriorityFee(client *ethclient.Client) (*big.Int, *big.Int) {
	// 获取最新区块号
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block number:", blockNumber)
	number := new(big.Int).SetUint64(blockNumber)
	// 获取当前建议的 base fee
	header, err := client.HeaderByNumber(context.Background(), number)
	if err != nil {
		log.Fatal(err)
	}
	baseFee := header.BaseFee
	fmt.Println("Current suggested base fee:", baseFee)
	// 获取当前网络建议的 maxPriorityFeePerGas
	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 输出 maxPriorityFeePerGas
	fmt.Printf("建议的 maxPriorityFeePerGas: %s wei\n", maxPriorityFeePerGas.String())
	return baseFee, maxPriorityFeePerGas
}
func gasLimitBaseEstimateGas(client *ethclient.Client, msg ethereum.CallMsg, multiFactor float64) uint64 {

	// 使用 EstimateGas 估算基础 gasLimit
	baseGasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}

	// 将 gasLimit 上调 20%
	adjustedGasLimit := uint64(float64(baseGasLimit) * multiFactor)

	fmt.Printf("基础 gasLimit: %d\n", baseGasLimit)
	fmt.Printf("上调 20%% 后的 gasLimit: %d\n", adjustedGasLimit)
	return adjustedGasLimit
}
func abiTxEip1559(contractAdr string) {
	contractABI := initAbi()
	// 连接到以太坊节点
	client := initClient()
	// 合约地址
	contractAddress := common.HexToAddress(contractAdr)

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 获取发送方的公钥地址
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	fmt.Println("发送方地址:", fromAddress.Hex())

	// 获取账户 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	msg := "hello rain!"
	// 打包调用数据
	data, err := contractABI.Pack("setGreeting", msg)
	if err != nil {
		log.Fatal(err)
	}

	baseFee, maxPriorityFeePerGas := getBaseFeeAndPriorityFee(client)
	maxFeePerGas := new(big.Int).Add(baseFee, maxPriorityFeePerGas)

	// 构建交易对象
	// 签名交易
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 创建消息调用的参数
	callMsg := ethereum.CallMsg{
		From:      fromAddress,
		GasFeeCap: maxFeePerGas,         // 设置 maxFeePerGas
		GasTipCap: maxPriorityFeePerGas, // 设置 maxPriorityFeePerGas
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      data, // 如果是合约调用，需要填写合约方法的编码数据
	}

	adjustedGasLimit := gasLimitBaseEstimateGas(client, callMsg, 1.2)
	// 构造 EIP-1559 交易
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId, // 主网链 ID
		Nonce:     nonce,
		GasFeeCap: maxFeePerGas,         // 设置 maxFeePerGas
		GasTipCap: maxPriorityFeePerGas, // 设置 maxPriorityFeePerGas
		Gas:       adjustedGasLimit,
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      data,
	})

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

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

func parseTxLog() {
	// TOOD
}

func abiCall(contractAdr string) {
	contractABI := initAbi()
	// 连接到以太坊节点
	client := initClient()
	// 合约地址
	contractAddress := common.HexToAddress(contractAdr)

	// 准备调用的合约方法
	data, err := contractABI.Pack("greet")
	if err != nil {
		log.Fatal(err)
	}

	// 调用合约
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 解析返回结果
	var msg string
	err = contractABI.UnpackIntoInterface(&msg, "greet", result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("msg: %s\n", msg)
}
