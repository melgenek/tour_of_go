package main

func main() {
	t1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
	}
	flatten(t1)

	t2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
	}
	flatten(t2)
	println(t2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	rec(root)
}

func rec(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	} else {
		oldRight := node.Right
		next := node

		if node.Left != nil {
			next = rec(node.Left)
			node.Right = node.Left
			node.Left = nil
		}

		if oldRight != nil {
			next.Right = oldRight
			next = rec(node.Right)
		}

		return next
	}
}
