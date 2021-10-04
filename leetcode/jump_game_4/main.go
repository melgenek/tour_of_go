package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}) == 3)
	fmt.Printf("%v\n", minJumps([]int{7, 6, 9, 6, 9, 6, 9, 7}) == 1)
	fmt.Printf("%v\n", minJumps([]int{7}) == 0)
	fmt.Printf("%v\n", minJumps([]int{2, 0, 2, 0}) == 2)
}

func minJumps(arr []int) int {
	g := make(map[int][]int)
	for i, v := range arr {
		g[v] = append(g[v], i)
	}
	n := len(arr)
	seen := make([]bool, len(arr))
	seen[0] = true

	dist := make([]int, len(arr))
	for i := range dist {
		dist[i] = 999999
	}
	dist[0] = 0

	q := list.New()
	q.PushFront(0)

	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)
		i := el.Value.(int)
		v := arr[i]

		adj := g[v]

		if i+1 < n {
			adj = append(adj, i+1)
		}
		if i-1 >= 0 {
			adj = append(adj, i-1)
		}
		for _, to := range adj {
			if !seen[to] {
				seen[to] = true
				dist[to] = dist[i] + 1
				q.PushBack(to)
			}
		}
		delete(g, v)
	}

	return dist[n-1]
}
