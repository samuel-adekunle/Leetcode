package main

import (
	"fmt"
)



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
