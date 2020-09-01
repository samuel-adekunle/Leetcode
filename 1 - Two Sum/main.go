/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

// brute force - O(n^2)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// alternative - O(n)
func twoSumUseMap(nums []int, target int) []int {
	mem := make(map[int]int)
	for i, v := range nums {
		if ci, ok := mem[target-v]; ok {
			return []int{ci, i}
		}
		mem[v] = i
	}
	return nil
}

// Won't work if nums isn't sorted
func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for right > left {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
