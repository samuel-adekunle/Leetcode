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
	"math"
)

// Brute force - O(n)
func arrangeCoins(n int) int {
	tri, i := 0, 0
	for tri < n {
		i++
		tri += i
	}
	if tri > n {
		return i - 1
	}
	return i
}

// Binary search - O(log n)
func binaryArrangeCoins(n int) int {
	left, right := 0, n
	for left <= right {
		k := left + (right-left)/2
		curr := k * (k + 1) / 2
		if curr == n {
			return k
		}
		if curr > n {
			right = k - 1
		} else {
			left = k + 1
		}
	}
	return right
}

// Math formula - O(1)
func mathArrangeCoins(n int) int {
	return int(math.Sqrt(float64(2*n)+0.25) - 0.5)
}
