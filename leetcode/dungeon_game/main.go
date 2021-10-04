package main

import "fmt"

func main() {
	fmt.Printf("%v\n", calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}) == 7)
	fmt.Printf("%v\n", calculateMinimumHP([][]int{
		{2, 3, 3},
		{5, 10, 1},
		{10, 30, 5},
	}) == 1)
	fmt.Printf("%v\n", calculateMinimumHP([][]int{
		{0},
	}) == 1)
	fmt.Printf("%v\n", calculateMinimumHP([][]int{
		{0, 0, 0},
		{1, 1, -1},
	}) == 1)
	fmt.Printf("%v\n", calculateMinimumHP([][]int{
		{1, -3, 3},
		{0, -2, 0},
		{-3, -3, -3},
	}) == 3)
	fmt.Printf("%v\n", calculateMinimumHP([][]int{
		{-3, 5},
	}) == 4)
}

func calculateMinimumHP(dungeon [][]int) int {
	height := len(dungeon)
	width := len(dungeon[0])

	board := make([][]int, height)
	for i := 0; i < height; i++ {
		board[i] = make([]int, width)
	}

	board[height-1][width-1] = min(0, dungeon[height-1][width-1])
	for j := width - 2; j >= 0; j-- {
		board[height-1][j] = min(0, board[height-1][j+1]+dungeon[height-1][j])
	}
	for i := height - 2; i >= 0; i-- {
		board[i][width-1] = min(0, board[i+1][width-1]+dungeon[i][width-1])
	}

	for i := height - 2; i >= 0; i-- {
		for j := width - 2; j >= 0; j-- {
			rightCell := board[i][j+1]
			bottomCell := board[i+1][j]
			current := dungeon[i][j]

			value := current + max(rightCell, bottomCell)
			board[i][j] = min(value, 0)
		}
	}

	res := board[0][0]
	if res <= 0 {
		return -res + 1
	} else {
		return 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}
