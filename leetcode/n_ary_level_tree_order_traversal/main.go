package main

import "container/list"

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

type LeveledNode struct {
	level int
	node  *Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := list.New()
	q.PushFront(&LeveledNode{0, root})

	res := make([][]int, 0)

	for q.Len() > 0 {
		el := q.Back()
		q.Remove(el)
		n := el.Value.(*LeveledNode)

		if len(res) == n.level {
			res = append(res, []int{})
		}
		res[n.level] = append(res[n.level], n.node.Val)

		for _, child := range n.node.Children {
			q.PushFront(&LeveledNode{n.level + 1, child})
		}
	}

	return res
}
