package main

import "fmt"

func main() {
	fmt.Printf("%v\n", nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1

	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {
		c2 := nums[i2] * 2
		c3 := nums[i3] * 3
		c5 := nums[i5] * 5
		minC := min(c2, min(c3, c5))

		nums[i] = minC
		if minC == c2 {
			i2++
		}
		if minC == c3 {
			i3++
		}
		if minC == c5 {
			i5++
		}
	}

	return nums[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
