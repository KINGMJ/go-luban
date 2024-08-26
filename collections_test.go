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

	result := luban.MapSlice(numbers, func(x int) int {
		return x * 2
	})
	assert.Equal(t, expected, result)
}

func TestMapIntToFloat(t *testing.T) {
	numbers := []int{1, 2, 3}
	expected := []float64{1.0, 2.0, 3.0}

	result := luban.MapSlice(numbers, func(x int) float64 {
		return float64(x)
	})
	assert.Equal(t, expected, result)
}

func TestMapEmptySlice(t *testing.T) {
	empty := []int{}
	expected := []string{}

	result := luban.MapSlice(empty, func(x int) string {
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

	result := luban.MapSlice(input, func(a A) B {
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
