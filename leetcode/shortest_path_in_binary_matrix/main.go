package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", shortestPathBinaryMatrix([][]int{
		{0, 1},
		{1, 0},
	}) == 2)
	fmt.Printf("%v\n", shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}) == 4)
	fmt.Printf("%v\n", shortestPathBinaryMatrix([][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}) == -1)

	fmt.Printf("%v\n", shortestPathBinaryMatrix([][]int{
		{0, 1, 1},
		{1, 1, 1},
		{1, 1, 0},
	}) == -1)
}

type Point struct {
	x int
	y int
}

type Direction struct {
	dx int
	dy int
}

var directions = []Direction{
	{0, 1},
	{1, 0},
	{1, 1},
	{0, -1},
	{-1, 1},
	{-1, -1},
	{-1, 0},
	{1, -1},
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	distances := make([][]int, len(grid))
	for i := range distances {
		distances[i] = make([]int, len(grid[i]))
		for j := range distances[i] {
			distances[i][j] = 9999999
		}
	}
	distances[0][0] = 1

	q := list.New()
	if grid[0][0] == 0 {
		q.PushBack(Point{0, 0})
	}

	for q.Len() > 0 {
		el := q.Front()
		q.Remove(el)

		point := el.Value.(Point)
		currentDistance := distances[point.x][point.y]

		for _, direction := range directions {
			nextPoint := Point{point.x + direction.dx, point.y + direction.dy}

			if nextPoint.x >= 0 && nextPoint.x < n &&
				nextPoint.y >= 0 && nextPoint.y < n &&
				grid[nextPoint.x][nextPoint.y] == 0 &&
				currentDistance+1 < distances[nextPoint.x][nextPoint.y] {
				distances[nextPoint.x][nextPoint.y] = currentDistance + 1
				q.PushBack(nextPoint)
			}
		}
	}

	if distances[n-1][n-1] == 9999999 {
		return -1
	} else {
		return distances[n-1][n-1]
	}

}
