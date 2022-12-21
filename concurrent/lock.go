package concurrent

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type value struct {
	memAccess sync.Mutex
	value     int
}

func DeadLockTest() {
	log.Println("concurrent DeadLockTest start")
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	sum := func(v1, v2 *value) {
		defer wg.Done()
		v1.memAccess.Lock()
		time.Sleep(2 * time.Second)
		v2.memAccess.Lock()
		fmt.Printf("concurrent DeadLockTest sum = %d\n", v1.value+v2.value)
		v2.memAccess.Unlock()
		v1.memAccess.Unlock()
	}
	product := func(v1, v2 *value) {
		defer wg.Done()
		v2.memAccess.Lock()
		time.Sleep(2 * time.Second)
		v1.memAccess.Lock()
		fmt.Printf("concurrent DeadLockTest product = %d\n", v1.value*v2.value)
		v1.memAccess.Unlock()
		v2.memAccess.Unlock()
	}
	var v1, v2 value
	v1.value = 1
	v2.value = 1
	wg.Add(2)
	go sum(&v1, &v2)
	go product(&v1, &v2)
	wg.Wait()
}

// 活锁
// 活锁是另一种形式的活跃性问题，该问题尽管不会阻塞线程，但也不能继续执行，因为线程将不断重复同样的操作，而且总会失败。

// 例如线程 1 可以使用资源，但它很礼貌，让其他线程先使用资源，线程 2 也可以使用资源，但它同样很绅士，也让其他线程先使用资源。就这样你让我，我让你，最后两个线程都无法使用资源。

// 活锁通常发生在处理事务消息中，如果不能成功处理某个消息，那么消息处理机制将回滚事务，并将它重新放到队列的开头。这样，错误的事务被一直回滚重复执行，这种形式的活锁通常是由过度的错误恢复代码造成的，因为它错误地将不可修复的错误认为是可修复的错误。

// 当多个相互协作的线程都对彼此进行相应而修改自己的状态，并使得任何一个线程都无法继续执行时，就导致了活锁。这就像两个过于礼貌的人在路上相遇，他们彼此让路，然后在另一条路上相遇，然后他们就一直这样避让下去。

// 要解决这种活锁问题，需要在重试机制中引入随机性。例如在网络上发送数据包，如果检测到冲突，都要停止并在一段时间后重发，如果都在 1 秒后重发，还是会冲突，所以引入随机性可以解决该类问题。
func IdleLockTest() {
	log.Println("concurrent IdleLockTest start")
	runtime.GOMAXPROCS(3)
	cv := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Second) { // 通过tick控制两个人的步调
			cv.Broadcast()
		}
	}()
	takeStep := func() {
		cv.L.Lock()
		cv.Wait()
		cv.L.Unlock()
	}
	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %+v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()                      //走上一步
		if atomic.LoadInt32(dir) == 1 { //走成功就返回
			fmt.Fprint(out, ". Success!")
			return true
		}
		takeStep() // 没走成功，再走回来
		atomic.AddInt32(dir, -1)
		return false
	}
	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool {
		return tryDir("向左走", &left, out)
	}
	tryRight := func(out *bytes.Buffer) bool {
		return tryDir("向右走", &right, out)
	}
	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer walking.Done()
		defer func() { fmt.Println(out.String()) }()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v is tried!", name)
	}
	var trail sync.WaitGroup
	trail.Add(2)
	go walk(&trail, "男人") // 男人在路上走
	go walk(&trail, "女人") // 女人在路上走
	trail.Wait()
}

// 饥饿
// 饥饿是指一个可运行的进程尽管能继续执行，但被调度器无限期地忽视，而不能被调度执行的情况。

// 与死锁不同的是，饥饿锁在一段时间内，优先级低的线程最终还是会执行的，比如高优先级的线程执行完之后释放了资源。

// 活锁与饥饿是无关的，因为在活锁中，所有并发进程都是相同的，并且没有完成工作。更广泛地说，饥饿通常意味着有一个或多个贪婪的并发进程，它们不公平地阻止一个或多个并发进程，以尽可能有效地完成工作，或者阻止全部并发进程。
func HungerLockTest() {
	log.Println("concurrent HungerLockTest start")
	runtime.GOMAXPROCS(3)
	var wg sync.WaitGroup
	const runtime = 1 * time.Second
	var sharedLock sync.Mutex
	greedyWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("concurrent HungerLockTest Greedy worker was able to execute %v work loops\n", count)
	}
	politeWorker := func() {
		defer wg.Done()
		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			sharedLock.Lock()
			time.Sleep(1 * time.Nanosecond)
			sharedLock.Unlock()
			count++
		}
		fmt.Printf("concurrent HungerLockTest Polite worker was able to execute %v work loops\n", count)
	}
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}

// 总结
// 1.不适用锁肯定会出问题。如果用了，虽然解了前面的问题，但是又出现了更多的新问题。
// 2.死锁：是因为错误的使用了锁，导致异常；
// 3.活锁：是饥饿的一种特殊情况，逻辑上感觉对，程序也一直在正常的跑，但就是效率低，逻辑上进行不下去；
// 4.饥饿：与锁使用的粒度有关，通过计数取样，可以判断进程的工作效率。
