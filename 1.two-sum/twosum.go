package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsLength := len(nums)
	for i := 0; i < numsLength; i++ {
		for j := i + 1; j < numsLength; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	storage := make(map[int]int)
	for idx, num := range nums {
		if value, ok := storage[target-num]; ok {
			return []int{value, idx}
		}
		storage[num] = idx
	}
	return []int{}
}

func main() {
	a := []int{2, 7, 11, 15}
	fmt.Println(twoSum(a, 9))
	b := []int{2, 7, 11, 15}
	fmt.Println(twoSum(b, 9))
}
