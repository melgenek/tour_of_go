package main

import "fmt"

func main() {
	fmt.Printf("%v\n", subarraySum([]int{1, 1, 1}, -2) == 0)
	fmt.Printf("%v\n", subarraySum([]int{-1, -1, -1}, -2) == 2)
	fmt.Printf("%v\n", subarraySum([]int{1, 1, 1}, 2) == 2)
	fmt.Printf("%v\n", subarraySum([]int{1, 2, 3}, 3) == 2)
	fmt.Printf("%v\n", subarraySum([]int{3, 2, 1}, 3) == 2)
	fmt.Printf("%v\n", subarraySum([]int{3, 2, 1, 0}, 3) == 3)
	fmt.Printf("%v\n", subarraySum([]int{2, 1, 0, 3}, 3) == 4)
	fmt.Printf("%v\n", subarraySum([]int{2, -2, 2, 1, 0, 3}, 3) == 6)
	fmt.Printf("%v\n", subarraySum([]int{1, 2, 3, 4}, 5) == 1)
	fmt.Printf("%v\n", subarraySum([]int{1, 2, 8, 4}, 5) == 0)
}

func subarraySum(nums []int, k int) int {
	n := len(nums)
	count := 0

	prefix := make([]int, n+1)
	prefix[1] = nums[0]
	for i := 2; i <= n; i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}

	for i := 0; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if prefix[j]-prefix[i] == k {
				count++
			}
		}
	}

	return count
}
