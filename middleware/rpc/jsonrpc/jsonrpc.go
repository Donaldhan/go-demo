package jsonrpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}
type Rect struct {
}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}
func jsonRpcServer() {
	rpc.Register(new(Rect))
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
func jsonRpcClient() {
	conn, err := jsonrpc.Dial("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Panicln(err2)
	}
	fmt.Println("面积：", ret)
	err3 := conn.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err3 != nil {
		log.Panicln(err3)
	}
	fmt.Println("周长：", ret)
}
