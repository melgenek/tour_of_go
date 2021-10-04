package main

import "fmt"

func main() {
	fmt.Printf("%v\n", canReach([]int{4, 2, 3, 0, 3, 1, 2}, 5) == true)
	fmt.Printf("%v\n", canReach([]int{4, 2, 3, 0, 3, 1, 2}, 0) == true)
	fmt.Printf("%v\n", canReach([]int{3, 0, 2, 1, 2}, 2) == false)
}

func canReach(arr []int, start int) bool {
	seen := make([]bool, len(arr))
	return dfs(arr, seen, start)
}

func dfs(arr []int, seen []bool, i int) bool {
	if arr[i] == 0 {
		return true
	} else {
		seen[i] = true

		res := false
		if i-arr[i] >= 0 && !seen[i-arr[i]] {
			res = res || dfs(arr, seen, i-arr[i])
		}

		if i+arr[i] < len(arr) && !seen[i+arr[i]] {
			res = res || dfs(arr, seen, i+arr[i])
		}
		return res
	}
}
