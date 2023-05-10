package ipfsapi

import (
	"context"
	"fmt"
	"io"
	"io/fs"
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
	//TODO 从配置文件读取
	token, ok := os.LookupEnv("WEB3_STORAGE_TOKEN")
	if !ok {
		fmt.Fprintln(os.Stderr, "No API token - set the WEB3_STORAGE_TOKEN environment var and try again.")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}
	filename := os.Args[1]

	// Create a new web3.storage client using the token
	client, err := w3s.NewClient(w3s.WithToken(token))
	if err != nil {
		panic(err)
	}

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

// Usage:
// TOKEN="API_TOKEN" go run ./main.go
// TODO 从配置文件读取
func web3StorageClient() {
	c, err := w3s.NewClient(
		w3s.WithEndpoint(os.Getenv("WEB3_STORAGE_ENDPOINT")),
		w3s.WithToken(os.Getenv("WEB3_STORAGE_TOKEN")),
	)
	if err != nil {
		panic(err)
	}

	// cid := putSingleFile(c)
	// getStatusForCid(c, cid)
	// getStatusForKnownCid(c)
	getFiles(c)
	// listUploads(c)
}

func putSingleFile(c w3s.Client) cid.Cid {
	file, err := os.Open("images/donotresist.jpg")
	if err != nil {
		panic(err)
	}
	return putFile(c, file)
}

func putMultipleFiles(c w3s.Client) cid.Cid {
	f0, err := os.Open("images/donotresist.jpg")
	if err != nil {
		panic(err)
	}
	f1, err := os.Open("images/pinpie.jpg")
	if err != nil {
		panic(err)
	}
	dir := w3fs.NewDir("comic", []fs.File{f0, f1})
	return putFile(c, dir)
}

func putMultipleFilesAndDirectories(c w3s.Client) cid.Cid {
	f0, err := os.Open("images/donotresist.jpg")
	if err != nil {
		panic(err)
	}
	f1, err := os.Open("images/pinpie.jpg")
	if err != nil {
		panic(err)
	}
	d0 := w3fs.NewDir("one", []fs.File{f0})
	d1 := w3fs.NewDir("two", []fs.File{f1})
	rootdir := w3fs.NewDir("comic", []fs.File{d0, d1})
	return putFile(c, rootdir)
}

func putDirectory(c w3s.Client) cid.Cid {
	dir, err := os.Open("images")
	if err != nil {
		panic(err)
	}
	return putFile(c, dir)
}

func putFile(c w3s.Client, f fs.File, opts ...w3s.PutOption) cid.Cid {
	cid, err := c.Put(context.Background(), f, opts...)
	if err != nil {
		panic(err)
	}
	fmt.Printf("https://%v.ipfs.dweb.link\n", cid)
	return cid
}

func getStatusForCid(c w3s.Client, cid cid.Cid) {
	s, err := c.Status(context.Background(), cid)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %+v", s)
}

func getStatusForKnownCid(c w3s.Client) {
	cid, _ := cid.Parse("bafybeiauyddeo2axgargy56kwxirquxaxso3nobtjtjvoqu552oqciudrm")
	getStatusForCid(c, cid)
}

func getFiles(c w3s.Client) {
	cid, _ := cid.Parse("bafybeide43vps6vt2oo7nbqfwn5zz6l2alyi64mym3sb7reqhmypjnmej4")

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
