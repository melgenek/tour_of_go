package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", eraseOverlapIntervals([][]int{}) == 0)
	fmt.Printf("%v\n", eraseOverlapIntervals([][]int{
		{1, 2},
	}) == 0)
	fmt.Printf("%v\n", eraseOverlapIntervals([][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	}) == 1)
	fmt.Printf("%v\n", eraseOverlapIntervals([][]int{
		{1, 2}, {1, 2}, {1, 2},
	}) == 2)
	fmt.Printf("%v\n", eraseOverlapIntervals([][]int{
		{1, 2}, {2, 3},
	}) == 0)
	fmt.Printf("%v\n", eraseOverlapIntervals([][]int{
		{1, 3}, {3, 5}, {2, 4},
	}) == 1)
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
			end = min(end, intervals[i][1])
		} else {
			end = intervals[i][1]
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
