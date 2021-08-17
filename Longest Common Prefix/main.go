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
