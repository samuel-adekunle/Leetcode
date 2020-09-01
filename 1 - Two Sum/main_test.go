/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
package main

import (
	"fmt"
	"testing"
)

func ExampleTwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	// Output:
	// [0 1]
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{2, 20, 22, 27, 28, 43, 1234, 43, 4345, 35465, 322, 3121, 12334, 23, 433, 243, 34, 344, 7, 11, 15}, 18)
	}
}

func ExampleTwoSumUseMap() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumUseMap(nums, target))
	// Output:
	// [0 1]
}

func BenchmarkTwoSumUseMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumUseMap([]int{2, 20, 22, 27, 28, 43, 1234, 43, 4345, 35465, 322, 3121, 12334, 23, 433, 243, 34, 344, 7, 11, 15}, 18)
	}
}

func ExampleTwoSumSorted() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumSorted(nums, target))
	// Output:
	// [0 1]
}

func BenchmarkTwoSumSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSumSorted([]int{2, 20, 22, 27, 28, 43, 1234, 43, 4345, 35465, 322, 3121, 12334, 23, 433, 243, 34, 344, 7, 11, 15}, 18)
	}
}
