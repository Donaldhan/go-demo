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
	// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
	//阻塞模式，一次性
	select {
	case <-conn: //读取通道数据, 阻塞接收数据后，忽略从通道返回的数据,执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步
		timer.Stop()
		fmt.Println("^^^^^^^^ TimerWaitChannel recive data from conn,stop timer")
		return true
	case <-timer.C: //超时，从通道内读取time， 有表示触发发定时任务,
		fmt.Println("^^^^^^^^ TimerWaitChannel timeout")
		return false
	}
}

// 同步等待，超时
func TimerWaitChannelDemoWithoutWrite() {
	ch := make(chan string, 3)
	TimerWaitChannel(ch)
	log.Println("================TimerWaitChannelDemoWithoutWrite done")
}

// 同步等待，超时
func TimerWaitChannelDemoAfterSelectWrite() {
	ch := make(chan string, 3)
	TimerWaitChannel(ch)
	ch <- "hello"
	log.Println("================TimerWaitChannelDemoAfterSelectWrite done")
}

// 同步等待，缓存区先写入数据，先取到数据，不会超时,
func TimerWaitChannelDemoBeforeSelectWrite() {
	ch := make(chan string, 3)
	ch <- "hello"
	TimerWaitChannel(ch)
	log.Println("================TimerWaitChannelDemoBeforeSelectWrite done")
}

// 异步后台协程等待，然后写入数据到连接通道，可能取到数据，可能超时，看go协程调度的时间
func TimerWaitChannelDemoGoSelectThenWrite() {
	ch := make(chan string, 3)
	go TimerWaitChannel(ch)
	ch <- "hello"
	log.Println("================TimerWaitChannelDemoGoSelectThenWrite done")
}

// 循环超时阻塞
func TimerChannelSelectForTimeoutMode(conn <-chan int) bool {
	timer := time.NewTimer(5 * time.Second)
	for {
		// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
		//循环阻塞模式
		select {
		case msg := <-conn: //读取通道数据
			fmt.Println("********* TimerChannelSelectForTimeMode recive msg from conn, msg:", msg)
			fmt.Println("+++++++ TimerChannelSelectForTimeMode conn  chan size:", len(conn))
			return true
		case <-timer.C: //超时，从通道内读取time， 有表示触发发定时任务
			fmt.Println("********* TimerChannelSelectForTimeMode timeout")
			timer.Stop()
			return false
		default: //设置一个可用的 case，让 select 变成非阻塞
			time.Sleep(time.Second)
			fmt.Println("********* TimerChannelSelectForTimeMode default invoke")
		}
	}
}

// 异步循环select,超时阻塞，后写入，读到数据，3秒后超时
func AsynChannelSelectForModeWithTimeoutDemo() {
	ch := make(chan int, 3)
	//异步阻塞等待，可以读到数据，可能有默认case处理，最后超时
	go TimerChannelSelectForTimeoutMode(ch)
	// ch <- "hello" //无法读取数据
	/*
		for i := 0; i < 5; i++ {
			ch <- i //等待缓冲区可用,阻塞
		}
	*/
	for i := 0; i < 3; i++ { //只能可以读到1个数据？？？？？？？？？？？？
		ch <- i
		time.Sleep(time.Second)
	}
	// 查看当前通道的大小
	// fmt.Println("+++++++ AsynChannelSelectForModeWithTimeoutDemo chan size:", len(ch))
	// for n := range ch {//阻塞模式
	// 	fmt.Println("+++++++ AsynChannelSelectForModeWithTimeoutDemo chan data:", n)
	// }
	log.Println("================ AsynChannelSelectForModeWithTimeoutDemo done")
}

