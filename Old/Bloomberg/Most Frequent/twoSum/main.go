package main

import (
	"fmt"
	"log"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func main() {
	nums := []int{2,7,11,15}
	target := 9

	xi := twoSum(nums, target)
	if xi != nil {
		fmt.Printf("%d + %d = %d\n", nums[xi[0]], nums[xi[1]], target)
	}
}

// unordered array nums
func twoSum(nums []int, target int) []int {
	mem := make(map[int]int, len(nums))

	for i, v := range nums {
		u := target - v
		if j, ok := mem[u]; ok {
			return []int{i, j}
		}
		mem[v] = i
	}
	log.Fatalf("Cannot find two sum for %d from %v", target, nums)
	return nil
}