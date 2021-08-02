package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	Solution(os.Stdin, os.Stdout)
}

type Edge struct {
	a      int
	b      int
	weight int
}

type Query struct {
	from int
	to   int
}

type Sets struct {
	parents    []int
	sizes      []int
	edgesCount int
}

func NewSets(n int) Sets {
	sets := Sets{
		make([]int, n),
		make([]int, n),
		0,
	}
	for i := 0; i < n; i++ {
		sets.parents[i] = i
		sets.sizes[i] = 1
	}
	return sets
}

func (sets *Sets) get(v int) int {
	if sets.parents[v] == v {
		return v
	} else {
		sets.parents[v] = sets.get(sets.parents[v])
		return sets.parents[v]
	}
}

func (sets *Sets) union(a int, b int) int {
	a = sets.get(a)
	b = sets.get(b)

	if a != b {
		if sets.sizes[a] < sets.sizes[b] {
			a, b = b, a
		}
		sets.edgesCount += sets.sizes[a] * sets.sizes[b]
		sets.parents[b] = a
		sets.sizes[a] += sets.sizes[b]
	}

	return sets.edgesCount
}

//
//func options(n int) int {
//	return n * (n - 1) / 2
//}

type WeightedCount struct {
	weight int
	count  int
}

func solve(n int, edges []Edge, queries []Query) []int {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	thresholds := make([]WeightedCount, len(edges))
	sets := NewSets(n)
	for i, edge := range edges {
		count := sets.union(edge.a, edge.b)
		thresholds[i] = WeightedCount{edge.weight, count}
	}

	answers := make([]int, len(queries))
	for i, query := range queries {
		toValue := count(thresholds, query.to)
		fromValue := count(thresholds, query.from-1)
		answers[i] = toValue - fromValue
	}

	return answers
}

func count(counts []WeightedCount, targetWeight int) int {
	upperBound := sort.Search(len(counts), func(i int) bool {
		return counts[i].weight > targetWeight
	})
	if upperBound == 0 {
		return 0
	}
	return counts[upperBound-1].count
}

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	in.Scan()
	queriesCount, _ := strconv.Atoi(in.Text())

	edges := make([]Edge, n-1)
	for i := 0; i < n-1; i++ {
		in.Scan()
		a, _ := strconv.Atoi(in.Text())
		in.Scan()
		b, _ := strconv.Atoi(in.Text())
		in.Scan()
		weight, _ := strconv.Atoi(in.Text())

		edges[i] = Edge{a - 1, b - 1, weight}
	}

	queries := make([]Query, queriesCount)
	for i := 0; i < queriesCount; i++ {
		in.Scan()
		from, _ := strconv.Atoi(in.Text())
		in.Scan()
		to, _ := strconv.Atoi(in.Text())

		queries[i] = Query{from, to}
	}

	result := solve(n, edges, queries)

	for i, value := range result {
		if i != len(result)-1 {
			fmt.Fprintf(out, "%d\n", value)
		} else {
			fmt.Fprintf(out, "%d", value)
		}
	}
}
