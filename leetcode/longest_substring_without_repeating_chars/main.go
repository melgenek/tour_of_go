package main

import "fmt"

func main() {
	//fmt.Printf("%v\n", lengthOfLongestSubstring("dvdf") == 3)
	fmt.Printf("%v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Printf("%v\n", lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Printf("%v\n", lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Printf("%v\n", lengthOfLongestSubstring("jbpnbwwd") == 4)
}

func lengthOfLongestSubstring(s string) int {
	current := make(map[uint8]int)

	maxN := 0

	for i, j := 0, 0; j < len(s); j++ {
		ch := s[j]
		existingIdx, found := current[ch]

		if found && existingIdx >= i {
			i = existingIdx + 1
		}

		current[ch] = j
		if maxN < j-i+1 {
			maxN = j - i + 1
		}
	}

	return maxN
}
