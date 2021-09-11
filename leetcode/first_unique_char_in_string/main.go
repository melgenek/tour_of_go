package main

import "fmt"

func main() {
	fmt.Printf("%v\n", firstUniqChar("loveleetcode"))
}
func firstUniqChar(s string) int {

	m := make(map[uint8][]int)

	for i := 0; i < len(s); i++ {
		arr := m[s[i]]
		if arr == nil {
			m[s[i]] = []int{1, i}
		} else {
			arr[0]++
		}
	}

	min := -1
	for _, arr := range m {
		if arr[0] == 1 {
			if min == -1 || arr[1] < min {
				min = arr[1]
			}
		}
	}

	return min
}
