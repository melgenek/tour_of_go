package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Printf("%v\n", dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Printf("%v\n", dailyTemperatures([]int{30, 60, 90}))
}

type Day struct {
	value int
	idx   int
}

func dailyTemperatures(temperatures []int) []int {
	stack := list.New()
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		current := temperatures[i]
		var nextDay *Day
		for stack.Len() > 0 {
			el := stack.Front()
			future := el.Value.(*Day)
			if future.value <= current {
				stack.Remove(el)
			} else {
				nextDay = future
				break
			}
		}
		if nextDay != nil {
			res[i] = nextDay.idx - i
		}
		stack.PushFront(&Day{current, i})
	}

	return res
}
