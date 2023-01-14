package mod

/**
第三方工具包测试
*/

import (
	"log"
	"math"
	"strconv"
	"strings"

	// "github.com/labstack/echo"
	// jsoniter "github.com/json-iterator/go"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func init() {
	log.Println("==============packagex mod package init")
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

func LoDemo() {
	// Interleave
	// Round-robin alternating input slices and sequentially appending value at index into result.

	interleaved := lo.Interleave[int]([]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9})
	// []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Println("LoDemo interleaved:", interleaved)

	interleavedx := lo.Interleave[int]([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
	// []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println("LoDemo interleavedx:", interleavedx)

	// 	Shuffle
	// Returns an array of shuffled values. Uses the Fisher-Yates shuffle algorithm.
	randomOrder := lo.Shuffle[int]([]int{0, 1, 2, 3, 4, 5})
	// []int{1, 4, 0, 3, 5, 2}
	log.Println("LoDemo randomOrder:", randomOrder)

	// 	Reverse
	// Reverses array so that the first element becomes the last, the second element becomes the second to last, and so on.
	// This helper is mutable. This behavior might change in v2.0.0. See #160.
	reverseOrder := lo.Reverse[int]([]int{0, 1, 2, 3, 4, 5})
	// []int{5, 4, 3, 2, 1, 0}
	log.Println("LoDemo reverseOrder:", reverseOrder)

}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

func LoDemoX() {
	// Fill
	// Fills elements of array with initial value.
	initializedSlice := lo.Fill[foo]([]foo{foo{"a"}, foo{"a"}}, foo{"b"})
	// []foo{foo{"b"}, foo{"b"}}
	log.Println("LoDemoX initializedSlice:", initializedSlice)
	// Repeat
	// Builds a slice with N copies of initial value.
	slice := lo.Repeat[foo](2, foo{"a"})
	// []foo{foo{"a"}, foo{"a"}}
	log.Println("LoDemoX slice:", slice)

	// 	RepeatBy
	// Builds a slice with values returned by N calls of callback.

	slicex := lo.RepeatBy[string](0, func(i int) string {
		return strconv.FormatInt(int64(math.Pow(float64(i), 2)), 10)
	})
	// []string{}
	log.Println("LoDemoX slicex:", slicex)
	slicex = lo.RepeatBy[string](5, func(i int) string {
		return strconv.FormatInt(int64(math.Pow(float64(i), 2)), 10)
	})
	// []string{"0", "1", "4", "9", "16"}
	log.Println("LoDemoX slicex:", slicex)
}

type Character struct {
	dir  string
	code int
}

// KeyBy
// Transforms a slice or an array of structs to a map based on a pivot callback.
func KeyBy() {

	m := lo.KeyBy[int, string]([]string{"a", "aa", "aaa"}, func(str string) int {
		return len(str)
	})
	// map[int]string{1: "a", 2: "aa", 3: "aaa"}
	log.Println("KeyBy m:", m)

	characters := []Character{
		{dir: "left", code: 97},
		{dir: "right", code: 100},
	}
	result := lo.KeyBy[string, Character](characters, func(char Character) string {
		return string(rune(char.code))
	})
	//map[a:{dir:left code:97} d:{dir:right code:100}]
	log.Println("KeyBy result:", result)
}

// Associate (alias: SliceToMap)
// Returns a map containing key-value pairs provided by transform function applied to elements of the given slice.
// If any of two pairs would have the same key the last one gets added to the map.

// The order of keys in returned map is not specified and is not guaranteed to be the same from the original array.

type foox struct {
	baz string
	bar int
}

func Associate() {
	in := []*foox{{baz: "apple", bar: 1}, {baz: "banana", bar: 2}}

	aMap := lo.Associate[*foox, string, int](in, func(f *foox) (string, int) {
		return f.baz, f.bar
	})
	// map[string][int]{ "apple":1, "banana":2 }
	log.Println("Associate aMap:", aMap)
}

func Drop() {
	// Drop
	// Drops n elements from the beginning of a slice or array.

	l := lo.Drop[int]([]int{0, 1, 2, 3, 4, 5}, 2)
	// []int{2, 3, 4, 5}
	log.Println("Drop l:", l)

	// DropRight
	// Drops n elements from the end of a slice or array.

	l = lo.DropRight[int]([]int{0, 1, 2, 3, 4, 5}, 2)
	// []int{0, 1, 2, 3}
	log.Println("Drop l:", l)

	// 	DropWhile
	// Drop elements from the beginning of a slice or array while the predicate returns true.

	lx := lo.DropWhile[string]([]string{"a", "aa", "aaa", "aa", "aa"}, func(val string) bool {
		return len(val) <= 2
	})
	// []string{"aaa", "aa", "aa"}
	log.Println("Drop lx:", lx)

	// 	DropRightWhile
	// Drop elements from the end of a slice or array while the predicate returns true.

	lx = lo.DropRightWhile[string]([]string{"a", "aa", "aaa", "aa", "aa"}, func(val string) bool {
		return len(val) <= 2
	})
	// []string{"a", "aa", "aaa"}
	log.Println("Drop lx:", lx)
}

// Reject
func Reject() {

}
