package main

import (
	"math"
)


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




func lengthOfLongestSubstring(s string) int {
	n := len(s)
	mem := map[byte]int{}
	acc := 0
	for i, j := 0, 0; j < n; j++ {
		if v, ok := mem[s[j]]; ok {
			
			i = int(math.Max(float64(i), float64(v))) 
		}
		acc = int(math.Max(float64(acc), float64(j-i+1)))
		mem[s[j]] = j + 1
	}
	return acc
}
