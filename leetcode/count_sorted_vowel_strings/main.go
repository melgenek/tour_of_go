package main

import "fmt"

func main() {
	fmt.Printf("%v\n", countVowelStrings(1) == 5)
	fmt.Printf("%v\n", countVowelStrings(2) == 15)
	fmt.Printf("%v\n", countVowelStrings(3) == 35)
	fmt.Printf("%v\n", countVowelStrings(4) == 70)
	fmt.Printf("%v\n", countVowelStrings(33) == 66045)
	fmt.Printf("%v\n", countVowelStrings(50) == 316251)
}

func countVowelStrings(n int) int {
	arr := []int{5, 4, 3, 2, 1}
	for i := 2; i <= n; i++ {
		newLevel := []int{}
		for j := 0; j < len(arr); j++ {
			for k := 1; k <= arr[j]; k++ {
				newLevel = append(newLevel, k)
			}
		}
		arr = newLevel
	}

	return len(arr)
}
