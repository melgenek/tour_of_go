package main

import "fmt"

func main() {
	fmt.Printf("%v\n", rob([]int{1}))
	fmt.Printf("%v\n", rob([]int{1, 1, 1, 2}))
	fmt.Printf("%v\n", rob([]int{1, 2, 1, 1}))
	fmt.Printf("%v\n", rob([]int{2, 3, 2}))
	fmt.Printf("%v\n", rob([]int{1, 2, 3, 1}))
	fmt.Printf("%v\n", rob([]int{1, 2, 3}))
	fmt.Printf("%v\n", rob([]int{1, 2, 3, 4, 5}))
}

func rob(cost []int) int {
	n := len(cost)
	if n == 1 {
		return cost[0]
	} else if n == 2 {
		return max(cost[0], cost[1])
	} else {
		return max(robSeq(cost[1:]), robSeq(cost[:len(cost)-1]))
	}
}

func robSeq(cost []int) int {
	n := len(cost)

	if n == 0 {
		return 0
	} else if n == 1 {
		return cost[0]
	} else if n == 2 {
		return max(cost[0], cost[1])
	} else {
		prePrevious := cost[0]
		previous := max(prePrevious, cost[1])

		for i := 2; i < n; i++ {
			prePrevious, previous = previous, max(previous, prePrevious+cost[i])
		}

		return previous
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
