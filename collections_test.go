package luban_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/KINGMJ/luban"
	"github.com/stretchr/testify/assert"
)

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

func TestMapStructToSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	persons := []Person{
		{"John", 25},
		{"Amy", 30},
		{"Bob", 40},
	}
	expected := []string{"John", "Amy", "Bob"}

	result := luban.Map(persons, func(p Person) string {
		return p.Name
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

func TestMapMapToMap(t *testing.T) {
	dict := map[string]int{"one": 1, "two": 2, "three": 3}
	expected := map[string]string{"one": "one=1", "two": "two=2", "three": "three=3"}

	result := luban.MapMap(dict, func(k string, v int) string {
		return fmt.Sprintf("%s=%d", k, v)
	})

	assert.Equal(t, expected, result)
}

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

func TestEachRight(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}

	var result []int
	luban.EachRight(numbers, func(x int) {
		result = append(result, x)
	})
	assert.Equal(t, expected, result)
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

func TestEvery(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := luban.Every(numbers, func(n int) bool {
		return n > 0
	})
	assert.True(t, result)

	numbers = []int{1, -2, 3, 4, 5}
	result = luban.Every(numbers, func(n int) bool {
		return n > 0
	})
	assert.False(t, result)

	numbers = []int{}
	result = luban.Every(numbers, func(n int) bool {
		return n > 0
	})
	assert.True(t, result)
}

func TestEveryMap(t *testing.T) {
	dict := map[string]int{"a": 1, "b": 2, "c": 3}
	result := luban.EveryMap(dict, func(k string, v int) bool {
		return v > 0
	})
	assert.True(t, result)

	dict = map[string]int{"a": 1, "b": -2, "c": 3}
	result = luban.EveryMap(dict, func(k string, v int) bool {
		return v > 0
	})
	assert.False(t, result)

	dict = map[string]int{}
	result = luban.EveryMap(dict, func(k string, v int) bool {
		return v > 0
	})
	assert.True(t, result)
}

func TestSome(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	isEven := func(x int) bool { return x%2 == 0 }
	isNegative := func(x int) bool { return x < 0 }

	assert.True(t, luban.Some(numbers, isEven))
	assert.False(t, luban.Some(numbers, isNegative))

	emptySlice := []int{}
	assert.False(t, luban.Some(emptySlice, isEven))
}

func TestSomeMap(t *testing.T) {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	containsTwo := func(k string, v int) bool { return v == 2 }
	containsZero := func(k string, v int) bool { return v == 0 }

	assert.True(t, luban.SomeMap(m, containsTwo))
	assert.False(t, luban.SomeMap(m, containsZero))

	emptyMap := map[string]int{}
	assert.False(t, luban.SomeMap(emptyMap, containsTwo))
}

func TestFind(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	index, value := luban.Find(numbers, func(x int) bool {
		return x > 3
	})
	assert.Equal(t, 3, index)
	assert.Equal(t, 4, value)

	index, value = luban.Find(numbers, func(x int) bool {
		return x > 10
	})
	assert.Equal(t, -1, index)
	assert.Equal(t, 0, value)
}

func TestFindMap(t *testing.T) {
	myMap := map[string]int{
		"apple":  5,
		"banana": 10,
		"cherry": 15,
	}
	key, value := luban.FindMap(myMap, func(k string, v int) bool {
		return v > 10
	})
	assert.Equal(t, "cherry", key)
	assert.Equal(t, 15, value)

	key, value = luban.FindMap(myMap, func(k string, v int) bool {
		return v > 20
	})
	assert.Equal(t, "", key)
	assert.Equal(t, 0, value)
}

func TestChunk(t *testing.T) {
	// Test case 1: Normal scenario with integer slice
	t.Run("Chunk integer slice", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := [][]int{
			{1, 2},
			{3, 4},
			{5},
		}
		result, err := luban.Chunk(input, 2)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	// Test case 2: Chunk size larger than slice length
	t.Run("Chunk size larger than slice", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := [][]int{
			{1, 2, 3},
		}
		result, err := luban.Chunk(input, 5)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	// Test case 3: Chunk size equal to slice length
	t.Run("Chunk size equal to slice length", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := [][]int{
			{1, 2, 3},
		}
		result, err := luban.Chunk(input, 3)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	// Test case 4: Chunk size of 1
	t.Run("Chunk size of 1", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := [][]int{
			{1},
			{2},
			{3},
		}
		result, err := luban.Chunk(input, 1)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	// Test case 5: Chunk with empty slice
	t.Run("Chunk empty slice", func(t *testing.T) {
		input := []int{}
		expected := [][]int{}
		result, err := luban.Chunk(input, 3)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	// Test case 6: Invalid chunk size (zero)
	t.Run("Invalid chunk size zero", func(t *testing.T) {
		input := []int{1, 2, 3}
		result, err := luban.Chunk(input, 0)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, errors.New("cannot be less than 1"), err)
	})

	// Test case 7: Invalid chunk size (negative)
	t.Run("Invalid chunk size negative", func(t *testing.T) {
		input := []int{1, 2, 3}
		result, err := luban.Chunk(input, -1)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, errors.New("cannot be less than 1"), err)
	})
}

func TestCompact(t *testing.T) {
	t.Run("Compact with int slice", func(t *testing.T) {
		input := []int{0, 1, 2, 2, 0, 3, 0, 4}
		expected := []int{1, 2, 2, 3, 4}
		result := luban.Compact(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Compact with string slice", func(t *testing.T) {
		input := []string{"", "foo", "", "bar", ""}
		expected := []string{"foo", "bar"}
		result := luban.Compact(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Compact with mixed slice", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		ch := make(chan int, 1)
		var x any
		var y any = 32
		var fn func(int) bool
		ptr1 := &x
		ptr2 := &y
		input := []any{0, 1, 1, "", nil, true, false, map[int]string{}, ch, Person{}, []*Person{}, []int64{}, x, y, fn, ptr1, ptr2}
		expected := []any{1, 1, true, y, ptr2}
		result := luban.Compact(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Compact with empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := luban.Compact(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Compact with slice of pointers", func(t *testing.T) {
		a := 1
		b := 0
		c := 3
		input := []*int{&b, nil, &a, &b, &c}
		expected := []*int{&a, &c}
		result := luban.Compact(input)
		assert.Equal(t, expected, result)
	})
}
