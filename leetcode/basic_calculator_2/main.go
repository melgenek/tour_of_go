package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", 1 == calculate("1-1+1"))
	fmt.Printf("%v\n", 7 == calculate("3+2*2"))
	fmt.Printf("%v\n", 1 == calculate("3-2"))
	fmt.Printf("%v\n", 1 == calculate("3/2"))
	fmt.Printf("%v\n", 1 == calculate(" 3/2 "))
	fmt.Printf("%v\n", 5 == calculate(" 3+5 / 2 "))
}

type Operation func(int, int) int

var Add Operation = func(a, b int) int {
	return a + b
}
var Sub Operation = func(a, b int) int {
	return a - b
}

var Mul Operation = func(a, b int) int {
	return a * b
}

var Div Operation = func(a, b int) int {
	return a / b
}

func getOperation(ch uint8) *Operation {
	if ch == '+' {
		return &Add
	} else if ch == '-' {
		return &Sub
	} else if ch == '/' {
		return &Div
	} else if ch == '*' {
		return &Mul
	} else {
		return nil
	}
}

func calculate(s string) int {
	stack := list.New()

	n := len(s)
	for i := 0; i < n; {
		if s[i] == ' ' {
			i++
			continue
		}
		operation := getOperation(s[i])

		if operation != nil {
			i++
			if operation == &Sub {
				v := 0
				for ; i < n && s[i] != ' ' && getOperation(s[i]) == nil; i++ {
					v = v*10 + int(s[i]-'0')
				}
				stack.PushFront(&Add)
				stack.PushFront(-v)
			} else {
				stack.PushFront(operation)
			}
		} else {
			v := 0
			for ; i < n && s[i] != ' ' && getOperation(s[i]) == nil; i++ {
				v = v*10 + int(s[i]-'0')
			}

			if stack.Len() > 0 {
				operationEl := stack.Front()
				operation = operationEl.Value.(*Operation)

				if operation == &Div || operation == &Mul {
					stack.Remove(operationEl)

					anotherOperandEL := stack.Front()
					stack.Remove(anotherOperandEL)

					v = (*operationEl.Value.(*Operation))(anotherOperandEL.Value.(int), v)
				}
			}

			stack.PushFront(v)
		}
	}

	for stack.Len() != 1 {
		operand1El := stack.Front()
		operand1 := operand1El.Value.(int)
		stack.Remove(operand1El)

		operationEl := stack.Front()
		operation := operationEl.Value.(*Operation)
		stack.Remove(operationEl)

		operand2El := stack.Front()
		operand2 := operand2El.Value.(int)
		stack.Remove(operand2El)

		stack.PushFront((*operation)(operand2, operand1))
	}

	return stack.Front().Value.(int)
}
