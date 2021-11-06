package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findMin([]int{1, 1}) == 1)
	fmt.Printf("%v\n", findMin([]int{1, 3, 3, 3}) == 1)
	fmt.Printf("%v\n", findMin([]int{1, 1, 3}) == 1)
	fmt.Printf("%v\n", findMin([]int{3, 3, 1}) == 1)
	fmt.Printf("%v\n", findMin([]int{3, 1, 3, 3}) == 1)
	fmt.Printf("%v\n", findMin([]int{3, 1, 3, 3, 3, 3, 3, 3}) == 1)
	fmt.Printf("%v\n", findMin([]int{3, 1, 3}) == 1)
	fmt.Printf("%v\n", findMin([]int{1, 3, 5}) == 1)
	fmt.Printf("%v\n", findMin([]int{1, 1, 1}) == 1)
	fmt.Printf("%v\n", findMin([]int{2, 2, 2, 0, 1}) == 0)
	fmt.Printf("%v\n", findMin([]int{2, 2, 2, 0, 2}) == 0)
	fmt.Printf("%v\n", findMin([]int{2, 2, 2, 2, 0, 2}) == 0)
	fmt.Printf("%v\n", findMin([]int{2, 2, 2, 2, 2, 0, 2}) == 0)
	fmt.Printf("%v\n", findMin([]int{4, 5, 6, 7, 0, 1, 4}) == 0)
	fmt.Printf("%v\n", findMin([]int{0, 1, 4, 4, 5, 6, 7}) == 0)
}

func findMin(nums []int) int {
	middle := len(nums) / 2

	if middle == 0 || nums[middle] < nums[middle-1] {
		return nums[middle]
	} else if nums[middle] < nums[len(nums)-1] {
		return findMin(nums[:middle])
	} else if nums[middle] > nums[len(nums)-1] {
		return findMin(nums[middle+1:])
	} else {
		return findMin(nums[:len(nums)-1])
	}
}
