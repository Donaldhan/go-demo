package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

// 定义常量，表示每个单位等于多少 wei
var (
	Wei    = big.NewInt(1)
	Kwei   = new(big.Int).Exp(big.NewInt(10), big.NewInt(3), nil)  // 10^3 wei
	Mwei   = new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)  // 10^6 wei
	Gwei   = new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)  // 10^9 wei
	Szabo  = new(big.Int).Exp(big.NewInt(10), big.NewInt(12), nil) // 10^12 wei
	Finney = new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil) // 10^15 wei
	Ether  = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 10^18 wei
)

// weiToEth 将 wei 转换为 ether
func WeiToEth(wei *big.Int) *big.Float {
	ethValue := new(big.Float).SetInt(wei)
	ethValue.Quo(ethValue, big.NewFloat(1e18)) // 除以 10^18
	return ethValue
}

// ethToWei 将 ether 转换为 wei
func EthToWei(eth *big.Float) *big.Int {
	weiValue := new(big.Float).Mul(eth, big.NewFloat(1e18)) // 乘以 10^18
	result := new(big.Int)
	weiValue.Int(result)
	return result
}

// waitForReceipt 等待并返回交易回执
func WaitForReceipt(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
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
func GetBaseFeeAndPriorityFee(client *ethclient.Client) (*big.Int, *big.Int) {
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
	fmt.Printf("Current maxPriorityFeePerGas: %s wei\n", maxPriorityFeePerGas.String())
	return baseFee, maxPriorityFeePerGas
}
func GasLimitBaseEstimateGas(client *ethclient.Client, msg ethereum.CallMsg, multiFactor float64) uint64 {

	// 使用 EstimateGas 估算基础 gasLimit
	baseGasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}

	// 将 gasLimit 上调
	adjustedGasLimit := uint64(float64(baseGasLimit) * multiFactor)

	fmt.Printf("基础 gasLimit: %d\n", baseGasLimit)
	fmt.Printf("上调 后的 gasLimit: %d\n", adjustedGasLimit)
	return adjustedGasLimit
}
