package main

import "fmt"

func main() {
	fmt.Printf("%v\n", jump([]int{2, 3, 1, 1, 4}) == 2)
	fmt.Printf("%v\n", jump([]int{2, 3, 0, 1, 4}) == 2)
}

func jump(nums []int) int {
	n := len(nums)
	cache := make([]int, n)
	for i := range cache {
		cache[i] = 9999999
	}
	cache[n-1] = 0

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n && j <= i+nums[i]; j++ {
			if cache[j] < cache[i] {
				cache[i] = cache[j] + 1
			}
		}
	}

	return cache[0]
}
