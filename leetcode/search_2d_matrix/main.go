package main

import "fmt"

func main() {

	fmt.Printf("%v\n", searchMatrix(
		[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
		3),
	)

}

func searchMatrix(matrix [][]int, target int) bool {
	from := 0
	height := len(matrix)
	width := len(matrix[0])
	to := height * width

	for from < to {
		middle := (from + to) / 2
		y := middle / width
		x := middle % width

		if matrix[y][x] == target {
			return true
		} else if matrix[y][x] > target {
			to = middle
		} else {
			from = middle + 1
		}
	}

	return false
}
