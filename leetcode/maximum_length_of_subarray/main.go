package main

import "fmt"

func main() {
	fmt.Printf("%v\n", getMaxLen([]int{1, -2, -3, 4}) == 4)
	fmt.Printf("%v\n", getMaxLen([]int{0, 1, -2, -3, -4}) == 3)
	fmt.Printf("%v\n", getMaxLen([]int{0, 1, -2, -3, 4}) == 4)
	fmt.Printf("%v\n", getMaxLen([]int{0, 1, -2, -3, -4, 1, 2, 3}) == 5)
	fmt.Printf("%v\n", getMaxLen([]int{-1, -2, -3, 0, 1}) == 2)
	fmt.Printf("%v\n", getMaxLen([]int{-1, 2}) == 1)
	fmt.Printf("%v\n", getMaxLen([]int{1, 2, 3, 5, -6, 4, 0, 10}) == 4)
	fmt.Printf("%v\n", getMaxLen([]int{1, 2, 3, -6, 3, 4, 5, 4, 0, 10}) == 4)
	fmt.Printf("%v\n", getMaxLen([]int{1}) == 1)
	fmt.Printf("%v\n", getMaxLen([]int{1, 2}) == 2)
	fmt.Printf("%v\n", getMaxLen([]int{0}) == 0)
}

func getMaxLen(nums []int) int {
	max := 0

	negatives := 0
	firstNegative := -1
	lastNegative := -1

	lastZero := -1
	for i := 0; i <= len(nums); i++ {
		if i == len(nums) || nums[i] == 0 {
			if negatives%2 == 0 {
				newMax := i - lastZero - 1
				if newMax > max {
					max = newMax
				}
			} else {
				newMaxPrefix := lastNegative - lastZero - 1
				newMaxSuffix := i - firstNegative - 1
				if newMaxPrefix > max {
					max = newMaxPrefix
				}
				if newMaxSuffix > max {
					max = newMaxSuffix
				}
			}
			lastZero = i
			negatives = 0
			firstNegative = -1
			lastNegative = -1
		} else if nums[i] < 0 {
			negatives++
			if firstNegative == -1 {
				firstNegative = i
				lastNegative = i
			} else {
				lastNegative = i
			}
		}
	}

	return max
}
