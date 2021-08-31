package main

import "fmt"

func main() {
	fmt.Printf("%v\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 5) == 1)
	fmt.Printf("%v\n", search([]int{5, 6, 7, 0, 1, 2, 4}, 1) == 4)
	fmt.Printf("%v\n", search([]int{5, 6, 7, 0, 1, 2, 4}, 5) == 0)
	fmt.Printf("%v\n", search([]int{5, 6, 7, 0, 1, 2, 4}, 4) == 6)
	fmt.Printf("%v\n", search([]int{0, 1, 2, 4, 5, 6, 7}, 3) == -1)
	fmt.Printf("%v\n", search([]int{1}, 0) == -1)
	fmt.Printf("%v\n", search([]int{4, 5, 6, 7, 0, 1, 2}, 0) == 4)
	fmt.Printf("%v\n", search([]int{3, 4, 5, 6, 1, 2}, 2) == 5)
}

func search(nums []int, target int) int {
	return searchRecursive(nums, target, 0, len(nums)-1)
}

func searchRecursive(nums []int, target int, from int, to int) int {
	if to < from {
		return -1
	}
	middle := (to + from) / 2
	if nums[middle] == target {
		return middle
	} else {
		right := false
		res := -1
		if target < nums[middle] {
			res = searchRecursive(nums, target, from, middle-1)
		} else {
			res = searchRecursive(nums, target, middle+1, to)
			right = true
		}
		if res == -1 {
			if right {
				res = searchRecursive(nums, target, from, middle-1)
			} else {
				res = searchRecursive(nums, target, middle+1, to)
			}

		}
		return res
	}
}
