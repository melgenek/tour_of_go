package main

import (
	"container/heap"
	"fmt"
)

func main() {
	room := Constructor(8)

	fmt.Printf("%v\n", room.Seat() == 0)
	fmt.Printf("%v\n", room.Seat() == 7)
	fmt.Printf("%v\n", room.Seat() == 3)
	room.Leave(0)
	room.Leave(7)
	fmt.Printf("%v\n", room.Seat() == 7)
	fmt.Printf("%v\n", room.Seat() == 0)
	fmt.Printf("%v\n", room.Seat() == 5)
	fmt.Printf("%v\n", room.Seat() == 1)
	fmt.Printf("%v\n", room.Seat() == 2)
	fmt.Printf("%v\n", room.Seat() == 4)
	fmt.Printf("%v\n", room.Seat() == 6)
}

type PQ []*Range

func (this PQ) Len() int {
	return len(this)
}

func (this PQ) Less(i, j int) bool {
	return this[i].priority > this[j].priority || this[i].priority == this[j].priority && this[i].start < this[j].start
}

func (this PQ) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *PQ) Push(x interface{}) {
	*this = append(*this, x.(*Range))
}

func (this *PQ) Pop() interface{} {
	last := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return last
}

type Range struct {
	start    int
	end      int
	priority int
}

func NewRange(start int, end int, n int) *Range {
	p := (end - start) / 2
	if start == -1 {
		p = end
	} else if end == n {
		p = n - 1 - start
	}
	return &Range{
		start:    start,
		end:      end,
		priority: p,
	}
}

type ExamRoom struct {
	n      int
	pq     *PQ
	starts map[int]*Range
	ends   map[int]*Range
}

func Constructor(n int) ExamRoom {
	pq := &PQ{}
	wholeRange := &Range{start: -1, end: n}
	heap.Push(pq, wholeRange)
	starts := make(map[int]*Range)
	ends := make(map[int]*Range)
	starts[wholeRange.start] = wholeRange
	ends[wholeRange.end] = wholeRange

	return ExamRoom{
		n:      n,
		pq:     pq,
		starts: starts,
		ends:   ends,
	}
}

func (this *ExamRoom) Seat() int {
	r := this.nextRange()

	seat := 0
	if r.start == -1 {
		seat = 0
	} else if r.end == this.n {
		seat = this.n - 1
	} else {
		seat = r.start + (r.end-r.start)/2
	}
	left := NewRange(r.start, seat, this.n)
	right := NewRange(seat, r.end, this.n)

	this.removeRange(r)
	this.addRange(left)
	this.addRange(right)
	return seat
}

func (this *ExamRoom) Leave(p int) {
	left := this.ends[p]
	right := this.starts[p]

	combinedRange := NewRange(left.start, right.end, this.n)

	this.removeRange(left)
	this.removeRange(right)
	this.addRange(combinedRange)
}

func (this *ExamRoom) addRange(r *Range) {
	this.starts[r.start] = r
	this.ends[r.end] = r
	heap.Push(this.pq, r)
}

func (this *ExamRoom) removeRange(r *Range) {
	delete(this.starts, r.start)
	delete(this.ends, r.end)
}

func (this *ExamRoom) nextRange() *Range {
	var r *Range
	for r == nil {
		candidate := heap.Pop(this.pq).(*Range)
		if this.starts[candidate.start] == candidate && this.ends[candidate.end] == candidate {
			r = candidate
		}
	}
	return r
}
