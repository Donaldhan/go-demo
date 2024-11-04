package rpc

import "testing"

// cd timer
// go test -v
func TestRpcServer(t *testing.T) {
	rpcServer()
}

func TestRpcClient(t *testing.T) {
	rpcClient()
}
