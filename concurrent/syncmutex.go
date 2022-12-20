package concurrent

import (
	"fmt"
	"sync"
)

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
	// 逻辑中使用的某个变量
	countRw int
	// 与变量对应的使用互斥锁
	countGuardRw sync.RWMutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}
func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

// Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex
func SyncMutexTest() {
	// 可以进行并发安全的设置
	SetCount(1)
	// 可以进行并发安全的获取
	fmt.Println("concurrent SyncMutexTest:", GetCount())
}

func GetCountX() int {
	// 锁定
	countGuardRw.RLock()
	// 在函数退出时解除锁定
	defer countGuardRw.RUnlock()
	return countRw
}
func SetCountX(c int) {
	countGuardRw.Lock()
	countRw = c
	countGuardRw.Unlock()
}
func SyncMutexRwTest() {
	// 可以进行并发安全的设置
	SetCountX(2)
	// 可以进行并发安全的获取
	fmt.Println("concurrent SyncMutexRwTest:", GetCountX())
}
