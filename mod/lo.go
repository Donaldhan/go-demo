package mod

/**
第三方工具包测试
*/

import (
	"fmt"
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
}
