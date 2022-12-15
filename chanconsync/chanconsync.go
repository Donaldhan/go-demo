package chanconsync

import (
	"fmt"
	"time"
)

// 通道做并发同步的写法
func TestChanconsync() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("chanconsync TestChanconsync start goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("chanconsync TestChanconsync exit goroutine")
	}()
	fmt.Println("chanconsync TestChanconsync wait goroutine")
	// 等待匿名goroutine
	<-ch
	fmt.Println("chanconsync TestChanconsync all done")
}

// 通道的数据接收可以借用 for range 语句进行多个元素的接收操作
func ChanFor() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i
			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()
	// 遍历接收通道数据, 阻塞模式接收数据时，将接收变量作为<-操作符的左值，格式如下；data := <-ch
	// 使用非阻塞方式从通道接收数据时，语句不会发生阻塞，格式如下：data, ok := <-ch
	for data := range ch {
		// 打印通道数据
		fmt.Println("chanconsync ChanFor data:", data)
		// 当遇到数据0时, 退出接收循环
		if data == 0 {
			break
		}
	}

}

func printer(c chan int) {

	// 开始无限循环等待数据
	for {
		// 使用非阻塞方式从通道接收数据时，语句不会发生阻塞，格式如下：data, ok := <-ch
		// 从channel中获取一个数据
		data := <-c
		// 将0视为数据结束
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println("chanconsync ConPrinter data:", data)
	}
	// 通知main已经结束循环(我搞定了!)
	c <- 0

}

// 创建的都是无缓冲通道。使用无缓冲通道往里面装入数据时，装入方将被阻塞，
// 直到另外通道在另外一个 goroutine 中被取出。同样，如果通道中没有放入任何数据，接收方试图从通道中获取数据时，同样也是阻塞。发送和接收的操作是同步完成的。
// 并发打印（借助通道实现）
func ConPrinter() {
	// 创建一个channel
	c := make(chan int)
	// 并发执行printer, 传入channel
	go printer(c)
	for i := 1; i <= 10; i++ {
		// 将数据通过channel投送给printer
		c <- i
	}
	// 通知并发的printer结束循环(没数据啦!)
	c <- 0
	// 等待printer结束(搞定喊我!)
	<-c
}
