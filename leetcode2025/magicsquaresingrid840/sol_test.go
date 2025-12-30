package main

import "testing"

func TestMyFunc(t *testing.T) {
	grid := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
	expected := 1

	if numMagicSquaresInside(grid) != expected {
		t.Error("The test failed")
	}
}
