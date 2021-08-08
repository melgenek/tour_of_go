package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	//Solution(os.Stdin, os.Stdout)
	input, _ := os.Open("input.txt")
	output, _ := os.Create("output.txt")
	Solution(input, output)
}

type Point struct {
	idx int
	x   int
	y   int
}

type Variant struct {
	k             int
	point         Point
	originalPoint Point
}

func positiveVariant(point Point) Variant {
	if point.x >= 0 && point.y >= 0 {
		return Variant{1, Point{point.idx, point.x, point.y}, point}
	} else if point.x < 0 && point.y >= 0 {
		return Variant{2, Point{point.idx, -point.x, point.y}, point}
	} else if point.x >= 0 && point.y < 0 {
		return Variant{3, Point{point.idx, point.x, -point.y}, point}
	} else {
		return Variant{4, Point{point.idx, -point.x, -point.y}, point}
	}
}

func negativeVariant(point Point) Variant {
	if point.x <= 0 && point.y <= 0 {
		return Variant{1, Point{point.idx, point.x, point.y}, point}
	} else if point.x > 0 && point.y <= 0 {
		return Variant{2, Point{point.idx, -point.x, point.y}, point}
	} else if point.x <= 0 && point.y > 0 {
		return Variant{3, Point{point.idx, point.x, -point.y}, point}
	} else {
		return Variant{4, Point{point.idx, -point.x, -point.y}, point}
	}
}

type Result struct {
	a        Variant
	b        Variant
	distance float64
}

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	str := in.Text()
	n, _ := strconv.Atoi(str)

	var xSorted []Variant
	var ySorted []Variant
	for i := 1; i <= n; i++ {
		in.Scan()
		x, _ := strconv.Atoi(in.Text())

		in.Scan()
		y, _ := strconv.Atoi(in.Text())

		variant := positiveVariant(Point{i, x, y})
		xSorted = append(xSorted, variant)
		ySorted = append(ySorted, variant)
	}

	sort.SliceStable(xSorted, func(i, j int) bool {
		return xSorted[i].point.x < xSorted[j].point.x
	})
	sort.SliceStable(ySorted, func(i, j int) bool {
		return ySorted[i].point.y < ySorted[j].point.y
	})
	result := solveFast(xSorted, ySorted)
	negativeB := negativeVariant(result.b.originalPoint)

	fmt.Fprintf(out, "%d %d %d %d", result.a.point.idx, result.a.k, negativeB.point.idx, negativeB.k)
	out.Flush()

}

func solveFast(xSorted []Variant, ySorted []Variant) *Result {
	if len(xSorted) <= 3 {
		return solve(xSorted)
	} else {
		middle := len(xSorted) / 2
		middleX := xSorted[middle].point.x

		xLeft := xSorted[:middle]
		xRight := xSorted[middle:]

		yLeft := make([]Variant, len(xLeft))
		yRight := make([]Variant, len(xRight))
		for il, ir, i := 0, 0, 0; i < len(ySorted); i++ {
			if ySorted[i].point.x <= middleX && il < len(yLeft) || ir == len(yRight) {
				yLeft[il] = ySorted[i]
				il++
			} else {
				yRight[ir] = ySorted[i]
				ir++
			}
		}

		leftResult := solveFast(xLeft, yLeft)
		rightResult := solveFast(xRight, yRight)

		min := &Result{distance: 99999999}
		if leftResult != nil && leftResult.distance < min.distance {
			min = leftResult
		}
		if rightResult != nil && rightResult.distance < min.distance {
			min = rightResult
		}

		strip := make([]Variant, 0)
		for _, v := range ySorted {
			if math.Abs(float64(v.point.x-middleX)) < min.distance {
				strip = append(strip, v)
			}
		}
		//sort.SliceStable(strip, func(i, j int) bool {
		//	return strip[i].point.y < strip[j].point.y
		//})

		for i := 0; i < len(strip); i++ {
			for j := i + 1; j < len(strip) && j-i <= 7; j++ {
				a := strip[i]
				b := strip[j]
				if a.point.idx != b.point.idx {
					dist := distance(a.point, b.point)
					if min.distance > dist {
						min = &Result{a, b, dist}
					}
				}
			}
		}

		return min
	}
}

func solve(vectors []Variant) *Result {
	var min *Result
	for _, a := range vectors {
		for _, b := range vectors {
			if a.point.idx != b.point.idx {
				dist := distance(a.point, b.point)
				if min == nil || min.distance > dist {
					min = &Result{a, b, dist}
				}
			}
		}
	}
	return min
}

func distance(a Point, b Point) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}
