package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", canPartition([]int{1, 5, 11, 5}) == true)
	fmt.Printf("%v\n", canPartition([]int{1, 2, 3, 5}) == false)
	fmt.Printf("%v\n", false == canPartition([]int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97}))
	fmt.Printf("%v\n", false == canPartition([]int{83, 12, 15, 68, 83, 71, 72, 99, 66, 75, 53, 74, 30, 65, 95, 40, 22, 4, 67, 61, 55, 63, 85, 81, 67, 10, 93, 24, 24, 43, 29, 88, 94, 97, 27, 87, 51, 12, 26, 47, 10, 21, 16, 2, 8, 20, 94, 19, 66, 6, 13, 68, 27, 45, 90, 20, 47, 53, 71, 89, 75, 88, 88, 92, 12, 85, 22, 74, 82, 38, 2, 74, 21, 16, 29, 9, 9, 24, 23, 76, 24, 70, 64, 89, 78, 84, 76, 84, 95, 9, 75, 62, 94, 84, 48, 57, 82, 26, 47, 95}))
}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	half := sum / 2
	n := len(nums)
	cache := make([][]bool, n+1)
	for i := range cache {
		cache[i] = make([]bool, half+1)
	}
	cache[0][0] = true

	for i := 1; i <= n; i++ {
		num := nums[i-1]
		for j := 0; j <= half; j++ {
			cache[i][j] = cache[i-1][j]
			if j-num >= 0 {
				cache[i][j] = cache[i][j] || cache[i-1][j-num]
			}
		}
	}

	return cache[n][half]
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	half := sum / 2
	cache := make([][]int8, len(nums))
	for i := range cache {
		cache[i] = make([]int8, half)
	}
	return rec(nums, half, cache, 0, 0)
}

func rec(nums []int, half int, cache [][]int8, i int, sum int) bool {
	if len(nums) == i || sum > half {
		return false
	} else if half == sum {
		return true
	} else if cached := cache[i][sum]; cached != 0 {
		if cached == -1 {
			return false
		} else {
			return true
		}
	} else {
		res := rec(nums, half, cache, i+1, sum+nums[i]) || rec(nums, half, cache, i+1, sum)
		if res {
			cache[i][sum] = 1
		} else {
			cache[i][sum] = -1
		}
		return res
	}
}
