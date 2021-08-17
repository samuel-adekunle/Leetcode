/* Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
			 Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package main

import (
	"fmt"
	"testing"
)

func Example() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("abcadadvcdsbcbb"))
	//Output:
	// 3
	// 5
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcadadvcdsbcbb")
	}
}
