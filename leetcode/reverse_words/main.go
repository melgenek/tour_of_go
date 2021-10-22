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
	var b bytes.Buffer

	start := len(s)
	end := len(s)

	atLeastOne := false
	for i := len(s) - 1; i >= 0; i-- {
		ch := s[i]
		if ch == ' ' {
			if start < end {
				if atLeastOne {
					b.WriteByte(' ')
				}
				b.WriteString(s[start:end])
				atLeastOne = true
			}
			start = i
			end = i
		} else {
			start = i
		}
	}

	if start < end {
		if atLeastOne {
			b.WriteByte(' ')
		}
		b.WriteString(s[start:end])
	}
	return b.String()
}
