package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%v\n", increasingTriplet([]int{1, 2, 3, 4, 5}) == true)
	fmt.Printf("%v\n", increasingTriplet([]int{5, 4, 3, 2, 1}) == false)
	fmt.Printf("%v\n", increasingTriplet([]int{2, 1, 5, 0, 4, 6}) == true)
	fmt.Printf("%v\n", increasingTriplet([]int{2, 5, 1, 0, 4, 6}) == true)
	fmt.Printf("%v\n", increasingTriplet([]int{20, 100, 10, 12, 5, 13}) == true)
	fmt.Printf("%v\n", increasingTriplet([]int{1, 1, 1, 1, 1}) == false)
}

func increasingTriplet(nums []int) bool {
	n := len(nums)
	smallest := math.MaxInt32
	secondSmallest := math.MaxInt32
	for i := 0; i < n; i++ {
		v := nums[i]
		if v > secondSmallest {
			return true
		} else if v > smallest {
			secondSmallest = v
		} else {
			smallest = v
		}
	}

	return false
}
