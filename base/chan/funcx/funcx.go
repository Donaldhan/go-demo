package funcx

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	log.Println("============funcx package init=========")
}

/*
ch := make(chan int)
// 声明一个只能写入数据的通道类型, 并赋值为ch
var chSendOnly chan<- int = ch
//声明一个只能读取数据的通道类型, 并赋值为ch
var chRecvOnly <-chan int = ch
*/

func CreateChan() chan int {
	// 构建一个通道
	ch := make(chan int, 3)
	return ch
}

func ReturnChan() {
	ch := CreateChan()
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("ReturnChan start goroutine")
		ch <- 1
		ch <- 2
		ch <- 3
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("ReturnChan exit goroutine")
	}()
	log.Println("ReturnChan wait goroutine")
	// 等待匿名goroutine
	<-ch
	log.Println("ReturnChan all done")
}
