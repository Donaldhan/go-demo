package concurrent

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func AtomicTest() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait() //等待goroutine结束
	fmt.Println("concurrent AtomicTest counter:", counter)
}
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) //安全的对counter加1
		runtime.Gosched()            //让出执行权
	}
}

var (
	shutdown int64
	wgx      sync.WaitGroup
)

func AtomicLoadStoreTest() {
	wgx.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second)
	fmt.Println("concurrent AtomicLoadStoreTest Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)
	wgx.Wait()
}
func doWork(name string) {
	defer wgx.Done()
	for {
		fmt.Printf("concurrent AtomicLoadStoreTest  Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("concurrent AtomicLoadStoreTest  Shutting %s Down\n", name)
			break
		}
	}
}
