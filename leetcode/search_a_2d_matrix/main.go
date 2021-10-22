package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", searchMatrix([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 1))
}

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	if n < m {
		for i := 0; i < n; i++ {
			idx := sort.Search(m, func(j int) bool {
				return matrix[i][j] >= target
			})
			if idx != m && matrix[i][idx] == target {
				return true
			}
		}
	} else {
		for j := 0; j < m; j++ {
			idx := sort.Search(n, func(i int) bool {
				return matrix[i][j] >= target
			})
			if idx != n && matrix[idx][j] == target {
				return true
			}
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	return rec(matrix, 0, len(matrix[0])-1, target)
}

func rec(matrix [][]int, i, j int, target int) bool {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return false
	}

	v := matrix[i][j]

	if v == target {
		return true
	} else if target > v {
		return rec(matrix, i+1, j, target)
	} else {
		return rec(matrix, i, j-1, target)
	}
}
