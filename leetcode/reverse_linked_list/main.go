package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", reverseList(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var b bytes.Buffer
	for cur := l; cur != nil; cur = cur.Next {
		b.WriteString(fmt.Sprintf("%v ", cur.Val))
	}
	return b.String()
}

func reverseList(head *ListNode) *ListNode {
	last, head := reverseRec(head)
	head.Next = nil
	return last
}

func reverseRec(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	} else {
		last, prev := reverseRec(head.Next)
		prev.Next = head
		return last, head
	}
}
