package main

import (
	"testing"
)

type addTest struct {
	arg1 string
	want int
}

var addTests = []addTest{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range addTests {
		if output := lengthOfLongestSubstring(test.arg1); output != test.want {
			t.Errorf("Output %v not equal to expected %v", output, test.want)

		}

	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("pwwkew")
	}
}
