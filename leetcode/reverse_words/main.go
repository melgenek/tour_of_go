package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", reverseWords("the sky is blue") == "blue is sky the")
	fmt.Printf("%v\n", reverseWords("   hello   world      ") == "world hello")
	fmt.Printf("%v\n", reverseWords("        ") == "")
}

func reverseWords(s string) string {
	arr := []byte(s)
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}

	var b bytes.Buffer
	start := 0
	end := 0
	for end < n {
		if arr[end] == ' ' {
			if end != start {
				if b.Len() > 0 {
					b.WriteByte(' ')
				}
				for k := end - 1; k >= start; k-- {
					b.WriteByte(arr[k])
				}
			}
			start = end + 1
			end = start
		} else {
			end++
		}
	}
	if end != start && b.Len() > 0 {
		b.WriteByte(' ')
	}
	for k := end - 1; k >= start; k-- {
		b.WriteByte(arr[k])
	}

	return b.String()
}
