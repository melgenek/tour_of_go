package main

import "fmt"

func main() {
	fmt.Printf("%v\n", matrixReshape([][]int{{1, 2}, {3, 4}}, 1, 4))
	fmt.Printf("%v\n", matrixReshape([][]int{{1, 2, 3}, {4, 5, 6}}, 3, 2))
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) == r*c {
		newMat := make([][]int, r)
		for i := range newMat {
			newMat[i] = make([]int, c)
		}
		oldC := len(mat[0])
		for i := 0; i < r*c; i++ {
			oldX := i / oldC
			oldY := i % oldC

			newX := i / c
			newY := i % c

			newMat[newX][newY] = mat[oldX][oldY]
		}
		return newMat
	} else {
		return mat
	}
}
