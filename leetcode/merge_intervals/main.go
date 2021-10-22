package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", merge([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
	fmt.Printf("%v\n", merge([][]int{
		{2, 6}, {15, 18}, {8, 10}, {1, 3},
	}))
	fmt.Printf("%v\n", merge([][]int{
		{1, 4}, {4, 5},
	}))
	fmt.Printf("%v\n", merge([][]int{
		{4, 5}, {1, 4},
	}))
	fmt.Printf("%v\n", merge([][]int{
		{1, 7}, {2, 4}, {4, 5},
	}))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	for i := 0; i < len(intervals); {
		start := intervals[i][0]
		end := intervals[i][1]

		for ; i < len(intervals) && end >= intervals[i][0]; i++ {
			end = max(end, intervals[i][1])
		}
		res = append(res, []int{start, end})
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
