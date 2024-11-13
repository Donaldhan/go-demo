package evm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	// 连接到以太坊节点
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 获取最新区块号
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block number:", blockNumber)
}
