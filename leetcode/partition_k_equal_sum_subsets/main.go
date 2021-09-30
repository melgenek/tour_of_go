package main

import "fmt"

func main() {
	fmt.Printf("%v\n", true == canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Printf("%v\n", false == canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2}, 4))
	fmt.Printf("%v\n", false == canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	fmt.Printf("%v\n", false == canPartitionKSubsets([]int{1, 1, 1, 1}, 3))
	fmt.Printf("%v\n", true == canPartitionKSubsets([]int{1, 1, 1}, 3))
	fmt.Printf("%v\n", true == canPartitionKSubsets([]int{815, 625, 3889, 4471, 60, 494, 944, 1118, 4623, 497, 771, 679, 1240, 202, 601, 883}, 3))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	kSum := sum / k

	used := make([]bool, len(nums))

	return m(nums, used, 0, 0, kSum, k)
}

func m(nums []int, used []bool, i int, current int, target int, k int) bool {
	if k == 0 {
		return true
	} else if current == target {
		return m(nums, used, 0, 0, target, k-1)
	} else if current > target {
		return false
	} else {
		for j := i; j < len(nums); j++ {
			if !used[j] {
				used[j] = true
				if m(nums, used, j+1, current+nums[j], target, k) {
					return true
				}
				used[j] = false
			}
		}
		return false
	}
}
