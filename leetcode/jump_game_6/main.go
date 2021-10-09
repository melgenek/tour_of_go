package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", maxResult([]int{1, -1, -2, 4, -7, 3}, 2) == 7)
	fmt.Printf("%v\n", maxResult([]int{10, -5, -2, 4, 0, 3}, 3) == 17)
	fmt.Printf("%v\n", maxResult([]int{1, -5, -20, 4, -1, 3, -6, -3}, 2) == 0)
	fmt.Printf("%v\n", maxResult([]int{1, -1, -1, 1000}, 2) == 1000)

}

func maxResult(nums []int, k int) int {
	n := len(nums)
	state := make([]int, n)
	state[n-1] = nums[n-1]

	q := list.New()
	q.PushFront(state[n-1])

	for i := n - 2; i >= 0; i-- {
		nextMax := q.Back().Value.(int)
		state[i] = nextMax + nums[i]
		if n-1-i >= k && q.Back().Value.(int) == state[i+k] {
			q.Remove(q.Back())
		}
		for q.Len() > 0 && q.Front().Value.(int) < state[i] {
			q.Remove(q.Front())
		}
		q.PushFront(state[i])
	}

	return state[0]
}
