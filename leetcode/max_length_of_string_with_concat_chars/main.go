package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", maxLength([]string{"a", "b", "c"}) == 3)
	fmt.Printf("%v\n", maxLength([]string{"un", "iq", "ue"}) == 4)
	fmt.Printf("%v\n", maxLength([]string{"cha", "r", "act", "ers"}) == 6)
	fmt.Printf("%v\n", maxLength([]string{"cear", "q", "r", "act", "es"}) == 7)
	fmt.Printf("%v\n", maxLength([]string{"r", "cear", "q", "act", "es"}) == 7)
	fmt.Printf("%v\n", maxLength([]string{"abcdefghijklmnopqrstuvwxyz"}) == 26)
	fmt.Printf("%v\n", maxLength([]string{"yy", "bkhwmpbiisbldzknpm"}) == 0)
	fmt.Printf("%v\n", maxLength([]string{"jnfbyktlrqumowxd", "mvhgcpxnjzrdei"}) == 16)
}

func maxLength(arr []string) int {
	bits := make([]bool, 26)

	var rec func([]string) int

	rec = func(arr []string) int {
		if len(arr) == 0 {
			return wordLen(bits)
		} else {
			head := arr[0]
			res := rec(arr[1:])

			with := 0
			addSuccess := add(bits, head)
			if addSuccess {
				with = rec(arr[1:])
				remove(bits, head)
			}

			if with > res {
				return with
			} else {
				return res
			}

		}
	}

	return rec(arr)
}

func add(bits []bool, word string) bool {
	i := 0
	for ; i < len(word); i++ {
		if bits[word[i]-'a'] {
			i--
			for ; i >= 0; i-- {
				bits[word[i]-'a'] = false
			}
			return false
		}
		bits[word[i]-'a'] = true
	}
	return true
}

func remove(bits []bool, word string) {
	for i := 0; i < len(word); i++ {
		bits[word[i]-'a'] = false
	}
}

func wordLen(bits []bool) int {
	count := 0
	for i := 0; i < 26; i++ {
		if bits[i] {
			count++
		}
	}
	return count
}
