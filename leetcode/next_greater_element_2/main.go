package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", nextGreaterElements([]int{1, 2, 1}))
	fmt.Printf("%v\n", nextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Printf("%v\n", nextGreaterElements([]int{4, 5, 3, 2, 1}))
}

func nextGreaterElements(nums []int) []int {
	stack := list.New()

	res := make([]int, len(nums))

	n := len(nums)

	for i := 2*n - 1; i >= 0; i-- {
		adjustedI := i % n
		current := nums[adjustedI]
		for stack.Len() > 0 && stack.Front().Value.(int) <= current {
			stack.Remove(stack.Front())
		}
		if stack.Len() > 0 {
			res[adjustedI] = stack.Front().Value.(int)
		} else {
			res[adjustedI] = -1
		}
		stack.PushFront(current)
	}

	return res
}
