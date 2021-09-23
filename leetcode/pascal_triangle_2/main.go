package main

import "fmt"

func main() {
	fmt.Printf("%v\n", getRow(4))
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	res := make([]int, rowIndex+1)
	res[0] = 1

	for j := 0; j < rowIndex; j++ {
		for i := rowIndex; i >= 1; i-- {
			res[i] = res[i-1] + res[i]
		}
	}

	return res
}
