package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", findMinHeightTrees(4, [][]int{
		{1, 0},
		{1, 2},
		{1, 3},
	}))
	fmt.Printf("%v\n", findMinHeightTrees(1, [][]int{}))
	fmt.Printf("%v\n", findMinHeightTrees(2, [][]int{
		{0, 1},
	}))
	fmt.Printf("%v\n", findMinHeightTrees(2, [][]int{
		{1, 0},
	}))
	fmt.Printf("%v\n", findMinHeightTrees(6, [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 4},
		{5, 4},
	}))
	fmt.Printf("%v\n", findMinHeightTrees(6, [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{3, 4},
		{4, 5},
	}))
}

type Node struct {
	adj    []*Node
	idx    int
	degree int
}

func findMinHeightTrees(n int, edges [][]int) []int {
	g := make(map[int]*Node)
	for i := 0; i < n; i++ {
		g[i] = &Node{make([]*Node, 0), i, 0}
	}
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		g[a].adj = append(g[a].adj, g[b])
		g[a].degree++
		g[b].adj = append(g[b].adj, g[a])
		g[b].degree++
	}

	for i := 0; i < n-2; {
		leaves := make([]*Node, 0)
		for _, v := range g {
			if v.degree == 1 {
				leaves = append(leaves, v)
			}
		}
		i += len(leaves)
		for _, v := range leaves {
			for _, to := range v.adj {
				to.degree--
				g[v.idx].degree--
			}
			delete(g, v.idx)
		}
	}

	result := make([]int, 0)
	for _, v := range g {
		result = append(result, v.idx)
	}

	return result
}
