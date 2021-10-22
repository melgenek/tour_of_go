package main

import "fmt"

func main() {
	fmt.Printf("%v\n", diameterOfBinaryTree(
		&TreeNode{1, &TreeNode{2, &TreeNode{Val: 4}, &TreeNode{Val: 5}}, &TreeNode{Val: 3}},
	))
	fmt.Printf("%v\n", diameterOfBinaryTree(
		&TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
	))
	fmt.Printf("%v\n", diameterOfBinaryTree(
		&TreeNode{Val: 1, Left: &TreeNode{2, &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}, &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}},
	))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res, _ := rec(root)
	return res
}

func rec(t *TreeNode) (int, int) {
	if t.Left == nil && t.Right == nil {
		return 0, 0
	} else if t.Left == nil {
		maxSoFar, depth := rec(t.Right)
		return max(maxSoFar, depth+1), depth + 1
	} else if t.Right == nil {
		maxSoFar, depth := rec(t.Left)
		return max(maxSoFar, depth+1), depth + 1
	} else {
		leftMaxSoFar, leftDepth := rec(t.Left)
		rightMaxSoFar, rightDepth := rec(t.Right)

		return max(max(leftMaxSoFar, rightMaxSoFar), leftDepth+rightDepth+2), max(leftDepth, rightDepth) + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
