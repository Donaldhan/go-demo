package evm

import "math/big"

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
func weiToEth(wei *big.Int) *big.Float {
	ethValue := new(big.Float).SetInt(wei)
	ethValue.Quo(ethValue, big.NewFloat(1e18)) // 除以 10^18
	return ethValue
}

// ethToWei 将 ether 转换为 wei
func ethToWei(eth *big.Float) *big.Int {
	weiValue := new(big.Float).Mul(eth, big.NewFloat(1e18)) // 乘以 10^18
	result := new(big.Int)
	weiValue.Int(result)
	return result
}
