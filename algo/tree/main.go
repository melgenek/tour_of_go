package main

import "fmt"

func main() {
	tree := NewTree()
	tree.Add(1)
	tree.Add(6)
	tree.Add(2)
	tree.Add(4)
	tree.Add(3)

	fmt.Printf("%v\n", tree.ToSlice())
	tree.Remove(1)
	tree.Add(1)
	fmt.Printf("%v\n", tree.ToSlice())
	tree.Add(6)
	tree.Remove(2)
	tree.Add(2)
	fmt.Printf("%v\n", tree.ToSlice())
	tree.Remove(4)
	tree.Add(4)
	fmt.Printf("%v\n", tree.ToSlice())

	fmt.Printf("%v\n", tree.Find(10) == false)
	fmt.Printf("%v\n", tree.Find(6) == true)

	fmt.Printf("%v\n", tree.Next(-1) == 1)
	fmt.Printf("%v\n", tree.Next(1) == 2)
	fmt.Printf("%v\n", tree.Next(2) == 3)
	fmt.Printf("%v\n", tree.Next(3) == 4)
	fmt.Printf("%v\n", tree.Next(4) == 6)
	fmt.Printf("%v\n", tree.Next(6) == -1)

}

type Tree struct {
	head *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Add(v int) {
	if t.head == nil {
		t.head = &Node{Val: v}
	}
	add(t.head, v)
}

func add(node *Node, v int) *Node {
	if node == nil {
		return &Node{Val: v}
	} else {
		if node.Val < v {
			node.Right = add(node.Right, v)
		} else if node.Val > v {
			node.Left = add(node.Left, v)
		}
		return node
	}
}

func (t *Tree) Remove(v int) {
	t.head = remove(t.head, v)
}

func remove(node *Node, v int) *Node {
	if node == nil {
	} else if node.Val == v {
		if node.Left == nil {
			node = node.Right
		} else if node.Right == nil {
			node = node.Left
		} else {
			nextMinimal := min(node.Right)
			node.Val = nextMinimal.Val
			node.Right = remove(node.Right, nextMinimal.Val)
		}
	} else if node.Val > v {
		node.Left = remove(node.Left, v)
	} else {
		node.Right = remove(node.Right, v)
	}
	return node
}

func min(node *Node) *Node {
	if node.Left == nil {
		return node
	} else {
		return min(node.Left)
	}
}

func (t *Tree) ToSlice() []int {
	return toSlice(t.head, []int{})
}

func toSlice(node *Node, arr []int) []int {
	if node == nil {
		return arr
	} else {
		leftArr := toSlice(node.Left, arr)
		leftArr = append(leftArr, node.Val)
		return toSlice(node.Right, leftArr)
	}
}

func (t *Tree) Find(v int) bool {
	return find(t.head, v)
}

func find(node *Node, v int) bool {
	if node == nil {
		return false
	} else if node.Val == v {
		return true
	} else if node.Val > v {
		return find(node.Left, v)
	} else {
		return find(node.Right, v)
	}
}

func (t *Tree) Next(v int) int {
	res := next(t.head, nil, v)
	if res == nil {
		return -1
	} else {
		return res.Val
	}
}

func next(node *Node, successor *Node, v int) *Node {
	if node == nil {
		return successor
	} else if node.Val > v {
		return next(node.Left, node, v)
	} else {
		return next(node.Right, successor, v)
	}
}
