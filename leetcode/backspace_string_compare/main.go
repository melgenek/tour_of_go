package main

import "fmt"

func main() {
	fmt.Printf("%v\n", backspaceCompare2("bxj##tw", "bxo#j##tw") == true)
	fmt.Printf("%v\n", backspaceCompare2("ab#c", "ad#c") == true)
	fmt.Printf("%v\n", backspaceCompare2("a#", "a") == false)
	fmt.Printf("%v\n", backspaceCompare2("ab##", "a#b#") == true)
	fmt.Printf("%v\n", backspaceCompare2("a##c", "#a#c") == true)
	fmt.Printf("%v\n", backspaceCompare2("a#c", "c") == true)
	fmt.Printf("%v\n", backspaceCompare2("a#c", "b") == false)
	fmt.Printf("%v\n", backspaceCompare2("b", "a#c") == false)
	fmt.Printf("%v\n", backspaceCompare2("b", "a#c") == false)
}

func backspaceCompare2(s1 string, s2 string) bool {
	j1, j2 := len(s1)-1, len(s2)-1

	skip1 := 0
	skip2 := 0
	for j1 >= 0 && j2 >= 0 {
		for ; j1 >= 0; j1-- {
			if s1[j1] == '#' {
				skip1++
			} else if skip1 > 0 {
				skip1--
			} else {
				break
			}
		}

		for ; j2 >= 0; j2-- {
			if s2[j2] == '#' {
				skip2++
			} else if skip2 > 0 {
				skip2--
			} else {
				break
			}
		}

		if j1 >= 0 && j2 >= 0 && s1[j1] != s2[j2] {
			return false
		}

		j1--
		j2--
	}

	for ; j1 >= 0; j1-- {
		if s1[j1] == '#' {
			skip1++
		} else if skip1 > 0 {
			skip1--
		} else {
			break
		}
	}

	for ; j2 >= 0; j2-- {
		if s2[j2] == '#' {
			skip2++
		} else if skip2 > 0 {
			skip2--
		} else {
			break
		}
	}

	return j1 == j2
}

func backspaceCompare(s1 string, s2 string) bool {
	a1 := make([]uint8, len(s1))
	a2 := make([]uint8, len(s2))

	j1 := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != '#' {
			a1[j1] = s1[i]
			j1++
		} else {
			j1--
			if j1 < 0 {
				j1 = 0
			}
		}
	}

	j2 := 0
	for i := 0; i < len(s2); i++ {
		if s2[i] != '#' {
			a2[j2] = s2[i]
			j2++
		} else {
			j2--
			if j2 < 0 {
				j2 = 0
			}
		}
	}

	if j1 != j2 {
		return false
	} else {
		for i := 0; i < j1; i++ {
			if a1[i] != a2[i] {
				return false
			}
		}
		return true
	}

}
