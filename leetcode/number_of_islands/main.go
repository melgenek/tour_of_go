package main

import "fmt"

func main() {
	fmt.Printf("%v\n", numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

	fmt.Printf("%v\n", numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}

type Direction struct {
	dx int
	dy int
}

var directions = []Direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func numIslands(grid [][]byte) int {
	used := make([][]bool, len(grid))
	for i := range used {
		used[i] = make([]bool, len(grid[i]))
	}

	count := 0

	for i, row := range grid {
		for j, cell := range row {
			if !used[i][j] && cell != '0' {
				dfs(i, j, grid, used)
				count++
			}
		}
	}

	return count
}

func dfs(x int, y int, grid [][]byte, used [][]bool) {
	used[x][y] = true

	for _, direction := range directions {
		nextX := x + direction.dx
		nextY := y + direction.dy
		if nextX >= 0 && nextX < len(grid) &&
			nextY >= 0 && nextY < len(grid[0]) &&
			!used[nextX][nextY] && grid[nextX][nextY] != '0' {
			dfs(nextX, nextY, grid, used)
		}
	}

}
