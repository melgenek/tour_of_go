package main

import (
	"container/heap"
)

func main() {

}

type IntHeap []int

func (this IntHeap) Len() int           { return len(this) }
func (this IntHeap) Less(i, j int) bool { return this[i] > this[j] }
func (this IntHeap) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

func (this *IntHeap) Push(x interface{}) { *this = append(*this, x.(int)) }
func (this *IntHeap) Pop() interface{} {
	res := (*this)[this.Len()-1]
	*this = (*this)[0 : this.Len()-1]
	return res
}

func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)

	for i := 0; i < k-1; i++ {
		heap.Pop(&h)
	}

	return heap.Pop(&h).(int)
}
