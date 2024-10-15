package mod

/**
第三方工具包测试
go 工具包，包含类似java相关的steam api， map， redis， foreach 和 filter，包括串行和并行模式

https://github.com/samber/lo
*/

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

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
func RejectAndCount() {
	// 	Reject
	// The opposite of Filter, this method returns the elements of collection that predicate does not return truthy for.
	dd := lo.Reject[int]([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	// []int{1, 3}
	log.Println("RejectAndCount dd:", dd)
	// 	Count
	// Counts the number of elements in the collection that compare equal to value.
	count := lo.Count[int]([]int{1, 5, 1}, 1)
	log.Println("RejectAndCount count:", count)
	// 2
	// 	CountBy
	// Counts the number of elements in the collection for which predicate is true.
	count = lo.CountBy[int]([]int{1, 5, 1}, func(i int) bool {
		return i < 4
	})
	log.Println("RejectAndCount CountBy count:", count)
	// 2

	// 	CountValues
	// Counts the number of each element in the collection.

	countValue := lo.CountValues([]int{})
	// map[int]int{}
	log.Println("RejectAndCount CountValues:", countValue)
	// 2
	countValue = lo.CountValues([]int{1, 2})
	// map[int]int{1: 1, 2: 1}
	log.Println("RejectAndCount CountValues:", countValue)
	countValue = lo.CountValues([]int{1, 2, 2})
	// map[int]int{1: 1, 2: 2}
	log.Println("RejectAndCount CountValues:", countValue)

	countValueX := lo.CountValues([]string{"foo", "bar", ""})
	// map[string]int{"": 1, "foo": 1, "bar": 1}
	log.Println("RejectAndCount countValueX:", countValueX)
	countValueX = lo.CountValues([]string{"foo", "bar", "bar"})
	// map[string]int{"foo": 1, "bar": 2}
	log.Println("RejectAndCount countValueX:", countValueX)
	// 	CountValuesBy
	// Counts the number of each element in the collection. It ss equivalent to chaining lo.Map and lo.CountValues.

	isEven := func(v int) bool {
		return v%2 == 0
	}

	countValuesBy := lo.CountValuesBy([]int{}, isEven)
	// map[bool]int{}
	log.Println("RejectAndCount CountValuesBy:", countValuesBy)

	countValuesBy = lo.CountValuesBy([]int{1, 2}, isEven)
	// map[bool]int{false: 1, true: 1}
	log.Println("RejectAndCount CountValuesBy:", countValuesBy)

	countValuesBy = lo.CountValuesBy([]int{1, 2, 2}, isEven)
	// map[bool]int{false: 1, true: 2}
	log.Println("RejectAndCount CountValuesBy:", countValuesBy)

	length := func(v string) int {
		return len(v)
	}

	countValuesByX := lo.CountValuesBy([]string{"foo", "bar", ""}, length)
	// map[int]int{0: 1, 3: 2}
	log.Println("RejectAndCount countValuesByX:", countValuesByX)
	countValuesByX = lo.CountValuesBy([]string{"foo", "bar", "bar"}, length)
	// map[int]int{3: 3}
	log.Println("RejectAndCount countValuesByX:", countValuesByX)
}

func SubsetAndSlice() {
	// Subset
	// Returns a copy of a slice from offset up to length elements. Like slice[start:start+length], but does not panic on overflow.

	in := []int{0, 1, 2, 3, 4}

	sub := lo.Subset(in, 2, 3)
	// []int{2, 3, 4}

	log.Println("SubsetAndSlice sub:", sub)

	sub = lo.Subset(in, -4, 3)
	// []int{1, 2, 3}
	log.Println("SubsetAndSlice sub:", sub)

	sub = lo.Subset(in, -2, math.MaxUint)
	// []int{3, 4}
	log.Println("SubsetAndSlice sub:", sub)

	// 	Slice
	// Returns a copy of a slice from start up to, but not including end. Like slice[start:end], but does not panic on overflow.

	slice := lo.Slice(in, 0, 5)
	// []int{0, 1, 2, 3, 4}
	log.Println("SubsetAndSlice slice:", slice)

	slice = lo.Slice(in, 2, 3)
	// []int{2}
	log.Println("SubsetAndSlice slice:", slice)

	slice = lo.Slice(in, 2, 6)
	// []int{2, 3, 4}
	log.Println("SubsetAndSlice slice:", slice)

	slice = lo.Slice(in, 4, 3)
	// []int{}
	log.Println("SubsetAndSlice slice:", slice)

	// 	Replace
	// Returns a copy of the slice with the first n non-overlapping instances of old replaced by new.

	inx := []int{0, 1, 0, 1, 2, 3, 0}

	slice = lo.Replace(inx, 0, 42, 1)
	// []int{42, 1, 0, 1, 2, 3, 0}
	log.Println("SubsetAndSlice Replace:", slice)

	slice = lo.Replace(inx, -1, 42, 1)
	// []int{0, 1, 0, 1, 2, 3, 0}
	log.Println("SubsetAndSlice Replace:", slice)

	slice = lo.Replace(inx, 0, 42, 2)
	// []int{42, 1, 42, 1, 2, 3, 0}
	log.Println("SubsetAndSlice Replace:", slice)

	slice = lo.Replace(inx, 0, 42, -1)
	// []int{42, 1, 42, 1, 2, 3, 42}
	log.Println("SubsetAndSlice Replace:", slice)

	// ReplaceAll
	// Returns a copy of the slice with all non-overlapping instances of old replaced by new.

	slice = lo.ReplaceAll(inx, 0, 42)
	// []int{42, 1, 42, 1, 2, 3, 42}
	log.Println("SubsetAndSlice ReplaceAll:", slice)

	slice = lo.ReplaceAll(inx, -1, 42)
	// []int{0, 1, 0, 1, 2, 3, 0}
	log.Println("SubsetAndSlice ReplaceAll:", slice)

	// Compact
	// Returns a slice of all non-zero elements.

	inxx := []string{"", "foo", "", "bar", ""}

	slicex := lo.Compact[string](inxx)
	// []string{"foo", "bar"}
	log.Println("SubsetAndSlice Compact:", slicex)

	// IsSorted
	// Checks if a slice is sorted.

	slicexx := lo.IsSorted([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	// true
	log.Println("SubsetAndSlice slicexx:", slicexx)
	// IsSortedByKey
	// Checks if a slice is sorted by iteratee.

	slicexx = lo.IsSortedByKey([]string{"a", "bb", "ccc"}, func(s string) int {
		return len(s)
	})
	log.Println("SubsetAndSlice IsSortedByKey:", slicexx)
	// true
}
func KeyValuePick() {
	// Keys
	// Creates an array of the map keys.

	keys := lo.Keys[string, int](map[string]int{"foo": 1, "bar": 2})
	// []string{"foo", "bar"}
	log.Println("KeyValuePick keys:", keys)

	// 	Values
	// Creates an array of the map values.

	values := lo.Values[string, int](map[string]int{"foo": 1, "bar": 2})
	// []int{1, 2}
	log.Println("KeyValuePick values:", values)

	// 	PickBy
	// Returns same map type filtered by given predicate.

	m := lo.PickBy[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})
	// map[string]int{"foo": 1, "baz": 3}
	log.Println("KeyValuePick PickBy:", m)

	// 	OmitBy
	// Returns same map type filtered by given predicate.

	m = lo.OmitBy[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(key string, value int) bool {
		return value%2 == 1
	})
	// map[string]int{"bar": 2}
	log.Println("KeyValuePick OmitBy:", m)
	// OmitByKeys
	// Returns same map type filtered by given keys.

	m = lo.OmitByKeys[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}, []string{"foo", "baz"})
	// map[string]int{"bar": 2}
	log.Println("KeyValuePick OmitByKeys:", m)

	// OmitByValues
	// Returns same map type filtered by given values.

	m = lo.OmitByValues[string, int](map[string]int{"foo": 1, "bar": 2, "baz": 3}, []int{1, 3})
	// map[string]int{"bar": 2}
	log.Println("KeyValuePick OmitByValues:", m)
	// Entries (alias: ToPairs)
	// Transforms a map into array of key/value pairs.

	entries := lo.Entries[string, int](map[string]int{"foo": 1, "bar": 2})
	// []lo.Entry[string, int]{
	//     {
	//         Key: "foo",
	//         Value: 1,
	//     },
	//     {
	//         Key: "bar",
	//         Value: 2,
	//     },
	// }
	log.Println("KeyValuePick entries:", entries)
	// FromEntries (alias: FromPairs)
	// Transforms an array of key/value pairs into a map.

	m = lo.FromEntries[string, int]([]lo.Entry[string, int]{
		{
			Key:   "foo",
			Value: 1,
		},
		{
			Key:   "bar",
			Value: 2,
		},
	})
	// map[string]int{"foo": 1, "bar": 2}
	log.Println("KeyValuePick FromEntries:", m)
	// Invert
	// Creates a map composed of the inverted keys and values. If map contains duplicate values, subsequent values overwrite property assignments of previous values.

	m1 := lo.Invert[string, int](map[string]int{"a": 1, "b": 2})
	// map[int]string{1: "a", 2: "b"}
	log.Println("KeyValuePick Invert m1:", m1)
	m2 := lo.Invert[string, int](map[string]int{"a": 1, "b": 2, "c": 1})
	// map[int]string{1: "c", 2: "b"}
	log.Println("KeyValuePick Invert m2:", m2)
	// Assign
	// Merges multiple maps from left to right.

	mergedMaps := lo.Assign[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)
	// map[string]int{"a": 1, "b": 3, "c": 4}
	log.Println("KeyValuePick mergedMaps：", mergedMaps)
	// MapKeys
	// Manipulates a map keys and transforms it to a map of another type.

	m2x := lo.MapKeys[int, int, string](map[int]int{1: 1, 2: 2, 3: 3, 4: 4}, func(_ int, v int) string {
		return strconv.FormatInt(int64(v), 10)
	})
	// map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
	log.Println("KeyValuePick MapKeys", m2x)
	// MapValues
	// Manipulates a map values and transforms it to a map of another type.

	m1x := map[int]int64{1: 1, 2: 2, 3: 3}

	m2xx := lo.MapValues[int, int64, string](m1x, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	log.Println("KeyValuePick MapValues", m2xx)
	// map[int]string{1: "1", 2: "2", 3: "3"}

	// MapEntries
	// Manipulates a map entries and transforms it to a map of another type.

	in := map[string]int{"foo": 1, "bar": 2}

	out := lo.MapEntries(in, func(k string, v int) (int, string) {
		return v, k
	})
	// map[int]string{1: "foo", 2: "bar"}
	log.Println("KeyValuePick MapEntries:", out)
	// MapToSlice
	// Transforms a map into a slice based on specific iteratee.

	my := map[int]int64{1: 4, 2: 5, 3: 6}

	s := lo.MapToSlice(my, func(k int, v int64) string {
		return fmt.Sprintf("%d_%d", k, v)
	})
	// []string{"1_4", "2_5", "3_6"}
	log.Println("KeyValuePick MapToSlice:", s)
}

func RangeClamp() {
	// Range / RangeFrom / RangeWithSteps
	// Creates an array of numbers (positive and/or negative) progressing from start up to, but not including end.

	result := lo.Range(4)
	// [0, 1, 2, 3]
	log.Println("RangeClamp Range:", result)

	result = lo.Range(-4)
	// [0, -1, -2, -3]
	log.Println("RangeClamp Range:", result)

	result = lo.RangeFrom(1, 5)
	// [1, 2, 3, 4, 5]
	log.Println("RangeClamp RangeFrom:", result)

	resultx := lo.RangeFrom[float64](1.0, 5)
	// [1.0, 2.0, 3.0, 4.0, 5.0]
	log.Println("RangeClamp RangeFrom:", resultx)

	result = lo.RangeWithSteps(0, 20, 5)
	// [0, 5, 10, 15]
	log.Println("RangeClamp RangeWithSteps:", result)

	resultxx := lo.RangeWithSteps[float32](-1.0, -4.0, -1.0)
	// [-1.0, -2.0, -3.0]
	log.Println("RangeClamp RangeWithSteps:", resultxx)

	result = lo.RangeWithSteps(1, 4, -1)
	// []
	log.Println("RangeClamp RangeWithSteps:", result)

	result = lo.Range(0)
	// []
	log.Println("RangeClamp Range:", result)

	// Clamp
	// Clamps number within the inclusive lower and upper bounds.

	r1 := lo.Clamp(0, -10, 10)
	// 0
	log.Println("RangeClamp Clamp:", r1)
	r2 := lo.Clamp(-42, -10, 10)
	// -10
	log.Println("RangeClamp Clamp:", r2)
	r3 := lo.Clamp(42, -10, 10)
	// 10
	log.Println("RangeClamp Clamp:", r3)
}

func SumString() {
	// Sum
	// Sums the values in a collection.
	// If collection is empty 0 is returned.

	list := []int{1, 2, 3, 4, 5}
	sum := lo.Sum(list)
	// 15
	log.Println("SumString sum:", sum)
	// SumBy
	// Summarizes the values in a collection using the given return value from the iteration function.
	// If collection is empty 0 is returned.

	strings := []string{"foo", "bar"}
	sum = lo.SumBy(strings, func(item string) int {
		return len(item)
	})
	// 6
	log.Println("SumString SumBy:", sum)
	// 	RandomString
	// Returns a random string of the specified length and made of the specified charset.

	str := lo.RandomString(5, lo.LettersCharset)
	// example: "eIGbt"
	log.Println("SumString RandomString:", str)
	// Substring
	// Return part of a string.

	sub := lo.Substring("hello", 2, 3)
	// "llo"
	log.Println("SumString Substring:", sub)
	sub = lo.Substring("hello", -4, 3)
	// "ell"
	log.Println("SumString Substring:", sub)
	sub = lo.Substring("hello", -2, math.MaxUint)
	// "lo"
	log.Println("SumString Substring:", sub)
	// ChunkString
	// Returns an array of strings split into groups the length of size. If array can't be split evenly, the final chunk will be the remaining elements.

	chukcString := lo.ChunkString("123456", 2)
	// []string{"12", "34", "56"}
	log.Println("SumString ChunkString:", chukcString)

	chukcString = lo.ChunkString("1234567", 2)
	// []string{"12", "34", "56", "7"}
	log.Println("SumString ChunkString:", chukcString)
	chukcString = lo.ChunkString("", 2)
	// []string{""}
	log.Println("SumString ChunkString:", chukcString)
	chukcString = lo.ChunkString("1", 2)
	// []string{"1"}
	log.Println("SumString ChunkString:", chukcString)
	// RuneLength
	// An alias to utf8.RuneCountInString which returns the number of runes in string.

	lenx := lo.RuneLength("hellô")
	// 5
	log.Println("SumString RuneLength:", lenx)
	lenx = len("hellô")
	// 6
	log.Println("SumString len:", lenx)
}

func example() (string, int) { return "y", 2 }
func TurpleZip() {
	// T2 -> T9
	// Creates a tuple from a list of values.

	tuple1 := lo.T2("x", 1)
	// Tuple2[string, int]{A: "x", B: 1}
	log.Println("TurpleZip tuple1:", tuple1)

	tuple2 := lo.T2(example())
	// Tuple2[string, int]{A: "y", B: 2}
	log.Println("TurpleZip tuple2:", tuple2)
	// Unpack2 -> Unpack9
	// Returns values contained in tuple.

	r1, r2 := lo.Unpack2[string, int](lo.Tuple2[string, int]{"a", 1})
	// "a", 1
	log.Println("TurpleZip Unpack2:", r1, r2)
	// Unpack is also available as a method of TupleX.

	tuple2 = lo.T2("a", 1)
	a, b := tuple2.Unpack()
	// "a" 1
	log.Println("TurpleZip Unpack:", a, b)
	// Zip2 -> Zip9
	// Zip creates a slice of grouped elements, the first of which contains the first elements of the given arrays,
	//  the second of which contains the second elements of the given arrays, and so on.

	// When collections have different size, the Tuple attributes are filled with zero value.

	tuples := lo.Zip2[string, int]([]string{"a", "b"}, []int{1, 2})
	// []Tuple2[string, int]{{A: "a", B: 1}, {A: "b", B: 2}}
	log.Println("TurpleZip Zip2:", tuples)
	// Unzip2 -> Unzip9
	// Unzip accepts an array of grouped elements and creates an array regrouping the elements to their pre-zip configuration.

	// ax, bx := lo.Unzip2[string, int]([]Tuple2[string, int]{{A: "a", B: 1}, {A: "b", B: 2}})
	// // []string{"a", "b"}
	// // []int{1, 2}
	// log.Println("TurpleZip Unzip2:", ax, bx)
}

// Many distributions strategies are available:

// lo.DispatchingStrategyRoundRobin: Distributes messages in a rotating sequential manner.
// lo.DispatchingStrategyRandom: Distributes messages in a random manner.
// lo.DispatchingStrategyWeightedRandom: Distributes messages in a weighted manner.
// lo.DispatchingStrategyFirst: Distributes messages in the first non-full channel.
// lo.DispatchingStrategyLeast: Distributes messages in the emptiest channel.
// lo.DispatchingStrategyMost: Distributes to the fullest channel.
/*
type Message struct {
	TenantID uuid.UUID
}

func hash(id uuid.UUID) int {
	h := fnv.New32a()
	h.Write([]byte(id.String()))
	return int(h.Sum32())
}
*/
func ChannelDispatcher() {
	ch := make(chan int, 42)
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	// Routes messages per TenantID.
	/*
		customStrategy := func(message pubsub.AMQPSubMessage, messageIndex uint64, channels []<-chan pubsub.AMQPSubMessage) int {
			destination := hash(message.TenantID) % len(channels)

			// check if channel is full
			if len(channels[destination]) < cap(channels[destination]) {
				return destination
			}

			// fallback when child channel is full
			return utils.DispatchingStrategyRoundRobin(message, uint64(destination), channels)
		}

		children := lo.ChannelDispatcher(ch, 5, 10, customStrategy)
		consumer := func(c <-chan int) {
			for {
				msg, ok := <-c
				if !ok {
					println("ChannelDispatcher closed")
					break
				}

				println("ChannelDispatcher msg:", msg)
			}
		}

		for i := range children {
			go consumer(children[i])
		}
	*/
}
func ChannleSlice() {
	// SliceToChannel
	// Returns a read-only channels of collection elements. Channel is closed after last element. Channel capacity can be customized.

	list := []int{1, 2, 3, 4, 5}

	for v := range lo.SliceToChannel(2, list) {
		log.Println("SliceToChannel v:", v)
	}
	// prints 1, then 2, then 3, then 4, then 5
	// ChannelToSlice
	// Returns a slice built from channels items. Blocks until channel closes.

	list = []int{1, 2, 3, 4, 5}
	ch := lo.SliceToChannel(2, list)

	items := lo.ChannelToSlice(ch)
	// []int{1, 2, 3, 4, 5}

	log.Println("ChannelToSlice items:", items)
	// Generator
	// Implements the generator design pattern. Channel is closed after last element. Channel capacity can be customized.

	generator := func(yield func(int)) {
		yield(1)
		yield(2)
		yield(3)
	}

	for v := range lo.Generator(2, generator) {
		println("Generator:", v)
	}
	// prints 1, then 2, then 3

	// 	Buffer
	// Creates a slice of n elements from a channel. Returns the slice, the slice length, the read time and the channel status (opened/closed).

	ch = lo.SliceToChannel(2, []int{1, 2, 3, 4, 5})
	log.Println("ChannelToSlice ch:", ch)

	items1, length1, duration1, ok1 := lo.Buffer(ch, 3)
	// []int{1, 2, 3}, 3, 0s, true
	log.Println("ChannelToSlice Buffer:", items1, length1, duration1, ok1)

	items2, length2, duration2, ok2 := lo.Buffer(ch, 3)
	// []int{4, 5}, 2, 0s, false
	log.Println("ChannelToSlice Buffer:", items2, length2, duration2, ok2)

	// 	BufferWithTimeout
	// Creates a slice of n elements from a channel, with timeout. Returns the slice, the slice length, the read time and the channel status (opened/closed).

	generator = func(yield func(int)) {
		for i := 0; i < 5; i++ {
			yield(i)
			time.Sleep(35 * time.Millisecond)
		}
	}

	ch = lo.Generator(0, generator)

	items1, length1, duration1, ok1 = lo.BufferWithTimeout(ch, 3, 100*time.Millisecond)
	// []int{1, 2}, 2, 100ms, true
	log.Println("ChannelToSlice BufferWithTimeout:", items1, length1, duration1, ok1)

	items2, length2, duration2, ok2 = lo.BufferWithTimeout(ch, 3, 100*time.Millisecond)
	// []int{3, 4, 5}, 3, 75ms, true
	log.Println("ChannelToSlice BufferWithTimeout:", items2, length2, duration2, ok2)

	items3, length3, duration2, ok3 := lo.BufferWithTimeout(ch, 3, 100*time.Millisecond)
	// []int{}, 0, 10ms, false
	log.Println("ChannelToSlice BufferWithTimeout:", items3, length3, duration2, ok3)
}

func SetStream() {
	// Contains
	// Returns true if an element is present in a collection.

	present := lo.Contains[int]([]int{0, 1, 2, 3, 4, 5}, 5)
	// true
	log.Println("SetStream Contains present:", present)

	// ContainsBy
	// Returns true if the predicate function returns true.

	present = lo.ContainsBy[int]([]int{0, 1, 2, 3, 4, 5}, func(x int) bool {
		return x == 3
	})
	// true
	log.Println("SetStream ContainsBy present:", present)

	// Every
	// Returns true if all elements of a subset are contained into a collection or if the subset is empty.

	ok := lo.Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	// true
	log.Println("SetStream Every:", ok)
	ok = lo.Every[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	// false
	log.Println("SetStream Every:", ok)

	// EveryBy
	// Returns true if the predicate returns true for all of the elements in the collection or if the collection is empty.

	b := lo.EveryBy[int]([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 5
	})
	// true
	log.Println("SetStream EveryBy:", b)

	// Some
	// Returns true if at least 1 element of a subset is contained into a collection. If the subset is empty Some returns false.

	ok = lo.Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	// true
	log.Println("SetStream Some:", ok)
	ok = lo.Some[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	// false
	log.Println("SetStream Some:", ok)

	// SomeBy
	// Returns true if the predicate returns true for any of the elements in the collection. If the collection is empty SomeBy returns false.

	b = lo.SomeBy[int]([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 3
	})
	// true
	log.Println("SetStream SomeBy:", b)
	// None
	// Returns true if no element of a subset are contained into a collection or if the subset is empty.

	b = lo.None[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	// false
	log.Println("SetStream None:", b)

	b = lo.None[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	// true
	log.Println("SetStream None:", b)

	// NoneBy
	// Returns true if the predicate returns true for none of the elements in the collection or if the collection is empty.

	b = lo.NoneBy[int]([]int{1, 2, 3, 4}, func(x int) bool {
		return x < 0
	})
	// true
	log.Println("SetStream NoneBy:", b)

	// Intersect
	// Returns the intersection between two collections.

	result1 := lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2})
	// []int{0, 2}
	log.Println("SetStream Intersect:", result1)

	result2 := lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 6})
	// []int{0}
	log.Println("SetStream Intersect:", result2)

	result3 := lo.Intersect[int]([]int{0, 1, 2, 3, 4, 5}, []int{-1, 6})
	// []int{}
	log.Println("SetStream Intersect:", result3)

	// Difference
	// Returns the difference between two collections.

	// The first value is the collection of element absent of list2.
	// The second value is the collection of element absent of list1.
	left, right := lo.Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2, 6})
	// []int{1, 3, 4, 5}, []int{6}
	log.Println("SetStream Difference:", left, right)

	left, right = lo.Difference[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5})
	// []int{}, []int{}
	log.Println("SetStream Difference:", left, right)

	// Union
	// Returns all distinct elements from given collections. Result will not change the order of elements relatively.

	union := lo.Union[int]([]int{0, 1, 2, 3, 4, 5}, []int{0, 2}, []int{0, 10})
	// []int{0, 1, 2, 3, 4, 5, 10}
	log.Println("SetStream Union:", union)

	// Without
	// Returns slice excluding all given values.

	subset := lo.Without[int]([]int{0, 2, 10}, 2)
	// []int{0, 10}
	log.Println("SetStream Without:", subset)

	subset = lo.Without[int]([]int{0, 2, 10}, 0, 1, 2, 3, 4, 5)
	// []int{10}
	log.Println("SetStream Without:", subset)

	// WithoutEmpty
	// Returns slice excluding empty values.

	subset = lo.WithoutEmpty[int]([]int{0, 2, 10})
	// []int{2, 10}
	log.Println("SetStream WithoutEmpty:", subset)

	// IndexOf
	// Returns the index at which the first occurrence of a value is found in an array or return -1 if the value cannot be found.

	found := lo.IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
	// 2
	log.Println("SetStream IndexOf:", found)
	notFound := lo.IndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)
	// -1
	log.Println("SetStream IndexOf:", notFound)

	// LastIndexOf
	// Returns the index at which the last occurrence of a value is found in an array or return -1 if the value cannot be found.

	found = lo.LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 2)
	// 4
	log.Println("SetStream LastIndexOf:", found)

	notFound = lo.LastIndexOf[int]([]int{0, 1, 2, 1, 2, 3}, 6)
	// -1
	log.Println("SetStream LastIndexOf:", notFound)

	// Find
	// Search an element in a slice based on a predicate. It returns element and true if element was found.

	str, ok := lo.Find[string]([]string{"a", "b", "c", "d"}, func(i string) bool {
		return i == "b"
	})
	// "b", true
	log.Println("SetStream Find:", str, ok)

	str, ok = lo.Find[string]([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})
	// "", false
	log.Println("SetStream Find:", str, ok)

	// FindIndexOf
	// FindIndexOf searches an element in a slice based on a predicate and returns the index and true. It returns -1 and false if the element is not found.

	str, index, ok := lo.FindIndexOf[string]([]string{"a", "b", "a", "b"}, func(i string) bool {
		return i == "b"
	})
	// "b", 1, true
	log.Println("SetStream FindIndexOf:", str, index, ok)

	str, index, ok = lo.FindIndexOf[string]([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})
	// "", -1, false
	log.Println("SetStream FindIndexOf:", str, index, ok)

	// FindLastIndexOf
	// FindLastIndexOf searches an element in a slice based on a predicate and returns the index and true. It returns -1 and false if the element is not found.

	str, index, ok = lo.FindLastIndexOf[string]([]string{"a", "b", "a", "b"}, func(i string) bool {
		return i == "b"
	})
	// "b", 4, true
	log.Println("SetStream FindLastIndexOf:", str, index, ok)

	str, index, ok = lo.FindLastIndexOf[string]([]string{"foobar"}, func(i string) bool {
		return i == "b"
	})
	// "", -1, false
	log.Println("SetStream FindLastIndexOf:", str, index, ok)

	// FindKey
	// Returns the key of the first value matching.

	resultx1, okx := lo.FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 2)
	// // "bar", true
	log.Println("SetStream FindKey:", resultx1, okx)

	resultx1, okx = lo.FindKey(map[string]int{"foo": 1, "bar": 2, "baz": 3}, 42)
	// // "", false
	log.Println("SetStream FindKey:", resultx1, okx)

	type test struct {
		foobar string
	}

	result3x, ok3x := lo.FindKey(map[string]test{"foo": test{"foo"}, "bar": test{"bar"}, "baz": test{"baz"}}, test{"foo"})
	// "foo", true
	log.Println("SetStream FindKey:", result3x, ok3x)

	// FindKeyBy
	// Returns the key of the first element predicate returns truthy for.

	resultx1, okx1 := lo.FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return k == "foo"
	})
	// "foo", true
	log.Println("SetStream FindKeyBy:", resultx1, okx1)

	resultx2, okx2 := lo.FindKeyBy(map[string]int{"foo": 1, "bar": 2, "baz": 3}, func(k string, v int) bool {
		return false
	})
	// "", false
	log.Println("SetStream FindKeyBy:", resultx2, okx2)

	// FindUniques
	// Returns a slice with all the unique elements of the collection. The order of result values is determined by the order they occur in the array.

	uniqueValues := lo.FindUniques[int]([]int{1, 2, 2, 1, 2, 3})
	// []int{3}
	log.Println("SetStream FindUniques:", uniqueValues)

	// FindUniquesBy
	// Returns a slice with all the unique elements of the collection.
	// The order of result values is determined by the order they occur in the array.
	// It accepts iteratee which is invoked for each element in array to generate the criterion by which uniqueness is computed.

	uniqueValues = lo.FindUniquesBy[int, int]([]int{3, 4, 5, 6, 7}, func(i int) int {
		return i % 3
	})
	// []int{5}
	log.Println("SetStream FindUniquesBy:", uniqueValues)

	// FindDuplicates
	// Returns a slice with the first occurrence of each duplicated elements of the collection. The order of result values is determined by the order they occur in the array.

	duplicatedValues := lo.FindDuplicates[int]([]int{1, 2, 2, 1, 2, 3})
	// []int{1, 2}
	log.Println("SetStream FindDuplicates:", duplicatedValues)

	// FindDuplicatesBy
	// Returns a slice with the first occurrence of each duplicated elements of the collection.
	// The order of result values is determined by the order they occur in the array.
	// It accepts iteratee which is invoked for each element in array to generate the criterion by which uniqueness is computed.

	duplicatedValues = lo.FindDuplicatesBy[int, int]([]int{3, 4, 5, 6, 7}, func(i int) int {
		return i % 3
	})
	// []int{3, 4}
	log.Println("SetStream FindDuplicatesBy:", duplicatedValues)

	// Min
	// Search the minimum value of a collection.

	// Returns zero value when collection is empty.

	min := lo.Min([]int{1, 2, 3})
	// 1
	log.Println("SetStream Min:", min)
	min = lo.Min([]int{})
	// 0
	log.Println("SetStream Min:", min)
	// MinBy
	// Search the minimum value of a collection using the given comparison function.

	// If several values of the collection are equal to the smallest value, returns the first such value.

	// Returns zero value when collection is empty.

	minx := lo.MinBy([]string{"s1", "string2", "s3"}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	// "s1"
	log.Println("SetStream MinBy:", minx)
	minx = lo.MinBy([]string{}, func(item string, min string) bool {
		return len(item) < len(min)
	})
	// ""
	log.Println("SetStream MinBy:", minx)
	// Max
	// Search the maximum value of a collection.

	// Returns zero value when collection is empty.

	max := lo.Max([]int{1, 2, 3})
	// 3
	log.Println("SetStream Max:", max)
	max = lo.Max([]int{})
	// 0
	log.Println("SetStream Max:", max)

	// MaxBy
	// Search the maximum value of a collection using the given comparison function.

	// If several values of the collection are equal to the greatest value, returns the first such value.

	// Returns zero value when collection is empty.

	maxx := lo.MaxBy([]string{"string1", "s2", "string3"}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	// "string1"
	log.Println("SetStream MaxBy:", maxx)

	maxx = lo.MaxBy([]string{}, func(item string, max string) bool {
		return len(item) > len(max)
	})
	// ""
	log.Println("SetStream MaxBy:", maxx)

	// Last
	// Returns the last element of a collection or error if empty.

	last, err := lo.Last[int]([]int{1, 2, 3})
	// 3
	log.Println("SetStream Last:", last, err)
	// Nth
	// Returns the element at index nth of collection. If nth is negative, the nth element from the end is returned. An error is returned when nth is out of slice bounds.

	nth, err := lo.Nth[int]([]int{0, 1, 2, 3}, 2)
	// 2
	log.Println("SetStream Nth:", nth, err)

	nth, err = lo.Nth[int]([]int{0, 1, 2, 3}, -2)
	// 2
	log.Println("SetStream Nth:", nth, err)

	// Sample
	// Returns a random item from collection.

	sample := lo.Sample[string]([]string{"a", "b", "c"})
	// a random string from []string{"a", "b", "c"}
	log.Println("SetStream Sample:", sample)

	sample = lo.Sample[string]([]string{})
	// ""
	log.Println("SetStream Sample:", sample)

	// Samples
	// Returns N random unique items from collection.

	samples := lo.Samples[string]([]string{"a", "b", "c"}, 3)
	// []string{"a", "b", "c"} in random order
	log.Println("SetStream Samples:", samples)
}

func Ternary() {
	// Ternary
	// A 1 line if/else statement.

	result := lo.Ternary[string](true, "a", "b")
	// "a"
	log.Println("Ternary result:", result)

	result = lo.Ternary[string](false, "a", "b")
	// "b"
	log.Println("Ternary result:", result)
	// TernaryF
	// A 1 line if/else statement whose options are functions.

	result = lo.TernaryF[string](true, func() string { return "a" }, func() string { return "b" })
	// "a"
	log.Println("Ternary result:", result)
	result = lo.TernaryF[string](false, func() string { return "a" }, func() string { return "b" })
	// "b"
	log.Println("Ternary result:", result)
}
func Chaos() {
	// 	ToPtr
	// Returns a pointer copy of value.

	ptr := lo.ToPtr[string]("hello world")
	// *string{"hello world"}
	log.Println("Chaos ToPtr:", ptr)
	// FromPtr
	// Returns the pointer value or empty.

	str := "hello world"
	value := lo.FromPtr[string](&str)
	// "hello world"
	log.Println("Chaos FromPtr:", value)

	value = lo.FromPtr[string](nil)
	// ""
	log.Println("Chaos FromPtr:", value)

	// FromPtrOr
	// Returns the pointer value or the fallback value.

	value = lo.FromPtrOr[string](&str, "empty")
	// "hello world"
	log.Println("Chaos FromPtrOr:", value)

	value = lo.FromPtrOr[string](nil, "empty")
	// "empty"
	log.Println("Chaos FromPtrOr:", value)

	// ToSlicePtr
	// Returns a slice of pointer copy of value.

	ptrx := lo.ToSlicePtr[string]([]string{"hello", "world"})
	// []*string{"hello", "world"}
	log.Println("Chaos ToSlicePtr:", ptrx)

	// ToAnySlice
	// Returns a slice with all elements mapped to any type.

	elements := lo.ToAnySlice[int]([]int{1, 5, 1})
	// []any{1, 5, 1}
	log.Println("Chaos ToAnySlice:", elements)

	// FromAnySlice
	// Returns an any slice with all elements mapped to a type. Returns false in case of type conversion failure.

	elementsx, okx := lo.FromAnySlice[string]([]any{"foobar", 42})
	// []string{}, false
	log.Println("Chaos FromAnySlice:", elementsx, okx)

	elementsx, okx = lo.FromAnySlice[string]([]any{"foobar", "42"})
	// []string{"foobar", "42"}, true
	log.Println("Chaos FromAnySlice:", elementsx, okx)

	// Empty
	// Returns an empty value.

	lo.Empty[int]()
	// 0
	lo.Empty[string]()
	lo.Empty[bool]()
	// false
	// IsEmpty
	// Returns true if argument is a zero value.

	lo.IsEmpty[int](0)
	// true
	lo.IsEmpty[int](42)
	// false

	lo.IsEmpty[string]("")
	// true

	// lo.IsEmpty[bool]("foobar")
	// false

	type test struct {
		foobar string
	}

	lo.IsEmpty[test](test{foobar: ""})
	// true
	lo.IsEmpty[test](test{foobar: "foobar"})
	// false
	// IsNotEmpty
	// Returns true if argument is a zero value.

	lo.IsNotEmpty[int](0)
	// false
	lo.IsNotEmpty[int](42)
	// true

	lo.IsNotEmpty[string]("")
	// false
	// lo.IsNotEmpty[bool]("foobar")
	// true

	lo.IsNotEmpty[test](test{foobar: ""})
	// false
	lo.IsNotEmpty[test](test{foobar: "foobar"})
	// true
	// Coalesce
	// Returns the first non-empty arguments. Arguments must be comparable.

	resultx, ok := lo.Coalesce(0, 1, 2, 3)
	// 1 true
	log.Println("Chaos Coalesce:", resultx, ok)
	// resultx, ok = lo.Coalesce("")
	// "" false
	log.Println("Chaos Coalesce:", resultx, ok)
	// var nilStr *string
	str = "foobar"
	// resultx, ok = lo.Coalesce[*string](nil, nilStr, &str)
	// &"foobar" true
	// Partial
	// Returns new function that, when called, has its first argument set to the provided value.

	add := func(x, y int) int { return x + y }
	f := lo.Partial(add, 5)

	f(10)
	// 15

	f(42)
	// 47
	// Partial2 -> Partial5
	// Returns new function that, when called, has its first argument set to the provided value.

	// addx := func(x, y, z int) int { return x + y + z }
	// f = lo.Partial2(addx, 42)

	// f(10, 5)
	// 57

	// f(42, -4)
	// 80
	// Attempt
	// Invokes a function N times until it returns valid output. Returning either the caught error or nil. When first argument is less than 1, the function runs until a successful response is returned.

	iter, err := lo.Attempt(42, func(i int) error {
		if i == 5 {
			return nil
		}

		return fmt.Errorf("failed")
	})
	// 6
	// nil
	log.Println("Chaos Attempt:", iter, err)

	iter, err = lo.Attempt(2, func(i int) error {
		if i == 5 {
			return nil
		}

		return fmt.Errorf("failed")
	})
	// 2
	// error "failed"
	log.Println("Chaos Attempt:", iter, err)
	iter, err = lo.Attempt(0, func(i int) error {
		if i < 42 {
			return fmt.Errorf("failed")
		}

		return nil
	})
	// 43
	// nil
	log.Println("Chaos Attempt:", iter, err)
	// For more advanced retry strategies (delay, exponential backoff...), please take a look on cenkalti/backoff.

	// AttemptWithDelay
	// Invokes a function N times until it returns valid output, with a pause between each call. Returning either the caught error or nil.

	// When first argument is less than 1, the function runs until a successful response is returned.

	iter, duration, err := lo.AttemptWithDelay(5, 2*time.Second, func(i int, duration time.Duration) error {
		if i == 2 {
			return nil
		}

		return fmt.Errorf("failed")
	})
	// 3
	// ~ 4 seconds
	// nil

	log.Println("Chaos AttemptWithDelay:", iter, duration, err)

	// For more advanced retry strategies (delay, exponential backoff...), please take a look on cenkalti/backoff.

	// AttemptWhile
	// Invokes a function N times until it returns valid output. Returning either the caught error or nil, and along with a bool value to identifying whether it needs invoke function continuously. It will terminate the invoke immediately if second bool value is returned with falsy value.

	// When first argument is less than 1, the function runs until a successful response is returned.

	// count1, err1 := lo.AttemptWhile(5, func(i int) (error, bool) {
	// 	err := doMockedHTTPRequest(i)
	// 	if err != nil {
	// 		if errors.Is(err, ErrBadRequest) { // lets assume ErrBadRequest is a critical error that needs to terminate the invoke
	// 			return err, false // flag the second return value as false to terminate the invoke
	// 		}

	// 		return err, true
	// 	}

	// 	return nil, false
	// })
	// For more advanced retry strategies (delay, exponential backoff...), please take a look on cenkalti/backoff.

	// AttemptWhileWithDelay
	// Invokes a function N times until it returns valid output, with a pause between each call. Returning either the caught error or nil, and along with a bool value to identifying whether it needs to invoke function continuously. It will terminate the invoke immediately if second bool value is returned with falsy value.

	// When first argument is less than 1, the function runs until a successful response is returned.

	// count1, time1, err1 := lo.AttemptWhileWithDelay(5, time.Millisecond, func(i int, d time.Duration) (error, bool) {
	// 	err := doMockedHTTPRequest(i)
	// 	if err != nil {
	// 		if errors.Is(err, ErrBadRequest) { // lets assume ErrBadRequest is a critical error that needs to terminate the invoke
	// 			return err, false // flag the second return value as false to terminate the invoke
	// 		}

	// 		return err, true
	// 	}

	// 	return nil, false
	// })
	// For more advanced retry strategies (delay, exponential backoff...), please take a look on cenkalti/backoff.

	// Debounce
	// NewDebounce creates a debounced instance that delays invoking functions given until after wait milliseconds have elapsed, until cancel is called.

	fx := func() {
		println("Called once after 100ms when debounce stopped invoking!")
	}

	debounce, cancel := lo.NewDebounce(100*time.Millisecond, fx)
	for j := 0; j < 10; j++ {
		debounce()
	}

	time.Sleep(1 * time.Second)
	cancel()

	// Synchronize
	// Wraps the underlying callback in a mutex. It receives an optional mutex.

	s := lo.Synchronize()

	for i := 0; i < 10; i++ {
		go s.Do(func() {
			println("will be called sequentially")
		})
	}
	// It is equivalent to:

	// mu := sync.Mutex{}

	// func foobar() {
	//     mu.Lock()
	//     defer mu.Unlock()

	//     // ...
	// }
	// Async
	// Executes a function in a goroutine and returns the result in a channel.

	// ch := lo.Async(func() error { time.Sleep(10 * time.Second); return nil })
	// chan error (nil)
	// Async{0->6}
	// Executes a function in a goroutine and returns the result in a channel.
	// For function with multiple return values, the results will be returned as a tuple inside the channel.
	// For function without return, struct{} will be returned in the channel.

	// ch = lo.Async0(func() { time.Sleep(10 * time.Second) })
	// chan struct{}

	// ch = lo.Async1(func() int {
	// 	time.Sleep(10 * time.Second)
	// 	return 42
	// })
	// chan int (42)

	// ch = lo.Async2(func() (int, string) {
	// 	time.Sleep(10 * time.Second)
	// 	return 42, "Hello"
	// })
	// chan lo.Tuple2[int, string] ({42, "Hello"})
	// Transaction
	// Implements a Saga pattern.

	// transaction := NewTransaction[int]().
	// 	Then(
	// 		func(state int) (int, error) {
	// 			fmt.Println("step 1")
	// 			return state + 10, nil
	// 		},
	// 		func(state int) int {
	// 			fmt.Println("rollback 1")
	// 			return state - 10
	// 		},
	// 	).
	// 	Then(
	// 		func(state int) (int, error) {
	// 			fmt.Println("step 2")
	// 			return state + 15, nil
	// 		},
	// 		func(state int) int {
	// 			fmt.Println("rollback 2")
	// 			return state - 15
	// 		},
	// 	).
	// 	Then(
	// 		func(state int) (int, error) {
	// 			fmt.Println("step 3")

	// 			if true {
	// 				return state, fmt.Errorf("error")
	// 			}

	// 			return state + 42, nil
	// 		},
	// 		func(state int) int {
	// 			fmt.Println("rollback 3")
	// 			return state - 42
	// 		},
	// 	)

	// _, _ = transaction.Process(-5)

	// Output:
	// step 1
	// step 2
	// step 3
	// rollback 2
	// rollback 1
	// Validate
	// Helper function that creates an error when a condition is not met.

	slice := []string{"a"}
	// val := lo.Validate(len(slice) == 0, "Slice should be empty but contains %v", slice)
	// error("Slice should be empty but contains [a]")

	slice = []string{}
	// val:
	lo.Validate(len(slice) == 0, "Slice should be empty but contains %v", slice)
	// nil

	// Must
	// Wraps a function call to panics if second argument is error or false, returns the value otherwise.

	// val = lo.Must(time.Parse("2006-01-02", "2022-01-15"))
	// 2022-01-15

	// val = lo.Must(time.Parse("2006-01-02", "bad-value"))
	// panics

	// Must{0->6}
	// Must* has the same behavior as Must, but returns multiple values.
	/*
	   func example0() (error)
	   func example1() (int, error)
	   func example2() (int, string, error)
	   func example3() (int, string, time.Date, error)
	   func example4() (int, string, time.Date, bool, error)
	   func example5() (int, string, time.Date, bool, float64, error)
	   func example6() (int, string, time.Date, bool, float64, byte, error)

	   lo.Must0(example0())
	   val1 := lo.Must1(example1())    // alias to Must
	   val1, val2 := lo.Must2(example2())
	   val1, val2, val3 := lo.Must3(example3())
	   val1, val2, val3, val4 := lo.Must4(example4())
	   val1, val2, val3, val4, val5 := lo.Must5(example5())
	   val1, val2, val3, val4, val5, val6 := lo.Must6(example6())
	*/
	// You can wrap functions like func (...) (..., ok bool).

	// math.Signbit(float64) bool
	// lo.Must0(math.Signbit(v))

	// bytes.Cut([]byte,[]byte) ([]byte, []byte, bool)
	// before, after := lo.Must2(bytes.Cut(s, sep))
	// You can give context to the panic message by adding some printf-like arguments.

	// val, ok := lo.Find(myString, func(i string) bool {
	// 	return i == requiredChar
	// })
	// lo.Must0(ok, "'%s' must always contain '%s'", myString, requiredChar)

	list := []int{0, 1, 2}
	item := 5
	lo.Must0(lo.Contains[int](list, item), "'%s' must always contain '%s'", list, item)

	// Try
	// Calls the function and return false in case of error and on panic.

	ok = lo.Try(func() error {
		panic("error")
		return nil
	})
	// false

	ok = lo.Try(func() error {
		return nil
	})
	// true

	ok = lo.Try(func() error {
		return fmt.Errorf("error")
	})
	// false

	// Try{0->6}
	// The same behavior than Try, but callback returns 2 variables.

	ok = lo.Try2(func() (string, error) {
		panic("error")
		return "", nil
	})
	// false

	// TryOr
	// Calls the function and return a default value in case of error and on panic.

	str, ok = lo.TryOr(func() (string, error) {
		panic("error")
		return "hello", nil
	}, "world")
	// world
	// false

	// TryOr{0->6}
	// The same behavior than TryOr, but callback returns 2 variables.

	str, nbr, ok := lo.TryOr2(func() (string, int, error) {
		panic("error")
		return "hello", 42, nil
	}, "world", 21)
	// world
	// 21
	// false
	log.Println("Chaos TryOr2:", str, nbr, ok)
	// TryWithErrorValue
	// The same behavior than Try, but also returns value passed to panic.

	// err, ok = lo.TryWithErrorValue(func() error {
	// 	panic("error")
	// 	return nil
	// })
	// "error", false

	// TryCatch
	// The same behavior than Try, but calls the catch function in case of error.

	// caught := false

	// ok = lo.TryCatch(func() error {
	// 	panic("error")
	// 	return nil
	// }, func() {
	// 	caught = true
	// })
	// false
	// caught == true

	// TryCatchWithErrorValue
	// The same behavior than TryWithErrorValue, but calls the catch function in case of error.

	// caught := false

	// ok = lo.TryCatchWithErrorValue(func() error {
	// 	panic("error")
	// 	return nil
	// }, func(val any) {
	// 	caught = val == "error"
	// })
	// false
	// caught == true

	// ErrorsAs
	// A shortcut for:

	// err := doSomething()

	// var rateLimitErr *RateLimitError
	// if ok := errors.As(err, &rateLimitErr); ok {
	// 	// retry later
	// }
	// // 1 line lo helper:

	// err := doSomething()

	// if rateLimitErr, ok := lo.ErrorsAs[*RateLimitError](err); ok {
	// 	// retry later
	// }
}
