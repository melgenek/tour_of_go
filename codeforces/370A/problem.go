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

type Point struct {
	X int
	Y int
}

const Inf = 99999

var (
	rookDirections = []Point{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	bishopDirections = []Point{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
	kingDirections = []Point{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}
)

func Solution(input io.Reader, output io.Writer) {
	for i := 1; i <= 8; i++ {
		for _, direction := range rookDirections {
			rookDirections = append(rookDirections,
				Point{
					direction.X * i,
					direction.Y * i,
				})
		}
	}

	for i := 1; i <= 8; i++ {
		for _, direction := range bishopDirections {
			bishopDirections = append(bishopDirections,
				Point{
					direction.X * i,
					direction.Y * i,
				})
		}
	}

	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	r1, _ := strconv.Atoi(in.Text())
	in.Scan()
	c1, _ := strconv.Atoi(in.Text())
	in.Scan()
	r2, _ := strconv.Atoi(in.Text())
	in.Scan()
	c2, _ := strconv.Atoi(in.Text())

	rookTurns := dfs(rookDirections, Point{r1, c1}, Point{r2, c2})
	bishopTurns := dfs(bishopDirections, Point{r1, c1}, Point{r2, c2})
	kingTurns := dfs(kingDirections, Point{r1, c1}, Point{r2, c2})

	fmt.Fprintf(out, "%d %d %d", rookTurns, bishopTurns, kingTurns)

}

func dfs(directions []Point, start Point, target Point) int {
	var inner func(point Point)

	distance := make([][]int, 9)
	for i := 0; i < 9; i++ {
		distance[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			distance[i][j] = Inf
		}
	}
	distance[start.X][start.Y] = 0

	inner = func(point Point) {
		for _, direction := range directions {
			next := Point{point.X + direction.X, point.Y + direction.Y}
			if next.X > 0 && next.X <= 8 &&
				next.Y > 0 && next.Y <= 8 &&
				distance[next.X][next.Y] > distance[point.X][point.Y]+1 {
				distance[next.X][next.Y] = distance[point.X][point.Y] + 1
				inner(next)
			}
		}
	}

	inner(start)

	if distance[target.X][target.Y] == Inf {
		return 0
	} else {
		return distance[target.X][target.Y]
	}

}
