package main

import "fmt"

func main() {
	fmt.Printf("%v\n", loudAndRich(
		[][]int{{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3}},
		[]int{3, 2, 5, 4, 6, 1, 7, 0}),
	)
	fmt.Printf("%v\n", loudAndRich(
		[][]int{},
		[]int{0}),
	)
}

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)
	g := make([][]int, n)
	for _, edge := range richer {
		less := edge[1]
		more := edge[0]
		g[less] = append(g[less], more)
	}

	cache := make([]int, n)
	for i := range cache {
		cache[i] = -1
	}
	for i := 0; i < n; i++ {
		dfs(g, quiet, i, cache)
	}

	return cache
}

func dfs(g [][]int, quiet []int, x int, cache []int) int {
	if cache[x] >= 0 {
		return cache[x]
	}

	quitest := x
	for _, to := range g[x] {
		newLoudest := dfs(g, quiet, to, cache)
		if quiet[quitest] > quiet[newLoudest] {
			quitest = newLoudest
		}
	}
	cache[x] = quitest

	return quitest
}
