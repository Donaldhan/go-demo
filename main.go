package main

// 《Go 语言系列教程》之包 | Go 主题月
// https://juejin.cn/post/6946209245637378062
// https://juejin.cn/post/7122730352023437343
// https://zhuanlan.zhihu.com/p/387419521
import (
	// chanTool "godemo/chan/chanbuff" //重命名
	// "godemo/chan/chanconsync"
	// "godemo/chan/channeldemo"
	// "godemo/chan/chanselect"

	"flag"
	"fmt"
	"log"

	// "godemo/concurrent"
	// "godemo/chan/chanwithoutbuf"
	// "godemo/function"
	"godemo/structx"
)

// import "godemo/channeldemo"
// import "log"
// import (
//     "fmt"
//     "godemo/channeldemo" // 导入自定义的包
//     "log"
// )
// 导入自定义的包, 无用的包，将会被移除

// var _ = channeldemo.TestChannel
// go run .\main.go
// go run .\main.go client
// go run .\main.go server
const pi = 3.14159 // 相当于 math.Pi 的近似值
// 定义命令行参数
// go run main.go --mode=fast
var mode = flag.String("mode", "", "process mode")

func main() {
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println("run mode:", *mode)
	// channeldemo.TestChannel()
	// log.Println("=====channeldemo TestChannel=========")

	// chanconsync.TestChanconsync()
	// log.Println("=====chanconsync TestChanconsync=========")

	// chanconsync.ChanFor()
	// log.Println("=====chanconsync ChanFor=========")

	// chanconsync.ConPrinter()
	// log.Println("=====chanconsync ConPrinter=========")

	// chanwithoutbuf.PlayerTest()
	// log.Println("=====chanconsync PlayerTest done=========")

	// chanwithoutbuf.RunnerTest()
	// log.Println("=====chanconsync RunnerTest done=========")

	// chanwithoutbuf.RunnerTest()
	// log.Println("=====chanconsync RunnerTest done=========")

	// chanTool.FixBuffChannel()
	// log.Println("=====chanbuff FixBuffChannel done=========")

	// chanselect.ChanSelectTest()
	// log.Println("=====chanselect ChanSelectTest done=========")

	// concurrent.AsyncFuncTest()
	// log.Println("=====concurrent AsyncFuncTest done=========")

	// concurrent.DoAll()
	// log.Println("=====concurrent DoAll done=========")

	// concurrent.NumCPUTest()
	// log.Println("=====concurrent NumCPUTest done=========")

	// concurrent.SyncMutexTest()
	// concurrent.SyncMutexRwTest()
	// log.Println("=====concurrent SyncMutexTest and Rw done=========")
	// concurrent.WaitGroupTest()
	// log.Println("=====concurrent WaitGroupTest  done=========")

	// concurrent.DeadLockTest()
	// concurrent.IdleLockTest()
	// concurrent.HungerLockTest()
	// log.Println("=====concurrent LockTest  done=========")

	// concurrent.CspSyncMutexTest()
	// concurrent.CspChannWithoutBufTest()
	// concurrent.CspChannWithBufTest()
	// concurrent.CspMultiChannTest()
	// concurrent.CspWaitGroupTest()
	// log.Println("=====concurrent csp test  done=========")
	/*
		argsWithProg := os.Args
		argsWithoutProg := os.Args[1:]
		talkMode := os.Args[1]
		fmt.Println(argsWithProg)
		fmt.Println(argsWithoutProg)
		fmt.Println(talkMode)
		if talkMode == "server" {
			// go run .\main.go server
			// talk.StartTalkServer()
		}
		if talkMode == "client" {
			// go run .\main.go client
			// talk.StartTalkClient()
		}
		log.Println("=====talk test done=========")
	*/
	// concurrent.BacktheadTest()
	// log.Println("=====BacktheadTest done=========")

	// concurrent.AtomicTest()
	// concurrent.AtomicLoadStoreTest()
	// log.Println("=====AtomicLoadStoreTest done=========")

	// function.FunctionTestX()
	// function.DeferTest()
	// function.ErrorTest()
	// // function.PanicMustCompileTest()
	// function.RecoverFromPanicTest()
	// function.TimeConsumeTest()
	// log.Println("=====function package test done=========")

	structx.AnonymousTest()
	log.Println("=====structx package test done=========")
}
