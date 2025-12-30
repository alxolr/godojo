package main

import (
	"slices"
)

func numMagicSquaresInside(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	if rows < 3 || cols < 3 {
		// this is invalid matrix
		return 0
	}

	magicSquareCount := 0

	for row := range rows - 2 {
		for col := range cols - 2 {
			if isMagicSquare(row, col, grid) {
				magicSquareCount++
			}
		}
	}

	return magicSquareCount
}

func isMagicSquare(rowstart int, colstart int, grid [][]int) bool {
	rows := len(grid)
	cols := len(grid[0])

	rowend := rowstart + 2
	colend := colstart + 2

	if rowend >= rows || colend >= cols {
		return false
	}

	// extract the numbers we should have all values from 1 to 9
	allnums := make([]bool, 9)

	for row := rowstart; row <= rowend; row++ {
		for col := colstart; col <= colend; col++ {
			val := grid[row][col]
			if val < 1 || val > 9 {
				return false
			}
			allnums[val-1] = true
		}
	}

	// check if we have all nums from 1 to 9
	if slices.Contains(allnums, false) {
		return false
	}

	// get rows
	rowlines := [][]int{}
	for row := rowstart; row <= rowend; row++ {
		rowline := []int{}
		for col := colstart; col <= colend; col++ {
			rowline = append(rowline, grid[row][col])
		}

		rowlines = append(rowlines, rowline)
	}

	// get cols
	collines := [][]int{}
	for col := colstart; col <= colend; col++ {
		colline := []int{}
		for row := rowstart; row <= rowend; row++ {
			colline = append(colline, grid[row][col])
		}
		collines = append(collines, colline)
	}

	// get primary and anti diagonals
	primary := []int{}
	anti := []int{}
	for row := rowstart; row <= rowend; row++ {
		primary = append(primary, grid[row][colstart+(row-rowstart)])
		anti = append(anti, grid[row][colend-(row-rowstart)])
	}

	lines := append(rowlines, collines...)
	lines = append(lines, primary, anti)

	return checkAllLinesFifteen(lines)
}

func checkAllLinesFifteen(lines [][]int) bool {
	for _, line := range lines {
		if sum(line) != 15 {
			return false
		}
	}

	return true
}

func sum(line []int) int {
	sum := 0
	for _, num := range line {
		sum += num
	}

	return sum
}
