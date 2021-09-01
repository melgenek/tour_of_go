package main

import "fmt"

func main() {
	fmt.Printf("%v\n", minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Printf("%v\n", minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	arr := make([]int, n+1)
	arr[0] = cost[0]
	arr[1] = cost[1]

	for i := 2; i < n; i++ {
		one := arr[i-1]
		two := arr[i-2]

		min := one
		if one > two {
			min = two
		}
		arr[i] = min + cost[i]
	}

	min := arr[n-1]
	if arr[n-2] < min {
		min = arr[n-2]
	}

	return min
}
