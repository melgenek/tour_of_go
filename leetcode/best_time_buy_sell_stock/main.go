package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}) == 0)
	fmt.Printf("%v\n", maxProfit([]int{1, 2, 3, 4, 5}) == 4)
	fmt.Printf("%v\n", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}) == 4)
}

const aLot = 99999999

func maxProfit(prices []int) int {
	res := 0

	minPrice := aLot
	for i := 0; i < len(prices); i++ {
		res = max(res, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return res
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
