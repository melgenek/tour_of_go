package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := Constructor()

	fmt.Printf("%v\n", s.Insert(0))
	fmt.Printf("%v\n", s.Insert(1))
	fmt.Printf("%v\n", s.Remove(0))
	fmt.Printf("%v\n", s.Insert(2))
	fmt.Printf("%v\n", s.Remove(1))
	fmt.Printf("%v\n", s.GetRandom())
}

type RandomizedSet struct {
	m    map[int]int
	last int
	arr  []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		make(map[int]int),
		-1,
		make([]int, 200000),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, found := this.m[val]
	if !found {
		this.last++
		this.arr[this.last] = val
		this.m[val] = this.last
	}
	return !found
}

func (this *RandomizedSet) Remove(val int) bool {
	old, found := this.m[val]
	if found {
		replaced := this.arr[this.last]
		this.arr[old], this.arr[this.last] = this.arr[this.last], this.arr[old]
		this.last--
		this.m[replaced] = old
		delete(this.m, val)
	}

	return found
}

func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(this.last + 1)
	return this.arr[idx]
}
