package main

import "fmt"

func main() {
	fmt.Printf("%v\n", divisorGame(1) == false)
	fmt.Printf("%v\n", divisorGame(2) == true)
	fmt.Printf("%v\n", divisorGame(3) == false)
	fmt.Printf("%v\n", divisorGame(4) == true)
}

func divisorGame(n int) bool {
	if n == 1 {
		return false
	} else {
		outcomes := make([]bool, n+1)
		outcomes[2] = true

		for i := 3; i <= n; i++ {
			for k := i / 2; k >= 1; k-- {
				if i%k == 0 && !outcomes[i-k] {
					outcomes[i] = true
					continue
				}
			}
		}

		return outcomes[n]
	}
}
