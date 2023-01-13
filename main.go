package main

// 《Go 语言系列教程》之包 | Go 主题月
// https://juejin.cn/post/6946209245637378062
// https://juejin.cn/post/7122730352023437343
// https://zhuanlan.zhihu.com/p/387419521
import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	// "github.com/labstack/echo"
	jsoniter "github.com/json-iterator/go"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"

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

	"godemo/mod"
	flagx "godemo/packagex/flag"
	model "godemo/packagex/model"
	osx "godemo/packagex/os"
	"godemo/packagex/time"
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

// go run main.go -name "aaa" -age=123 -flagname=999
var Input_pstrName = flag.String("name", "gerry", "input ur name")
var Input_piAge = flag.Int("age", 20, "input ur age")
var Input_flagvar int

func Init() {
	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")
	log.Println("=======Main package init========")
}

// go run main.go -name "aaa" -age=123 -flagname=999
func FlagTest() {
	// After parsing, the arguments after the flag are available as the slice flag.Args()
	// or individually as flag.Arg(i). The arguments are indexed from 0 through flag.NArg()-1
	// Args returns the non-flag command-line arguments
	// NArg is the number of arguments remaining after flags have been processed
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *Input_pstrName)
	fmt.Println("age=", *Input_piAge)
	fmt.Println("flagname=", Input_flagvar)
}

/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
go run main.go -slice go,php,java -name "aaa" -age=123 -flagname=999
*/
func FlagTestX() {
	var languages []string
	flag.Var(flagx.NewSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
	flag.Parse()
	//打印结果slice接收到的值
	fmt.Println(languages)

}

// 定义一个类型，用于增加该类型方法
type sliceValue []string

// new一个存放命令行参数值的slice
func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

/*
Value接口：

	type Value interface {
	    String() string
	    Set(string) error
	}

实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
*/
func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

// flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's none of my business"
}

func main() {
	Init()
	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	// go run main.go --mode=fast
	fmt.Println("run mode:", *mode)

	// [martini] listening on :3000 (development)
	// reflectx.MartiniDemo()

	// FlagTest()

	// go run main.go -slice go,php,java
	/*
		var languages []string
		flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
		flag.Parse()
		//打印结果slice接收到的值
		fmt.Println(languages)
	*/

	// FlagTestX()

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
	// IntefaceTypeTest()

	// EncapseTest()

	// TimeTest()

	// Go语言的 os 包中提供了操作系统函数的接口，是一个比较重要的包。顾名思义，os 包的作用主要是在服务器上进行系统的基本操作，
	// 如文件操作、目录操作、执行命令、信号与中断、进程、系统状态等等。
	// OSTest()

	// OSSignallTest()

	// ModTest()
	// JsonTest()
	// LoTest()

	//第三方包测试
	mod.JsonTest()
	mod.LoTest()
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

/**
 * mod 第三方引用包测试
 */
// https://github.com/json-iterator/go
func JsonTest() {
	//转换对象为json字符串
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := jsoniter.Marshal(group)
	if err == nil {
		log.Println("JsonTest group:", string(b))
	} else {
		log.Println("JsonTest err:", err.Error())
	}
	// 获取JSON对象
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	result := jsoniter.Get(val, "Colors", 0).ToString()
	log.Println("JsonTest result:", result)
}

/**
 * 工具包测试
 */
// https://github.com/samber/lo
func LoTest() {
	// 	Uniq
	// Returns a duplicate-free version of an array, in which only the first occurrence of each element is kept.
	// The order of result values is determined by the order they occur in the array.

	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	// []string{"Samuel", "John"}
	log.Println("LoTest names:", names)

	// Manipulates a slice of one type and transforms it into a slice of another type:
	intToString := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		return strconv.FormatInt(x, 10)
	})
	// []string{"1", "2", "3", "4"}
	log.Println("LoTest intToString:", intToString)

	// Parallel processing: like lo.Map(), but the mapper function is called in a goroutine. Results are returned in the same order.
	intToString = lop.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	// []string{"1", "2", "3", "4"}
	log.Println("LoTest Parallel intToString:", intToString)

	even := lo.Filter[int]([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0
	})
	// []int{2, 4}
	log.Println("LoTest Filter even:", even)

	// FilterMap
	// Returns a slice which obtained after both filtering and mapping using the given callback function.
	// The callback function should return two values: the result of the mapping operation and whether the result element should be included or not.
	matching := lo.FilterMap[string, string]([]string{"cpu", "gpu", "mouse", "keyboard"}, func(x string, _ int) (string, bool) {
		if strings.HasSuffix(x, "pu") {
			return "xpu", true
		}
		return "", false
	})
	// []string{"xpu", "xpu"}
	log.Println("LoTest matching:", matching)

	// ForEach
	// Iterates over elements of a collection and invokes the function over each element.
	lo.ForEach[string]([]string{"hello", "world"}, func(x string, _ int) {
		println("LoTest ForEach :", x)
	})
	// prints "hello\nworld\n"

	// Parallel processing: like lo.ForEach(), but the callback is called as a goroutine.
	lop.ForEach[string]([]string{"hello", "world"}, func(x string, _ int) {
		println("LoTest ForEach Parallel :", x)
	})
	// prints "hello\nworld\n" or "world\nhello\n"

	// 	Reduce
	// Reduces a collection to a single value. The value is calculated by accumulating the result of running each element in the collection
	// through an accumulator function. Each successive invocation is supplied with the return value returned by the previous call.

	sum := lo.Reduce[int, int]([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	// 10
	println("LoTest Reduce sum :", sum)

}

func ModTest() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(":1323"))
}

func OSSignallTest() {
	osx.SignalTest()
	// osx.SignalStopTest()
}
func OSTest() {
	// exec 包可以执行外部命令，它包装了 os.StartProcess 函数以便更容易的修正输入和输出，使用管道连接 I/O，以及作其它的一些调整。
	f, err := exec.LookPath("main")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	// os/user 获取当前用户信息
	u, _ := user.Current()
	log.Println("用户名：", u.Username)
	log.Println("用户id", u.Uid)
	log.Println("用户主目录：", u.HomeDir)
	log.Println("主组id：", u.Gid)
	// 用户所在的所有的组的id
	s, _ := u.GroupIds()
	log.Println("用户所在的所有组：", s)
}
func TimeTest() {
	// time.TimerTest()
	time.TimeFormat()
	time.TimeNowTest()
}

func EncapseTest() {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, " age =", p.GetAge(), " sal = ", p.GetSal())
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
