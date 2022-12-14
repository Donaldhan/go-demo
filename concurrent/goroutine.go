package concurrent

import (
	"fmt"
	"time"
)

// C#、Lua、Python 语言都支持 coroutine 特性。coroutine 与 goroutine 在名字上类似，都可以将函数或者语句在独立的环境中运行，但是它们之间有两点不同：
// goroutine 可能发生并行执行；
// 但 coroutine 始终顺序执行。
// goroutines 意味着并行（或者可以以并行的方式部署），coroutines 一般来说不是这样的，goroutines 通过通道来通信；coroutines 通过让出和恢复操作来通信，goroutines 比 coroutines 更强大，也很容易从 coroutines 的逻辑复用到 goroutines。
// 狭义地说，goroutine 可能发生在多线程环境下，goroutine 无法控制自己获取高优先度支持；coroutine 始终发生在单线程，coroutine 程序需要主动交出控制权，宿主才能获得控制权并将控制权交给其他 coroutine。
// goroutine 间使用 channel 通信，coroutine 使用 yield 和 resume 操作。
// goroutine 和 coroutine 的概念和运行机制都是脱胎于早期的操作系统。
// coroutine 的运行机制属于协作式任务处理，早期的操作系统要求每一个应用必须遵守操作系统的任务处理规则，应用程序在不需要使用 CPU 时，会主动交出 CPU 使用权。如果开发者无意间或者故意让应用程序长时间占用 CPU，操作系统也无能为力，表现出来的效果就是计算机很容易失去响应或者死机。
// goroutine 属于抢占式任务处理，已经和现有的多线程和多进程任务处理非常类似。应用程序对 CPU 的控制最终还需要由操作系统来管理，操作系统如果发现一个应用程序长时间大量地占用 CPU，那么用户有权终止这个任务。

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("concurrent running tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
func BacktheadTest() {
	// 并发执行程序
	go running()
	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
}
