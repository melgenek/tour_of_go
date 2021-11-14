package main

import "fmt"

func main() {
	fmt.Printf("%v\n", change(4, []int{1, 2, 5}) == 3)
	fmt.Printf("%v\n", change(5, []int{1, 2, 5}) == 4)
	fmt.Printf("%v\n", change(3, []int{2}) == 0)
	fmt.Printf("%v\n", change(3, []int{}) == 0)
	fmt.Printf("%v\n", change(0, []int{2}) == 1)
	fmt.Printf("%v\n", change(10, []int{10}) == 1)
}

func change(amount int, coins []int) int {
	n := len(coins)
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, amount+1)
	}
	cache[0][0] = 1

	for i := 1; i <= n; i++ {
		coin := coins[i-1]
		for j := 0; j <= amount; j++ {
			cache[i][j] = cache[i-1][j]
			if j-coin >= 0 {
				cache[i][j] += cache[i][j-coin]
			}
		}
	}

	return cache[n][amount]
}
