package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", 110 == twoCitySchedCost([][]int{
		{10, 20},
		{30, 200},
		{400, 50},
		{30, 20},
	}))
	fmt.Printf("%v\n", 1859 == twoCitySchedCost([][]int{
		{259, 770},
		{448, 54},
		{926, 667},
		{184, 139},
		{840, 118},
		{577, 469},
	}))
	fmt.Printf("%v\n", 3086 == twoCitySchedCost([][]int{
		{515, 563},
		{451, 713},
		{537, 709},
		{343, 819},
		{855, 779},
		{457, 60},
		{650, 359},
		{631, 42},
	}))

}

func twoCitySchedCostGreedy(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][1]-costs[i][0] > costs[j][1]-costs[j][0]
	})
	n := len(costs) / 2
	sum := 0
	for i := 0; i < n; i++ {
		sum += costs[i][0]
	}
	for i := n; i < 2*n; i++ {
		sum += costs[i][1]
	}
	return sum
}

func twoCitySchedCost(costs [][]int) int {
	n := len(costs) / 2
	table := make([][]int, n+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		table[i][0] = costs[i-1][0] + table[i-1][0]
	}
	for j := 1; j <= n; j++ {
		table[0][j] = costs[j-1][1] + table[0][j-1]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			k := i + j - 1
			table[i][j] = min(
				costs[k][0]+table[i-1][j],
				costs[k][1]+table[i][j-1],
			)
		}
	}

	return table[n][n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func twoCitySchedCostRec(costs [][]int) int {
	cache := make(map[string]int)
	n := len(costs) / 2
	return rec(costs, cache, n, n)
}

func rec(costs [][]int, cache map[string]int, firstN int, secondN int) int {
	i := len(costs) - firstN - secondN
	key := fmt.Sprintf("%d_%d", firstN, secondN)
	cached, found := cache[key]
	if found {
		return cached
	} else if i == len(costs) {
		return 0
	} else {
		res := 9999999

		if firstN > 0 {
			firstRes := rec(costs, cache, firstN-1, secondN) + costs[i][0]
			res = min(res, firstRes)
		}

		if secondN > 0 {
			secondRes := rec(costs, cache, firstN, secondN-1) + costs[i][1]
			res = min(res, secondRes)
		}

		cache[key] = res
		return res
	}
}
