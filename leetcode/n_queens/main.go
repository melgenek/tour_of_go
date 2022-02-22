package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	used := make([][]bool, n)
	for i := range used {
		used[i] = make([]bool, n)
	}

	selections := rec(n, used, 0)

	return selections
}

func rec(n int, selection [][]bool, k int) [][]string {
	if k == n {
		rows := make([]string, n)
		for rowIdx, row := range selection {
			rowStr := ""
			for _, isQueen := range row {
				if isQueen {
					rowStr += "Q"
				} else {
					rowStr += "."
				}
			}
			rows[rowIdx] = rowStr
		}
		return [][]string{rows}
	} else {
		result := [][]string{}
		for j := 0; j < n; j++ {
			canUse := true
			for row := 0; row < k; row++ {
				if selection[row][j] {
					canUse = false
					break
				}
			}

			for row, column := k-1, j-1; row >= 0 && column >= 0; row-- {
				if selection[row][column] {
					canUse = false
					break
				}
				column--
			}

			for row, column := k-1, j+1; row >= 0 && column < n; row-- {
				if selection[row][column] {
					canUse = false
					break
				}
				column++
			}

			if canUse {
				selection[k][j] = true
				result = append(result, rec(n, selection, k+1)...)
				selection[k][j] = false
			}
		}

		return result
	}
}
