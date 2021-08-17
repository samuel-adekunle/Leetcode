package main

import "fmt"

func main() {
	fmt.Println(countComponents(5, [][]int{
		{0, 1},
		{1, 2},
		
		{3, 4},
	}))
}

func buildVertexMap(n int, edges [][]int) map[int][]int {
	vmap := map[int][]int{}
	for i := 0; i < n; i++ {
		vmap[i] = []int{}
	}

	for _, edge := range edges {
		vmap[edge[0]] = append(vmap[edge[0]], edge[1])
		vmap[edge[1]] = append(vmap[edge[1]], edge[0])
	}

	return vmap
}

func countComponents(n int, edges [][]int) int {
	vmap := buildVertexMap(n, edges)

	visited := map[int]bool{}
	for i := 0; i < n; i++ {
		visited[i] = false
	}

	var dfs func(int)
	dfs = func(vertex int) {
		if visited[vertex] {
			return
		} else {
			visited[vertex] = true
		}
		for _, v := range vmap[vertex] {
			dfs(v)
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if val := visited[i]; !val {
			count++
			dfs(i)
		}
	}


	return count
}
