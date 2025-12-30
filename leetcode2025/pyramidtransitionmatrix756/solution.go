package solution

func backtrack(idx int, currentPath []rune, solutionPool [][]rune, results *[]string) {
	if len(currentPath) == len(solutionPool) {
		*results = append(*results, string(currentPath))
		return
	}

	if idx < len(solutionPool) {
		for _, ch := range solutionPool[idx] {
			currentPath = append(currentPath, ch)
			backtrack(idx+1, currentPath, solutionPool, results)
			currentPath = currentPath[:len(currentPath)-1]
		}
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func pyramidTransition(bottom string, allowed []string) bool {
	type QueueItem struct {
		base  string
		level int
	}

	queue := []QueueItem{{base: bottom, level: len(bottom) - 1}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		base := current.base
		level := current.level

		// If level is 1, check if there's one allowed exit
		if level == 1 {
			for top := 'A'; top <= 'F'; top++ {
				pyramid := base + string(top)
				if contains(allowed, pyramid) {
					return true
				}
			}
		}

		baseAsRunes := []rune(base)
		var combos [][]rune

		// Find all valid top characters for each adjacent pair
		for i := 0; i < len(baseAsRunes)-1; i++ {
			bottom := string(baseAsRunes[i : i+2])
			var validTops []rune

			for top := 'A'; top <= 'F'; top++ {
				pyramid := bottom + string(top)
				if contains(allowed, pyramid) {
					validTops = append(validTops, top)
				}
			}
			combos = append(combos, validTops)
		}

		// Generate all possible bases using backtracking
		var newBases []string
		backtrack(0, []rune{}, combos, &newBases)

		for _, newBase := range newBases {
			if !visited[newBase] {
				queue = append(queue, QueueItem{base: newBase, level: level - 1})
				visited[newBase] = true
			}
		}
	}

	return false
}
