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
	"math"
)

func divByNUntilRem(num int, n ...int) int {
	for _, v := range n {
		for num%v == 0 {
			num /= v
		}
	}
	return num
}

func sieve(n int) bool {
	if n == 1 {
		return true
	}
	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		if divByNUntilRem(n, 2, 3, 5) == 1 {
			return true
		}
	}
	return false
}

// Brute force
func nthUglyNumber(n int) int {
	i := 0
	curr := 0
	for i < n {
		curr++
		if sieve(curr) {
			i++
		}
	}
	return curr
}

func multiplySlice(arr []int, n int) []int {
	newArr := make([]int, len(arr))
	for i, v := range arr {
		newArr[i] = v * n
	}
	return newArr
}

func combineSlices(arr1, arr2 []int) []int {
	merge := []int{}
	i, i1, i2 := -1, 0, 0
	for i1 < len(arr1) && i2 < len(arr2) {
		if arr1[i1] < arr2[i2] {
			if i != -1 && merge[i] == arr1[i1] {
				i1++
				continue
			}
			merge = append(merge, arr1[i1])
			i1++
		} else {
			if i != -1 && merge[i] == arr2[i2] {
				i2++
				continue
			}
			merge = append(merge, arr2[i2])
			i2++
		}
		i++
	}
	if i1 < len(arr1) {
		merge = append(merge, arr1[i1:]...)
	}
	if i2 < len(arr2) {
		merge = append(merge, arr2[i2:]...)
	}
	return merge
}

// Merge sort - failed
func triples(n int) int {
	ugly := []int{0, 1, 2, 3, 5}
	start := 2
	for len(ugly) <= n {
		two := multiplySlice(ugly[start:], 2)
		three := multiplySlice(ugly[start:], 3)
		four := multiplySlice(ugly[start:], 4)
		five := multiplySlice(ugly[start:], 5)
		merge := combineSlices(combineSlices(two, three), combineSlices(four, five))
		start = len(ugly)
		ugly = combineSlices(ugly, merge)
	}
	return ugly[n]
}

func genUgly(numbers, starts, factors []int) (int, []int) {
	min := math.MaxInt64
	for i, v := range starts {
		min = int(math.Min(float64(min), float64(numbers[v]*factors[i])))
	}
	for i, v := range starts {
		if min == numbers[v]*factors[i] {
			starts[i]++
		}
	}
	return min, starts
}

// Dynamic Programming
func dp(n int) int {
	factors := []int{2, 3, 5}
	numbers := make([]int, n)
	starts := []int{0, 0, 0}
	nextNum := 1
	numbers[0] = nextNum
	for i := 1; i < n; i++ {
		nextNum, starts = genUgly(numbers, starts, factors)
		numbers[i] = nextNum
	}
	return numbers[n-1]
}
