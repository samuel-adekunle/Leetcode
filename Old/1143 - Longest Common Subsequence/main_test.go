/* Given two strings text1 and text2, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.



If there is no common subsequence, return 0.



Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.
Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.
Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0. */

package main

import (
	"fmt"
	"testing"
)

func Example() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("bsbininm", "jmjkbkjkv"))
	fmt.Println(longestCommonSubsequence("bsbin", "jkbkj"))
	fmt.Println(longestCommonSubsequence("aspsd", "papkj"))

	//Output:
	// 3
	// 1
	// 1
	// 2
}

func BenchmarkDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonSubsequence("bsbininm", "jmjkbkjkv")
	}
}

func BenchmarkOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keepTwoRows("bsbininm", "jmjkbkjkv")
	}
}
