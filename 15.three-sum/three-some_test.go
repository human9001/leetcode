package main

import (
	"reflect"
	"testing"
)

type addTest struct {
	arg1 []int
	want [][]int
}

var result [][]int

var addTests = []addTest{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	{[]int{0, 1, 1}, result},
}

func TestThreeSum(t *testing.T) {
	for idx, test := range addTests {
		if output := threeSum(test.arg1); !reflect.DeepEqual(output, test.want) {
			t.Errorf("Test %d. Output %v (type %T) not equal to expected %v (type %T)", idx, output, output, test.want, test.want)

		}

	}
}

func BenchmarkThreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSum([]int{-1, 0, 1, 2, -1, -4})
	}
}
