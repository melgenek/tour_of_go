package main

import "fmt"

func main() {
	fmt.Printf("%v\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}) == 2.5)
	fmt.Printf("%v\n", findMedianSortedArrays([]int{1, 3, 5, 7, 9}, []int{2, 4, 8, 10, 11}) == 6)
	fmt.Printf("%v\n", findMedianSortedArrays([]int{1, 3, 5, 7}, []int{2, 4, 8, 10}))
	fmt.Printf("%v\n", findMedianSortedArrays([]int{1, 3, 5, 7}, []int{2, 4, 8, 10}) == 4.5)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
}
