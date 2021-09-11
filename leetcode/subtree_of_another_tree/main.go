package main

import "fmt"

func main() {
	fmt.Printf("%v\n", isSubtree(
		&TreeNode{3,
			&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
	))
	fmt.Printf("%v\n", isSubtree(
		&TreeNode{3,
			&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, &TreeNode{0, nil, nil}, nil}},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}},
	))

	fmt.Printf("%v\n", isSubtree(
		&TreeNode{1,
			&TreeNode{1, nil, nil},
			nil,
		},
		&TreeNode{1, nil, nil},
	))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	} else {
		return compare(root, subRoot) || isSubtree(root.Right, subRoot) || isSubtree(root.Left, subRoot)
	}
}

func compare(first *TreeNode, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	} else if first == nil && second != nil || first != nil && second == nil {
		return false
	} else {
		return first.Val == second.Val &&
			compare(first.Left, second.Left) &&
			compare(first.Right, second.Right)
	}
}
