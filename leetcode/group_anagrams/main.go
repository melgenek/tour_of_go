package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", groupAnagrams([]string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}))
	fmt.Printf("%v\n", groupAnagrams([]string{
		"", "",
	}))
	fmt.Printf("%v\n", groupAnagrams([]string{
		"a",
	}))
	fmt.Printf("%v\n", groupAnagrams([]string{
		"ac", "d",
	}))
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		h := hash(str)
		groups[h] = append(groups[h], str)
	}

	res := make([][]string, len(groups))
	k := 0
	for _, v := range groups {
		res[k] = v
		k++
	}
	return res
}

func hash(str string) string {
	arr := make([]int, 26)
	for _, v := range str {
		arr[v-'a']++
	}
	var b bytes.Buffer
	for _, v := range arr {
		b.WriteRune(int32(v))
	}
	return b.String()
}
