package main

// 《Go 语言系列教程》之包 | Go 主题月
// https://juejin.cn/post/6946209245637378062
// https://juejin.cn/post/7122730352023437343
// https://zhuanlan.zhihu.com/p/387419521
import (
	"godemo/chanconsync"
	"godemo/channeldemo"
	"log"
)

// "fmt"

// "godemo/channel"
// channeldemo "./channel"

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
func main() {
	channeldemo.TestChannel()
	log.Println("=====channeldemo TestChannel=========")
	chanconsync.TestChanconsync()
	log.Println("=====chanconsync TestChanconsync=========")
	chanconsync.ChanFor()
	log.Println("=====chanconsync ChanFor=========")
	chanconsync.ConPrinter()
	log.Println("=====chanconsync ConPrinter=========")

}
