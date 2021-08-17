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
	"fmt"
	"testing"
)

func Example() {
	fmt.Println(fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
	// Output:
	// [[-5 -4 -3 1]]
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11)
	}
}
