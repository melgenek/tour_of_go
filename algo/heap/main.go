package main

import "fmt"

type Heap struct {
	values []int
	last   int
}

func NewHeap() *Heap {
	return &Heap{
		make([]int, 1),
		-1,
	}
}

func (this *Heap) Add(v int) {
	this.last++
	if this.last >= len(this.values) {
		newArr := make([]int, len(this.values)*2)
		copy(newArr, this.values)
		this.values = newArr
	}
	this.values[this.last] = v

	for i := this.last; i >= this.last/2; {
		parent := (i - 1) / 2
		if this.values[parent] > this.values[i] {
			this.values[parent], this.values[i] = this.values[i], this.values[parent]
			i = parent
		} else {
			break
		}
	}
}

func (this *Heap) Pop() int {
	if this.last == -1 {
		return 9999999
	} else {
		result := this.values[0]
		this.values[0], this.values[this.last] = this.values[this.last], this.values[0]
		this.last--

		for i := 0; i < this.last; {
			childI := i*2 + 1
			child := 999999
			if childI <= this.last {
				child = this.values[childI]
				if childI+1 <= this.last && this.values[childI+1] < child {
					child = this.values[childI+1]
					childI++
				}
			}

			if child < this.values[i] {
				this.values[childI], this.values[i] = this.values[i], this.values[childI]
				i = childI
			} else {
				break
			}
		}
		return result
	}
}

func (this *Heap) Peek() int {
	if this.last == -1 {
		return 9999999
	} else {
		return this.values[0]
	}
}

func main() {
	h := NewHeap()
	fmt.Printf("%v\n", h.Peek() == 9999999)
	h.Add(10)
	fmt.Printf("%v\n", h.Peek() == 10)
	h.Add(1)
	fmt.Printf("%v\n", h.Peek() == 1)
	h.Add(9)
	fmt.Printf("%v\n", h.Peek() == 1)
	h.Add(12)
	fmt.Printf("%v\n", h.Peek() == 1)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", h.Pop() == 1)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", h.Peek() == 9)
	fmt.Printf("%v\n", h.Pop() == 9)
	fmt.Printf("%v\n", h)
	h.Add(3)
	fmt.Printf("%v\n", h.Pop() == 3)
	fmt.Printf("%v\n", h.Peek())
	fmt.Printf("%v\n", h.Peek() == 10)
	fmt.Printf("%v\n", h.Pop() == 10)
	fmt.Printf("%v\n", h.Pop() == 12)
	fmt.Printf("%v\n", h.Pop() == 9999999)
}
