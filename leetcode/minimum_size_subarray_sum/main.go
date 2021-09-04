package main

import "fmt"

func main() {
	fmt.Printf("%v\n", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}) == 2)
	fmt.Printf("%v\n", minSubArrayLen(4, []int{1, 4, 4}) == 1)
	fmt.Printf("%v\n", minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}) == 0)
	fmt.Printf("%v\n", minSubArrayLen(4, []int{1, 3, 4}) == 1)
}

const aLot = 999999

func minSubArrayLen(target int, nums []int) int {
	res := aLot

	sum := 0
	start, end := 0, 0
	for end < len(nums) {
		if sum >= target {
			if end-start < res {
				res = end - start
			}
			sum -= nums[start]
			start++
		} else {
			sum += nums[end]
			end++
		}
	}

	for ; start < end && sum >= target; start++ {
		if end-start < res {
			res = end - start
		}
		sum -= nums[start]
	}

	if res == aLot {
		return 0
	} else {
		return res
	}
}
