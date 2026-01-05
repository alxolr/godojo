package solution

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	minAbsVal := int(math.MaxInt)
	negativeCount := 0
	positiveSum := 0

	for _, row := range matrix {
		for _, val := range row {
			if val < 0 {
				negativeCount++
			}
			positiveSum += abs(val)
			minAbsVal = min(minAbsVal, abs(val))
		}
	}

	if negativeCount%2 != 0 {
		positiveSum -= 2 * minAbsVal
	}

	return int64(positiveSum)
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}
