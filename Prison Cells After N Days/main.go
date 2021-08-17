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


func prisonAfterNDays(cells []int, N int) []int {
	for i := 0; i < N; i++ {
		cells = nextDay(cells)
	}
	return cells
}


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
