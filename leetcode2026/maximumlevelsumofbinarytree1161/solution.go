package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Item struct {
	Level int
	Node  *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	sums := make(map[int]int)

	queue := []Item{{Level: 1, Node: root}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		sums[item.Level] += item.Node.Val

		if item.Node.Left != nil {
			queue = append(queue, Item{Level: item.Level + 1, Node: item.Node.Left})
		}

		if item.Node.Right != nil {
			queue = append(queue, Item{Level: item.Level + 1, Node: item.Node.Right})
		}
	}

	maxSum := math.MinInt
	maxPos := math.MaxInt

	for key, val := range sums {
		if val > maxSum || (val == maxSum && key < maxPos) {
			maxSum = val
			maxPos = key
		}
	}

	return maxPos
}
