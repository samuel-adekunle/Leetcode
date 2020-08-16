/* Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z. */

package main

import (
	"fmt"
	"testing"
)

func Example() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{}))
	//Output:
	// fl
	//
	//
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestCommonPrefix([]string{"flower", "flow", "flight"})
	}
}
