package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))

	fmt.Printf("%v\n", findCircleNum([][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}))
}

func findCircleNum(isConnected [][]int) int {
	used := make([]bool, len(isConnected))
	count := 0

	for city := 0; city < len(isConnected); city++ {
		if !used[city] {
			dfs(city, isConnected, used)
			count++
		}
	}

	return count
}

func dfs(city int, isConnected [][]int, used []bool) {
	used[city] = true

	for nextCity, isConnectedWithNextCity := range isConnected[city] {
		if isConnectedWithNextCity == 1 && !used[nextCity] {
			dfs(nextCity, isConnected, used)
		}
	}

}
