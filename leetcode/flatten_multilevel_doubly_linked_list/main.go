package main

import (
	"container/list"
	"fmt"
)

func main() {
	level2 := &Node{Val: 3}
	level1 := &Node{Val: 1, Child: level2, Next: &Node{Val: 2}}
	//level3 := &Node{Val: 11, Next: &Node{Val: 12}}
	//level2 := &Node{Val: 7, Next: &Node{Val: 8, Child: level3, Next: &Node{Val: 9, Next: &Node{Val: 10}}}}
	//level1 := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Child: level2, Next: &Node{Val: 4, Next: &Node{Val: 5, Next: &Node{Val: 6}}}}}}
	a := flatten(level1)
	fmt.Printf("%v\n", a)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	} else {
		rec(nil, root)
		return root
	}
}

func rec(prev *Node, cur *Node) *Node {
	if cur == nil {
		return prev
	}

	if prev != nil {
		prev.Next = cur
		cur.Prev = prev
	}

	next := cur.Next

	cur, cur.Child = rec(cur, cur.Child), nil
	cur = rec(cur, next)

	return cur
}

func flatten2(root *Node) *Node {
	stack := list.New()

	var prev *Node
	for cur := root; cur != nil; {
		if prev != nil {
			prev.Next = cur
			cur.Prev = prev
		}
		prev = cur

		if cur.Child != nil {
			if cur.Next != nil {
				stack.PushFront(cur.Next)
			}
			cur, cur.Child = cur.Child, nil
		} else if cur.Next != nil {
			cur = cur.Next
		} else if stack.Len() > 0 {
			topEl := stack.Front()
			stack.Remove(topEl)
			cur = topEl.Value.(*Node)
		} else {
			cur = nil
		}
	}
	return root
}
