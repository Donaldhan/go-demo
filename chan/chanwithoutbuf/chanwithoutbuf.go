// 这个示例程序展示如何用无缓冲的通道来模拟
// 2 个goroutine 间的网球比赛
package chanwithoutbuf

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
	log.Println("============chanwithoutbuf package init=========")
}

// main 是所有Go 程序的入口
func PlayerTest() {
	// 创建一个无缓冲的通道
	court := make(chan int)
	// 计数加 2，表示要等待两个goroutine
	wg.Add(2)
	// 启动两个选手
	go player("============chanwithoutbuf PlayerTest Nadal", court)
	go player("============chanwithoutbuf PlayerTest Djokovic", court)
	// 发球
	court <- 1
	// 等待游戏结束
	wg.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg.Done()

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("============chanwithoutbuf PlayerTest Player %s Won\n", name)
			return
		}
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("============chanwithoutbuf PlayerTest Player %s Missed\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}
		// 显示击球数，并将击球数加1
		fmt.Printf("============chanwithoutbuf PlayerTest Player %s Hit %d\n", name, ball)
		ball++
		// 将球打向对手
		court <- ball
	}
}

// 这个示例程序展示如何用无缓冲的通道来模拟
// 4 个goroutine 间的接力比赛

// wg 用来等待程序结束
var wgx sync.WaitGroup

func RunnerTest() {
	// 创建一个无缓冲的通道
	baton := make(chan int)
	// 为最后一位跑步者将计数加1
	wgx.Add(1)
	// 第一位跑步者持有接力棒
	go Runner(baton)
	// 开始比赛
	baton <- 1
	// 等待比赛结束
	wgx.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int
	// 等待接力棒
	runner := <-baton
	// 开始绕着跑道跑步
	fmt.Printf("============chanwithoutbuf RunnerTest  Runner %d Running With Baton\n", runner)
	// 创建下一位跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("============chanwithoutbuf RunnerTest Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}
	// 围绕跑道跑
	time.Sleep(100 * time.Millisecond)
	// 比赛结束了吗？
	if runner == 4 {
		fmt.Printf("============chanwithoutbuf RunnerTest  Runner %d Finished, Race Over\n", runner)
		wgx.Done()
		return
	}
	// 将接力棒交给下一位跑步者
	fmt.Printf("============chanwithoutbuf RunnerTest  Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)

	baton <- newRunner
}
