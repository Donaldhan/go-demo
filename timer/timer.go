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
	timer.Reset()
	timer.Stop()
}

// 比如在一个连接中等待数据，设定一个超时时间，当时间到来还是没有数据获取到，则为超时。
func TimerWaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(3 * time.Second)
	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: //超时
		fmt.Println("TimerWaitChannel timeout")
		return false
	}
}

func TimerWaitChannelDemo() {
	ch := make(chan string, 3)
	TimerWaitChannel(ch)
}

// 希望某个方法在今后的某个时刻执行
func TimerDelayFunction() {
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("TimerDelayFunction Delayed 5s，...")
	}
}

// 有时我们就是想等指定的时间，没有需求提前停止定时器，也没有需求复用该定时器，那么可以使用匿名的定时器
func TimerAfterDemo() {
	log.Println("TimerAfterDemo now", time.Now())
	<-time.After(1 * time.Second)
	log.Println("TimerAfterDemo After", time.Now())
}

// 我们可以使用 AfterFunc 更加简洁的实现延迟一个方法的调用
func TimerAfterFuncDemo() {
	log.Println("TimerAfterFuncDemo start", time.Now())
	time.AfterFunc(1*time.Second, func() {
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

	for range ticker.C {
		log.Println("ticker...")
	}
}
