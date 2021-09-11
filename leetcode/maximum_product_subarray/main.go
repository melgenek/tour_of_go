package main

import "fmt"

func main() {
	//fmt.Printf("%v\n", maxProduct([]int{2, 3, -2, 4}) == 6)
	//fmt.Printf("%v\n", maxProduct([]int{2, 3, -2, 4, 2}) == 8)
	//fmt.Printf("%v\n", maxProduct([]int{2, -2, 3, -2, 4, 2}) == 192)
	fmt.Printf("%v\n", maxProduct([]int{2, -2, 3, -2, -2, 4, 2}) == 96)
	fmt.Printf("%v\n", maxProduct([]int{2, -2, 3, -2, -2, 4, 2}))
	fmt.Printf("%v\n", maxProduct([]int{-2, 0, -1}) == 0)
}

// 2, -2, 3, -1, -2, 4, -3, 2

// 2, -2, 3, -2, -2, 4, 2

// 2, -2, 3, -2, 4, 2
func maxProduct(nums []int) int {
	max := 0

	return max
}
