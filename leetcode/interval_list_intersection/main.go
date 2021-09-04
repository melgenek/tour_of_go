package main

import "fmt"

func main() {
	fmt.Printf("%v\n", intervalIntersection2([][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	fmt.Printf("%v\n", intervalIntersection2([][]int{{1, 3}, {5, 9}}, [][]int{}))
	fmt.Printf("%v\n", intervalIntersection2([][]int{}, [][]int{{4, 8}, {10, 12}}))
	fmt.Printf("%v\n", intervalIntersection2([][]int{}, [][]int{}))
	fmt.Printf("%v\n", intervalIntersection2([][]int{{1, 7}}, [][]int{{3, 10}}))
	fmt.Printf("%v\n", intervalIntersection2([][]int{{1, 7}, {8, 10}}, [][]int{{9, 10}}))
}

//[][]int{{0,  2}, {5,     10},   {13,   23}, {24, 25}},
//[][]int{  {1,      5}, {8,   12},  {15,   24}, {25, 26}}

func intervalIntersection2(first [][]int, second [][]int) [][]int {
	res := make([][]int, 0)

	for i, j := 0, 0; i < len(first) && j < len(second); {
		f := first[i]
		s := second[j]

		start := max(f[0], s[0])
		end := min(f[1], s[1])

		if start <= end {
			res = append(res, []int{start, end})
		}

		if f[1] < s[1] {
			i++
		} else {
			j++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func intervalIntersection(ii [][]int, jj [][]int) [][]int {
	f := 0
	s := 0

	res := make([][]int, 0)

	for f < len(ii) && s < len(jj) {
		i := ii[f]
		j := jj[s]

		if i[0] == j[0] && i[1] == j[1] {
			res = append(res, []int{i[0], i[1]})
			f++
			s++
		} else if i[0] < j[0] {
			if i[1] < j[0] {
				f++
			} else {
				if i[1] < j[1] {
					res = append(res, []int{j[0], i[1]})
					f++
				} else {
					res = append(res, []int{j[0], j[1]})
					s++
				}
			}
		} else {
			if j[1] < i[0] {
				s++
			} else {
				if i[1] < j[1] {
					res = append(res, []int{i[0], i[1]})
					f++
				} else {
					res = append(res, []int{i[0], j[1]})
					s++
				}
			}
		}
	}

	return res
}
