package main

import "fmt"

func main() {
	fmt.Printf("%v\n", sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Printf("%v\n", sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}) == 58)
	fmt.Printf("%v\n", sumOddLengthSubarrays([]int{1, 2}) == 3)
	fmt.Printf("%v\n", sumOddLengthSubarrays([]int{10, 11, 12}) == 66)
}

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if (j-i)%2 == 1 {
				sum += prefix[j] - prefix[i]
			}
		}
	}

	return sum
}
