package main

import (
	"reflect"
	"slices"
	"testing"
)

type addTest struct {
	arg1 string
	want []string
}

var result [][]int

var addTests = []addTest{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{"", []string{}},
	{"2", []string{"a", "b", "c"}},
}

func TestLetterCombinations(t *testing.T) {
	for idx, test := range addTests {

		output := letterCombinations(test.arg1)
		slices.Sort(output)
		slices.Sort(test.want)
		if !reflect.DeepEqual(output, test.want) {
			t.Errorf("Test %d. Output %v (type %T) not equal to expected %v (type %T)", idx, output, output, test.want, test.want)

		}

	}
}

func BenchmarkLetterCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations("23")
	}
}
