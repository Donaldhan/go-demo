package evm

import "testing"

// go test -v
func TestChainDemo(t *testing.T) {
	chainDemo()
}

func TestTransfer(t *testing.T) {
	transfer()
}
func TestTransferBaseNewTx(t *testing.T) {
	transferBaseNewTx()
}

// 0xa280a67d03ced253d6922cba6839a57381a32c8f17f5a95ec1e528d3801caec4
func TestTransactionReceipt(t *testing.T) {
	txHash := "0x6f7408b015b2b96c6ca6ef75d3a2315b8d6bed14a56623d85b0f417cbc7157de"
	transactionReceipt(txHash)
}
func TestTransactionInfo(t *testing.T) {
	txHash := "0x6f7408b015b2b96c6ca6ef75d3a2315b8d6bed14a56623d85b0f417cbc7157de"
	transactionInfo(txHash)
}
