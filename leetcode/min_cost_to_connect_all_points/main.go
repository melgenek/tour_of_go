package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Printf("%v\n", 20 == minCostConnectPoints([][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}))
	fmt.Printf("%v\n", 0 == minCostConnectPoints([][]int{
		{0, 0},
	}))
	fmt.Printf("%v\n", 4000000 == minCostConnectPoints([][]int{
		{-1000000, -1000000}, {1000000, 1000000},
	}))
	fmt.Printf("%v\n", 4 == minCostConnectPoints([][]int{
		{0, 0}, {1, 1}, {1, 0}, {-1, 1},
	}))
	fmt.Printf("%v\n", 18 == minCostConnectPoints([][]int{
		{3, 12}, {-2, 5}, {-4, 1},
	}))
}

func minCostConnectPoints(points [][]int) int {
	return primDense(points)
}

type Edge struct {
	i    int
	j    int
	dist int
}

func kruskalSimple(points [][]int) int {
	n := len(points)

	edges := make([]Edge, n*(n-1)/2)
	k := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges[k] = Edge{i, j, manhattan(points[i], points[j])}
			k++
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	trees := make([]int, n)
	for i := range trees {
		trees[i] = i
	}

	cost := 0
	for _, edge := range edges {
		oldId := trees[edge.i]
		newId := trees[edge.j]

		if oldId != newId {
			cost += edge.dist

			for i, tree := range trees {
				if tree == oldId {
					trees[i] = newId
				}
			}
		}
	}
	return cost
}

type UnionSet struct {
	parents []int
	sizes   []int
}

func NewUnionSet(n int) *UnionSet {
	parents := make([]int, n)
	sizes := make([]int, n)
	for i := range parents {
		parents[i] = i
		sizes[i] = 1
	}
	return &UnionSet{
		parents,
		sizes,
	}
}

func (this *UnionSet) Find(v int) int {
	if this.parents[v] == v {
		return v
	} else {
		res := this.Find(this.parents[v])
		this.parents[v] = res
		return res
	}
}

func (this *UnionSet) Union(a, b int) {
	a = this.Find(a)
	b = this.Find(b)
	if a == b {
		return
	}

	aSize := this.sizes[a]
	bSize := this.sizes[b]

	if aSize > bSize {
		this.parents[b] = a
		this.sizes[a] += this.sizes[b]
	} else {
		this.parents[a] = b
		this.sizes[b] += this.sizes[a]
	}
}

func kruskalUnionSet(points [][]int) int {
	n := len(points)

	edges := make([]Edge, n*(n-1)/2)
	k := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges[k] = Edge{i, j, manhattan(points[i], points[j])}
			k++
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	trees := NewUnionSet(n)

	cost := 0
	for _, edge := range edges {
		if trees.Find(edge.i) != trees.Find(edge.j) {
			cost += edge.dist
			trees.Union(edge.i, edge.j)
		}
	}
	return cost
}

const Inf = math.MaxInt32

func primDense(points [][]int) int {
	n := len(points)

	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				g[i][j] = manhattan(points[i], points[j])
			} else {
				g[i][j] = Inf
			}
		}
	}

	selected := make([]bool, n)
	minEdges := make([]int, n)
	for i := 1; i < n; i++ {
		minEdges[i] = Inf
	}

	cost := 0
	for i := 0; i < n; i++ {
		v := -1
		for j := 0; j < n; j++ {
			if !selected[j] && (v == -1 || minEdges[j] < minEdges[v]) {
				v = j
			}
		}

		selected[v] = true
		cost += minEdges[v]

		for to := 0; to < n; to++ {
			if g[v][to] < minEdges[to] {
				minEdges[to] = g[v][to]
			}
		}
	}
	return cost
}

func manhattan(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
