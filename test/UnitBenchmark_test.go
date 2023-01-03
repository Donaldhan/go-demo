package UnitBenchmark

import "testing"

// http://c.biancheng.net/view/5409.html
// Go语言自带了 testing 测试包，可以进行自动化的单元测试，输出结果验证，并且可以测试性能。
// 为什么需要测试
// 完善的测试体系，能够提高开发的效率，当项目足够复杂的时候，想要保证尽可能的减少 bug，有两种有效的方式分别是代码审核和测试，Go语言中提供了 testing 包来实现单元测试功能。
// 测试规则
// 要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时文件名必须以_test.go结尾，单元测试源码文件可以由多个测试用例（可以理解为函数）组成，每个测试用例的名称需要以 Test 为前缀，例如：
// func TestXxx( t *testing.T ){
//     //......
// }

// 编写测试用例有以下几点需要注意：
// 测试用例文件不会参与正常源码的编译，不会被包含到可执行文件中；
// 测试用例的文件名必须以_test.go结尾；
// 需要使用 import 导入 testing 包；
// 测试函数的名称要以Test或Benchmark开头，后面可以跟任意字母组成的字符串，但第一个字母必须大写，例如 TestAbc()，一个测试用例文件中可以包含多个测试函数；
// 单元测试则以(t *testing.T)作为参数，性能测试以(t *testing.B)做为参数；
// 测试用例文件使用go test命令来执行，源码中不需要 main() 函数作为入口，所有以_test.go结尾的源码文件内以Test开头的函数都会自动执行。

// Go语言的 testing 包提供了三种测试方式，分别是单元（功能）测试、性能（压力）测试和覆盖率测试。

// 单元（功能）测试
// go test -v
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}

func TestGetAreaX(t *testing.T) {
	area := GetArea(40, 50)
	if area != 3000 {
		t.Error("TestGetAreaX 测试失败")
	}
}

// go test -bench="."
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}

// go test -cover
// 覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码），可以的话最好是覆盖100%。
