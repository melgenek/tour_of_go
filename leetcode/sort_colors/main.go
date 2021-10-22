package main

import "fmt"

func main() {
	arr1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr1)
	fmt.Printf("%v\n", arr1)

	arr2 := []int{2, 0, 1}
	sortColors(arr2)
	fmt.Printf("%v\n", arr2)

	arr3 := []int{0}
	sortColors(arr3)
	fmt.Printf("%v\n", arr3)
}

func sortColors(nums []int) {
	red := 0
	white := 0
	blue := 0
	for _, v := range nums {
		if v == 0 {
			red++
		}
		if v == 1 {
			white++
		}
		if v == 2 {
			blue++
		}
	}

	for i := 0; i < red; i++ {
		nums[i] = 0
	}
	for i := red; i < red+white; i++ {
		nums[i] = 1
	}
	for i := red + white; i < red+white+blue; i++ {
		nums[i] = 2
	}
}
