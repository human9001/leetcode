package main

import (
	"reflect"
	"testing"
)

type addTest struct {
	arg1 []int
	arg2 int
	want []int
}

var addTests = []addTest{
	{[]int{1, 2, 3, 4, 5}, 4, []int{0, 2}},
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, test := range addTests {
		if output := twoSum(test.arg1, test.arg2); !reflect.DeepEqual(output, test.want) {
			t.Errorf("Output %v not equal to expected %v", output, test.want)

		}

	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{2, 7, 11, 15}, 9)
	}
}

func TestTwoSum2(t *testing.T) {
	for _, test := range addTests {
		if output := twoSum(test.arg1, test.arg2); !reflect.DeepEqual(output, test.want) {
			t.Errorf("Output %q not equal to expected %q", output, test.want)

		}

	}
}

func BenchmarkTwoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum2([]int{2, 7, 11, 15}, 9)
	}
}
