package routeguide

import (
	"godemo/middleware/rpc/gorpc/routeguide/client"
	"godemo/middleware/rpc/gorpc/routeguide/server"
	"testing"
)

// cd timer
// go test -v
func TestRpcServer(t *testing.T) {
	server.StartServer()
}

func TestRpcClient(t *testing.T) {
	client.StartClient()
}
