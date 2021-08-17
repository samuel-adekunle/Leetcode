/*
There are 8 prison cells in a row, and each cell is either occupied or vacant.

Each day, whether the cell is occupied or vacant changes according to the following rules:

If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell becomes occupied.
Otherwise, it becomes vacant.
(Note that because the prison is a row, the first and the last cells in the row can't have two adjacent neighbors.)

We describe the current state of the prison in the following way: cells[i] == 1 if the i-th cell is occupied, else cells[i] == 0.

Given the initial state of the prison, return the state of the prison after N days (and N such changes described above.)



Example 1:

Input: cells = [0,1,0,1,1,0,0,1], N = 7
Output: [0,0,1,1,0,0,0,0]
Explanation:
The following table summarizes the state of the prison on each day:
Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

Example 2:

Input: cells = [1,0,0,1,0,0,1,0], N = 1000000000
Output: [0,0,1,1,1,1,1,0]
*/

package main

func nextDay(cells []int) []int {
	newCells := make([]int, len(cells))
	copy(newCells, cells)
	for j := 1; j < len(cells)-1; j++ {
		if cells[j-1] == cells[j+1] {
			newCells[j] = 1
		} else {
			newCells[j] = 0
		}
	}
	newCells[0], newCells[len(cells)-1] = 0, 0
	return newCells
}

// Brute force
func prisonAfterNDays(cells []int, N int) []int {
	for i := 0; i < N; i++ {
		cells = nextDay(cells)
	}
	return cells
}

// find the pattern
func useMap(cells []int, N int) []int {
	order := make(map[int][]int)
	for i := 0; i < N; i++ {
		if i == 0 {
			cells = nextDay(cells)
			continue
		}
		if i == 15 {
			return order[(N-1)%14]
		}
		order[i-1] = cells
		cells = nextDay(cells)
	}
	return cells
}
