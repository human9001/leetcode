package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	b := []rune(fmt.Sprintf("%d", x))
	c := []rune{}
	for i := len(b) - 1; i >= 0; i-- {
		c = append(c, b[i])
	}

	return string(c) == string(b)
}

func isPalindrome2(x int) bool {
	if 0 <= x && x < 10 {
		return true
	}
	numberStr := fmt.Sprintf("%d", x)
	limit := len(numberStr) / 2
	for idx := 0; idx <= limit; idx++ {
		if string(numberStr[idx]) != numberStr[len(numberStr)-(idx+1):len(numberStr)-idx] {
			return false
		}
	}
	return true
}
