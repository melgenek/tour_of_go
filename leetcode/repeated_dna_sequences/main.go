package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Printf("%v\n", findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Printf("%v\n", findRepeatedDnaSequences("AAAAAAAAAAA"))
}

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	res := make([]string, 0)
	for i := 0; i < len(s)-10+1; i++ {
		subStr := s[i : i+10]
		m[subStr]++
		if m[subStr] == 2 {
			res = append(res, subStr)
		}
	}
	return res
}
