package main

import "container/list"

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type LayerNode struct {
	Node  *Node
	Layer int
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := list.New()
	q.PushBack(LayerNode{root, 0})

	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)
		node := el.Value.(LayerNode)

		nextEl := q.Front()

		if nextEl != nil {
			next := nextEl.Value.(LayerNode)
			if next.Layer == node.Layer {
				node.Node.Next = next.Node
			}

		}

		if node.Node.Left != nil {
			q.PushBack(LayerNode{node.Node.Left, node.Layer + 1})
		}
		if node.Node.Right != nil {
			q.PushBack(LayerNode{node.Node.Right, node.Layer + 1})
		}
	}

	return root
}
