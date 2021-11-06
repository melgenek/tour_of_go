package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}) == 7)
	fmt.Printf("%v\n", maxProfit([]int{1, 2, 3, 4, 5}) == 4)
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}) == 0)
	fmt.Printf("%v\n", maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}) == 8)
}

func maxProfit(prices []int) int {
	n := len(prices)
	state := make([][]int, n)
	state[n-1] = []int{prices[n-1], 0}

	for i := n - 2; i >= 0; i-- {
		state[i] = []int{
			max(prices[i], state[i+1][0]),
			max(state[i+1][1], state[i+1][0]-prices[i]),
		}
	}

	return state[0][1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
