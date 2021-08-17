package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDistance("sea", "ate"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i%2][j] = dp[(i-1)%2][j-1] + 1
			} else {
				dp[i%2][j] = int(math.Max(float64(dp[(i-1)%2][j]), float64(dp[i%2][j-1])))
			}
		}
	}

	return dp[n%2][m]
}

func minDistance(word1 string, word2 string) (distance int) {
	return len(word1) + len(word2) - 2*longestCommonSubsequence(word1, word2)
}
