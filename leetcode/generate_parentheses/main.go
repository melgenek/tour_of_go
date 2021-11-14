package main

import "fmt"

func main() {
	fmt.Printf("%v\n", generateParenthesis(1))
	fmt.Printf("%v\n", generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	return rec(n, 0, "")
}

func rec(n int, open int, cur string) []string {
	if n == 0 {
		if open == 0 {
			return []string{cur}
		} else {
			return rec(n, open-1, cur+")")
		}
	} else {
		var res []string
		if open > 0 {
			res = append(res, rec(n, open-1, cur+")")...)
		}
		res = append(res, rec(n-1, open+1, cur+"(")...)
		return res
	}
}
