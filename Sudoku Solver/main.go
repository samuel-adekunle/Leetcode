package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	for _, row := range board {
		for _, element := range row {
			fmt.Printf("%v ", byteToDigit(element))
		}
		fmt.Println()
	}
	fmt.Println()
	solveSudoku(board)
	for _, row := range board {
		for _, element := range row {
			fmt.Printf("%v ", byteToDigit(element))
		}
		fmt.Println()
	}
}

const ORDER = 9
const SECTION_SIZE = 3
const SECTION_ORDER = ORDER / SECTION_SIZE

func solveSudoku(board [][]byte) {
	rows, cols, sections := initCache(board)

	getMoves := func(row, col, section int) []int {
		moves := []int{}
		for move := 1; move <= ORDER; move++ {
			if !rows[row][move] && !cols[col][move] && !sections[section][move] {
				moves = append(moves, move)
			}
		}
		return moves
	}

	solved := false
	for !solved {
		solved = true
		for row := 0; row < ORDER; row++ {
			for col := 0; col < ORDER; col++ {
				element := board[row][col]
				if element == '.' {
					section := getSection(row, col)
					moves := getMoves(row, col, section)
					if len(moves) == 1 {
						move := moves[0]
						rows[row][move] = true
						cols[col][move] = true
						sections[section][move] = true
						board[row][col] = digitToByte(move)
					} else {
						solved = false
					}
				}
			}
		}
	}
}

func initCache(board [][]byte) ([ORDER]map[int]bool, [ORDER]map[int]bool, [ORDER]map[int]bool) {
	generateCache := func(size int) [ORDER]map[int]bool {
		cache := [ORDER]map[int]bool{}
		for i := 0; i < ORDER; i++ {
			cache[i] = make(map[int]bool)
		}
		return cache
	}
	rows, cols, sections := generateCache(ORDER), generateCache(ORDER), generateCache(ORDER)

	for row := 0; row < ORDER; row++ {
		for col := 0; col < ORDER; col++ {
			if element := board[row][col]; element != '.' {
				section := getSection(row, col)
				element := byteToDigit(element)
				rows[row][element] = true
				cols[col][element] = true
				sections[section][element] = true
			}
		}
	}

	return rows, cols, sections
}

func getSection(row, col int) int {
	return SECTION_ORDER*int(row/SECTION_ORDER) + int(col/SECTION_ORDER)
}

const BYTE_OFFSET = 48

func byteToDigit(b byte) int {
	return int(b) - BYTE_OFFSET
}

func digitToByte(d int) byte {
	return byte(d + BYTE_OFFSET)
}
