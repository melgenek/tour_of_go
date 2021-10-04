package main

import "fmt"

func main() {
	fmt.Printf("%v\n", canJump([]int{2, 3, 1, 1, 4}) == true)
	fmt.Printf("%v\n", canJump([]int{3, 2, 1, 0, 4}) == false)
}

func canJump(nums []int) bool {
	n := len(nums)
	cache := make([]bool, n)
	cache[n-1] = true

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n && j <= i+nums[i]; j++ {
			if cache[j] {
				cache[i] = true
				break
			}
		}
	}

	return cache[0]
}

func canJumpRec(nums []int) bool {
	cache := make(map[int]bool)
	return rec(nums, cache, 0)
}

func rec(nums []int, cache map[int]bool, i int) bool {
	cached, found := cache[i]
	if found {
		return cached
	} else if i >= len(nums)-1 {
		return true
	} else {
		for j := i + 1; j <= i+nums[i]; j++ {
			if rec(nums, cache, j) {
				cache[i] = true
				return true
			}
		}
		cache[i] = false
		return false
	}
}
