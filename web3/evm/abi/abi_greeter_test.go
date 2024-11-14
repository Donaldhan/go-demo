package abi

import "testing"

// go test -v
// https://sepolia.etherscan.io/address/0x995656a896fada3ff7c51bfb27e688c7772944b2
var contractAddr = "0x995656a896Fada3ff7c51Bfb27E688c7772944b2"

func TestAbiCall(t *testing.T) {
	abiCall(contractAddr)
}
func TestAbiTransact(t *testing.T) {
	abiTransaction(contractAddr)
	abiCall(contractAddr)
}

func TestAbiTxEip1559(t *testing.T) {
	abiTxEip1559(contractAddr)
	abiCall(contractAddr)
}
