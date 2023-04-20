package ipfsapi

import (
	"context"
	"fmt"
	"log"

	ipfsClient "github.com/ipfs/go-ipfs-http-client"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

func init() {
	log.Println("==============ipfsapi package init")
}

// # guide
// https://github.com/ipfs/go-ipfs-http-client
// https://pkg.go.dev/github.com/ipfs/go-ipfs-http-client?tab=versions
// https://github.com/ipfs/go-ipfs-api
// https://docs.ipfs.tech/reference/kubo/rpc/#api-v0-object-stat

func ipfsClientAdd() {
	// "Connect" to local node
	node, err := ipfsClient.NewLocalApi()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	// Pin a given file by its CID
	ctx := context.Background()
	cid := "bafkreidtuosuw37f5xmn65b3ksdiikajy7pwjjslzj2lxxz2vc4wdy3zku"
	p := path.New(cid)
	err = node.Pin().Add(ctx, p)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	return
}