// 异步循环select,超时阻塞，后写入，读到数据，3秒后超时
func AsynChannelSelectForModeWithTimeoutDemoX() {
	ch := make(chan int, 3)
	// ch <- "hello" //无法读取数据
	/*
		for i := 0; i < 5; i++ {
			ch <- i //等待缓冲区可用,阻塞
		}
	*/
	go func() {
		for i := 0; i < 3; i++ { //只能可以读到1个数据？？？？？？？？？？？？
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	//异步阻塞等待，可以读到数据，可能有默认case处理，最后超时
	TimerChannelSelectForTimeoutMode(ch)
	// 查看当前通道的大小
	// fmt.Println("+++++++ AsynChannelSelectForModeWithTimeoutDemoX chan size:", len(ch))
	// for n := range ch {//阻塞模式
	// 	fmt.Println("+++++++ AsynChannelSelectForModeWithTimeoutDemoX chan data:", n)
	// }
	log.Println("================ AsynChannelSelectForModeWithTimeoutDemoX done")
}

// 同步循环select,非阻塞，后写入，不会读到数据，3秒后超时
func SyncChannelSelectForModeWithTimeoutDemo() {
	ch := make(chan int, 3)
	//同步阻塞等待，直至超时
	TimerChannelSelectForTimeoutMode(ch)
	// ch <- "hello" //无法读取数据
	/*
		for i := 0; i < 5; i++ {
			ch <- i //等待缓冲区可用,阻塞
		}
	*/
	for i := 0; i < 3; i++ {
		ch <- i
	}
	// fmt.Println("+++++++ SyncChannelSelectForModeWithTimeoutDemo chan size:", len(ch))
	// for n := range ch { //阻塞模式
	// 	fmt.Println("+++++++ SyncChannelSelectForModeWithTimeoutDemo chan data:", n)
	// }
	log.Println("================SyncChannelSelectForModeWithTimeoutDemo done")
}

// 循环非阻塞模式,设置一个可用的 case，让 select 变成非阻塞
func TimerChannelSelectForDefaultCaseMode(conn <-chan string) bool {
	for {
		// 虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况
		//循环非阻塞模式
		select {
		case msg := <-conn: //读取通道数据
			fmt.Println("############ TimerChannelSelectForDefaultCaseMode recive data from conn,stop timer msg:", msg)
			fmt.Println("+++++++ TimerChannelSelectForDefaultCaseMode conn  chan size:", len(conn))
			return true
		default: //设置一个可用的 case，让 select 变成非阻塞
			time.Sleep(time.Second)
			fmt.Println("############ TimerChannelSelectForDefaultCaseMode default invoke")
		}
	}
}

// 异步循环select,非阻塞无超时，后写入，不会读的数据，最后不断执行默认case
func AsyncChannelSelectForDefaultCaseModeDemo() {
	ch := make(chan string, 3)
	//异步阻塞等待，直至超时
	go TimerChannelSelectForDefaultCaseMode(ch)
	// ch <- "hello" //无法读取数据
	for i := 0; i < 3; i++ { //只能可以读到1个数据？？？？？？？？？？？？
		ch <- "hello"
		time.Sleep(time.Second)
	}
	// fmt.Println("+++++++ AsyncChannelSelectForDefaultCaseModeDemo chan size:", len(ch))
	// for n := range ch {//阻塞模式
	// 	fmt.Println("+++++++ AsyncChannelSelectForDefaultCaseModeDemo chan data:", n)
	// }
	log.Println("================AsyncChannelSelectForDefaultCaseModeDemo done")
}

// 异步循环select,非阻塞无超时，后写入，不会读的数据，最后不断执行默认case
func AsyncChannelSelectForDefaultCaseModeDemoX() {
	ch := make(chan string, 3)
	go func() {
		// ch <- "hello" //无法读取数据
		for i := 0; i < 3; i++ { //只能可以读到1个数据？？？？？？？？？？？？
			ch <- "hello"
			time.Sleep(time.Second)
		}
	}()
	//异步阻塞等待，直至超时
	TimerChannelSelectForDefaultCaseMode(ch)

	// fmt.Println("+++++++ AsyncChannelSelectForDefaultCaseModeDemoX chan size:", len(ch))
	// for n := range ch {//阻塞模式
	// 	fmt.Println("+++++++ AsyncChannelSelectForDefaultCaseModeDemoX chan data:", n)
	// }
	log.Println("================AsyncChannelSelectForDefaultCaseModeDemoX done")
}

// 同步循环select,非阻塞无超时，后写入，不会读的数据，，不断执行默认case
func SyncChannelSelectForDefaultCaseModeDemo() {
	ch := make(chan string, 3)
	//同步阻塞等待，直至超时
	TimerChannelSelectForDefaultCaseMode(ch)
	// ch <- "hello" //无法读取数据
	for i := 0; i < 3; i++ {
		ch <- "hello"
	}
	// fmt.Println("+++++++ SyncChannelSelectForDefaultCaseModeDemo chan size:", len(ch))
	// // for n := range ch {
	// 	fmt.Println("+++++++ SyncChannelSelectForDefaultCaseModeDemo chan data:", n)
	// }
	log.Println("================SyncChannelSelectForDefaultCaseModeDemo done")
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
