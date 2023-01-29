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
func ChanSelectBase() {
	taskCh1 := make(chan bool)
	taskCh2 := make(chan bool)
	taskCh3 := make(chan bool)

	go func() {
		for {
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
	taskCh2 <- false
	taskCh3 <- false
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

	select {
	case data := <-ch:
		fmt.Printf("ChanSelectBaseX case invoke %+v\n", data)
	case <-quitCh:
		fmt.Println("ChanSelectBaseX program quit")

	}
}

func ChanSelectTimeout() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
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

	<-quit
	fmt.Println("程序结束")
}

// 设置一个可用的 case，让 select 变成非阻塞
func ChanSelectWithDefalut() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("ChannSelectWithDefalut case invoke")
			case num := <-ch:
				fmt.Println("ChannSelectWithDefalut case invoke num:", num)
			default: // 设置一个可用的 case，让 select 变成非阻塞
				time.Sleep(time.Second)
				fmt.Println("ChannSelectWithDefalut default invoke")
			}
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
}
