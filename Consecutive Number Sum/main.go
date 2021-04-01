package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func consecutiveNumbersSum(N int) int {
	ways := 1
	var acc int
	for i := 1; i < (N+1)/2; i++ {
		acc = 0
		for j := i; j <= (N+1)/2; j++ {
			acc += j
			if acc == N {
				ways += 1
			}
			if acc >= N {
				break
			}
		}
	}
	return ways
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()
	arg, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(consecutiveNumbersSum(int(arg)))
}
