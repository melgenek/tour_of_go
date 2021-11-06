package main

import "fmt"

func main() {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}

	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}

	a := cloneGraph(n4)
	fmt.Printf("%v", a)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	cache := make(map[*Node]*Node)
	dfs(node, cache)
	return cache[node]
}

func dfs(node *Node, cache map[*Node]*Node) {
	cache[node] = &Node{Val: node.Val}

	neighbors := make([]*Node, len(node.Neighbors))
	for i, neighbor := range node.Neighbors {
		if cache[neighbor] == nil {
			dfs(neighbor, cache)
		}
		neighbors[i] = cache[neighbor]
	}

	cache[node].Neighbors = neighbors
}
