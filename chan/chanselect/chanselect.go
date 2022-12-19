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

func ChanSelectTest() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("======chanselect ChanSelectTest 超时")
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
