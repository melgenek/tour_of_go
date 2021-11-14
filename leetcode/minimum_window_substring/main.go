package main

import "fmt"

func main() {
	fmt.Printf("%v\n", "BANC" == minWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("%v\n", "a" == minWindow("a", "a"))
	fmt.Printf("%v\n", "" == minWindow("a", "aa"))
}

func contains(fullMap map[uint8]int, subMap map[uint8]int) bool {
	for k, v := range subMap {
		if fullMap[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	n := len(s)

	tMap := make(map[uint8]int)
	for i := range t {
		ch := t[i]
		tMap[ch]++
	}

	subMap := make(map[uint8]int)

	res := ""
	found := false

	i, j := 0, 0
	for ; j < n; j++ {
		ch := s[j]
		subMap[ch]++

		for ; i <= j && contains(subMap, tMap); i++ {
			if !found || j-i+1 < len(res) {
				res = s[i : j+1]
				found = true
			}
			subMap[s[i]]--
		}
	}

	for ; i <= j && contains(subMap, tMap); i++ {
		if !found || j-i+1 < len(res) {
			res = s[i : j+1]
			found = true
		}
		subMap[s[i]]--
	}

	return res
}
