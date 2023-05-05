package ipfsapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os" // v0.0.8

	files "github.com/ipfs/go-ipfs-files" // v0.0.8
	ipfClient "github.com/ipfs/go-ipfs-http-client"
	ma "github.com/multiformats/go-multiaddr"
)

func init() {
	log.Println("==============ipfsapi package init")
}

// # guide
// https://github.com/ipfs/go-ipfs-http-client
// https://pkg.go.dev/github.com/ipfs/go-ipfs-http-client?tab=versions
// https://github.com/ipfs/go-ipfs-api
// https://docs.ipfs.tech/reference/kubo/rpc/#api-v0-object-stat
// https://sourcegraph.com/github.com/ipfs/go-ipfs-http-client/-/blob/unixfs.go

func ipfsClientAdd() {
	//多播地址，构造api
	addr, err := ma.NewMultiaddr("http://localhost:5001")
	api, err := ipfClient.NewApi(addr)
	if err != nil {
		panic(err)
	}
	// 上传文件
	fileName := "example.txt"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileNodeSource := files.NewReaderFile(file)
	cid, err := api.Unixfs().Add(context.Background(), fileNodeSource)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uploaded file %s with CID %s\n", fileName, cid)

	// 下载文件
	node, err := api.Unixfs().Get(context.Background(), cid)
	// 创建读取器
	reader, err := node.(files.File).Reader()
	if err != nil {
		// 创建读取器时发生错误
		return err
	}
	defer reader.Close()
	defer reader.Close()
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Downloaded content: %s\n", string(content))
}
