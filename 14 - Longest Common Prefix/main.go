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

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longest := strs[0]
	for _, v := range strs[1:] {
		next := ""
		for i := 0; i < len(longest); i++ {
			if i >= len(v) {
				break
			}
			if longest[i] == v[i] {
				next += string(longest[i])
			} else {
				break
			}
		}
		longest = next
		if len(longest) == 0 {
			return longest
		}
	}
	return longest
}
