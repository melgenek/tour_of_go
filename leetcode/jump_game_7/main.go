package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", canReach("011001110001000", 3, 5) == true)
	fmt.Printf("%v\n", canReach("000100100010", 1, 2) == true)
	fmt.Printf("%v\n", canReach("00111010", 3, 5) == false)
	fmt.Printf("%v\n", canReach("000", 1, 1) == true)
	fmt.Printf("%v\n", canReach("0000000000", 1, 1) == true)
	fmt.Printf("%v\n", canReach("0000000000", 2, 5) == true)
	fmt.Printf("%v\n", canReach("011010", 2, 3) == true)
	fmt.Printf("%v\n", canReach("011011", 2, 3) == false)
	fmt.Printf("%v\n", canReach("01101110", 2, 3) == false)
	fmt.Printf("%v\n", canReach("011111111111111111111111111110", 1, 100000) == true)
	fmt.Printf("%v\n", canReach("011011111101111111110111101100111111110011111111101110", 1, 10) == true)

}

func canReach(s string, minJump int, maxJump int) bool {

	return false
}
