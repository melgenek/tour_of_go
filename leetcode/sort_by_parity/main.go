package main

import "fmt"

func main() {
	fmt.Printf("%v\n", sortArrayByParityII([]int{4, 2, 5, 7}))
	fmt.Printf("%v\n", sortArrayByParityII([]int{2, 3}))
	fmt.Printf("%v\n", sortArrayByParityII([]int{3, 1, 4, 2}))
	fmt.Printf("%v\n", sortArrayByParityII([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", sortArrayByParityII([]int{888, 505, 627, 846}))
	fmt.Printf("%v\n", sortArrayByParityII([]int{3, 1, 4, 2}))
}

func sortArrayByParityII(nums []int) []int {
	nextOdd := 0
	nextEven := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 && nums[i]%2 != 0 {
			for ; nums[nextEven]%2 != 0; nextEven++ {
			}
			nums[i], nums[nextEven] = nums[nextEven], nums[i]
			nextOdd = i + 1
		} else if i%2 != 0 && nums[i]%2 == 0 {
			for ; nums[nextOdd]%2 == 0; nextOdd++ {
			}
			nums[i], nums[nextOdd] = nums[nextOdd], nums[i]
			nextEven = i + 1
		} else {
			nextOdd++
			nextEven++
		}
	}

	return nums
}
