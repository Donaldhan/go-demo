package channeldemo

// package main

import (
	"fmt"
	// "testing"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//传数据到通道
	c <- sum // send sum to c
}

// Go语言没有像其它语言一样有public、protected、private等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、
// 函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）
// 通道测试
func TestChannel() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

// 初始化函数, 引用包的时候，回调用函数
func init() {
	fmt.Println("channeldemo package init done.......")
}
