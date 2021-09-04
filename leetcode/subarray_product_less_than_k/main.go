package main

import "fmt"

func main() {
	fmt.Printf("%v\n", numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Printf("%v\n", numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
	fmt.Printf("%v\n", numSubarrayProductLessThanK([]int{101, 2, 2, 2, 3}, 100))
	fmt.Printf("%v\n", numSubarrayProductLessThanK([]int{1, 1, 1}, 5))
}

// [2] [2] [2] [3] [2 2] [2 2] [2 3] [2 2 2] [2 2 3] [2 2 2 3]

// [1] [1] [1] [1 1] [1 1 1] [1 1]

// 10 5 2 6

// j=0 i=0
// mul=10 j++

// i=0 j=1
// mul=50 j++

// i=0 j=2
// mul=100 j++

// i=0 j=3
// mul=10 i++

//

// 10 5 2 6
func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0

	mul := nums[0]
	i, j := 0, 1
	for i < len(nums) && j < len(nums) {
		if mul < k {
			mul *= nums[j]
			count += j - i
			j++
		} else {
			mul /= nums[i]
			i++
		}

	}

	for ; i < j; i++ {
		mul /= nums[i]
		count++
	}

	return count
}
