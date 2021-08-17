package main

import "fmt"




func isIdealPermutation(A []int) bool {
	local, global := 0, 0
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			local++
		}
		for j := i+1; j < len(A); j++ {
			if A[i] > A[j] {
				global++
			}
		}
	}
	fmt.Println(local, global)
	return local == global
}

func main() {
	A := []int{1,2,0}
	fmt.Println(isIdealPermutation(A))
}
