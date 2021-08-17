/*
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Given two integers x and y, calculate the Hamming distance.

Note:
0 ≤ x, y < 231.

Example:

Input: x = 1, y = 4

Output: 2

Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

The above arrows point to positions where the corresponding bits are different. */

package main

import (
	"fmt"
	"testing"
)

func Example() {
	fmt.Println(hammingDistance(1, 32))
	fmt.Println(xorDistance(1, 32))
	// Output:
	// 2
	// 2
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingDistance(1, 32)
	}
}

func BenchmarkXorDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		xorDistance(1, 32)
	}
}
