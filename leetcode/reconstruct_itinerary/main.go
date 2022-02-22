package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v\n", findItinerary([][]string{
		{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"},
	}))
	fmt.Printf("%v\n", findItinerary([][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
	}))
}

func findItinerary(tickets [][]string) []string {
	g := make(map[string][]string)
	for _, edge := range tickets {
		g[edge[0]] = append(g[edge[0]], edge[1])
	}

	for k := range g {
		sort.Sort(sort.Reverse(sort.StringSlice(g[k])))
	}

	var rec func(string)

	result := make([]string, 0)
	rec = func(current string) {
		for len(g[current]) > 0 {
			next := g[current][len(g[current])-1]
			g[current] = g[current][:len(g[current])-1]
			rec(next)
		}
		result = append(result, current)
	}

	rec("JFK")

	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result
}
