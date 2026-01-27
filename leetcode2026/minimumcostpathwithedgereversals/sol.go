package minimumcostpathwithedgereversals

import "container/heap"

type Item struct {
	node int
	cost int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func minCost(n int, edges [][]int) int {
	adjacencyList := make(map[int][]Item)

	for _, edge := range edges {
		u, v, cost := edge[0], edge[1], edge[2]
		adjacencyList[u] = append(adjacencyList[u], Item{node: v, cost: cost})
		adjacencyList[v] = append(adjacencyList[v], Item{node: u, cost: 2 * cost})
	}

	costs := make([]int, n)
	for i := range costs {
		costs[i] = 1 << 31
	}
	costs[0] = 0

	visited_nodes := make([]bool, n)

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{node: 0, cost: 0})

	for pq.Len() > 0 {
		currentItem := heap.Pop(pq).(Item)
		currentNode := currentItem.node

		if visited_nodes[currentNode] {
			continue
		}
		visited_nodes[currentNode] = true

		for _, neighbor := range adjacencyList[currentNode] {
			newCost := costs[currentNode] + neighbor.cost
			if newCost < costs[neighbor.node] && !visited_nodes[neighbor.node] {
				costs[neighbor.node] = newCost
				heap.Push(pq, Item{node: neighbor.node, cost: newCost})
			}
		}
	}

	if costs[n-1] != 1<<31 {
		return costs[n-1]
	}

	return -1
}
