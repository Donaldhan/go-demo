package flag

import (
	"log"
	"strings"
)

func init() {
	log.Println("==============packagex flag package init")
}

// 定义一个类型，用于增加该类型方法
type sliceValue []string

// new一个存放命令行参数值的slice
func NewSliceValue(vals []string, p *[]string) *sliceValue {
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
