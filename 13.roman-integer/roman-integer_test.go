package main

import (
	"testing"
)

type addTest struct {
	arg1 string
	want int
}

var addTests = []addTest{
	{"III", 3},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
	{"V", 5},
}

func TestRomanToInt(t *testing.T) {
	for _, test := range addTests {
		if output := romanToInt(test.arg1); output != test.want {
			t.Errorf("Output %v not equal to expected %v", output, test.want)

		}

	}
}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt("MCMXCIV")
	}
}
