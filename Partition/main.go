package main

import (
	"fmt"
)

func isSubsetSum(arr []int, dp []map[int]bool, sum int, i int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || i >= len(arr) {
		return false
	}

	if val, ok := dp[i][sum]; ok {
		return val
	}

	val := isSubsetSum(arr, dp, sum-arr[i], i+1) || isSubsetSum(arr, dp, sum, i+1)
	dp[i][sum] = val
	return val
}

func findPartition(arr []int) bool {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	// dp[i][sum] - store solved subproblem
	dp := make([]map[int]bool, len(arr))
	for i := range arr {
		dp[i] = make(map[int]bool)
	}

	return isSubsetSum(arr, dp, sum, 0)
}

func main() {
	fmt.Println(findPartition([]int{1, 1, 4, 1, 2, 3}))
}
