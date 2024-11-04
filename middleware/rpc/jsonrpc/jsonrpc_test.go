package jsonrpc

import "testing"

// cd timer
// go test -v
func TestRpcServer(t *testing.T) {
	jsonRpcServer()
}

func TestRpcClient(t *testing.T) {
	jsonRpcClient()
}
