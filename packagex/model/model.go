package mode

import (
	"fmt"
	"log"
)

func init() {
	log.Println("==============packagex mode package init")
}

// 在Go语言中封装就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在内部，程序的其它包只能通过被授权的方法，才能对字段进行操作。

// 封装的好处：
// 隐藏实现细节；
// 可以对数据进行验证，保证数据安全合理。

// 如何体现封装：
// 对结构体中的属性进行封装；
// 通过方法，包，实现封装。

// 封装的实现步骤：
// 1.将结构体、字段的首字母小写；
// 2.给结构体所在的包提供一个工厂模式的函数，首字母大写，类似一个构造函数；
// 3.提供一个首字母大写的 Set 方法（类似其它语言的 public），用于对属性判断并赋值；
// 4.提供一个首字母大写的 Get 方法（类似其它语言的 public），用于获取属性的值。

type person struct {
	Name string
	age  int //其它包不能直接访问..
	sal  float64
}

// 写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age 和 sal 我们编写一对SetXxx的方法和GetXxx的方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确..")
		//给程序员给一个默认值
	}
}
func (p *person) GetAge() int {
	return p.age
}
func (p *person) SetSal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确..")
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}
