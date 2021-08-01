package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	Solution(os.Stdin, os.Stdout)
}

type Edge struct {
	from   int32
	to     int32
	weight int32
}

func kruskals(nodesCount int32, from []int32, to []int32, weights []int32) int32 {
	edges := make([]Edge, len(from))
	for i := range edges {
		edges[i] = Edge{from[i] - 1, to[i] - 1, weights[i]}
	}
	sort.Slice(edges, func(i, j int) bool {
		edgeI, edgeJ := edges[i], edges[j]
		weightI, weightJ := edgeI.weight, edgeJ.weight
		if weightI != weightJ {
			return weightI < weightJ
		} else {
			return (edgeI.from + edgeI.to + weightI) < (edgeJ.from + edgeJ.to + weightJ)
		}
	})

	trees := make([]int32, nodesCount)
	for i := range trees {
		trees[i] = int32(i)
	}

	var cost int32 = 0
	for _, edge := range edges {
		oldTree, newTree := trees[edge.from], trees[edge.to]
		if trees[edge.from] != trees[edge.to] {
			cost += edge.weight
			for j, tree := range trees {
				if tree == oldTree {
					trees[j] = newTree
				}
			}
		}
	}

	return cost
}

func Solution(input io.Reader, output io.Writer) {
	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	gNodesEdges := strings.Split(readLine(reader), " ")

	gNodes, err := strconv.ParseInt(gNodesEdges[0], 10, 64)
	checkError(err)

	gEdges, err := strconv.ParseInt(gNodesEdges[1], 10, 64)
	checkError(err)

	var gFrom, gTo, gWeight []int32

	for i := 0; i < int(gEdges); i++ {
		edgeFromToWeight := strings.Split(readLine(reader), " ")

		edgeFrom, err := strconv.ParseInt(edgeFromToWeight[0], 10, 64)
		checkError(err)

		edgeTo, err := strconv.ParseInt(edgeFromToWeight[1], 10, 64)
		checkError(err)

		edgeWeight, err := strconv.ParseInt(edgeFromToWeight[2], 10, 64)
		checkError(err)

		gFrom = append(gFrom, int32(edgeFrom))
		gTo = append(gTo, int32(edgeTo))
		gWeight = append(gWeight, int32(edgeWeight))
	}

	res := kruskals(int32(gNodes), gFrom, gTo, gWeight)

	fmt.Fprintf(writer, "%d", res)

	writer.Flush()
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
