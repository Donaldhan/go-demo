package server

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"greeterdemo/greet"
)

func StartServer() {
	var serverConf zrpc.RpcServerConf
	conf.MustLoad("etc/greet-server.yaml", &serverConf)
	s := zrpc.MustNewServer(serverConf, func(server *grpc.Server) {
		greet.RegisterGreetServer(server, &exampleServer{})
	})

	defer s.Stop()
	s.Start()
}

type exampleServer struct {
	greet.UnimplementedGreetServer
}

func (e *exampleServer) Ping(ctx context.Context, request *greet.Request) (*greet.Response, error) {
	// fill your logic here
	return &greet.Response{Pong: "pong"}, nil
}
