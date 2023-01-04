package structx

import (
	"fmt"
	"log"
)

func init() {
	log.Println("==============structx package init")
}

// 打印消息类型, 传入匿名结构体
func printMsgType(msg *struct {
	id   int
	data string
}) {
	// 使用动词%T打印msg的类型
	fmt.Printf("==============structx package printMsgType %T\n", msg)
}

// 匿名结构体
func AnonymousTest() {
	// 实例化一个匿名结构体
	msg := &struct { // 定义部分
		id   int
		data string
	}{ // 值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}

// Go语言的类型或结构体没有构造函数的功能，但是我们可以使用结构体初始化的过程来模拟实现构造函数。

// 其他编程语言构造函数的一些常见功能及特性如下：
// 每个类可以添加构造函数，多个构造函数使用函数重载实现。
// 构造函数一般与类名同名，且没有返回值。
// 构造函数有一个静态构造函数，一般用这个特性来调用父类的构造函数。
// 对于 C++ 来说，还有默认构造函数、拷贝构造函数等。

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

// 带有父子关系的结构体的构造和初始化——模拟父级构造调用
// 黑猫是一种猫，猫是黑猫的一种泛称，同时描述这两种概念时，就是派生，黑猫派生自猫的种类，使用结构体描述猫和黑猫的关系时，
// 将猫（Cat）的结构体嵌入到黑猫（BlackCat）中，表示黑猫拥有猫的特性，然后再使用两个不同的构造函数分别构造出黑猫和猫两个结构体实例，参考下面的代码

type BlackCat struct {
	Cat // 嵌入Cat, 类似于派生
}

// “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

// 车轮
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

// 车
type Car struct {
	Wheel
	Engine
}

// 结构体内嵌初始化时，将结构体内嵌的类型作为字段名像普通结构体一样进行初始化
func InitStructTest() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("InitStructTest:%+v\n", c)
}

// / 车轮
type WheelX struct {
	Size int
}

// 车
type CarX struct {
	WheelX
	// 引擎
	EngineX struct {
		Power int    // 功率
		Type  string // 类型
	}
}

// 初始化内嵌匿名结构体
func InitInnerStructTest() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.5T",
			Power: 158,
		},
	}
	fmt.Printf("InitInnerStructTest:%+v\n", c)
}
