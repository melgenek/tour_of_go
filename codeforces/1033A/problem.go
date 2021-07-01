package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"math"
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

var (
	directions = []Point{
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
	}
)

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	in.Scan()
	a1, _ := strconv.Atoi(in.Text())
	in.Scan()
	a2, _ := strconv.Atoi(in.Text())

	in.Scan()
	b1, _ := strconv.Atoi(in.Text())
	in.Scan()
	b2, _ := strconv.Atoi(in.Text())

	in.Scan()
	c1, _ := strconv.Atoi(in.Text())
	in.Scan()
	c2, _ := strconv.Atoi(in.Text())

	seen := make([][]bool, n+1)
	for i := 1; i < n+1; i++ {
		seen[i] = make([]bool, n+1)
	}

	queue := list.New()
	queue.PushFront(Point{b1, b2})
	for queue.Len() > 0 {
		el := queue.Front()
		point := el.Value.(Point)
		seen[point.X][point.Y] = true
		if point.X == c1 && point.Y == c2 {
			break
		}
		for _, direction := range directions {
			next := Point{point.X + direction.X, point.Y + direction.Y}
			if next.X <= n && next.X > 0 &&
				next.Y <= n && next.Y > 0 &&
				next.X != a1 &&
				next.Y != a2 &&
				math.Abs(float64(next.X-a1)) != math.Abs(float64(next.Y-a2)) &&
				!seen[next.X][next.Y] {
				queue.PushFront(Point{point.X + direction.X, point.Y + direction.Y})
			}
		}

		queue.Remove(el)
	}

	if seen[c1][c2] {
		fmt.Fprintf(out, "YES")
	} else {
		fmt.Fprintf(out, "NO")
	}
	out.Flush()

}
