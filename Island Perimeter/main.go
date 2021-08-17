package main

var globalGrid [][]int

var dirs [][]int = [][]int{
	{0, -1}, 
	{-1, 0}, 
	{0, 1},  
	{1, 0}, 
}

func helper(x, y int) int {
	if x < 0 || y < 0 || x > len(globalGrid)-1 || y > len(globalGrid[0])-1 || globalGrid[x][y] == 0 {
		return 1
	}
	if globalGrid[x][y] == -1 {
		return 0
	}
	per := 0
	globalGrid[x][y] = -1
	for _, dir := range dirs {
		per += helper(x+dir[0], y+dir[1])
	}
	return per
}

func islandPerimeter(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < (len(grid[0])); j++ {
			if grid[i][j] == 1 {
				globalGrid = grid
				return helper(i, j)
			}
		}
	}
	return 0
}
