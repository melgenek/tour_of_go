package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1}, 1))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, -1}, 1))
	fmt.Printf("%v\n", maxSlidingWindow([]int{9, 11}, 2))
	fmt.Printf("%v\n", maxSlidingWindow([]int{4, -2}, 2))
	fmt.Printf("%v\n", maxSlidingWindow([]int{2, 2, 1}, 2))
}

type MonotonicQueue struct {
	q *list.List
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		list.New(),
	}
}

func (q *MonotonicQueue) max() int {
	return q.q.Front().Value.(int)
}

func (q *MonotonicQueue) add(a int) {
	for q.q.Len() > 0 && q.q.Back().Value.(int) < a {
		q.q.Remove(q.q.Back())
	}

	q.q.PushBack(a)
}

func (q *MonotonicQueue) remove(a int) {
	if q.q.Len() > 0 && q.q.Front().Value.(int) == a {
		q.q.Remove(q.q.Front())
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)

	q := NewMonotonicQueue()

	for i := 0; i < k-1; i++ {
		q.add(nums[i])
	}

	for i := k - 1; i < n; i++ {
		q.add(nums[i])
		resI := i - k + 1
		res[resI] = q.max()
		q.remove(nums[resI])
	}

	return res
}
