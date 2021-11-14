package main

import "fmt"

func main() {
	fmt.Printf("%v\n", 4 == longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Printf("%v\n", 9 == longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Printf("%v\n", 1 == longestConsecutive([]int{0}))
	fmt.Printf("%v\n", 4 == longestConsecutive([]int{-7, -1, 3, -9, -4, 7, -3, 2, 4, 9, 4, -9, 8, -7, 5, -1, -7}))
}

type Set struct {
	m map[int][]int
}

func NewSet() *Set {
	return &Set{
		make(map[int][]int),
	}
}

func (this *Set) Add(v int) {
	if this.m[v] == nil {
		this.m[v] = []int{v, v}
	}
}

func (this *Set) Find(v int) *int {
	if this.m[v] == nil {
		return nil
	} else if this.m[v][0] == v {
		return &v
	} else {
		res := this.Find(this.m[v][0])
		this.m[v][0] = *res
		return res
	}
}

func (this *Set) Union(a, b int) int {
	aKey := this.Find(a)
	bKey := this.Find(b)
	if aKey == nil || bKey == nil {
		return 0
	}
	a = *aKey
	b = *bKey

	aRange := this.m[a]
	bRange := this.m[b]

	if aRange[1]+1 == bRange[0] {
		this.m[b][0] = a
		this.m[a][1] = bRange[1]
	} else if bRange[1]+1 == aRange[0] {
		this.m[a][0] = b
		this.m[b][1] = aRange[1]
	}

	return max(aRange[1]-aRange[0], bRange[1]-bRange[0]) + 1
}

func longestConsecutive(nums []int) int {
	s := NewSet()
	res := 0
	for _, v := range nums {
		s.Add(v)
		res = max(res, 1)
		res = max(res, s.Union(v, v-1))
		res = max(res, s.Union(v, v+1))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
