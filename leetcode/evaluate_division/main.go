package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Printf("%v\n", calcEquation([][]string{
	//	{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"},
	//}, []float64{3.0, 4.0, 5.0, 6.0}, [][]string{
	//	{"x5", "x2"},
	//}))

	fmt.Printf("%v\n", calcEquation([][]string{
		{"a", "b"}, {"c", "d"},
	}, []float64{1.0, 1.0}, [][]string{
		{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"},
	}))
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := make(map[string][]Edge)
	for i, numbers := range equations {
		g[numbers[0]] = append(g[numbers[0]], Edge{numbers[1], values[i]})
		g[numbers[1]] = append(g[numbers[1]], Edge{numbers[0], 1.0 / values[i]})
	}
	n := len(g)
	results := make([]float64, len(queries))
	for queryIdx, query := range queries {
		start := query[0]
		end := query[1]
		if len(g[start]) == 0 || len(g[end]) == 0 {
			results[queryIdx] = -1.0
			continue
		}

		distances := make(map[string]float64, n)
		for k := range g {
			distances[k] = math.MaxFloat64
		}
		distances[start] = 1.0

		used := make(map[string]bool)

		for i := 0; i < n; i++ {
			v := ""
			minDist := math.MaxFloat64
			for node, dist := range distances {
				if !used[node] && minDist > dist {
					v = node
					minDist = dist
				}
			}

			used[v] = true

			if minDist == math.MaxFloat64 {
				break
			}

			for _, edge := range g[v] {
				if !used[edge.to] &&
					(distances[edge.to] == -1 || distances[edge.to] > distances[v]*edge.cost) {
					distances[edge.to] = distances[v] * edge.cost
				}
			}
		}
		if distances[end] == math.MaxFloat64 {
			results[queryIdx] = -1.0
		} else {
			results[queryIdx] = distances[end]
		}
	}

	return results
}

type Edge struct {
	to   string
	cost float64
}
