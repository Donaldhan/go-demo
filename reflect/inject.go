package reflect

import (
	"fmt"
	"log"

	"github.com/codegangsta/inject"
)

// 正常情况下，对函数或方法的调用是我们的主动直接行为，在调用某个函数之前我们需要清楚地知道被调函数的名称是什么，参数有哪些类型等等。

// 所谓的控制反转就是将这种主动行为变成间接的行为，我们不用直接调用函数或对象，而是借助框架代码进行间接的调用和初始化，这种行为称作“控制反转”，库和框架能很好的解释控制反转的概念。

// 依赖注入是实现控制反转的一种方法，如果说控制反转是一种设计思想，那么依赖注入就是这种思想的一种实现，通过注入参数或实例的方式实现控制反转。
// 如果没有特殊说明，我们可以认为依赖注入和控制反转是一个东西。

// 控制反转的价值在于解耦，有了控制反转就不需要将代码写死，可以让控制反转的的框架代码读取配置，动态的构建对象，这一点在 Java 的 Spring 框架中体现的尤为突出。
// inject 实践
// inject 是依赖注入的Go语言实现，它能在运行时注入参数，调用方法，是 Martini 框架（Go语言中著名的 Web 框架）的基础核心。
// http: //c.biancheng.net/view/5132.html

// https://github.com/go-martini/martini

type S1 interface{}
type S2 interface{}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("Format name ＝ %s, company=%s, level=%s, age ＝ %d!\n", name, company, level, age)
}
func InjectInvoke() {
	//控制实例的创建
	inj := inject.New()
	//实参注入
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)
	//函数反转调用
	inj.Invoke(Format)
	log.Println("InjectInvoke Format done==============")
}

type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

// http://c.biancheng.net/view/5132.html
func InjectInvokeStruct() {
	//创建被注入实例
	s := Staff{}
	//控制实例的创建
	inj := inject.New()
	//初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)
	//实现对 struct 注入
	inj.Apply(&s)
	//打印结果
	fmt.Printf("InjectInvokeStruct:s = %v\n", s)
	log.Println("InjectInvokeStruct done==============")
}
