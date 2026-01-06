package solution

import (
	"testing"
)

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name: "Example 1: [1,7,0,7,-8,null,null]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: -8},
				},
				Right: &TreeNode{Val: 0},
			},
			expected: 2,
		},
		{
			name: "Example 2: [989,null,10250,98693,-89388,null,null,null,-32127]",
			root: &TreeNode{
				Val:  989,
				Left: nil,
				Right: &TreeNode{
					Val: 10250,
					Left: &TreeNode{
						Val:   98693,
						Right: &TreeNode{Val: -32127},
					},
					Right: &TreeNode{Val: -89388},
				},
			},
			expected: 2,
		},
		{
			name:     "Single node",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name: "Two levels with negative values",
			root: &TreeNode{
				Val:   -10,
				Left:  &TreeNode{Val: -5},
				Right: &TreeNode{Val: -3},
			},
			expected: 2,
		},
		{
			name: "All same values",
			root: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 5},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxLevelSum(tt.root)
			if result != tt.expected {
				t.Errorf("maxLevelSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}
