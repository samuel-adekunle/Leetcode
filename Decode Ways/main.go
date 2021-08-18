package main

import "fmt"

func main() {
	fmt.Println(numDecodings("11106"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("06"))
}

// TLE
func numDecodings(s string) int {
	numOfDecoding := 0
	var helper func(int, string)
	helper = func(i int, sub string) {
		if !isValidCode(sub) {
			return
		}
		if i >= len(s) {
			numOfDecoding += 1
			return
		}
		nextOne := s[i : i+1]
		helper(i+1, nextOne)
		if i+2 <= len(s) {
			nextTwo := s[i : i+2]
			helper(i+2, nextTwo)
		}
	}
	helper(0, "")
	return numOfDecoding
}

func isValidCode(s string) bool {
	if len(s) == 0 {
		return true
	} else if len(s) == 1 && s != "0" {
		return true
	} else if len(s) == 2 && ((s[0] == '1') || (s[0] == '2' && s[1] != '7' && s[1] != '8' && s[1] != '9')) {
		return true
	}
	return false
}
