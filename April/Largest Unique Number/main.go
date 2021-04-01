package main

import "fmt"

func largestUniqueNumber(A []int) (max int) {
	count := map[int]int{}
	for _, v := range A {
		count[v]++
	}
	max = -1
	for k, v := range count {
		if v == 1 && k > max {
			max = k
		}
	}
	return
}

func main() {
	arr := []int{5, 7, 3, 9, 4, 9, 8, 3, 1}
	fmt.Println(largestUniqueNumber(arr))
}
