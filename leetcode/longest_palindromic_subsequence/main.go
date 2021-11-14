package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", longestPalindromeSubseq("bbbab") == 4)
	fmt.Printf("%v\n", longestPalindromeSubseq("cbbd") == 2)
	fmt.Printf("%v\n", longestPalindromeSubseq("cbbbb") == 4)
	fmt.Printf("%v\n", longestPalindromeSubseq("c") == 1)
	fmt.Printf("%v\n", longestPalindromeSubseq("cb") == 1)
}

func longestPalindromeSubseq(s string) int {
	n := len(s)
	cache := make([][]int, n+1)
	for i := range cache {
		cache[i] = make([]int, n+1)
	}

	reversed := reverseStr(s)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			ch1 := s[i-1]
			ch2 := reversed[j-1]

			if ch1 == ch2 {
				cache[i][j] = cache[i-1][j-1] + 1
			} else {
				cache[i][j] = max(cache[i-1][j], cache[i][j-1])
			}
		}
	}

	return cache[n][n]
}

func reverseStr(s string) string {
	var b bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
	}
	return b.String()
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPalindromeSubseqRec(s string) int {
	n := len(s)
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, n)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return rec(s, 0, n-1, cache)
}

func rec(s string, i, j int, cache [][]int) int {
	if j-i+1 <= 1 {
		return j - i + 1
	} else if cache[i][j] >= 0 {
		return cache[i][j]
	} else {
		var res int
		if s[i] == s[j] {
			res = rec(s, i+1, j-1, cache) + 2
		} else {
			res = max(rec(s, i, j-1, cache), rec(s, i+1, j, cache))
		}
		cache[i][j] = res
		return res
	}
}
