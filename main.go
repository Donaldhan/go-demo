package main

// 《Go 语言系列教程》之包 | Go 主题月
// https://juejin.cn/post/6946209245637378062
// https://juejin.cn/post/7122730352023437343
// https://zhuanlan.zhihu.com/p/387419521
import (
	"flag"
	"fmt"
	chanTool "godemo/chan/chanbuff" //重命名
	"godemo/chan/chanconsync"
	"godemo/chan/channeldemo"
	"godemo/chan/chanselect"
	"godemo/chan/chanwithoutbuf"
	"godemo/concurrent"
	"godemo/function"
	"godemo/interfacex"
	runtimex "godemo/runtime"
	"godemo/structx"

	"log"
)

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

	// chanTest()

	// 聊天室
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

	// concurrentPackageTest();

	// TestFunctionPackage()

	// structTest()

	// runtimeTest()
	// IntefacexTest()
	// IntefacexSortTest()
	IntefaceTypeTest()

}
func IntefaceTypeTest() {
	interfacex.NullTest()
	interfacex.NestingTest()
	// interfacex.TypeswitchTest()
	interfacex.SwitchTypeTest()
	interfacex.SwitchTypeTestX()
	interfacex.ErrorTest()
	interfacex.ErrorTestX()
	interfacex.WebTest()
	log.Println("=====IntefaceTypeTest package test done=========")
}

func IntefacexSortTest() {
	interfacex.SorterTest()
	interfacex.StringSliceTest()
	interfacex.IntSliceTest()
	interfacex.StructSortTest()
	interfacex.StructSortTestX()
	interfacex.NestingTest()
	interfacex.TypeswitchTest()
	log.Println("=====IntefacexSortTest package test done=========")
}
func IntefacexTest() {
	interfacex.IntefacexTest()
	interfacex.SocketTest()
	interfacex.ServiceTest()
	interfacex.AssertTest()
	log.Println("=====IntefacexTest package test done=========")
}

// 垃圾回收
func runtimeTest() {
	runtimex.TestGC()
	log.Println("=====runtimex package test done=========")
}

// 结构体测试
func structTest() {
	structx.AnonymousTest()
	structx.InitStructTest()
	structx.InitInnerStructTest()
	structx.InsertFrontLinkListTest()
	structx.InsertTailLinkListTest()
	log.Println("=====structx package test done=========")
}

// 通道测试
func chanTest() {
	channeldemo.TestChannel()
	log.Println("=====channeldemo TestChannel=========")

	chanconsync.TestChanconsync()
	log.Println("=====chanconsync TestChanconsync=========")

	chanconsync.ChanFor()
	log.Println("=====chanconsync ChanFor=========")

	chanconsync.ConPrinter()
	log.Println("=====chanconsync ConPrinter=========")

	chanwithoutbuf.PlayerTest()
	log.Println("=====chanconsync PlayerTest done=========")

	chanwithoutbuf.RunnerTest()
	log.Println("=====chanconsync RunnerTest done=========")

	chanwithoutbuf.RunnerTest()
	log.Println("=====chanconsync RunnerTest done=========")

	chanTool.FixBuffChannel()
	log.Println("=====chanbuff FixBuffChannel done=========")
	chanselect.ChanSelectTest()
	log.Println("=====chanselect ChanSelectTest done=========")
}

// 并发包测试
func concurrentPackageTest() {

	concurrent.AsyncFuncTest()
	log.Println("=====concurrent AsyncFuncTest done=========")

	concurrent.DoAll()
	log.Println("=====concurrent DoAll done=========")

	concurrent.NumCPUTest()
	log.Println("=====concurrent NumCPUTest done=========")

	concurrent.SyncMutexTest()
	concurrent.SyncMutexRwTest()
	log.Println("=====concurrent SyncMutexTest and Rw done=========")
	concurrent.WaitGroupTest()
	log.Println("=====concurrent WaitGroupTest  done=========")

	concurrent.DeadLockTest()
	concurrent.IdleLockTest()
	concurrent.HungerLockTest()
	log.Println("=====concurrent LockTest  done=========")

	concurrent.CspSyncMutexTest()
	concurrent.CspChannWithoutBufTest()
	concurrent.CspChannWithBufTest()
	concurrent.CspMultiChannTest()
	concurrent.CspWaitGroupTest()
	log.Println("=====concurrent csp test  done=========")

	concurrent.BacktheadTest()
	log.Println("=====BacktheadTest done=========")

	concurrent.AtomicTest()
	concurrent.AtomicLoadStoreTest()
	log.Println("=====AtomicLoadStoreTest done=========")
}

// 函数包测试
func TestFunctionPackage() {
	function.FunctionTestX()
	function.DeferTest()
	function.ErrorTest()
	// function.PanicMustCompileTest()
	function.RecoverFromPanicTest()
	function.TimeConsumeTest()
	log.Println("=====function package test done=========")
}
