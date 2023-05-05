package ipfsapi

import (
	"bufio"
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
// 上载下载文件
func ipfsClientAddAndGet() {
	//多播地址，构造api
	addr, err := ma.NewMultiaddr("/ip4/127.0.0.1/tcp/5001")
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
	// /ipfs/QmQ8BsiU8Vtv97PtcWZiP8bL7BjQmnGjqMZDj4ZXQy6W1w
	fmt.Printf("Uploaded file %s with CID %s\n", fileName, cid)

	// 下载文件
	node, err := api.Unixfs().Get(context.Background(), cid)
	reader := bufio.NewReader(node.(files.File))
	// for {
	// 	str, err := reader.ReadString('\n') //读到一个换行就结束
	// 	if err == io.EOF {                  //io.EOF 表示文件的末尾
	// 		break
	// 	}
	// 	fmt.Print(str)
	// }
	// reader, err = node.(files.File).Read()
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Downloaded content: %s\n", string(content))
}
