package main

import "container/list"

func main() {
	t := bstFromPreorder([]int{8, 5, 1, 7, 10, 12})
	println(t)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}

	stack := list.New()
	stack.PushFront(root)

	for i := 1; i < len(preorder); i++ {
		v := preorder[i]

		newEl := &TreeNode{Val: v}

		lastEl := stack.Front()

		if lastEl.Value.(*TreeNode).Val > v {
			lastEl.Value.(*TreeNode).Left = newEl
		} else {
			preLast := lastEl
			for ; stack.Len() > 0 && v > lastEl.Value.(*TreeNode).Val; preLast, lastEl = lastEl, stack.Front() {
				stack.Remove(lastEl)
			}
			preLast.Value.(*TreeNode).Right = newEl
		}

		stack.PushFront(newEl)
	}

	return root
}
