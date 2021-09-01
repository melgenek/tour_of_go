package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findMin([]int{1, 1}))
	fmt.Printf("%v\n", findMin([]int{3, 4, 5, 1, 2}))
	fmt.Printf("%v\n", findMin([]int{4, 5, 6, 7, 0, 1, 4}))
	fmt.Printf("%v\n", findMin([]int{0, 1, 4, 4, 5, 6, 7}))
	fmt.Printf("%v\n", findMin([]int{10, 10, 10, 10, 10, 1, 10}))
}

func findMin(nums []int) int {
	if len(nums) < 3 {
		min := nums[0]
		for _, v := range nums {
			if min > v {
				min = v
			}
		}
		return min
	} else {
		middle := len(nums) / 2

		if middle == 0 || nums[middle] < nums[middle-1] {
			return nums[middle]
		} else if nums[middle] == nums[0] && nums[middle] == nums[len(nums)-1] {
			l := findMin(nums[:middle])
			r := findMin(nums[middle+1:])

			if l > r {
				l = r
			}
			return l
		}
		if nums[middle] <= nums[len(nums)-1] {
			return findMin(nums[:middle])
		} else {
			return findMin(nums[middle+1:])
		}
	}
}
