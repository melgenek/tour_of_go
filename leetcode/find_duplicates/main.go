package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Printf("%v\n", findDuplicates([]int{1, 1, 2}))
	fmt.Printf("%v\n", findDuplicates([]int{1, 2, 1}))
	fmt.Printf("%v\n", findDuplicates([]int{2, 1, 2}))
	fmt.Printf("%v\n", findDuplicates([]int{1}))
	fmt.Printf("%v\n", findDuplicates([]int{1, 1}))
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); {
		v := nums[i]

		if v > 0 {
			if nums[v-1] < 0 {
				res = append(res, v)
				nums[i] = 0
				i++
			} else {
				nums[i], nums[v-1] = nums[v-1], nums[i]
				nums[v-1] = -1
			}
		} else {
			i++
		}
		//fmt.Printf("%v> %v\n", i, nums)
	}

	return res
}
