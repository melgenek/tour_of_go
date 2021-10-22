package main

import "fmt"

func main() {
	fmt.Printf("%v\n", rangeBitwiseAnd(5, 7) == 4)
	fmt.Printf("%v\n", rangeBitwiseAnd(0, 0) == 0)
	fmt.Printf("%v\n", rangeBitwiseAnd(1, 2147483647) == 0)
}

func rangeBitwiseAnd(left int, right int) int {
	var res = ^0
	for i := left; i <= right; i++ {
		res = res & i
	}
	return res
}
