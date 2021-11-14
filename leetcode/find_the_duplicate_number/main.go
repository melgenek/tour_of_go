package main

import "fmt"

func main() {
	fmt.Printf("%v\n", 2 == findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Printf("%v\n", 3 == findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Printf("%v\n", 1 == findDuplicate([]int{1, 1}))
}

func findDuplicate(nums []int) int {
	for _, v := range nums {
		if v < 0 {
			v = -v
		}

		if nums[v] < 0 {
			return v
		}
		nums[v] = -nums[v]
	}
	return 0
}
