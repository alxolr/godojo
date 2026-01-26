package minimumabsolutedifference1200

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	diffs := make(map[int][][]int)
	minDiff := math.MaxInt32

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]

		diffs[diff] = append(diffs[diff], []int{arr[i-1], arr[i]})
		if diff < minDiff {
			minDiff = diff
		}
	}

	return diffs[minDiff]
}
