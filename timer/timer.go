package timer

//单元测试demo
import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.Println("==============timer package init")
}

// [Go高阶20，定时器的使用](https://juejin.cn/post/7028191842629845029)

// Timer 是一种单一事件定时器，就是说 Timer 只执行一次就会结束。
// 创建：
// time.NewTimer(d Duration) ：创建一个 timer
// 参数为等待事件
// 时间到来后立即触发一个事件

// type Timer struct {
// 	C <-chan Time
// 	r runtimeTimer
// }

func TimerDemo() {
	timer := time.NewTimer(3 * time.Second)
	// timer.Reset(3 * time.Second)
	defer timer.Stop()
}

// 比如在一个连接中等待数据，设定一个超时时间，当时间到来还是没有数据获取到，则为超时。
// 一次性超时等待
func TimerWaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(3 * time.Second)
	//阻塞模式，一次性
	select {
	case <-conn: //读取通道数据
		timer.Stop()
		fmt.Println("TimerWaitChannel recive data from conn,stop timer")
		return true
	case <-timer.C: //超时，从通道内读取time， 有表示触发发定时任务
		fmt.Println("TimerWaitChannel timeout")
		return false
	}
}

// 循环等待
func TimerWaitChannelX(conn <-chan int) bool {
	timer := time.NewTimer(3 * time.Second)
	for {
		//循环非阻塞模式
		select {
		case msg := <-conn: //读取通道数据
			timer.Stop()
			fmt.Println("TimerWaitChannelX recive data from conn,stop timer msg:", msg)
			return true
		case <-timer.C: //超时，从通道内读取time， 有表示触发发定时任务
			fmt.Println("TimerWaitChannelX timeout")
			return false
		default: //设置一个可用的 case，让 select 变成非阻塞
			time.Sleep(time.Second)
			fmt.Println("TimerWaitChannelX default invoke")
		}
	}
}
func TimerWaitChannelXX(conn <-chan string) bool {
	timer := time.NewTimer(3 * time.Second)
	for {
		//循环非阻塞模式
		select {
		case msg := <-conn: //读取通道数据
			timer.Stop()
			fmt.Println("TimerWaitChannelXX recive data from conn,stop timer msg:", msg)
			return true
		default: //设置一个可用的 case，让 select 变成非阻塞
			time.Sleep(time.Second)
			fmt.Println("TimerWaitChannelXX default invoke")
		}
	}
}

// 超时
func TimerWaitChannelDemo() {
	ch := make(chan string, 3)
	TimerWaitChannel(ch)
}

// 超时
func TimerWaitChannelDemoX() {
	ch := make(chan string, 3)
	TimerWaitChannel(ch)
	ch <- "hello"
}

// 缓存区先写入数据，先取到数据，不会超时,
func TimerWaitChannelDemoXX() {
	ch := make(chan string, 3)
	ch <- "hello"
	TimerWaitChannel(ch)

}

// 后台协程等待，同时写入数据到连接通道，可能取到数据，不会超时
func TimerWaitChannelDemoXXX() {
	ch := make(chan string, 3)
	go TimerWaitChannel(ch)
	ch <- "hello"
}

// 循环select,非阻塞，后写入，不会读到数据，3秒后超时
func TimerWaitChannelDemoXXXX() {
	ch := make(chan int, 3)
	TimerWaitChannelX(ch)
	// ch <- "hello" //无法读取数据
	/*
		for i := 0; i < 5; i++ {
			ch <- i //等待缓冲区可用,阻塞
		}
	*/
	for i := 0; i < 3; i++ {
		ch <- i
	}
}

// 循环select,非阻塞，后写入，不会读的数据，无超时，不断执行默认case
func TimerWaitChannelDemoXXXXX() {
	ch := make(chan string, 3)
	TimerWaitChannelXX(ch)
	// ch <- "hello" //无法读取数据
	for i := 0; i < 3; i++ {
		ch <- "hello"
	}
}

// 希望某个方法在今后的某个时刻执行
func TimerDelayFunction() {
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C: //从通道内读取time， 有表示触发发定时任务
		fmt.Println("TimerDelayFunction Delayed 5s，...")
	}
}

// 有时我们就是想等指定的时间，没有需求提前停止定时器，也没有需求复用该定时器，那么可以使用匿名的定时器
func TimerAfterDemo() {
	log.Println("TimerAfterDemo now", time.Now())
	timex := <-time.After(2 * time.Second)
	// <-time.After(1 * time.Second)
	log.Println("TimerAfterDemo After timex", timex.String())
	log.Println("TimerAfterDemo After", time.Now())
}

// 我们可以使用 AfterFunc 更加简洁的实现延迟一个方法的调用
func TimerAfterFuncDemo() {
	log.Println("TimerAfterFuncDemo start", time.Now())
	time.AfterFunc(2*time.Second, func() {
		log.Println("TimerAfterFuncDemo end", time.Now())
	})
	time.Sleep(2 * time.Second) //等待协程退出
}

// Ticker是周期性定时器，即周期性的触发一个事件。其数据结构和 Timer 完全一致：
// type Ticker struct {
// 	C <-chan Time // The channel on which the ticks are delivered.
// 	r runtimeTimer
// }

// 在创建Ticker时会指定一个时间，作为事件触发的周期。这也是Ticker与Timer的最主要的区别。
func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	// ticker.Reset(1 * time.Second)
	for range ticker.C {
		log.Println("ticker...")
	}
}

// 如果我们需要一个定时轮询任务，可以使用一个简单的Tick函数来获取定时器的管道，函数原型如下：
// func TIck(d Durtion) <-chan Time :
// 这个函数内部实际还是创建一个 Ticker，但并不会返回出来，所以没有手段来停止该 Ticker。所以，一定要考虑具体的使用场景。
func TickerTask(t interface{}) {
	for {
		//不断从tick周期定时器通道中，读取定时触发时间，周期执行
		select {
		case <-t.(*time.Ticker).C: //从通道内读取time， 有表示触发发定时任务
			println("TickerTask 1s timer")
		}
	}
}

func TickerTaskDemo() {
	t := time.NewTicker(time.Second * 1)
	go TickerTask(t)
	select {}
}
