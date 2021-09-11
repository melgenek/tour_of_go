package main

import "fmt"

func main() {
	fmt.Printf("X=%v O=%v\n", 'X', 'O')
	fmt.Printf("%v\n", s([][]byte{
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
		{'O', 'O', 'O'},
	}))
	fmt.Printf("%v\n", s([][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}))
	fmt.Printf("%v\n", s([][]byte{
		{'X'},
	}))
	fmt.Printf("%v\n", s([][]byte{
		{'O'},
	}))
	fmt.Printf("%v\n", s([][]byte{
		{'X', 'O', 'X', 'O', 'X', 'O', 'O', 'O', 'X', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X', 'X'},
		{'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O', 'O'},
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'O'},
		{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'O', 'X', 'O', 'O', 'X'},
		{'O', 'O', 'O', 'O', 'X', 'O', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'O', 'X', 'X', 'X', 'X', 'O', 'O', 'O'},
	}))
}

func s(board [][]byte) [][]byte {
	solve(board)

	return board
}

type Direction struct {
	dx int
	dy int
}

var directions = []Direction{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func solve(board [][]byte) {
	seen := make([][]bool, len(board))
	for i := range seen {
		seen[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for _, j := range []int{0, len(board[i]) - 1} {
			if !seen[i][j] && board[i][j] == 'O' {
				dfs(board, seen, i, j, true)
			}
		}
	}

	for _, i := range []int{0, len(board) - 1} {
		for j := range board[i] {
			if !seen[i][j] && board[i][j] == 'O' {
				dfs(board, seen, i, j, true)
			}
		}
	}

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if !seen[i][j] && board[i][j] == 'O' {
				dfs(board, seen, i, j, false)
			}
		}
	}

}

func dfs(board [][]byte, seen [][]bool, x int, y int, seenEdge bool) bool {
	seen[x][y] = true
	if !seenEdge {
		board[x][y] = 'X'
	}

	for _, direction := range directions {
		nextX := x + direction.dx
		nextY := y + direction.dy
		if nextX >= 0 && nextX < len(board) &&
			nextY >= 0 && nextY < len(board[nextX]) &&
			!seen[nextX][nextY] && board[nextX][nextY] == 'O' {
			dfs(board, seen, nextX, nextY, seenEdge)
		}
	}

	return seenEdge
}
