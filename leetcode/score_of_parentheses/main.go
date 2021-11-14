package main

import (
	"container/list"
	"fmt"
)

func main() {

	fmt.Printf("%v\n", scoreOfParentheses("(()(()))") == 6)
	fmt.Printf("%v\n", scoreOfParentheses("(())()") == 3)
	fmt.Printf("%v\n", scoreOfParentheses("(())(())") == 4)
	fmt.Printf("%v\n", scoreOfParentheses("(())(())") == 4)
	fmt.Printf("%v\n", scoreOfParentheses("()") == 1)
	fmt.Printf("%v\n", scoreOfParentheses("(())") == 2)
	fmt.Printf("%v\n", scoreOfParentheses("((()))") == 4)
	fmt.Printf("%v\n", scoreOfParentheses("(((())))") == 8)
	fmt.Printf("%v\n", scoreOfParentheses("()()") == 2)
}

type Paren struct {
	open  bool
	value int
}

func scoreOfParentheses(s string) int {
	stack := list.New()

	for _, ch := range s {
		if ch == '(' {
			stack.PushFront(Paren{true, 0})
		} else {
			sum := 0
			for el := stack.Front(); stack.Len() > 0 && !el.Value.(Paren).open; el = stack.Front() {
				sum += el.Value.(Paren).value
				stack.Remove(el)
			}
			stack.Remove(stack.Front())

			if sum != 0 {
				sum = 2 * sum
			} else {
				sum = 1
			}

			stack.PushFront(Paren{false, sum})
		}
	}

	totalSum := 0

	for stack.Len() > 0 {
		totalSum += stack.Remove(stack.Front()).(Paren).value
	}

	return totalSum
}
