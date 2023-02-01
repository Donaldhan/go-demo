package chanselect

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.Println("==============ChanSelectTest package init")
}

// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况。

// 如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有如下两种可能的情况：
// 1. 如果给出了 default 语句，那么就会执行 default 语句，同时程序的执行会从 select 语句后的语句中恢复；
// 2. 如果没有 default 语句，那么 select 语句将被阻塞，直到至少有一个通信可以进行下去。
// Go并发编程-channel多路复用：https://juejin.cn/post/6970302437571706911
// go 语言管道多路复用：https://www.lesscode.work/sections/629d7f8bc65be.html

/*
Golang并发：一招掌握无阻塞通道读写:https://segmentfault.com/a/1190000017537297
Go: Select 语句的执行顺序:https://studygolang.com/articles/28990
无缓冲和有缓冲通道:https://studygolang.com/articles/23538
Using Go Channels - Part 3 - Read/Write Channels & Select:https://g14a.dev/posts/using-Go-Channels-p3/

通道的数据接收一共有以下 4 种写法。
1) 阻塞接收数据
阻塞模式接收数据时，将接收变量作为<-操作符的左值，格式如下：
data := <-ch

执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量。
2) 非阻塞接收数据
使用非阻塞方式从通道接收数据时，语句不会发生阻塞，格式如下：
data, ok := <-ch

data：表示接收到的数据。未接收到数据时，data 为通道类型的零值。
ok：表示是否接收到数据。

非阻塞的通道接收方法可能造成高的 CPU 占用，因此使用非常少。如果需要实现接收超时检测，可以配合 select 和计时器 channel 进行，可以参见后面的内容。
3) 接收任意数据，忽略接收的数据
阻塞接收数据后，忽略从通道返回的数据，格式如下：
<-ch

执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步。

Go语言中有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。
这种类型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。通道会阻塞发送和接收动作的条件也会不同。只有在通道中没有要接收的值时，接收动作才会阻塞。只有在通道没有可用缓冲区容纳被发送的值时，发送动作才会阻塞。

这导致有缓冲的通道和无缓冲的通道之间的一个很大的不同：无缓冲的通道保证进行发送和接收的
goroutine 会在同一时间进行数据交换；有缓冲的通道没有这种保证。

在无缓冲通道的基础上，为通道增加一个有限大小的存储空间形成带缓冲通道。带缓冲通道在发送时无需等待接收方接收即可完成发送过程，
并且不会发生阻塞，只有当存储空间满时才会发生阻塞。同理，如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞。

无缓冲通道保证收发过程同步。无缓冲收发过程类似于快递员给你电话让你下楼取快递，整个递交快递的过程是同步发生的，你和快递员不见不散。
但这样做快递员就必须等待所有人下楼完成操作后才能完成所有投递工作。如果快递员将快递放入快递柜中，并通知用户来取，快递员和用户就成了异步收发过程，效率可以有明显的提升。带缓冲的通道就是这样的一个“快递柜”。

http://c.biancheng.net/view/100.html

阻塞条件
带缓冲通道在很多特性上和无缓冲通道是类似的。无缓冲通道可以看作是长度永远为 0 的带缓冲通道。因此根据这个特性，带缓冲通道在下面列举的情况下依然会发生阻塞：
1. 带缓冲通道被填满时，尝试再次发送数据时发生阻塞。
2. 带缓冲通道为空时，尝试接收数据时发生阻塞。

为什么Go语言对通道要限制长度而不提供无限长度的通道？

我们知道通道（channel）是在两个 goroutine 间通信的桥梁。使用 goroutine 的代码必然有一方提供数据，一方消费数据。
当提供数据一方的数据供给速度大于消费方的数据处理速度时，如果通道不限制长度，那么内存将不断膨胀直到应用崩溃。
因此，限制通道的长度有利于约束数据提供方的供给速度，供给数据量必须在消费方处理量+通道长度的范围内，才能正常地处理数据
*/
func ChanSelectBase() {
	taskCh1 := make(chan bool)
	taskCh2 := make(chan bool)
	taskCh3 := make(chan bool)

	go func() {
		for {
			// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
			select {
			// 接收通道 1 的结果
			case r := <-taskCh1:
				fmt.Printf("ChanSelectBase task1 result %+v\n", r)
			// 接收通道 2 的结果
			case r := <-taskCh2:
				fmt.Printf("ChanSelectBase task2 result %+v\n", r)
			// 接收通道 3 的结果
			case r := <-taskCh3:
				fmt.Printf("ChanSelectBase task3 result %+v\n", r)
			}
		}
	}() //匿名方法
	taskCh1 <- true
	// taskCh2 <- false
	taskCh3 <- false
	fmt.Println("ChanSelectBase 程序结束")
}

func ChanSelectBaseX() {
	ch := make(chan int, 1)

	quitCh := make(chan bool, 1)

	go func(ch chan bool) {
		var quit string
		fmt.Printf("ChanSelectBaseX quit? are you sure?: ")
		fmt.Scanln(&quit)
		quitCh <- true
	}(quitCh)
	// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
	select {
	case data := <-ch:
		fmt.Printf("ChanSelectBaseX case invoke %+v\n", data)
	case <-quitCh:
		fmt.Println("ChanSelectBaseX program quit")

	}
	fmt.Println("ChanSelectBaseX 程序结束")
}

func ChanSelectTimeout() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
			select {
			case num := <-ch: //从通道读取数据
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second): //超时
				fmt.Println("======ChanSelectTimeout 超时")
				quit <- true
			}
		}

	}() //别忘了()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit //阻塞等待，忽略结果
	fmt.Println("ChanSelectTimeout 程序结束")
}

// 设置一个可用的 case，让 select 变成非阻塞
func ChanSelectWithDefalut() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
			select {
			case <-ch:
				fmt.Println("ChannSelectWithDefalut case invoke")
			case num := <-ch:
				fmt.Println("ChannSelectWithDefalut case invoke num:", num)
			default: // 设置一个可用的 case，让 select 变成非阻塞
				time.Sleep(time.Second)
				fmt.Println("ChannSelectWithDefalut default invoke")
				quit <- true
			}
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	<-quit
	fmt.Println("ChanSelectWithDefalut 程序结束")
}

// https://g14a.dev/posts/using-Go-Channels-p3/
// 带缓冲buffer
func BufferSelect() {
	helloChan := make(chan string, 5)
	worldChan := make(chan string, 5)
	timeoutChan := make(chan bool, 1)

	go hello(helloChan)
	go world(worldChan)

	go func() {
		time.Sleep(time.Second * 10)
		timeoutChan <- true
	}()

	for {
		select {
		case msg := <-helloChan:
			fmt.Println(msg)
		case msg := <-worldChan:
			fmt.Println(msg)
		case <-timeoutChan:
			fmt.Println("BufferSelect It has been 3 seconds. Timeout!")
			return
		}
	}
}

func hello(helloChan chan<- string) {
	for {
		time.Sleep(time.Second * 1)
		helloChan <- "Hello"
	}
}

func world(worldChan chan<- string) {
	for {
		time.Sleep(time.Second * 2)
		worldChan <- "World!"
	}
}
