package main

import "fmt"

func main() {
	fmt.Printf("%v\n", maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}) == 11)
	fmt.Printf("%v\n", maxScoreSightseeingPair([]int{1, 2}) == 2)
	fmt.Printf("%v\n", maxScoreSightseeingPair([]int{5, 2, 5}) == 8)
	fmt.Printf("%v\n", maxScoreSightseeingPair([]int{10, 4, 6, 4, 10}) == 16)
	fmt.Printf("%v\n", 192 == maxScoreSightseeingPair([]int{30, 13, 28, 32, 16, 8, 11, 78, 83, 5, 22, 93, 61, 60, 100, 8, 6, 48, 87, 43, 41, 86, 93, 5, 19, 29, 59, 31, 7, 51, 99, 47, 40, 24, 20, 98, 41, 42, 81, 92, 55, 85, 51, 92, 84, 21, 84, 92, 1, 73, 93, 51, 44, 27, 23, 54, 32, 57, 60, 9, 69, 14, 28, 86, 15, 92, 47, 63, 12, 99, 54, 6, 16, 52, 28, 86, 38, 73, 16, 52, 37, 30, 84, 81, 46, 97, 84, 17, 21, 14, 52, 19, 74, 20, 20, 56, 89, 7, 34, 21}))
}

func maxScoreSightseeingPair(values []int) int {
	res := 0

	i, j := 0, 1
	for j < len(values) {
		distance := j - i
		value := values[i] + values[j] - distance
		if value > res {
			res = value
		}

		if values[i]-distance > values[j] || distance == 1 {
			j++
		} else {
			i++
		}
	}

	return res
}
