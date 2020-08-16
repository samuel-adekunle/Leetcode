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
	"fmt"
	"testing"
)

func Example() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))

	arr = make([]int, 1000)
	fmt.Println(threeSum(arr))

	// Output:
	// [[-1 -1 2] [-1 0 1]]
	// [[0 0 0]]
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSum([]int{-1, 0, 1, 2, -1, -4})
	}
}
