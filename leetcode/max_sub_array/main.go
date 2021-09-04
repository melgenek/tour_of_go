package main

import "fmt"

func main() {
	fmt.Printf("%v\n", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	fmt.Printf("%v\n", maxSubArray([]int{-1}) == -1)
	fmt.Printf("%v\n", maxSubArray([]int{1}) == 1)
	fmt.Printf("%v\n", maxSubArray([]int{5, 4, -1, 7, 8}) == 23)
}

func maxSubArray(nums []int) int {
	maxSum := -999999

	currentSum := 0
	for i := 0; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
