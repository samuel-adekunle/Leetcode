package main

import "fmt"



func main() {
	grid := [][]byte{
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Printf("Grid: %v\n", grid)
	fmt.Printf("Num Islands: %d\n", numIslands(grid))

}

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	acc := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				acc++
				helper(&grid, i, j)
			}
		}
	}
	return acc
}

func helper(grid *[][]byte, i, j int) {
	m, n := len(*grid), len((*grid)[0])
	if i < 0 || j < 0 || i >= m || j >= n || (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	helper(grid, i-1, j) 
	helper(grid, i+1, j) 
	helper(grid, i, j-1) 
	helper(grid, i, j+1)
}
