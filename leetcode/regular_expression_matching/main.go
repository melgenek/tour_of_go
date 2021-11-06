package main

import "fmt"

func main() {
	fmt.Printf("%v\n", isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*") == true)
	fmt.Printf("%v\n", isMatch("aaa", "a*a") == true)
	fmt.Printf("%v\n", isMatch("ab", ".*c") == false)
	fmt.Printf("%v\n", isMatch("aa", "aa") == true)
	fmt.Printf("%v\n", isMatch("aa", "a") == false)
	fmt.Printf("%v\n", isMatch("aa", "a*") == true)
	fmt.Printf("%v\n", isMatch("ab", ".*") == true)
	fmt.Printf("%v\n", isMatch("aab", "c*a*b") == true)
	fmt.Printf("%v\n", isMatch("mississippi", "mis*is*p*.") == false)
}

func isMatch(s string, p string) bool {
	cache := make([][]int8, len(s)+1)
	for i := range cache {
		cache[i] = make([]int8, len(p)+1)
	}
	return rec(s, p, 0, 0, cache)
}

func rec(s string, p string, i, j int, cache [][]int8) bool {
	res := false
	cached := cache[i][j]
	if cached != 0 {
		res = cached == 1
	} else if i == len(s) {
		for ; j < len(p)-1; j += 2 {
			if p[j+1] != '*' {
				break
			}
		}
		res = j == len(p)
	} else if j >= len(p) {
		res = false
	} else {
		if j != len(p)-1 && p[j+1] == '*' {
			res = res ||
				(s[i] == p[j] || p[j] == '.') && rec(s, p, i+1, j, cache) ||
				rec(s, p, i, j+2, cache)
		} else {
			res = res ||
				(s[i] == p[j] || p[j] == '.') && rec(s, p, i+1, j+1, cache)
		}
	}

	if res {
		cache[i][j] = 1
	} else {
		cache[i][j] = -1
	}

	return res
}
