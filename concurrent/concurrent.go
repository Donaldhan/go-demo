package concurrent

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

var total = 160000000000

func init() {
	log.Println("============concurrent package init=========")
}
func AsyncFuncTest() {
	//秒
	// now := time.Now().Unix()
	//毫秒
	now := time.Now().UnixMilli()
	for i := 0; i < 16; i++ {
		go AsyncFunc(total*i/16, total*(i+1)/16)
	}
	end := time.Now().UnixMilli()
	log.Println("concurrent AsyncFuncTest consume time:", end-now)
	log.Println("concurrent AsyncFuncTest time:", end, now)
}

func AsyncFunc(index int, end int) {

	sum := 0
	for ; index < end; index++ {
		sum += index
	}
	fmt.Printf("============concurrent AsyncFunc  线程%d, sum为:%d\n", index*16/total, sum)
}

type Vector []float64

// 分配给每个CPU的计算任务
func DoSome(index int, end int, c chan int) {
	sum := 0
	for ; index < end; index++ {
		sum += index
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
	fmt.Printf("============concurrent DoSome  线程%d, sum为:%d\n", index*16/total, sum)
}

const NCPU = 16 // 假设总共有16核
func DoAll() {
	now := time.Now().UnixMilli()
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go DoSome(i*total/NCPU, (i+1)*total/NCPU, c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
	end := time.Now().UnixMilli()
	log.Println("concurrent DoAll time:", end-now)
	log.Println("concurrent DoAll time:", end, now)
}

func NumCPUTest() {
	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量
	now := time.Now().UnixMilli()
	DoAll()
	// 到这里表示所有计算已经结束
	end := time.Now().UnixMilli()
	log.Println("concurrent NumCPUTest time:", end-now)
}
