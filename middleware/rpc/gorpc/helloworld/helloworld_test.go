package hellworld

import "testing"

// cd timer
// go test -v
func TestRpcServer(t *testing.T) {
	startServer()
}

func TestRpcClient(t *testing.T) {
	startClient()
}
