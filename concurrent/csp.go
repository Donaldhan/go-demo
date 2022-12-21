package concurrent

import (
	"fmt"
	"log"
	"sync"
)

// Go实现了两种并发形式，第一种是大家普遍认知的多线程共享内存，其实就是 Java 或 C++ 等语言中的多线程开发；另外一种是Go语言特有的，也是Go语言推荐的 CSP（communicating sequential processes）并发模型。
// CSP 并发模型是上个世纪七十年代提出的，用于描述两个独立的并发实体通过共享 channel（管道）进行通信的并发模型。
// Go语言就是借用 CSP 并发模型的一些概念为之实现并发的，但是Go语言并没有完全实现了 CSP 并发模型的所有理论，仅仅是实现了 process 和 channel 这两个概念。
// process 就是Go语言中的 goroutine，每个 goroutine 之间是通过 channel 通讯来实现数据共享。
// 这里我们要明确的是“并发不是并行”。并发更关注的是程序的设计层面，并发的程序完全是可以顺序执行的，只有在真正的多核 CPU 上才可能真正地同时运行；并行更关注的是程序的运行层面，并行一般是简单的大量重复，例如 GPU 中对图像处理都会有大量的并行运算。
// 为了更好地编写并发程序，从设计之初Go语言就注重如何在编程语言层级上设计一个简洁安全高效的抽象模型，让开发人员专注于分解问题和组合方案，而且不用被线程管理和信号互斥这些烦琐的操作分散精力。
// 在并发编程中，对共享资源的正确访问需要精确地控制，在目前的绝大多数语言中，都是通过加锁等线程同步方案来解决这一困难问题，而Go语言却另辟蹊径，它将共享的值通过通道传递（实际上多个独立执行的线程很少主动共享资源）。
func CspSyncMutexTest() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("concurrent CspSyncMutexTest Printer")
		mu.Unlock()
	}()
	mu.Lock()
	log.Println("concurrent CspSyncMutexTest done")
}
func CspChannWithoutBufTest() {
	done := make(chan int) //无缓冲区
	go func() {
		fmt.Println("concurrent CspChannWithoutBufTest Printer")
		<-done
	}()
	done <- 1
	log.Println("concurrent CspChannWithoutBufTest done")
}
func CspChannWithBufTest() {
	done := make(chan int, 1) // 带缓存通道
	go func() {
		fmt.Println("concurrent CspChannWithBufTest Printer")
		done <- 1
	}()
	<-done
	log.Println("concurrent CspChannWithBufTest done")
}
func CspMultiChannTest() {
	done := make(chan int, 10) // 带10个缓存
	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("concurrent CspMultiChannTest Printer")
			done <- 1
		}()
	}
	// 等待N个后台线程完成
	for i := 0; i < cap(done); i++ {
		<-done
	}
	log.Println("concurrent CspMultiChannTest done")
}
func CspWaitGroupTest() {
	var wg sync.WaitGroup
	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("concurrent CspWaitGroupTest Printer")
			wg.Done()
		}()
	}
	// 等待N个后台线程完成
	wg.Wait()
	log.Println("concurrent CspWaitGroupTest done")
}
