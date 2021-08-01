package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	Solution(os.Stdin, os.Stdout)
}

type UnionSets struct {
	parents []int32
	ranks   []int32
}

func NewUnionSet(size int32) *UnionSets {
	return &UnionSets{
		make([]int32, size),
		make([]int32, size),
	}
}

func (sets *UnionSets) make(v int32) {
	sets.parents[v] = v
	sets.ranks[v] = 0
}

func (sets *UnionSets) find(v int32) int32 {
	if sets.parents[v] == v {
		return v
	} else {
		sets.parents[v] = sets.find(sets.parents[v])
		return sets.parents[v]
	}
}

func (sets *UnionSets) union(a, b int32) bool {
	a = sets.find(a)
	b = sets.find(b)

	if a != b {
		if sets.ranks[a] < sets.ranks[b] {
			a, b = b, a
		}
		sets.parents[b] = a
		if sets.ranks[a] == sets.ranks[b] {
			sets.ranks[a]++
		}
		return true
	} else {
		return false
	}
}

func roadsAndLibraries(n int32, libraryCost int32, roadCost int32, cities [][]int32) int64 {
	if libraryCost < roadCost {
		return int64(n) * int64(libraryCost)
	} else {
		trees := NewUnionSet(n)
		for i := 0; int32(i) < n; i++ {
			trees.make(int32(i))
		}

		var totalRoadsCost int64 = 0
		librariesCount := n

		for _, road := range cities {
			a, b := road[0]-1, road[1]-1
			if trees.union(a, b) {
				totalRoadsCost += int64(roadCost)
				librariesCount -= 1
			}
		}

		return int64(totalRoadsCost) + int64(librariesCount)*int64(libraryCost)
	}
}

func Solution(input io.Reader, output io.Writer) {
	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)
		fmt.Fprintf(writer, "%d", result)
		if qItr != int(q)-1 {
			fmt.Fprintf(writer, "\n")
		}
		writer.Flush()
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
