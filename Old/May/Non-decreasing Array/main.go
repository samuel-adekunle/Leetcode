package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkPossibility([]int{5, 7, 1, 8}))
}

func checkPossibility(nums []int) bool {
	violations := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if violations == 1 {
				return false
			}
			violations++
			if i == 1 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return true
}
