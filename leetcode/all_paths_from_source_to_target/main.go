package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", allPathsSourceTarget([][]int{
		{1, 2},
		{3},
		{3},
		{},
	}))
	fmt.Printf("%v\n", allPathsSourceTarget([][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}))
	fmt.Printf("%v\n", allPathsSourceTarget([][]int{
		{1, 2, 3},
		{2},
		{3},
		{},
	}))
	fmt.Printf("%v\n", allPathsSourceTarget([][]int{
		{1, 3},
		{2},
		{3},
		{},
	}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := dfs(graph, 0, len(graph)-1)
	for _, path := range paths {
		for i := 0; i < len(path)/2; i++ {
			path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
		}
	}
	return paths
}

func dfs(graph [][]int, x int, target int) [][]int {
	if x == target {
		return [][]int{{target}}
	} else {
		paths := make([][]int, 0)
		for _, next := range graph[x] {
			paths = append(paths, dfs(graph, next, target)...)
		}
		for i, _ := range paths {
			paths[i] = append(paths[i], x)
		}
		return paths
	}
}
