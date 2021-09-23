package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", subsetsWithDup([]int{1, 2, 2}))
	fmt.Printf("%v\n", subsetsWithDup([]int{1, 1, 2, 2}))
	fmt.Printf("%v\n", subsetsWithDup([]int{4, 4, 4, 1, 4}))
}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})

	sort.Sort(sort.IntSlice(nums))

	seen := make(map[string]bool)
	for i := len(nums) - 1; i >= 0; i-- {
		for _, v := range res {
			newArr := append([]int{nums[i]}, v...)
			asStr := arrAsString(newArr)
			if !seen[asStr] {
				res = append(res, newArr)
				seen[asStr] = true
			}
		}
	}
	return res
}

func arrAsString(arr []int) string {
	return fmt.Sprintf("%v", arr)
}
