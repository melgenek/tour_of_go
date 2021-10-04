package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{100, 200},
		{200, 1300},
		{1000, 1250},
		{2000, 3200},
	}) == 3)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{1, 2},
	}) == 1)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{3, 2},
		{4, 3},
	}) == 0)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{1, 2},
		{2, 3},
	}) == 2)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{5, 15}, {3, 19}, {6, 7}, {2, 10}, {5, 16}, {8, 14}, {10, 11}, {2, 19},
	}) == 5)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{7, 17}, {3, 12}, {10, 20}, {9, 10}, {5, 20}, {10, 19}, {4, 18},
	}) == 4)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{10, 20}, {4, 13}, {4, 4}, {3, 11}, {3, 5}, {3, 5},
	}) == 4)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{10, 12}, {6, 15}, {1, 12}, {3, 20}, {10, 19},
	}) == 4)
	fmt.Printf("%v\n", scheduleCourse([][]int{
		{914, 9927}, {333, 712}, {163, 5455}, {835, 5040}, {905, 8433}, {417, 8249}, {921, 9553}, {913, 7394}, {303, 7525}, {582, 8658}, {86, 957}, {40, 9152}, {600, 6941}, {466, 5775}, {718, 8485}, {34, 3903}, {380, 9996}, {316, 7755},
	}))
}

func scheduleCourse(input [][]int) int {
	used := make([]bool, len(input))
	cache := make(map[string]int)

	return rec(input, used, cache, 0, 0, 0)
}

func rec(input [][]int, used []bool, cache map[string]int, i int, time int, courses int) int {
	key := fmt.Sprintf("%d_%d", i, time)
	cached, found := cache[key]
	if found {
		return cached
	} else if i == len(input) {
		return courses
	} else {
		res := courses
		for j := 0; j < len(input); j++ {
			if !used[j] && input[j][0]+time <= input[j][1] {
				used[j] = true
				res = max(res, rec(input, used, cache, i+1, time+input[j][0], courses+1))
				used[j] = false
			}
		}
		cache[key] = res
		return res
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
