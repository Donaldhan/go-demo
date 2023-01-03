package function

import (
	"fmt"
	"runtime"
	"time"
)

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string // 所在函数
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("================function ProtectRun runtime error:", err)
		default: // 非运行时错误
			fmt.Println("================function ProtectRun error:", err)
		}
	}()
	entry()
}

// Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover
// 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。

// 通常来说，不应该对进入 panic 宕机的程序做任何处理，但有时，需要我们可以从宕机中恢复，至少我们可以在程序崩溃前，做一些操作，举个例子，
// 当 web 服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭，如果不做任何处理，会使得客户端一直处于等待状态，如果 web 服务器还在开发阶段，
// 服务器甚至可以将异常信息反馈到客户端，帮助调试。
// 提示
// 在其他语言里，宕机往往以异常的形式存在，底层抛出异常，上层逻辑通过 try/catch 机制捕获异常，没有被捕获的严重异常会导致宕机，捕获的异常可以被忽略，让代码继续运行。

// Go语言没有异常系统，其使用 panic 触发宕机类似于其他语言的抛出异常，recover 的宕机恢复机制就对应其他语言中的 try/catch 机制。

// panic 和 recover 的关系
// panic 和 recover 的组合有如下特性：
// 有 panic 没 recover，程序宕机。
// 有 panic 也有 recover，程序不会宕机，执行完对应的 defer 后，从宕机点退出当前函数后继续执行。
// 提示
// 虽然 panic/recover 能模拟其他语言的异常机制，但并不建议在编写普通函数时也经常性使用这种特性。

// 在 panic 触发的 defer 函数内，可以继续调用 panic，进一步将错误外抛，直到程序整体崩溃。

// 如果想在捕获错误时设置当前函数的返回值，可以对返回值使用命名返回值方式直接进行设置。

func RecoverFromPanicTest() {
	fmt.Println("================function RecoverFromPanicTest 运行前")
	// 允许一段手动触发的错误
	ProtectRun(func() {
		fmt.Println("================function RecoverFromPanicTest 手动宕机前")
		// 使用panic传递上下文
		panic(&panicContext{
			"RecoverFromPanicTest 手动触发panic",
		})
		fmt.Println("================function RecoverFromPanicTest 手动宕机后")
	})
	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("================function RecoverFromPanicTest 赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("================function RecoverFromPanicTest 赋值宕机后")
	})
	fmt.Println("================function RecoverFromPanicTest 运行后")
}

// 函数的运行时间的长短是衡量这个函数性能的重要指标，特别是在对比和基准测试中，要得到函数的运行时间，最简单的办法就是在函数执行之前设置一个起始时间，并在函数运行结束时获取从起始时间到现在的时间间隔，这个时间间隔就是函数的运行时间。

// 在Go语言中我们可以使用 time 包中的 Since() 函数来获取函数的运行时间，Go语言官方文档中对 Since() 函数的介绍是这样的

func TimeConsumeTest() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时 Since：", elapsed)
	elapsed = time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时 Sub：", elapsed)
}
