package main

import "fmt"

func main() {
	fmt.Printf("%v\n", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Printf("%v\n", maxSubArray([]int{-1}))
}

func maxSubArray(nums []int) int {
	_, _, v, arr := max(nums)
	fmt.Printf("%v\n", arr)
	return v
}

func max(nums []int) (int, []int, int, []int) {
	if len(nums) == 0 {
		return -99999999, []int{}, -99999999, []int{}
	} else {
		first := nums[0]
		currentMax, currentArr, max, arr := max(nums[1:])

		if first >= currentMax && currentMax <= 0 {
			currentMax = first
			currentArr = []int{first}
		} else {
			currentMax += first
			currentArr = append([]int{first}, currentArr...)
		}

		if currentMax > max {
			return currentMax, currentArr, currentMax, currentArr
		} else {
			return currentMax, currentArr, max, arr
		}
	}
}
