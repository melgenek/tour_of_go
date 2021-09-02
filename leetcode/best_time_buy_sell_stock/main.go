package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Printf("%v\n", maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}))
}

const aLot = 99999999

func maxProfit(prices []int) int {
	cache := make(map[string]int)
	return findProfit(prices, 0, aLot, cache)
}

func findProfit(prices []int, day int, bidPrice int, cache map[string]int) int {
	key := fmt.Sprintf("%d_%d", day, bidPrice)
	cached, found := cache[key]
	if found {
		return cached
	} else if day == len(prices) {
		return 0
	} else {
		dayProfit := max(0, prices[day]-bidPrice)
		without := findProfit(prices, day+1, min(bidPrice, prices[day]), cache)
		result := without
		if dayProfit > 0 {
			with := dayProfit + findProfit(prices, day+1, aLot, cache)
			result = max(with, without)
		}
		cache[key] = result
		return result
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
