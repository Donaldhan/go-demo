package time

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.Println("==============packagex timer package init")
}

// 返回通Time类型的道
// func Tick(d Duration) <-chan Time {
// 	if d <= 0 {
// 		return nil
// 	}
// 	return NewTicker(d).C
// }

// 使用 time.Tick(时间间隔) 可以设置定时器，定时器的本质上是一个通道（channel）
func TimerTest() {
	ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
	for i := range ticker {
		fmt.Println("timer TimerTest", i) //每秒都会执行的任务
	}
}

func TimeFormat() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func TimeNowTest() {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("现在的时间戳：%v\n", timestamp1)
	fmt.Printf("现在的纳秒时间戳：%v\n", timestamp2)

	timestamp := now.Unix()            //时间戳
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year = timeObj.Year()     //年
	month = timeObj.Month()   //月
	day = timeObj.Day()       //日
	hour = timeObj.Hour()     //小时
	minute = timeObj.Minute() //分钟
	second = timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//时间戳
	t := time.Now()
	fmt.Println(t.Weekday().String())
	later := now.Add(time.Hour) // 当前时间加1小时后的时间
	fmt.Println(later)

	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj2)
}
