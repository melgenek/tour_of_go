package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10))
	fmt.Printf("%v\n", kSmallestPairs([]int{1, 70, 71}, []int{1, 2}, 4))
	fmt.Printf("%v\n", kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
	fmt.Printf("%v\n", kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2))
	fmt.Printf("%v\n", kSmallestPairs([]int{1, 2}, []int{3}, 2))
	fmt.Printf("%v\n", kSmallestPairs([]int{1, 2}, []int{3}, 3))
}

type Pair struct {
	v1 int
	v2 int
	i2 int
}

type PQ []Pair

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this PQ) Less(i, j int) bool {
	return this[i].v1+this[i].v2 < this[j].v1+this[j].v2
}

func (this *PQ) Push(v interface{}) {
	*this = append(*this, v.(Pair))
}

func (this *PQ) Pop() interface{} {
	last := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return last
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1 := len(nums1)
	n2 := len(nums2)

	if k > n1*n2 {
		k = n1 * n2
	}

	pq := &PQ{}
	for i := 0; i < n1; i++ {
		heap.Push(pq, Pair{nums1[i], nums2[0], 0})
	}

	result := make([][]int, k)

	for i := 0; i < k; i++ {
		pair := heap.Pop(pq).(Pair)
		result[i] = []int{pair.v1, pair.v2}

		if pair.i2 != n2-1 {
			heap.Push(pq, Pair{pair.v1, nums2[pair.i2+1], pair.i2 + 1})
		}
	}

	return result
}
