package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", findOrder(2, [][]int{
		{1, 0},
	}))
	fmt.Printf("%v\n", findOrder(2, [][]int{
		{0, 1},
	}))
	fmt.Printf("%v\n", findOrder(1, [][]int{}))
	fmt.Printf("%v\n", findOrder(4, [][]int{
		{1, 0},
		{2, 0},
		{3, 1},
		{3, 2},
	}))
	fmt.Printf("%v\n", findOrder(4, [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
		{0, 2},
	}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)

	for _, prerequisite := range prerequisites {
		to := prerequisite[0]
		from := prerequisite[1]
		g[from] = append(g[from], to)
	}

	used := make([]int, numCourses)
	res := make([]int, 0)

	var dfs func(int) bool

	dfs = func(i int) bool {
		used[i] = 1
		for _, to := range g[i] {
			if used[to] == 1 {
				return true
			} else if used[to] == 0 {
				if dfs(to) {
					return true
				}
			}
		}
		used[i] = 2
		res = append(res, i)
		return false
	}

	for i := 0; i < numCourses; i++ {
		if used[i] == 0 {
			if dfs(i) {
				return []int{}
			}
		}
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
