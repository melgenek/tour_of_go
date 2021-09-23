package main

func main() {

}

func longestIncreasingPath(matrix [][]int) int {
	maxLength := 0

	cache := make([][]int, len(matrix))
	for i := range cache {
		cache[i] = make([]int, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			maxLength = max(maxLength, dfs(matrix, i, j, cache))
		}
	}
	return maxLength
}

type Direction struct {
	dx int
	dy int
}

var directions = []Direction{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func dfs(g [][]int, x int, y int, cache [][]int) int {
	if cache[x][y] != 0 {
		return cache[x][y]
	}

	length := 1
	for _, direction := range directions {
		nextX := x + direction.dx
		nextY := y + direction.dy

		if nextX >= 0 && nextX < len(g) &&
			nextY >= 0 && nextY < len(g[nextX]) &&
			g[x][y] < g[nextX][nextY] {
			length = max(length, dfs(g, nextX, nextY, cache)+1)
		}
	}
	cache[x][y] = length

	return length
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
