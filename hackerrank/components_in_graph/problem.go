package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	Solution(os.Stdin, os.Stdout)
}

func componentsInGraph(edges [][]int) (int, int) {
	g := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	min := 99999999
	max := -1

	used := make(map[int]bool)

	for v := range g {
		if !used[v] {
			depth := dfs(v, g, used)

			if depth < min {
				min = depth
			}
			if depth > max {
				max = depth
			}
		}
	}

	return min, max
}

func dfs(v int, g map[int][]int, used map[int]bool) int {
	used[v] = true
	depth := 1
	for _, to := range g[v] {
		if !used[to] {
			depth += dfs(to, g, used)
			//if newDepth > depth {
			//	depth = newDepth
			//}
		}
	}
	return depth
}

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	m, _ := strconv.Atoi(in.Text())

	edges := make([][]int, m)

	for i := 0; i < m; i++ {
		in.Scan()
		a, _ := strconv.Atoi(in.Text())
		in.Scan()
		b, _ := strconv.Atoi(in.Text())

		edges[i] = []int{a, b}
	}

	min, max := componentsInGraph(edges)

	fmt.Fprintf(out, "%d %d", min, max)
}
