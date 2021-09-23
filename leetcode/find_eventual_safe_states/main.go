package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", eventualSafeNodes([][]int{
		{1, 2},
		{2, 3},
		{5},
		{0},
		{5},
		{},
		{},
	}))

	fmt.Printf("%v\n", eventualSafeNodes([][]int{
		{1, 2, 3, 4},
		{1, 2},
		{3, 4},
		{0, 4},
		{},
	}))
}

func eventualSafeNodes(input [][]int) []int {
	n := len(input)
	incoming := make([][]int, n)
	outgoing := make([]int, n)

	var q list.List

	for from, tos := range input {
		for _, to := range tos {
			incoming[to] = append(incoming[to], from)
		}
		outgoing[from] = len(tos)
		if outgoing[from] == 0 {
			q.PushBack(from)
		}
	}

	safe := make([]bool, n)
	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)
		v := el.Value.(int)
		safe[v] = true

		for _, from := range incoming[v] {
			outgoing[from]--
			if outgoing[from] == 0 {
				q.PushBack(from)
			}
		}
	}

	res := make([]int, 0)
	for i, isSafe := range safe {
		if isSafe {
			res = append(res, i)
		}
	}
	return res
}
