package main

import "fmt"

func main() {
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) == 6)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{-1}) == -1)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{1}) == 1)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{5, 4, -1, 7, 8}) == 24)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{1, -2, 3, -2}) == 3)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{5, -3, 5}) == 10)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{3, -1, 2, -1}) == 4)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{3, -2, 2, -3}) == 3)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{-2, -3, -1}) == -1)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{4, 4, -3, 5}) == 13)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{4, -3, 4, 5}) == 13)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{4, -3, 5, -3, 4, 5}) == 15)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{4, -3, 5, -4, 4, 5}) == 15)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{4, 4, -3, 5, 5}) == 18)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{6, 9, -3}) == 15)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{-1, 3, -3, 9, -6, 8, -5, -5, -6, 10}) == 20)
	fmt.Printf("%v\n", maxSubarraySumCircular([]int{-9, 14, 24, -14, 12, 18, -18, -10, -10, -23, -2, -23, 11, 12, 18, -9, 9, -29, 12, 4, -8, 15, 11, -12, -16, -9, 19, -12, 22, 16}) == 99)
}

func maxSubarraySumCircular(nums []int) int {
	return 1
}
