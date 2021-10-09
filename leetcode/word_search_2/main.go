package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findWords([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain"}))
	fmt.Printf("%v\n", findWords([][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}, []string{"abcb"}))
}

type Trie struct {
	letters [26]*Trie
	word    *string
}

func (this *Trie) Add(word string) {
	t := this
	for i := range word {
		v := word[i] - 'a'
		if t.letters[v] == nil {
			t.letters[v] = &Trie{}
		}
		t = t.letters[v]
	}
	t.word = &word
}

func findWords(board [][]byte, words []string) []string {
	t := &Trie{}
	for _, word := range words {
		t.Add(word)
	}

	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[i]))
	}

	res := &Result{}
	for i, row := range board {
		for j := range row {
			used[i][j] = true
			v := board[i][j] - 'a'
			subTrie := t.letters[v]
			dfs(board, used, i, j, subTrie, res)
			used[i][j] = false
		}
	}

	return res.words
}

type Result struct {
	words []string
}

type Direction struct {
	dx int
	dy int
}

var directions = []Direction{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func dfs(board [][]byte, used [][]bool, x int, y int, t *Trie, result *Result) {
	if t == nil {
		return
	}
	if t.word != nil {
		result.words = append(result.words, *t.word)
		t.word = nil
	}

	for _, direction := range directions {
		newX := x + direction.dx
		newY := y + direction.dy
		if newX >= 0 && newX < len(board) &&
			newY >= 0 && newY < len(board[newX]) &&
			!used[newX][newY] {
			used[newX][newY] = true
			v := board[newX][newY] - 'a'
			subTrie := t.letters[v]
			dfs(board, used, newX, newY, subTrie, result)
			used[newX][newY] = false
		}
	}

}
