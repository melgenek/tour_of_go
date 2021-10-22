package main

import "fmt"

func main() {
	m1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(m1)
	fmt.Printf("%v\n", m1)

	m2 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(m2)
	fmt.Printf("%v\n", m2)
}

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			newI, newJ := newCoords(n, i, j)
			prevValue := matrix[i][j]
			for k := 0; k < 4; k++ {
				matrix[newI][newJ], prevValue = prevValue, matrix[newI][newJ]
				newI, newJ = newCoords(n, newI, newJ)
			}
		}
	}
}

func newCoords(n, i, j int) (int, int) {
	return j, n - 1 - i
}
