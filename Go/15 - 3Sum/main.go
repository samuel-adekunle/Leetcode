/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
] */

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

func threeSum(nums []int) [][]int {
	var arr [][]int
	target := 0
	mem := make(map[[3]int]bool)
	sort.Ints(nums)

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
