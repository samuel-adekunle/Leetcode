/* Write a program to find the n-th ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Example:

Input: n = 10
Output: 12
Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
Note:

1 is typically treated as an ugly number.
n does not exceed 1690. */

package main

import (
	"fmt"
	"testing"
)

func ExampleBruteForce() {
	n := 100
	fmt.Println(nthUglyNumber(n))

	// Output:
	// 1536
}

func BenchmarkBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nthUglyNumber(100)
	}
}

func ExampleDP() {
	n := 100
	fmt.Println(nthUglyNumber(n))
	fmt.Println(dp(n))

	// Output:
	// 1536
	// 1536
}

func BenchmarkDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dp(100)
	}
}
