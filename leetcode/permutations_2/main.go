package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", permute([]int{1, 2, 3}))
	fmt.Printf("%v\n", permute([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", permuteUnique([]int{1, 2, 3}))
	fmt.Printf("%v\n", permuteUnique([]int{1, 1, 2}))
	fmt.Printf("%v\n", permute([]int{1, 1, 2}))
	fmt.Printf("%v\n", permuteUnique([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", permuteUnique([]int{1, -1, 1, 2, -1, 2, 2, -1}))
}

func permute(nums []int) [][]int {
	var m func(map[int]bool, []int)

	res := make([][]int, 0)

	m = func(used map[int]bool, selection []int) {
		if len(selection) == len(nums) {
			res = append(res, selection)
		} else {
			for _, n := range nums {
				if used[n] {
					continue
				}
				used[n] = true
				m(used, append([]int{n}, selection...))
				delete(used, n)
			}
		}
	}

	m(make(map[int]bool), make([]int, 0))

	return res
}

func permuteUnique(input []int) [][]int {
	counts := make(map[int]int)
	for _, v := range input {
		counts[v]++
	}

	var m func([]int)

	res := make([][]int, 0)

	m = func(selection []int) {
		if len(selection) == len(input) {
			res = append(res, selection)
		} else {
			for k, v := range counts {
				if counts[k] == 0 {
					continue
				}
				counts[k] = v - 1
				m(append([]int{k}, selection...))
				counts[k] = v
			}
		}
	}

	m([]int{})

	return res
}
