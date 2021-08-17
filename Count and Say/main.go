package main

import "fmt"

func say(seq string) string {
	res, c, n := "", seq[0], 1
	for _, v := range seq[1:] {
		if c == byte(v) {
			n++
			continue
		}
		res += fmt.Sprintf("%d%c", n, c)
		c, n = byte(v), 1
	}
	res += fmt.Sprintf("%d%c", n, c)
	return res
}

func countAndSay(n int) string {
	count := "1"
	for i := 1; i < n; i++ {
		count = say(count)
	}
	return count
}
