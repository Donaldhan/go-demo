package interfacex

import (
	"io"
	"log"
)

type Socket struct {
}

func (s *Socket) Write(p []byte) (n int, err error) {
	log.Println("====Socket Write====")
	return 0, nil
}
func (s *Socket) Close() error {
	log.Println("====Socket Close====")
	return nil
}

type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser(closer io.Closer) {
	closer.Close()
}

// 一个类型可以实现多个接口
func SocketTest() {
	// 实例化Socket
	s := new(Socket)
	usingWriter(s)
	usingCloser(s)
}

// 一个服务需要满足能够开启和写日志的功能
type Service interface {
	Start()     // 开启服务
	Log(string) // 日志输出
}

// 日志器
type Logger struct {
}

// 实现Service的Log()方法
func (g *Logger) Log(loginfo string) {
	log.Println("Logger Log", loginfo)
}

// 游戏服务
type GameService struct {
	Logger // 嵌入日志器
}

// 实现Service的Start()方法
func (g *GameService) Start() {
	log.Println("====GameService Start====")
}

// 多个类型可以实现相同的接口
func ServiceTest() {
	var s Service = new(GameService)
	s.Start()
	s.Log("ServiceTest hello")
}
