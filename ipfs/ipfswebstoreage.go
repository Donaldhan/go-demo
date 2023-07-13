package ipfsapi

import (
	"context"
	"fmt"
	"godemo/config"
	"io"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/ipfs/go-cid"
	"github.com/web3-storage/go-w3s-client"
	w3fs "github.com/web3-storage/go-w3s-client/fs"
)

// http://web3.storage.ipns.localhost:8080/docs/how-tos/store/?lang=go#uploading-to-web3storage
//
// Uploading to web3.storage
func upLoad() {
	log.Printf("start upload............")
	//TODO 从配置文件读取
	// token, ok := os.LookupEnv("WEB3_STORAGE_TOKEN")
	// if !ok {
	// 	fmt.Fprintln(os.Stderr, "No API token - set the WEB3_STORAGE_TOKEN environment var and try again.")
	// 	os.Exit(1)
	// }
	token := config.GetWebStorageToken()
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "usage: %s <filename>\n", os.Args[0])
	// 	os.Exit(1)
	// }
	// filename := os.Args[1]

	// Create a new web3.storage client using the token
	client, err := w3s.NewClient(w3s.WithToken(token))
	if err != nil {
		panic(err)
	}
	filename := "./example.txt"
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	basename := path.Base(filename)
	// Upload to web3.storage
	fmt.Printf("Storing %s ...\n", basename)
	cid, err := client.Put(context.Background(), file)
	if err != nil {
		panic(err)
	}

	gatewayURL := fmt.Sprintf("https://%s.ipfs.dweb.link/%s\n", cid.String(), basename)
	fmt.Printf("Stored %s with web3.storage! View it at: %s\n", basename, gatewayURL)
}

// WithEndpoint sets the URL of the root API when making requests (default https://api.web3.storage).
// Usage:
// TOKEN="API_TOKEN" go run ./main.go
// TODO 从配置文件读取
func web3StorageClient() {
	c, err := w3s.NewClient(
		w3s.WithEndpoint(config.GetWebStorageEndPoint()),
		w3s.WithToken(config.GetWebStorageToken()),
	)
	if err != nil {
		panic(err)
	}
	//上传单个文件
	// cid := putSingleFile(c)
	// 上传多文件
	// cid := putMultipleFiles(c)
	//上传多文件多目录
	// cid := putMultipleFilesAndDirectories(c)
	//上传目录
	cid := putDirectory(c)
	getStatusForCid(c, cid)
	// getStatusForKnownCid(c)
	//获取文件
	// getFiles(c)
	getFilesByCid(c, cid)
	// 列举当前节点下的上传cid文件
	listUploads(c)
}

// /上传单个文件
func putSingleFile(c w3s.Client) cid.Cid {
	file, err := os.Open("./images/baby.jpg")
	if err != nil {
		panic(err)
	}
	return putFile(c, file)
}

// 上传多文件
func putMultipleFiles(c w3s.Client) cid.Cid {
	f0, err := os.Open("./images/baby.jpg")
	if err != nil {
		panic(err)
	}
	f1, err := os.Open("./images/dance.gif")
	if err != nil {
		panic(err)
	}
	dir := w3fs.NewDir("comic", []fs.File{f0, f1})
	return putFile(c, dir)
}

// 上传多文件多目录
func putMultipleFilesAndDirectories(c w3s.Client) cid.Cid {
	f0, err := os.Open("./images/baby.jpg")
	if err != nil {
		panic(err)
	}
	f1, err := os.Open("./images/dance.gif")
	if err != nil {
		panic(err)
	}
	d0 := w3fs.NewDir("one", []fs.File{f0})
	d1 := w3fs.NewDir("two", []fs.File{f1})
	rootdir := w3fs.NewDir("comicxdir", []fs.File{d0, d1})
	return putFile(c, rootdir)
}

// 上传目录
func putDirectory(c w3s.Client) cid.Cid {
	dir, err := os.Open("./images")
	if err != nil {
		panic(err)
	}
	return putFile(c, dir)
}

// 上传文件
func putFile(c w3s.Client, f fs.File, opts ...w3s.PutOption) cid.Cid {
	cid, err := c.Put(context.Background(), f, opts...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("https://%v.ipfs.dweb.link\n", cid)
	return cid
}

// 获取cid状态
func getStatusForCid(c w3s.Client, cid cid.Cid) {
	s, err := c.Status(context.Background(), cid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %+v\n", s)
}

// 根据cid 获取状态
func getStatusForKnownCid(c w3s.Client) {
	cid, _ := cid.Parse("bafybeibcdos7ga53b7jzllmrqywi5xcojpny74q5s6xnqjqxtqgn2ku2wm")
	getStatusForCid(c, cid)
}

// 获取文件
func getFiles(c w3s.Client) {
	cid, _ := cid.Parse("bafybeierxlzrcfvmd2gveovzbmyskhadova3fxg3bzkaxxc644kkk5jfty")
	getFilesByCid(c, cid)
}

// 获取cid下的文件
func getFilesByCid(c w3s.Client, cid cid.Cid) {

	res, err := c.Get(context.Background(), cid)
	if err != nil {
		panic(err)
	}

	f, fsys, err := res.Files()
	if err != nil {
		panic(err)
	}

	info, err := f.Stat()
	if err != nil {
		panic(err)
	}

	if info.IsDir() {
		err = fs.WalkDir(fsys, "/", func(path string, d fs.DirEntry, err error) error {
			info, _ := d.Info()
			fmt.Printf("%s (%d bytes)\n", path, info.Size())
			return err
		})
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("%s (%d bytes)\n", cid.String(), info.Size())
	}
}

// 列举当前节点下的上传cid文件
func listUploads(c w3s.Client) {
	uploads, err := c.List(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		u, err := uploads.Next()
		if err != nil {
			// finished successfully
			if err == io.EOF {
				break
			}
			panic(err)
		}

		fmt.Printf("%s	%s	Size: %d	Deals: %d	Pins: %d\n", u.Created.Format("2006-01-02 15:04:05"), u.Cid, u.DagSize, len(u.Deals), len(u.Pins))
	}
}
