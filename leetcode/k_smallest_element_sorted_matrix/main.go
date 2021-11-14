package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", kthSmallest([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 3) == 9)
	fmt.Printf("%v\n", kthSmallest([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 8) == 13)
	fmt.Printf("%v\n", kthSmallest([][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}, 1) == 1)
	fmt.Printf("%v\n", kthSmallest([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 9) == 15)
	fmt.Printf("%v\n", kthSmallest([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 2) == 5)
	fmt.Printf("%v\n", kthSmallest([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 6) == 12)
}

type Cell struct {
	v int
	x int
	y int
}

type PQ []Cell

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this PQ) Less(i, j int) bool {
	return this[i].v < this[j].v
}

func (this *PQ) Push(v interface{}) {
	*this = append(*this, v.(Cell))
}

func (this *PQ) Pop() interface{} {
	last := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return last
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	pq := &PQ{}

	for j := 0; j < n; j++ {
		heap.Push(pq, Cell{matrix[0][j], 0, j})
	}

	i := 0
	for ; i < k-1; i++ {
		cell := heap.Pop(pq).(Cell)
		if cell.x != n-1 {
			heap.Push(pq, Cell{matrix[cell.x+1][cell.y], cell.x + 1, cell.y})
		}
	}

	return heap.Pop(pq).(Cell).v
}
