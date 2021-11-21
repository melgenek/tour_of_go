package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.5v\n", maxProbability(3, [][]int{
		{0, 1}, {1, 2}, {0, 2},
	}, []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Printf("%.5v\n", maxProbability(3, [][]int{
		{1, 0}, {2, 1}, {2, 0},
	}, []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Printf("%.5v\n", maxProbability(5, [][]int{
		{1, 4}, {2, 4}, {0, 4}, {0, 3}, {0, 2}, {2, 3},
	}, []float64{0.37, 0.17, 0.93, 0.23, 0.39, 0.04}, 3, 4))
}

const Inf = 1000000

type DijkstraEdge struct {
	dist float64
	to   int
}

type Edge struct {
	to     int
	weight float64
}

type PQ []DijkstraEdge

func (this PQ) Len() int           { return len(this) }
func (this PQ) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this PQ) Less(i, j int) bool { return this[i].dist < this[j].dist }
func (this *PQ) Pop() interface{} {
	last := (*this)[this.Len()-1]
	*this = (*this)[:this.Len()-1]
	return last
}
func (this *PQ) Push(v interface{}) {
	*this = append(*this, v.(DijkstraEdge))
}

// O(m*log(m))
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	for i, v := range succProb {
		succProb[i] = -math.Log2(v)
	}

	g := make([][]Edge, n)
	for i, edge := range edges {
		g[edge[0]] = append(g[edge[0]], Edge{edge[1], succProb[i]})
		g[edge[1]] = append(g[edge[1]], Edge{edge[0], succProb[i]})
	}

	dist := make([]float64, n)
	for i := range dist {
		dist[i] = Inf
	}
	dist[start] = 0

	pq := &PQ{}
	heap.Push(pq, DijkstraEdge{0, start})

	for pq.Len() > 0 {
		v := heap.Pop(pq).(DijkstraEdge)
		if dist[v.to] != v.dist {
			continue
		}
		from := v.to

		for _, adj := range g[from] {
			if dist[from]+adj.weight < dist[adj.to] {
				dist[adj.to] = dist[from] + adj.weight
				heap.Push(pq, DijkstraEdge{dist[adj.to], adj.to})
			}
		}
	}

	if dist[end] != Inf {
		return math.Pow(2, -dist[end])
	} else {
		return 0
	}
}

// O(n^2 +m)
func maxProbabilitySimpleDijkstra(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	for i, v := range succProb {
		succProb[i] = -math.Log2(v)
	}

	g := make([][]Edge, n)
	for i, edge := range edges {
		g[edge[0]] = append(g[edge[0]], Edge{edge[1], succProb[i]})
		g[edge[1]] = append(g[edge[1]], Edge{edge[0], succProb[i]})
	}

	dist := make([]float64, n)
	for i := range dist {
		dist[i] = Inf
	}
	dist[start] = 0
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		v := -1
		for j := 0; j < n; j++ {
			if !used[j] && (v == -1 || dist[j] < dist[v]) {
				v = j
			}
		}

		if dist[v] == Inf {
			break
		}

		used[v] = true
		for _, edge := range g[v] {
			if dist[edge.to] > dist[v]+edge.weight {
				dist[edge.to] = dist[v] + edge.weight
			}
		}
	}

	if dist[end] != Inf {
		return math.Pow(2, -dist[end])
	} else {
		return 0
	}
}

// O(n*m)
func maxProbabilityBellman(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	for i, v := range succProb {
		succProb[i] = -math.Log2(v)
	}

	dist := make([]float64, n)
	for i := range dist {
		dist[i] = Inf
	}
	dist[start] = 0

	for i := 0; i < n-1; i++ {
		for j, edge := range edges {
			if dist[edge[0]] < Inf {
				prob := succProb[j]
				if dist[edge[0]]+prob < dist[edge[1]] {
					dist[edge[1]] = dist[edge[0]] + prob
				}
			}
		}
		for j, edge := range edges {
			if dist[edge[1]] < Inf {
				prob := succProb[j]
				if dist[edge[1]]+prob < dist[edge[0]] {
					dist[edge[0]] = dist[edge[1]] + prob
				}
			}
		}
	}

	if dist[end] != Inf {
		return math.Pow(2, -dist[end])
	} else {
		return 0
	}
}
