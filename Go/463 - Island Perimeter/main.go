/* You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.



Example:

Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:
*/

// Note - decided to go DFS route, but simpler mathematical solution exists

package main

var globalGrid [][]int

var dirs [][]int = [][]int{
	{0, -1}, // left
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
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
