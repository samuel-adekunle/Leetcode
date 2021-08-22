package main

import "fmt"

func main() {
	fmt.Println(validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}))
	fmt.Println(validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}))
}

func validTree(n int, edges [][]int) bool {
	parents := []map[int]bool{}
	for i := 0; i < n; i++ {
		parents = append(parents, map[int]bool{})
	}
	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		parents[dest][src] = true
		if len(parents[dest]) > 1 {
			return false
		}
	}
	return true
}
