/* Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
Accept */

package main

import (
	"sort"
)

func twoSumSorted(target int, nums []int) [][]int {
	var arr [][]int
	left, right := 0, len(nums)-1
	for right > left {
		sum := nums[left] + nums[right]
		if sum == target {
			arr = append(arr, []int{nums[left], nums[right]})
			left++
			right--
			continue
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}
	return arr
}

func threeSumSorted(target int, nums []int) [][]int {
	var arr [][]int
	mem := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] > target && nums[i] > 0 {
			break
		}
		twoSums := twoSumSorted(target-nums[i], nums[i+1:])
		for _, v := range twoSums {
			sub := [3]int{nums[i], v[0], v[1]}
			if !mem[sub] {
				arr = append(arr, []int{nums[i], v[0], v[1]})
			}
			mem[sub] = true
		}
	}
	return arr
}

func fourSum(nums []int, target int) [][]int {
	var arr [][]int
	mem := make(map[[4]int]bool)
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > target && nums[i] > 0 {
			break
		}
		threeSums := threeSumSorted(target-nums[i], nums[i+1:])
		for _, v := range threeSums {
			sub := [4]int{nums[i], v[0], v[1], v[2]}
			if !mem[sub] {
				arr = append(arr, []int{nums[i], v[0], v[1], v[2]})
			}
			mem[sub] = true
		}
	}
	return arr
}
