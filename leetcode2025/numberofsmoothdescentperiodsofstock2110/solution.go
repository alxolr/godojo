package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	var countDays int
	pricesLen := len(prices)

	left := 0

	for left < pricesLen {
		right := left

		for right+1 < pricesLen && prices[right+1] == prices[right]-1 {
			right++
		}

		countDays += combos((right - left) + 1)
		left = right + 1

	}

	return int64(countDays)
}

func combos(n int) int {
	return n * (n + 1) / 2
}

func main() {
	prices := []int{4, 3, 2, 1}
	result := getDescentPeriods(prices)
	fmt.Println("Hello from my module", result)
}
