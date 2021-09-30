package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", matrixRankTransform([][]int{
		{1, 2},
		{3, 4},
	}))
	fmt.Printf("%v\n", matrixRankTransform([][]int{
		{7, 7},
		{7, 7},
	}))
	fmt.Printf("%v\n", matrixRankTransform([][]int{
		{20, -21, 14},
		{-19, 4, 19},
		{22, -47, 24},
		{-19, 4, 19},
	}))
	fmt.Printf("%v\n", matrixRankTransform([][]int{
		{7, 3, 6},
		{1, 4, 5},
		{9, 8, 2},
	}))
	fmt.Printf("%v\n", matrixRankTransform([][]int{
		{-37, -50, -3, 44},
		{-37, 46, 13, -32},
		{47, -42, -3, -40},
		{-17, -22, -39, 24},
	}))
}

type UnionSet struct {
	parents []int
	sizes   []int
}

func NewUnionSet(n int) UnionSet {
	parents := make([]int, n)
	sizes := make([]int, n)
	for i := range parents {
		parents[i] = -1
	}

	return UnionSet{
		parents,
		sizes,
	}
}

func (s *UnionSet) make(v int) {
	s.parents[v] = v
	s.sizes[v] = 1
}

func (s *UnionSet) find(v int) int {
	if s.parents[v] == v {
		return v
	} else {
		res := s.find(s.parents[v])
		s.parents[v] = res
		return res
	}
}

func (s *UnionSet) union(v1, v2 int) {
	a := s.find(v1)
	b := s.find(v2)

	if s.sizes[a] < s.sizes[b] {
		a, b = b, a
	}
	s.parents[b] = a
	s.sizes[a] += s.sizes[b]
}

func (s *UnionSet) groups() map[int][]int {
	res := make(map[int][]int)
	for i, v := range s.parents {
		if v != -1 {
			group := s.find(i)
			res[group] = append(res[group], i)
		}
	}
	return res
}

type Cell struct {
	x     int
	y     int
	value int
}

func matrixRankTransform(matrix [][]int) [][]int {
	rowsCount := len(matrix)
	columnsCount := len(matrix[0])

	cells := make([]Cell, rowsCount*columnsCount)
	res := make([][]int, rowsCount)

	ranks := make([]int, rowsCount+columnsCount)

	k := 0
	for i, row := range matrix {
		res[i] = make([]int, len(row))
		for j, v := range row {
			cells[k] = Cell{i, j, v}
			k++
		}
	}

	sort.Slice(cells, func(i, j int) bool {
		return cells[i].value < cells[j].value
	})

	for i := 0; i < len(cells); {
		current := cells[i]
		valueCells := []Cell{}
		for ; i < len(cells) && current.value == cells[i].value; i++ {
			valueCells = append(valueCells, cells[i])
		}

		sets := NewUnionSet(rowsCount + columnsCount)
		for _, cell := range valueCells {
			sets.make(cell.x)
			sets.make(cell.y + rowsCount)
		}
		for _, cell := range valueCells {
			sets.union(cell.x, cell.y+rowsCount)
		}

		for _, group := range sets.groups() {
			maxRank := 0
			for _, rankIdx := range group {
				maxRank = max(maxRank, ranks[rankIdx])
			}
			for _, rankIdx := range group {
				ranks[rankIdx] = maxRank + 1
			}
		}

		for _, cell := range valueCells {
			res[cell.x][cell.y] = ranks[cell.x]
		}
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
