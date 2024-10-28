package os

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func init() {
	log.Println("==============packagex os package init")
}

// os/signal 信号处理
// 一个运行良好的程序在退出（正常退出或者强制退出，如 Ctrl+C，kill 等）时是可以执行一段清理代码的，将收尾工作做完后再真正退出。一般采用系统 Signal 来通知系统退出，如 kill pid，在程序中针对一些系统信号设置了处理函数，当收到信号后，会执行相关清理程序或通知各个子进程做自清理。

// Go语言中对信号的处理主要使用 os/signal 包中的两个方法，一个是 Notify 方法用来监听收到的信号，一个是 stop 方法用来取消监听。

// 使用 Notify 方法来监听收到的信号
func SignalTest() {
	c := make(chan os.Signal, 0)
	go func() {
		signal.Notify(c)
	}()
	// Block until a signal is received.
	s := <-c
	fmt.Println("SignalTest Got signal:", s)

}

// 因为使用 Stop 方法取消了 Notify 方法的监听，所以运行程序没有输出结果
func SignalStopTest() {
	c := make(chan os.Signal, 0)
	signal.Notify(c)
	signal.Stop(c) //不允许继续往c中存入内容
	s := <-c       //c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
	fmt.Println("SignalStopTest Got signal:", s)
}
