package main

import (
	"testing"
)

type addTest struct {
	arg1 []string
	want string
}

var addTests = []addTest{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range addTests {
		if output := longestCommonPrefix(test.arg1); output != test.want {
			t.Errorf("Output %v not equal to expected %v", output, test.want)

		}

	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}
