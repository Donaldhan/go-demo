package reflect

//单元测试demo
import (
	"fmt"
	"log"
	"reflect"
)

func init() {
	log.Println("==============reflect package init")
}

// 在Go语言程序中，使用 reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息
func Reflect() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println("ReflectTest:", typeOfA.Name(), typeOfA.Kind())
}

// 反射的类型（Type）与种类（Kind）
// 在使用反射时，需要首先理解类型（Type）和种类（Kind）的区别。编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。
// 1) 反射种类（Kind）的定义
// Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。

// 种类（Kind）指的是对象归属的品种

// / 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

func ReflectTypeKind() {
	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println("ReflectTypeKindTest:", typeOfCat.Name(), typeOfCat.Kind())
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println("ReflectTypeKindTest:", typeOfA.Name(), typeOfA.Kind())
}

// 指针与指针指向的元素
// Go语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
func ReflectElem() {
	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("ReflectElemTest name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("ReflectElemTest element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}

// 使用反射获取结构体的成员类型
// 任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息。

// 与成员获取相关的 reflect.Type 的方法如下表所示。

// 结构体成员访问的方法列表
// 方法	说明
// Field(i int) StructField	根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机
// NumField() int	返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机
// FieldByName(name string) (StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机
// FieldByIndex(index []int) StructField	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机
// FieldByNameFunc(match func(string) bool) (StructField,bool)	根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机
// 1) 结构体字段类型
// reflect.Type 的 Field() 方法返回 StructField 结构，这个结构描述结构体的成员信息，通过这个信息可以获取成员与结构体的关系，
// 如偏移、索引、是否为匿名字段、结构体标签（StructTag）等，而且还可以通过 StructField 的 Type 字段进一步获取结构体成员的类型信息。

// StructField 的结构如下：
//
//	type StructField struct {
//	    Name string          // 字段名
//	    PkgPath string       // 字段路径
//	    Type      Type       // 字段反射类型对象
//	    Tag       StructTag  // 字段的结构体标签
//	    Offset    uintptr    // 字段在结构体中的相对偏移
//	    Index     []int      // Type.FieldByIndex中的返回的索引值
//	    Anonymous bool       // 是否为匿名字段
//	}
//
// 字段说明如下：
// Name：为字段名称。
// PkgPath：字段在结构体中的路径。
// Type：字段本身的反射类型对象，类型为 reflect.Type，可以进一步获取字段的类型信息。
// Tag：结构体标签，为结构体字段标签的额外信息，可以单独提取。
// Index：FieldByIndex 中的索引顺序。
// Anonymous：表示该字段是否为匿名字段。

// 获取成员反射信息
// 下面代码中，实例化一个结构体并遍历其结构体成员，再通过 reflect.Type 的 FieldByName() 方法查找结构体中指定名称的字段，直接获取其类型信息。
func ReflectField() {
	// 声明一个空结构体
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("ReflectFieldTest name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println("ReflectFieldTest:", catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

// 结构体标签（Struct Tag）
// 通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。

// JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称 ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到。
// 1) 结构体标签的格式
// Tag 在结构体字段后方书写的格式如下：
// `key1:"value1" key2:"value2"`

// 结构体标签由一个或多个键值对组成；键与值使用冒号分隔，值用双引号括起来；键值对之间使用一个空格分隔。
// 2) 从结构体标签中获取值
// StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：
// func (tag StructTag) Get(key string) string：根据 Tag 中的键获取对应的值，例如`key1:"value1" key2:"value2"`的 Tag 中，可以传入“key1”获得“value1”。
// func (tag StructTag) Lookup(key string) (value string, ok bool)：根据 Tag 中的键，查询值是否存在。
// 3) 结构体标签格式错误导致的问题
// 编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，示例代码如下

func ReflectStructTagError() {
	type cat struct {
		Name string
		Type int `json: "type" id:"100"`
	}
	typeOfCat := reflect.TypeOf(cat{})
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println("ReflectStructTagErrorTest:", catType.Tag.Get("json"))
	}
}

// 反射第一定律：反射可以将“接口类型变量”转换为“反射类型对象”
// 注：这里反射类型指 reflect.Type 和 reflect.Value。

// 从使用方法上来讲，反射提供了一种机制，允许程序在运行时检查接口变量内部存储的 (value, type) 对。

// 类型 reflect.Type 和 reflect.Value 都有很多方法，我们可以检查和使用它们，这里我们举几个例子。

// 类型 reflect.Value 有一个方法 Type()，它会返回一个 reflect.Type 类型的对象。

// Type 和 Value 都有一个名为 Kind 的方法，它会返回一个常量，表示底层数据的类型，常见值有：Uint、Float64、Slice 等。

// Value 类型也有一些类似于 Int、Float 的方法，用来提取底层的数据：
// Int 方法用来提取 int64
// Float 方法用来提取 float64，示例代码如下：
func ReflectValue() {
	var x float64 = 3.4
	fmt.Println("ReflectValue value:", reflect.ValueOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("ReflectValue type:", v.Type())
	fmt.Println("ReflectValue kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("ReflectValue value:", v.Float())
}

// 虽然变量 v 的静态类型是 MyInt，而不是 int，但 Kind 方法仍然会返回 reflect.Int。换句话说 Kind 方法不会像 Type 方法一样区分 MyInt 和 int。
func ReflectKindX() {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("ReflectKindX type:", v.Type())                            // uint8.
	fmt.Println("ReflectKindX kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	x = uint8(v.Uint())                                                    // v.Uint returns a uint64.
}

// 反射第二定律：反射可以将“反射类型对象”转换为“接口类型变量”
// 和物理学中的反射类似，Go语言中的反射也能创造自己反面类型的对象。

// 根据一个 reflect.Value 类型的变量，我们可以使用 Interface 方法恢复其接口类型的值。事实上，这个方法会把 type 和 value 信息打包并填充到一个接口变量中，然后返回。

// 反射第三定律：如果要修改“反射类型对象”其值必须是“可写的”
// 这条定律很微妙，也很容易让人迷惑，但是如果从第一条定律开始看，应该比较容易理解。

// http://c.biancheng.net/view/5131.html

func CanSet() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("CanSet settability of v:", v.CanSet())

	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("CanSet type of p:", p.Type())
	fmt.Println("CanSet settability of p:", p.CanSet())
	// 反射对象 p 是不可写的，但是我们也不像修改 p，事实上我们要修改的是 *p。为了得到 p 指向的数据，可以调用 Value 类型的 Elem 方法。
	// Elem 方法能够对指针进行“解引用”，然后将结果存储到反射 Value 类型对象 v 中：
	// 反射不太容易理解，reflect.Type 和 reflect.Value 会混淆正在执行的程序，但是它做的事情正是编程语言做的事情。
	// 只需要记住：只要反射对象要修改它们表示的对象，就必须获取它们表示的对象的地址。
	e := p.Elem()
	fmt.Println("CanSet settability of e:", e.CanSet())
	e.SetFloat(7.1)
	fmt.Println("CanSet:", e.Interface())
	fmt.Println("CanSet:", e)
}

func StrutTypeOfT() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("StrutTypeOfT,%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func StructElem() {
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

// 总结
// 反射规则可以总结为如下几条：
// 反射可以将“接口类型变量”转换为“反射类型对象”；
// 反射可以将“反射类型对象”转换为“接口类型变量”；
// 如果要修改“反射类型对象”，其值必须是“可写的”。

func InterfaceType() {
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	// 获取interface{}类型的值, 通过类型断言转换
	var getA int = valueOfA.Interface().(int)
	// 获取64位的值, 强制类型转换为int类型
	var getA2 int = int(valueOfA.Int())
	fmt.Println("InterfaceType:", getA, getA2)
}

// 反射值对象（reflect.Value）提供一系列方法进行零值和空判定，如下表所示。

// 反射值对象的零值和有效性判断方法
// 方 法	说 明
// IsNil() bool	返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的v== nil操作
// IsValid() bool	判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等。

func NinValid() {
	// *int的空指针
	var a *int
	fmt.Println("NinValid var a *int:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("NinValid nil:", reflect.ValueOf(nil).IsValid())
	// *int类型的空指针
	fmt.Println("NinValid (*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	// 实例化一个结构体
	s := struct{}{}
	// 尝试从结构体中查找一个不存在的字段
	fmt.Println("NinValid 不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
	// 尝试从结构体中查找一个不存在的方法
	fmt.Println("NinValid 不存在的结构体方法:", reflect.ValueOf(s).MethodByName("").IsValid())
	// 实例化一个map
	m := map[int]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("NinValid 不存在的键：", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}

// 每当我们通过指针间接地获取的 reflect.Value 都是可取地址的，即使开始的是一个不可取地址的 Value。在反射机制中，所有关于是否支持取地址的规则都是类似的。
// 例如，slice 的索引表达式 e[i]将隐式地包含一个指针，它就是可取地址的，即使开始的e表达式不支持也没有关系。

// 以此类推，reflect.ValueOf(e).Index(i) 对于的值也是可取地址的，即使原始的 reflect.ValueOf(e) 不支持也没有关系。

// 使用 reflect.Value 对包装的值进行修改时，需要遵循一些规则。如果没有按照规则进行代码设计和编写，轻则无法修改对象值，重则程序在运行时会发生宕机。

// 其中 a 对应的变量则不可取地址。因为 a 中的值仅仅是整数 2 的拷贝副本。b 中的值也同样不可取地址。
// c 中的值还是不可取地址，它只是一个指针 &x 的拷贝。实际上，所有通过 reflect.ValueOf(x) 返回的 reflect.Value 都是不可取地址的。
// 但是对于 d，它是 c 的解引用方式生成的，指向另一个变量，因此是可取地址的。我们可以通过调用 reflect.ValueOf(&x).Elem()，来获取任意变量x对应的可取地址的 Value。

func CanAddrX() {
	x := 2                                  // value type variable?
	a := reflect.ValueOf(2)                 // 2 int no
	b := reflect.ValueOf(x)                 // 2 int no
	c := reflect.ValueOf(&x)                // &x *int no
	d := c.Elem()                           // 2 int yes (x)
	fmt.Println("CanAddrX a:", a.CanAddr()) // "false"
	fmt.Println("CanAddrX b:", b.CanAddr()) // "false"
	fmt.Println("CanAddrX c:", c.CanAddr()) // "false"
	fmt.Println("CanAddrX d:", d.CanAddr()) // "true"
}

// 值可修改条件之一：可被寻址
// 通过反射修改变量值的前提条件之一：这个值必须可以被寻址
// 提示
// 当 reflect.Value 不可寻址时，使用 Addr() 方法也是无法取到值的地址的，同时会发生宕机。
// 虽然说 reflect.Value 的 Addr() 方法类似于语言层的&操作；Elem() 方法类似于语言层的*操作，但并不代表这些方法与语言层操作等效。
func ReflectValueSet() {
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象(a的地址)
	valueOfA := reflect.ValueOf(&a)
	// 取出a地址的元素(a的值)
	valueOfA = valueOfA.Elem()
	// 修改a的值为1
	valueOfA.SetInt(1)
	// 打印a的值
	fmt.Println("ReflectValueSet:", valueOfA.Int())

	// 获取变量a的反射值对象
	valueOfB := reflect.ValueOf(a)
	// 尝试将a修改为1(此处会发生崩溃)
	valueOfB.SetInt(1)
}

// 值可修改条件之一：被导出
// 结构体成员中，如果字段没有被导出，即便不使用反射也可以被访问，但不能通过反射修改
func ReflectValueStructSet() {
	type dog struct {
		legCount int
	}
	// 获取dog实例的反射值对象
	valueOfDog := reflect.ValueOf(dog{})
	// 获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("legCount")
	// 尝试设置legCount的值(这里会发生崩溃)
	vLegCount.SetInt(4)
}

// 值的修改从表面意义上叫可寻址，换一种说法就是值必须“可被设置”。那么，想修改变量值，一般的步骤是：
// 取这个变量的地址或者这个变量所在的结构体已经是指针类型。
// 使用 reflect.ValueOf 进行值包装。
// 通过 Value.Elem() 获得指针值指向的元素值对象（Value），因为值对象（Value）内部对象为指针时，使用 set 设置时会报出宕机错误。
// 使用 Value.Set 设置值。
func ReflectValueStructSetX() {
	type dog struct {
		LegCount int
	}
	// 获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})
	// 取出dog实例地址的元素
	valueOfDog = valueOfDog.Elem()
	// 获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")
	// 尝试设置legCount的值(这里会发生崩溃)
	vLegCount.SetInt(4)
	fmt.Println("ReflectValueStructSetX:", vLegCount.Int())
}

// Go语言通过类型信息创建实例
func ReflectNew() {
	var a int
	// 取变量a的反射类型对象
	typeOfA := reflect.TypeOf(a)
	// 根据反射类型对象创建类型实例
	aIns := reflect.New(typeOfA)
	// 输出Value的类型和种类
	fmt.Println("ReflectNew:", aIns.Type(), aIns.Kind())
}

// / 普通函数
func add(a, b int) int {
	return a + b
}

// 如果反射值对象（reflect.Value）中值的类型为函数时，可以通过 reflect.Value 调用该函数。
// 使用反射调用函数时，需要将参数使用反射值对象的切片 []reflect.Value 构造后传入 Call() 方法中，调用完成时，函数的返回值通过 []reflect.Value 返回。

// 下面的代码声明一个加法函数，传入两个整型值，返回两个整型值的和。将函数保存到反射值对象（reflect.Value）中，
// 然后将两个整型值构造为反射值对象的切片（[]reflect.Value），使用 Call() 方法进行调用。

// 提示
// 反射调用函数的过程需要构造大量的 reflect.Value 和中间变量，对函数参数值进行逐一检查，还需要将调用参数复制到调用函数的参数内存中。
// 调用完毕后，还需要将返回值转换为 reflect.Value，用户还需要从中取出调用值。因此，反射调用函数的性能问题尤为突出，不建议大量使用反射函数调用

func ReflectFunctionCall() {
	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	// 构造函数参数, 传入两个整型值
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// 反射调用函数
	retList := funcValue.Call(paramList)
	// 获取第一个返回值, 取整数值
	fmt.Println("ReflectFunctionCall", retList[0].Int())
}
