package main

import (
	"fmt"
	"math/bits"
	"strings"
)

func letterCombinations(digits string) []string {

	letters := map[string][]string{
		"1": {},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
		"0": {" "},
	}
	if len(digits) == 1 {
		return letters[digits]
	}
	if len(digits) == 0 {
		return letters["1"]
	}
	allLetters := make([]string, 0)
	letterCombinations := make([]string, 0)
	lettersIntervals := make([]int, 0)
	dl := len(digits)
	for i := 0; i <= dl-1; i++ {
		allLetters = append(allLetters, letters[string(digits[i])]...)
		lettersIntervals = append(lettersIntervals, len(letters[string(digits[i])]))
	}
	ddl := len(allLetters)
	totalCombinations := (1 << ddl) - 1

MAIN:
	for i := 1; i <= totalCombinations; i++ {
		onesCount := bits.OnesCount(uint(i))
		binaryStr := fmt.Sprintf("%016b", i)
		if onesCount == dl {
			cutSizeStart := 0
			cutSizeEnd := 0
			letterCombination := ""
			for j := 0; j <= len(lettersIntervals)-1; j++ {
				cutSizeEnd = len(binaryStr) - cutSizeStart

				cutSizeStart += lettersIntervals[j]
				currectSlice := binaryStr[len(binaryStr)-cutSizeStart : cutSizeEnd]
				if strings.Count(currectSlice, "1") > 1 || strings.Count(currectSlice, "1") <= 0 {
					continue MAIN
				}
				onePos := strings.Index(currectSlice, "1")
				letterCombination += allLetters[cutSizeStart-(onePos+1)]

			}
			letterCombinations = append(letterCombinations, letterCombination)

		}
	}
	return letterCombinations
}
