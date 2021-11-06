package main

import "fmt"

func main() {
	n := &ListNode{2, &ListNode{}}
	cycle := &ListNode{0, &ListNode{4, n}}
	n.Next = cycle
	l1 := &ListNode{3, n}
	fmt.Printf("%v\n", detectCycle(l1))

	fmt.Printf("%v\n", detectCycle(&ListNode{1, &ListNode{2, nil}}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
