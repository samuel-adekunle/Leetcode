package main

import (
	"fmt"
	"strconv"
)

func hammingDistance(x int, y int) int {
	return hammingWeight(booleanSum(x, y))
}

func booleanSum(x int, y int) int {
	bx, by := strconv.FormatInt(int64(x), 2), strconv.FormatInt(int64(y), 2)
	bx, by = applyEqualPadding(bx, by)
	bsum := ""
	for i := 0; i < len(bx); i++ {
		if bx[i] == by[i] {
			bsum = bsum + "0"
		} else {
			bsum = bsum + "1"
		}
	}
	sum, _ := strconv.ParseInt(bsum, 2, 64)
	return int(sum)
}

func applyEqualPadding(x, y string) (string, string) {
	for len(x) > len(y) {
		y = "0" + y
	}
	for len(y) > len(x) {
		x = "0" + x
	}
	return x, y
}

func hammingWeight(x int) (i int) {
	for x > 0 {
		i += x & 1
		x = x >> 1
	}
	return
}

func oneBitCount(x int) int {
	var i int
	for i = 0; x > 0; i++ {
		x &= (x - 1)
	}
	return i
}

func xorDistance(x int, y int) int {
	dist := x ^ y
	return oneBitCount(dist)
}

func main() {
	fmt.Println(hammingDistance(3, 1))
}
