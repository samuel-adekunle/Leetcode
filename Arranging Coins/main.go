package main

import (
	"math"
)

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

func mathArrangeCoins(n int) int {
	return int(math.Sqrt(float64(2*n)+0.25) - 0.5)
}
