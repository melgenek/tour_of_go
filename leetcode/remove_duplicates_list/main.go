package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", deleteDuplicates(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}))
	fmt.Printf("%v\n", deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	var b bytes.Buffer
	b.WriteString("[")
	for cur := head; cur != nil; cur = cur.Next {
		b.WriteString(fmt.Sprintf("%d,", cur.Val))
	}
	b.WriteString("]")
	return b.String()
}

func deleteDuplicates(head *ListNode) *ListNode {
	var first *ListNode
	var last *ListNode
	for cur := head; cur != nil; {
		n := 1
		node := cur
		for ; node != nil && node.Next != nil && node.Next.Val == node.Val; node = node.Next {
			n++
		}

		if n == 1 {
			if first == nil {
				first = &ListNode{node.Val, nil}
				last = first
			} else {
				last.Next = &ListNode{node.Val, nil}
				last = last.Next
			}
		}
		cur = node.Next
	}
	return first
}
