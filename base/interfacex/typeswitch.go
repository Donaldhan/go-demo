package interfacex

import "fmt"

// 定义飞行动物接口
type Flyer interface {
	Fly()
}

// 定义行走动物接口
type Walker interface {
	Walk()
}

// 定义鸟类
type bird struct {
}

// 实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 为鸟添加Walk()方法, 实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// 定义猪
type pig struct {
}

// 为猪添加Walk()方法, 实现行走动物接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

// interface{} 类型表示空接口，意思就是这种接口可以保存为任意类型
// 接口和其他类型的转换可以在Go语言中自由进行，前提是已经完全实现。
// 接口断言类似于流程控制中的 if。但大量类型断言出现时，应使用更为高效的类型分支 switch 特性。
func TypeswitchTest() {
	// 创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	// 遍历映射
	for name, obj := range animals {
		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer)
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)
		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}
		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}

	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)
	fmt.Printf("TypeswitchTest p1=%p p2=%p\n", p1, p2)

	p3 := a.(*bird)
	fmt.Println("TypeswitchTest p1=%p p3=%p", p1, p3)
}

// 空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。
// 提示
// 空接口类型类似于 C# 或 Java 语言中的 Object、C语言中的 void*、C++ 中的 std::any。在泛型和模板出现前，空接口是一种非常灵活的数据抽象保存和使用的方法。

// 空接口的内部实现保存了对象的类型和指针。使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。因此在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。
func NullTest() {
	var any interface{}
	any = 1
	fmt.Println("NullTest:", any)
	any = "hello"
	fmt.Println("NullTest:", any)
	any = false
	fmt.Println("NullTest:", any)
}

func printType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	}
}

// type-switch 流程控制的语法或许是Go语言中最古怪的语法。 它可以被看作是类型断言的增强版。它和 switch-case 流程控制代码块有些相似

func SwitchTypeTest() {
	printType(1024)
	printType("pig")
	printType(true)
}

// 电子支付方式
type Alipay struct {
}

// 为Alipay添加CanUseFaceID()方法, 表示电子支付方式支持刷脸
func (a *Alipay) CanUseFaceID() {
}

// 现金支付方式
type Cash struct {
}

// 为Cash添加Stolen()方法, 表示现金支付方式会出现偷窃情况
func (a *Cash) Stolen() {
}

// 具备刷脸特性的接口
type CantainCanUseFaceID interface {
	CanUseFaceID()
}

// 具备被偷特性的接口
type ContainStolen interface {
	Stolen()
}

// 打印支付方式具备的特点
func print(payMethod interface{}) {
	switch payMethod.(type) {
	case CantainCanUseFaceID: // 可以刷脸
		fmt.Printf("%T can use faceid\n", payMethod)
	case ContainStolen: // 可能被偷
		fmt.Printf("%T may be stolen\n", payMethod)
	}
}
func SwitchTypeTestX() {
	// 使用电子支付判断
	print(new(Alipay))
	// 使用现金判断
	print(new(Cash))
}
