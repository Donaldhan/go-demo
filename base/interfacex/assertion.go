package interfacex

import (
	"fmt"
)

// 需要注意如果不接收第二个参数也就是上面代码中的 ok，断言失败时会直接造成一个 panic。如果 x 为 nil 同样也会 panic。
func AssertTest() {
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println(value, ",", ok)

	var a int
	a = 10
	getType(a)

	var y interface{}
	y = "Hello"
	valuey := y.(int)
	fmt.Println(valuey)

}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
