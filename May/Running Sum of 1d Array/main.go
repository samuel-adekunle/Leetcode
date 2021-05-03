package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{3,1,2,10,1}))
}

func runningSum(nums []int) []int {
	arr, sum := []int{}, 0
	for _, v := range nums {
		sum += v
		arr = append(arr, sum)
	}
	return arr
}