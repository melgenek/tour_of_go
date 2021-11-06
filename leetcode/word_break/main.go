package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", wordBreak("leetcode", []string{"leet", "code"}) == true)
	fmt.Printf("%v\n", wordBreak("applepenapple", []string{"apple", "pen"}) == true)
	fmt.Printf("%v\n", wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) == false)
	fmt.Printf("%v\n", wordBreak("catsanddog", []string{"cats", "dog", "sand", "and", "cat"}) == true)
	fmt.Printf("%v\n", wordBreak("catsanddog", []string{"cats", "dog", "and", "cat"}) == true)
	fmt.Printf("%v\n", wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}) == false)
}

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	maxWordLen := 0
	for _, v := range wordDict {
		dict[v] = true
		if len(v) > maxWordLen {
			maxWordLen = len(v)
		}
	}

	n := len(s)
	cache := make([]bool, n+1)
	cache[n] = true
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= n && j <= i+maxWordLen; j++ {
			subStr := s[i:j]
			if dict[subStr] && cache[j] {
				cache[i] = true
				break
			}
		}
	}

	return cache[0]
}
