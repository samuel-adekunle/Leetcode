package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func jump(nums []int) int {
	curr, target := 0, len(nums)-1
	var helper func(curr int) int

	cache := map[int]int{}
	helper = func(curr int) int {
		if curr == target {
			return 0
		}
		if val, ok := cache[curr]; ok {
			return  val
		}
		minSteps := 100
		for i := 1; i <= nums[curr] && curr+i <= target; i++ {
			minSteps = min(minSteps, 1+helper(curr+i))
		}
		cache[curr] = minSteps
		return minSteps
	}
	return helper(curr)
}
