package mod

/**
第三方工具包测试
*/

import (
	"log"
	"strconv"
	"strings"

	// "github.com/labstack/echo"
	jsoniter "github.com/json-iterator/go"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func init() {
	log.Println("==============packagex mod package init")
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

	groups := lo.GroupBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	// 	GroupBy
	// Returns an object composed of keys generated from the results of running each element of collection through iteratee.
	// map[int][]int{0: []int{0, 3}, 1: []int{1, 4}, 2: []int{2, 5}}
	println("LoTest groups :", groups)

	// Parallel processing: like lo.GroupBy(), but callback is called in goroutine.

	groups = lop.GroupBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})

	println("LoTest Parallel groups :", groups)
	// map[int][]int{0: []int{0, 3}, 1: []int{1, 4}, 2: []int{2, 5}}

	// 	Chunk
	// Returns an array of elements split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.

	chunk := lo.Chunk[int]([]int{0, 1, 2, 3, 4, 5}, 2)
	// [][]int{{0, 1}, {2, 3}, {4, 5}}
	println("LoTest chunk :", chunk)
	chunk = lo.Chunk[int]([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	// [][]int{{0, 1}, {2, 3}, {4, 5}, {6}}
	println("LoTest chunk :", chunk)
	chunk = lo.Chunk[int]([]int{}, 2)
	// [][]int{}
	println("LoTest chunk :", chunk)
	chunk = lo.Chunk[int]([]int{0}, 2)
	// [][]int{{0}}
	println("LoTest chunk :", chunk)

	// PartitionBy
	// Returns an array of elements split into groups. The order of grouped values is determined by the order they occur in collection.
	// The grouping is generated from the results of running each element of collection through iteratee.

	partitions := lop.PartitionBy[int, string]([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	// [][]int{{-2, -1}, {0, 2, 4}, {1, 3, 5}}
	println("LoTest partitions :", partitions)

	// 	Flatten
	// Returns an array a single level deep.

	flat := lo.Flatten[int]([][]int{{0, 1}, {2, 3, 4, 5}})
	// []int{0, 1, 2, 3, 4, 5}
	println("LoTest flat :", flat)
}
