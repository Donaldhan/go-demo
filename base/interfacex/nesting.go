package interfacex

import (
	"io"
	"log"
)

// 声明一个设备结构
type device struct {
}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	log.Println("device Write")
	return 0, nil
}

// 实现io.Closer的Close()方法
func (d *device) Close() error {
	log.Println("device Close")
	return nil
}

// 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用。
func NestingTest() {
	// 声明写入关闭器, 并赋予device的实例
	var wc io.WriteCloser = new(device)
	// 写入数据
	wc.Write(nil)
	// 关闭设备
	wc.Close()
	// 声明写入器, 并赋予device的新实例
	var writeOnly io.Writer = new(device)
	// 写入数据
	writeOnly.Write(nil)
}
