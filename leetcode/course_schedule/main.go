package main

import "fmt"

func main() {
	fmt.Printf("%v\n", canFinish(2, [][]int{
		{1, 0},
	}) == true)
	fmt.Printf("%v\n", canFinish(2, [][]int{
		{0, 1},
	}) == true)
	fmt.Printf("%v\n", canFinish(2, [][]int{
		{1, 0},
		{0, 1},
	}) == false)
	fmt.Printf("%v\n", canFinish(4, [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
	}) == true)
	fmt.Printf("%v\n", canFinish(4, [][]int{
		{1, 0},
		{2, 1},
		{3, 2},
		{0, 2},
	}) == false)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)

	for _, prerequisite := range prerequisites {
		to := prerequisite[0]
		from := prerequisite[1]
		g[from] = append(g[from], to)
	}

	used := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if used[i] == 0 {
			if dfs(g, i, used) {
				return false
			}
		}
	}

	return true
}

func dfs(g [][]int, i int, used []int) bool {
	used[i] = 1

	for _, to := range g[i] {
		if used[to] == 1 {
			return true
		} else if used[to] == 0 {
			if dfs(g, to, used) {
				return true
			}
		}
	}

	used[i] = 2
	return false
}
