package main

import "fmt"

func main() {
	fmt.Printf("%v\n", subsets([]int{9, 0, 3, 5, 7}))
	fmt.Printf("%v\n", subsets2([]int{9, 0, 3, 5, 7}))
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for i := len(nums) - 1; i >= 0; i-- {
		for _, v := range res {
			res = append(res, append([]int{nums[i]}, v...))
		}
	}
	return res
}

func subsets2(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	} else {
		head := nums[0]
		sub := subsets(nums[1:])

		res := make([][]int, 0)
		for _, v := range sub {
			res = append(res, v)
			res = append(res, append([]int{head}, v...))
		}
		return res
	}
}
