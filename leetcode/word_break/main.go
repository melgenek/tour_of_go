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
	longestWord := 0
	for _, v := range wordDict {
		dict[v] = true

		if len(v) > longestWord {
			longestWord = len(v)
		}
	}

	n := len(s)
	memory := make([]bool, n+1)
	memory[n] = true
	for i := n - 1; i >= 0; i-- {
		subRes := false
		for j := min(i+longestWord, n); j > i; j-- {
			substring := s[i:j]
			if dict[substring] && memory[j] {
				subRes = true
				break
			}
		}
		memory[i] = subRes
	}

	return memory[0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func wordBreak2(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	return wb(s, dict)
}

func wb(s string, dict map[string]bool) bool {
	if len(s) == 0 {
		return true
	} else {
		cached, found := dict[s]
		if found {
			return cached
		}

		res := false
		for i := 1; i < len(s); i++ {
			if dict[s[:i]] {
				res = wb(s[i:], dict)
				if res {
					break
				}
			}
		}
		dict[s] = res
		return res
	}
}
