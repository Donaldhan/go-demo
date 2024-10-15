package interfacex

import (
	"fmt"
	"log"
)

func init() {
	log.Println("==============interface package init")
}

// 当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用。
// 实现接口的方法签名不一致导致的报错
// 函数名不一致导致的报错

// 定义一个数据写入器
type DataWriter interface {
	// WriteData()，输入一个 interface{} 类型的 data，返回一个 error 结构表示可能发生的错误。
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData方法

// file 的 WriteData() 方法使用指针接收器。输入一个 interface{} 类型的 data，返回 error。
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

// 接口被实现的条件一：接口的方法与实现接口的类型方法格式一致
func IntefacexTest() {
	// 实例化file
	f := new(file)
	// 声明一个DataWriter的接口
	var writer DataWriter
	// 将接口赋值f，也就是*file类型
	writer = f
	// 使用DataWriter接口进行数据写入
	writer.WriteData("data")
}
