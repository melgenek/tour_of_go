package main

import (
	"container/list"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v\n", 3 == largestRectangleArea([]int{2, 1, 2}))
	fmt.Printf("%v\n", 3 == largestRectangleArea([]int{2, 1, 2}))
	fmt.Printf("%v\n", 10 == largestRectangleArea([]int{3, 2, 6, 5, 1, 2}))
	fmt.Printf("%v\n", 4 == largestRectangleArea([]int{2, 4}))
	fmt.Printf("%v\n", 4 == largestRectangleArea([]int{4, 2}))
	fmt.Printf("%v\n", 7 == largestRectangleArea([]int{1, 1, 1, 1, 1, 2, 3}))
	fmt.Printf("%v\n", 7 == largestRectangleArea([]int{3, 2, 1, 1, 1, 1, 1}))
	fmt.Printf("%v\n", 7 == largestRectangleArea([]int{2, 3, 1, 1, 1, 1, 1}))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]Entry, n)
	right := make([]Entry, n)

	left[0] = Entry{height: heights[0], count: 1}
	right[n-1] = Entry{height: heights[n-1], count: 1}

	for i := 1; i < n; i++ {
		previous := left[i-1]
		left[i] = Entry{height: min(previous.height, heights[i]), count: previous.count + 1}
	}
	for i := n - 2; i >= 0; i-- {
		next := right[i+1]
		right[i] = Entry{height: min(next.height, heights[i]), count: next.count + 1}
	}

	result := 0
	for i := 0; i < n; i++ {
		leftMin := left[i]
		rightMin := right[i]
		height := min(leftMin.height, rightMin.height)
		newArea := height * (leftMin.count + rightMin.count - 1)
		if newArea > result {
			result = newArea
		}
	}
	return result
}

type Entry struct {
	height int
	count  int
}

func largestRectangleAreaStack(heights []int) int {
	n := len(heights)
	result := 0

	stack := list.New()

	for i := 0; i < n; i++ {
		currentHeight := heights[i]

		minHeight := math.MaxInt32
		count := 0
		for el := stack.Front(); stack.Len() > 0 && el.Value.(Entry).height >= currentHeight; el = stack.Front() {
			entry := el.Value.(Entry)
			minHeight = min(minHeight, entry.height)
			count += entry.count
			stack.Remove(el)

			newArea := minHeight * count
			if newArea > result {
				result = newArea
			}
		}
		stack.PushFront(Entry{height: currentHeight, count: count + 1})
	}

	minHeight := math.MaxInt32
	count := 0
	for el := stack.Front(); stack.Len() > 0; el = stack.Front() {
		entry := el.Value.(Entry)
		minHeight = min(minHeight, entry.height)
		count += entry.count
		stack.Remove(el)

		newArea := minHeight * count
		if newArea > result {
			result = newArea
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
