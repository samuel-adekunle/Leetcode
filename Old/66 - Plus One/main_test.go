/* Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321. */

package main

import (
	"fmt"
	"testing"
)

func Example() {
	fmt.Println(plusOne([]int{9, 9, 9}))

	// Output:
	// [1 0 0 0]
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusOne([]int{9, 9, 9})
	}
}
