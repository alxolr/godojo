package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{4, 3, 8},
		{9, 5, 1},
		{2, 7, 6},
	}
	result := numMagicSquaresInside(grid)
	fmt.Println(result)
}
