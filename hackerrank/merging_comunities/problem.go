package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	Solution(os.Stdin, os.Stdout)
}

type UnionSets struct {
	parents []int
	sizes   []int
}

func NewUnionSet(size int) *UnionSets {
	return &UnionSets{
		make([]int, size),
		make([]int, size),
	}
}

func (sets *UnionSets) make(v int) {
	sets.parents[v] = v
	sets.sizes[v] = 1
}

func (sets *UnionSets) find(v int) int {
	if sets.parents[v] == v {
		return v
	} else {
		sets.parents[v] = sets.find(sets.parents[v])
		return sets.parents[v]
	}
}

func (sets *UnionSets) union(a, b int) {
	a = sets.find(a)
	b = sets.find(b)

	if a != b {
		if sets.sizes[a] < sets.sizes[b] {
			a, b = b, a
		}
		sets.parents[b] = a
		sets.sizes[a] += sets.sizes[b]
	}
}

func (sets *UnionSets) size(v int) int {
	return sets.sizes[sets.find(v)]
}

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	sets := NewUnionSet(n)
	for i := 0; i < n; i++ {
		sets.make(i)
	}

	in.Scan()
	queriesCount, _ := strconv.Atoi(in.Text())

	for i := 0; i < queriesCount; i++ {
		in.Scan()
		operation := in.Text()

		if operation == "M" {
			in.Scan()
			a, _ := strconv.Atoi(in.Text())
			in.Scan()
			b, _ := strconv.Atoi(in.Text())

			sets.union(a-1, b-1)
		} else {
			in.Scan()
			v, _ := strconv.Atoi(in.Text())

			size := sets.size(v - 1)

			if i != queriesCount-1 {
				fmt.Fprintf(out, "%d\n", size)
			} else {
				fmt.Fprintf(out, "%d", size)
			}
		}
	}
}
