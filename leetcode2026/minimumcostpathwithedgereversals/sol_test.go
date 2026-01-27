package minimumcostpathwithedgereversals

import "testing"

func Test_minCost(t *testing.T) {
	tests := []struct {
		edgesCount int
		edges      [][]int
		expected   int
	}{

		{4, [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}}, 5},
		{4, [][]int{{0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 3}}, 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minCost(tt.edgesCount, tt.edges); got != tt.expected {
				t.Errorf("minCost() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
