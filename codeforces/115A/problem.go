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

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	str := in.Text()
	n, _ := strconv.Atoi(str)

	graph := make([][]int, n)
	var roots []int
	for i := 0; i < n; i++ {
		in.Scan()
		managerStr := in.Text()
		manager, _ := strconv.Atoi(managerStr)
		if manager == -1 {
			roots = append(roots, i)
		} else {
			graph[manager-1] = append(graph[manager-1], i)
		}
	}

	max := -1
	used := make([]bool, n)
	for _, root := range roots {
		newDepth := depth(graph, used, root)
		if newDepth > max {
			max = newDepth
		}
	}

	fmt.Fprintf(out, "%v", max)
	out.Flush()

}

func depth(graph [][]int, used []bool, i int) int {
	if used[i] {
		return 1
	} else {
		used[i] = true
		max := 1
		for _, element := range graph[i] {
			elementDepth := depth(graph, used, element) + 1
			if elementDepth > max {
				max = elementDepth
			}
		}
		return max
	}
}
