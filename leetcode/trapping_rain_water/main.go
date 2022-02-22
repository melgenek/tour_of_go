package main

import "fmt"

func main() {
	fmt.Printf("%v\n", 9 == trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Printf("%v\n", 6 == trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	result := 0
	n := len(height)

	maxFromRight := make([]int, n)
	maxFromRight[n-1] = height[n-1]

	for i := n - 2; i >= 0; i-- {
		maxFromRight[i] = max(height[i], maxFromRight[i+1])
	}

	maxFromLeft := make([]int, n)
	maxFromLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxFromLeft[i] = max(height[i], maxFromLeft[i-1])
	}

	for i := 0; i < n; i++ {
		boundary := min(maxFromLeft[i], maxFromRight[i])
		amount := boundary - height[i]
		result += max(amount, 0)
	}

	return result
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
