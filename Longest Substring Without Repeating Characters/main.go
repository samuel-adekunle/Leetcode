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
	"math"
)

// brute force
func bruteForce(s string) int {
	acc := math.MinInt32
	for i := 0; i < len(s); i++ {
		cnt := 1
		xs := map[byte]bool{s[i]: true}
		for j := i + 1; j < len(s); j++ {
			if ok := xs[s[j]]; !ok {
				cnt++
				xs[s[j]] = true
			}
		}
		acc = int(math.Max(float64(acc), float64(cnt)))
	}
	return acc
}

// abcabcbb
// remove char at i and move on (makes a lot of sense)
// e.g
// abc(hit a)
// bc
// bca
// bca (hit b)
// ca
// cab (no matter what we hit, we keep reducing and adding in linear time)
func slidingWindow(s string) int {
	n := len(s)
	mem := map[byte]bool{}
	acc, i, j := 0, 0, 0
	for i < n && j < n {
		if ok := mem[s[j]]; !ok {
			mem[s[j]] = true
			acc = int(math.Max(float64(acc), float64(j-i+1)))
			j++
		} else {
			mem[s[i]] = false
			i++
		}
	}
	return acc
}

// abcabcbb
// abc
//
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	mem := map[byte]int{}
	acc := 0
	for i, j := 0, 0; j < n; j++ {
		if v, ok := mem[s[j]]; ok {
			// we've encountered this char before (okay move i to the point in front point)
			i = int(math.Max(float64(i), float64(v))) // we don't want to move backwards
		}
		acc = int(math.Max(float64(acc), float64(j-i+1)))
		mem[s[j]] = j + 1
	}
	return acc
}
