package greeter

import (
	"context"
	"time"

	"go-micro.dev/v4/util/log"
	hello "godemo/micro/gomicro/greeter"
	"google.golang.org/grpc"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	go func() {
		for {
			grpc.DialContext(context.TODO(), "127.0.0.1:9091")
			time.Sleep(time.Second)
		}
	}()

	service := micro.NewService(
		micro.Name("godemo.micro.greeter"),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	hello.RegisterSayHandler(service.Server(), new(Say))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
