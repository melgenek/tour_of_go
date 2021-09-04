package main

import "fmt"

func main() {
	fmt.Printf("%v\n", generate(5))
}

func generate(numRows int) [][]int {
	t := make([][]int, numRows)

	for i := range t {
		t[i] = make([]int, i+1)
	}
	t[0][0] = 1

	for i := 1; i < numRows; i++ {
		row := t[i]
		previousRow := t[i-1]

		row[0] = 1
		for j := 1; j < i; j++ {
			row[j] = previousRow[j-1] + previousRow[j]
		}
		row[i] = 1
	}

	return t
}
