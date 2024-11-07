package client

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"golang.org/x/net/context"
	"greeterdemo/greet"
	"log"
)

func StartClient() {
	var clientConf zrpc.RpcClientConf
	conf.MustLoad("etc/greet-client.yaml", &clientConf)
	conn := zrpc.MustNewClient(clientConf)
	client := greet.NewGreetClient(conn.Conn())
	resp, err := client.Ping(context.Background(), &greet.Request{})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp)
}
