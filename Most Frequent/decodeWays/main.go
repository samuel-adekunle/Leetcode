package main

import (
	"fmt"
)

/*
A message containing letters from A-Z can be encoded into numbers using the following mapping:

	'A' -> "1"
	'B' -> "2"
	...
	'Z' -> "26"

To decode an encoded message, all the digits must be mapped back into letters using the reverse of the mapping above (there may be multiple ways). For example, "111" can have each of its "1"s be mapped into 'A's to make "AAA", or it could be mapped to "11" and "1" ('K' and 'A' respectively) to make "KA". Note that "06" cannot be mapped into 'F' since "6" is different from "06".

Given a string s containing only digits, return the number of ways to decode it.

The answer is guaranteed to fit in a 32-bit integer.

*/

func main() {
	s := "1111111111111111111111"
	fmt.Printf("Number of decoding for %s: %d\n", s, numDecodings(s))
}

func numDecodings(s string) int {
	mem := make(map[string]int, 0)
	return helper(&mem, s)
}

func helper(mem *map[string]int, rest string) int {
	n := len(rest)
	if n == 0 {
		return 0
	}
	if v, ok := (*mem)[rest]; ok {
		return v
	}

	ways := 0

	if validDecoding(string(rest[0])) {
		if n == 1 {
			ways++
		} else {
			ways += helper(mem, rest[1:])
		}
	}
	if n > 1 && validDecoding(rest[0:2]) {
		if n == 2 {
			ways++
		} else {
			ways += helper(mem, rest[2:])
		}
	}

	(*mem)[rest] = ways
	return ways
}

func validDecoding(s string) bool {
	n := len(s)

	if n == 1 && s != "0" {
		return true
	}

	if n == 2 && (s[0] == '1') || (s[0] == '2' && s[1] <= '6') {
		return true
	}

	return false
}
