package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", nextGreater([]int{1, 3, 4, 5, 6, 8, 11}, 0) == 1)
	fmt.Printf("%v\n", nextGreater([]int{1, 3, 4, 5, 6, 8, 11}, 1) == 3)
	fmt.Printf("%v\n", nextGreater([]int{1, 3, 4, 5, 6, 8, 11}, 3) == 4)
	fmt.Printf("%v\n", nextGreater([]int{1, 3, 4, 5, 6, 8, 11}, 4) == 5)
	fmt.Printf("%v\n", nextGreater([]int{1, 3, 4, 5, 6, 8, 11}, 5) == 6)
	fmt.Printf("%v\n", nextGreaterCustom([]int{1, 3, 4, 5, 6, 8, 11}, 6) == 8)
	fmt.Printf("%v\n", nextGreaterCustom([]int{1, 3, 4, 5, 6, 8, 11}, 7) == 8)
	fmt.Printf("%v\n", nextGreaterCustom([]int{1, 3, 4, 5, 6, 8, 11}, 8) == 11)
	fmt.Printf("%v\n", nextGreaterCustom([]int{1, 3, 4, 5, 6, 8, 11}, 10) == 11)
	fmt.Printf("%v\n", nextGreaterCustom([]int{1, 3, 4, 5, 6, 8, 11}, 12) == -1)
	fmt.Printf("%v\n", nextGreaterCustom([]int{1, 3, 4, 4, 4, 5, 6, 8, 11}, 4) == 5)
	fmt.Printf("~~~~~~~~\n")

	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 4, 5, 6, 8, 11}, 12) == 11)
	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 4, 5, 6, 8, 11}, 11) == 8)
	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 4, 5, 6, 8, 11}, 8) == 6)
	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 4, 5, 6, 8, 11}, 7) == 6)
	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 4, 5, 6, 8, 11}, 4) == 3)
	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 4, 5, 6, 8, 11}, 1) == -1)
	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 4, 5, 6, 8, 11}, 0) == -1)
	fmt.Printf("%v\n", previousSmaller([]int{1, 3, 3, 5, 6, 8, 11}, 3) == 1)
}

func previousSmaller(arr []int, target int) int {
	idx := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if idx-1 < 0 {
		return -1
	} else {
		return arr[idx-1]
	}
}

func nextGreater(arr []int, target int) int {
	idx := sort.Search(len(arr), func(i int) bool {
		return arr[i] > target
	})

	if idx != len(arr) {
		return arr[idx]
	} else {
		return -1
	}
}

func nextGreaterCustom(arr []int, target int) int {
	return rec(arr, target, -1)
}

func rec(arr []int, target int, next int) int {
	if len(arr) == 0 {
		return next
	} else {
		middleIdx := len(arr) / 2
		middle := arr[middleIdx]

		if middle > target {
			return rec(arr[:middleIdx], target, middle)
		} else {
			return rec(arr[middleIdx+1:], target, next)
		}
	}
}
