package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Printf("%v\n", nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	stack := list.New()

	for i := len(nums2) - 1; i >= 0; i-- {
		current := nums2[i]
		for stack.Len() > 0 && stack.Front().Value.(int) <= current {
			stack.Remove(stack.Front())
		}
		if stack.Len() > 0 {
			cache[current] = stack.Front().Value.(int)
		}
		stack.PushFront(current)
	}

	res := make([]int, len(nums1))

	for i, v := range nums1 {
		cached, found := cache[v]
		if found {
			res[i] = cached
		} else {
			res[i] = -1
		}

	}

	return res
}
