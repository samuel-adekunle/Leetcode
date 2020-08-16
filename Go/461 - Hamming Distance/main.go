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

// Using conversion to calculate hamming distance
func hammingDistance(x int, y int) int {
	dist := 0
	for x > 0 || y > 0 {
		if x%2 != y%2 {
			dist++
		}
		x /= 2
		y /= 2
	}
	return dist
}

func oneBitCount(x int) int {
	var i int
	for i = 0; x > 0; i++ {
		x &= (x - 1)
	}
	return i
}

// Xor Alternative
func xorDistance(x int, y int) int {
	dist := x ^ y
	return oneBitCount(dist)
}
