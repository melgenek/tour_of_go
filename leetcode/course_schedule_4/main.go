package main

import "fmt"

func main() {
	fmt.Printf("%v\n", checkIfPrerequisite(2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}))
}

const aLot = 999999

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, n)
		for j := range table[i] {
			table[i][j] = aLot
		}
	}

	for _, edge := range prerequisites {
		from := edge[0]
		to := edge[1]
		table[from][to] = 1
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				table[i][j] = min(table[i][j], table[i][k]+table[k][j])
			}
		}
	}

	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = table[q[0]][q[1]] != aLot
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
