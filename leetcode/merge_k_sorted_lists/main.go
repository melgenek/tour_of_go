package main

import "container/heap"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type PQ []*ListNode

func (this PQ) Len() int           { return len(this) }
func (this PQ) Less(i, j int) bool { return this[i].Val < this[j].Val }
func (this PQ) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }

func (this *PQ) Push(x interface{}) { *this = append(*this, x.(*ListNode)) }

func (this *PQ) Pop() interface{} {
	last := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	pq := &PQ{}
	for _, el := range lists {
		if el != nil {
			heap.Push(pq, el)
		}
	}

	var head *ListNode
	var last *ListNode
	for pq.Len() > 0 {
		top := heap.Pop(pq).(*ListNode)

		if last == nil {
			head = top
			last = top
		} else {
			last.Next = top
			last = last.Next
		}
		if top.Next != nil {
			heap.Push(pq, top.Next)
		}
	}
	return head
}
