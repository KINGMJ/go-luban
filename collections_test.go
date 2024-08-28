package luban_test

import (
	"fmt"
	"testing"

	"github.com/KINGMJ/luban"
	"github.com/stretchr/testify/assert"
)

// -----------------------
// map slice test
// -----------------------
func TestMapIntToInt(t *testing.T) {
	numbers := []int{1, 2, 3}
	expected := []int{2, 4, 6}

	result := luban.Map(numbers, func(x int) int {
		return x * 2
	})
	assert.Equal(t, expected, result)
}

func TestMapIntToFloat(t *testing.T) {
	numbers := []int{1, 2, 3}
	expected := []float64{1.0, 2.0, 3.0}

	result := luban.Map(numbers, func(x int) float64 {
		return float64(x)
	})
	assert.Equal(t, expected, result)
}

func TestMapEmptySlice(t *testing.T) {
	empty := []int{}
	expected := []string{}

	result := luban.Map(empty, func(x int) string {
		return fmt.Sprintf("%d", x)
	})

	assert.Equal(t, expected, result)
}

func TestMapStructToStruct(t *testing.T) {
	type A struct {
		Num int
	}
	type B struct {
		Str string
	}

	input := []A{{1}, {2}, {3}}
	expected := []B{{"1"}, {"2"}, {"3"}}

	result := luban.Map(input, func(a A) B {
		return B{fmt.Sprintf("%d", a.Num)}
	})

	assert.Equal(t, expected, result)
}

// -----------------------
// map map test
// -----------------------
func TestMapMapToMap(t *testing.T) {
	dict := map[string]int{"one": 1, "two": 2, "three": 3}
	expected := map[string]string{"one": "one=1", "two": "two=2", "three": "three=3"}

	result := luban.MapMap(dict, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})

	assert.Equal(t, expected, result)
}

// -----------------------
// filter slice test
// -----------------------

func TestFilterSliceInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	result := luban.Filter(numbers, func(x int) bool {
		return x%2 == 0
	})

	assert.Equal(t, expected, result)
}

func TestFilterSliceString(t *testing.T) {
	words := []string{"apple", "banana", "cherry", "date", "fig"}
	expected := []string{"banana", "cherry"}

	result := luban.Filter(words, func(s string) bool {
		return len(s) == 6
	})

	assert.Equal(t, expected, result)
}

func TestFilterMap(t *testing.T) {
	dict := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
	expected := map[string]int{"two": 2, "four": 4}

	result := luban.FilterMap(dict, func(k string, v int) bool {
		return v%2 == 0
	})

	assert.Equal(t, expected, result)
}

func TestFilterMapStringKey(t *testing.T) {
	dict := map[string]string{"apple": "fruit", "carrot": "vegetable", "banana": "fruit"}
	expected := map[string]string{"apple": "fruit", "banana": "fruit"}

	result := luban.FilterMap(dict, func(k, v string) bool {
		return v == "fruit"
	})

	assert.Equal(t, expected, result)
}

func TestReduceSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := 15

	result := luban.Reduce(numbers, func(acc, x int) int {
		return acc + x
	}, 0)

	assert.Equal(t, expected, result)
}

func TestReduceConcat(t *testing.T) {
	words := []string{"Hello", "World", "!"}
	expected := "HelloWorld!"

	result := luban.Reduce(words, func(acc, s string) string {
		return acc + s
	}, "")

	assert.Equal(t, expected, result)
}

func TestReduceMapToSum(t *testing.T) {
	dict := map[string]int{"one": 1, "two": 2, "three": 3}
	expected := 6

	result := luban.ReduceMap(dict, func(acc int, k string, v int) int {
		return acc + v
	}, 0)

	assert.Equal(t, expected, result)
}

func TestReduceMapToSlice(t *testing.T) {
	dict := map[string]int{"one": 1, "two": 2, "three": 3}
	expected := []string{"one:1", "two:2", "three:3"}

	result := luban.ReduceMap(dict, func(acc []string, k string, v int) []string {
		value := fmt.Sprintf("%s:%d", k, v)
		if len(acc) == 0 {
			return []string{value}
		}
		return append(acc, value)
	}, nil)

	assert.Equal(t, expected, result)
}

func TestEach(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	luban.Each(numbers, func(x int) {
		sum += x
	})

	assert.Equal(t, 15, sum)
}

func TestEachMap(t *testing.T) {
	dict := map[string]int{"one": 1, "two": 2, "three": 3}
	keys := []string{}
	values := []int{}

	luban.EachMap(dict, func(k string, v int) {
		keys = append(keys, k)
		values = append(values, v)
	})

	expectedKeys := []string{"one", "two", "three"}
	expectedValues := []int{1, 2, 3}

	assert.ElementsMatch(t, expectedKeys, keys)
	assert.ElementsMatch(t, expectedValues, values)
}
