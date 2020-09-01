/* You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of full staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1:

n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.
Example 2:

n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

Because the 4th row is incomplete, we return 3. */

package main

import (
	"fmt"
	"testing"
)

func ExampleArrangeCoins() {
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(10))

	// Output:
	// 2
	// 3
	// 4
}

func BenchmarkArrangeCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arrangeCoins(10000)
	}
}

func ExampleBinary() {
	fmt.Println(arrangeCoins(20))
	fmt.Println(binaryArrangeCoins(20))

	// Output:
	// 5
	// 5
}

func BenchmarkBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryArrangeCoins(10000)
	}
}

func ExampleMath() {
	fmt.Println(arrangeCoins(20))
	fmt.Println(binaryArrangeCoins(20))
	fmt.Println(mathArrangeCoins(20))

	// Output:
	// 5
	// 5
	// 5
}

func BenchmarkMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathArrangeCoins(10000)
	}
}
