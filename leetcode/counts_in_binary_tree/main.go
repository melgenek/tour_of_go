package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	d1, p1 := depth(root, nil, 0, x)
	d2, p2 := depth(root, nil, 0, y)

	return d1 == d2 && p1 != p2
}

func depth(t *TreeNode, parent *TreeNode, d int, target int) (int, *TreeNode) {
	if t.Val == target {
		return d, parent
	} else {
		var dep int = -1
		var par *TreeNode = t
		if t.Left != nil {
			dep, par = depth(t.Left, t, d+1, target)
		}
		if dep == -1 && t.Right != nil {
			dep, par = depth(t.Right, t, d+1, target)
		}

		return dep, par
	}
}
