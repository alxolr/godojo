package solution

import (
	"testing"
)

func TestMaxMatrixSum(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int64
	}{
		{
			name:     "2x2 all positive",
			matrix:   [][]int{{1, 2}, {3, 4}},
			expected: 10,
		},
		{
			name:     "2x2 with two negative",
			matrix:   [][]int{{1, -1}, {-1, 1}},
			expected: 4,
		},
		{
			name:     "2x2 with two negatives adjacent",
			matrix:   [][]int{{1, 2}, {-3, -4}},
			expected: 10,
		},
		{
			name:     "2x2 all negative",
			matrix:   [][]int{{-1, -2}, {-3, -4}},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxMatrixSum(tt.matrix)
			if result != tt.expected {
				t.Errorf("maxMatrixSum(%v) = %d; expected %d", tt.matrix, result, tt.expected)
			}
		})
	}
}
