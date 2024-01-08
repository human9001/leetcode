package main

import (
	"testing"
)

type addTest struct {
	arg1 int
	want bool
}

var addTests = []addTest{
	{1221, true},
	{12231, false},
	{-1221, false},
	{1345431, true},
	{10000021, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range addTests {
		if output := isPalindrome(test.arg1); output != test.want {
			t.Errorf("Output %v not equal to expected %v", output, test.want)

		}

	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(123321)
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, test := range addTests {
		if output := isPalindrome2(test.arg1); output != test.want {
			t.Errorf("Output %v not equal to expected %v", output, test.want)

		}

	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2(123321)
	}
}
