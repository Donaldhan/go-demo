package interfacex

import (
	"fmt"
	"sort"
)

// 将[]string定义为MyStringList类型
type MyStringList []string

// 实现sort.Interface接口的获取元素数量方法
func (m MyStringList) Len() int {
	return len(m)
}

// 实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// 实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// 序操作和字符串格式化一样是很多程序经常使用的操作。尽管一个最短的快排程序只要 15 行就可以搞定，但是一个健壮的实现需要更多的代码，并且我们不希望每次我们需要的时候都重写或者拷贝这些代码。

// 幸运的是，sort 包内置的提供了根据一些排序函数来对任何序列排序的功能。它的设计非常独到。在很多语言中，排序算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。

// 相比之下，Go语言的 sort.Sort 函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型 sort.Interface 来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

// 一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 sort.Interface 的三个方法：
// package sort
//
//	type Interface interface {
//	    Len() int            // 获取元素数量
//	    Less(i, j int) bool // i，j是序列元素的指数。
//	    Swap(i, j int)        // 交换元素
//	}
//
// http://c.biancheng.net/view/81.html
func SorterTest() {
	// 准备一个内容被打乱顺序的字符串切片
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用sort包进行排序
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("SorterTest:%s\n", v)
	}
}

// type StringSlice []string

// func (p StringSlice) Len() int           { return len(p) }
// func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
// func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// // Sort is a convenience method.
// func (p StringSlice) Sort() { sort.Sort(p) }

func StringSliceTest() {
	names := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("StringSliceTest:%s\n", v)
	}
}

// type IntSlice []int
// func (p IntSlice) Len() int           { return len(p) }
// func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
// func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// // Sort is a convenience method.
// func (p IntSlice) Sort() { sort.Sort(p) }

func IntSliceTest() {
	names := sort.IntSlice{
		3, 2, 6, 8, 1,
	}
	sort.Sort(names)
	// 遍历打印结果
	for _, v := range names {
		fmt.Printf("IntSliceTest:%s\n", v)
	}
}

// 声明英雄的分类
type HeroKind int

// 定义HeroKind常量, 类似于枚举
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

// 定义英雄名单的结构
type Hero struct {
	Name string   // 英雄的名字
	Kind HeroKind // 英雄的种类
}

// 将英雄指针的切片定义为Heros类型
type Heros []*Hero

// 实现sort.Interface接口取元素数量方法
func (s Heros) Len() int {
	return len(s)
}

// 实现sort.Interface接口比较元素方法
func (s Heros) Less(i, j int) bool {
	// 如果英雄的分类不一致时, 优先对分类进行排序
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	// 默认按英雄名字字符升序排列
	return s[i].Name < s[j].Name
}

// 实现sort.Interface接口交换元素方法
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func StructSortTest() {
	// 准备英雄列表
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	// 使用sort包进行排序
	sort.Sort(heros)
	// 遍历英雄列表打印排序结果
	for _, v := range heros {
		fmt.Printf("StructSortTest:%+v\n", v)
	}
}

// 声明英雄的分类
type HeroKindX int

// 定义HeroKind常量, 类似于枚举
const (
	NoneX HeroKindX = iota
	TankX
	AssassinX
	MageX
)

// 定义英雄名单的结构
type HeroX struct {
	Name string    // 英雄的名字
	Kind HeroKindX // 英雄的种类
}

// 从 Go 1.8 开始，Go语言在 sort 包中提供了 sort.Slice() 函数进行更为简便的排序方法。
// sort.Slice() 函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类型为 func(i,j int)bool，sort.Slice() 函数的定义如下：
func StructSortTestX() {
	heros := []*HeroX{
		{"吕布", TankX},
		{"李白", AssassinX},
		{"妲己", MageX},
		{"貂蝉", AssassinX},
		{"关羽", TankX},
		{"诸葛亮", MageX},
	}
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	for _, v := range heros {
		fmt.Printf("StructSortTestX:%+v\n", v)
	}
}
