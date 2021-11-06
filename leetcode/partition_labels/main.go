package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Printf("%v\n", partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {
	counts := make([]int, 26)
	for _, v := range s {
		counts[v-'a']++
	}

	res := make([]int, 0)
	partition := make(map[int]bool)

	currentPartitionLength := 0
	for _, v := range s {
		ch := int(v - 'a')

		if !partition[ch] {
			partition[ch] = true
		}

		counts[ch]--
		currentPartitionLength++

		leftForPartition := 0
		for chInPartition := range partition {
			leftForPartition += counts[chInPartition]
		}
		if leftForPartition == 0 {
			res = append(res, currentPartitionLength)
			currentPartitionLength = 0
			partition = make(map[int]bool)
		}
	}

	return res
}

func partitionLabels2(s string) []int {
	intervals := make(map[uint8][]int)
	for i := range s {
		ch := s[i]
		interval := intervals[ch]
		if interval != nil {
			intervals[ch][1] = i
		} else {
			intervals[ch] = []int{i, i}
		}
	}

	intervalsArr := make([][]int, len(intervals))
	k := 0
	for _, v := range intervals {
		intervalsArr[k] = v
		k++
	}
	sort.Slice(intervalsArr, func(i, j int) bool {
		return intervalsArr[i][0] < intervalsArr[j][0]
	})

	mergedIntervals := [][]int{intervalsArr[0]}
	for i := 1; i < len(intervalsArr); i++ {
		lastMerged := len(mergedIntervals) - 1
		if mergedIntervals[lastMerged][1] > intervalsArr[i][0] {
			mergedIntervals[lastMerged][1] = max(mergedIntervals[lastMerged][1], intervalsArr[i][1])
		} else {
			mergedIntervals = append(mergedIntervals, intervalsArr[i])
		}
	}

	res := make([]int, len(mergedIntervals))
	for i, v := range mergedIntervals {
		res[i] = v[1] - v[0] + 1
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
