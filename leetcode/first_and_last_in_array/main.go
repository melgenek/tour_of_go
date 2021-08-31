package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", searchRange([]int{1}, 1))
	fmt.Printf("%v\n", searchRange([]int{}, 1))
	fmt.Printf("%v\n", searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Printf("%v\n", searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func searchRange(nums []int, target int) []int {
	return []int{first(nums, target), last(nums, target)}
}

func first(nums []int, target int) int {
	from := 0
	to := len(nums)

	for from < to {
		middle := (from + to) / 2
		if nums[middle] == target && (middle == 0 || target > nums[middle-1]) {
			return middle
		} else if target > nums[middle] {
			from = middle + 1
		} else {
			to = middle
		}
	}

	return -1
}

func last(nums []int, target int) int {
	from := 0
	to := len(nums)

	for from < to {
		middle := (from + to) / 2
		if nums[middle] == target && (middle == len(nums)-1 || target < nums[middle+1]) {
			return middle
		} else if target < nums[middle] {
			to = middle
		} else {
			from = middle + 1
		}
	}

	return -1
}
