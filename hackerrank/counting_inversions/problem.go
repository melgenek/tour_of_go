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

func solve(arr []int) int {
	_, inversionsCount := mergeSort(arr)
	return inversionsCount
}

func mergeSort(arr []int) ([]int, int) {
	if len(arr) <= 1 {
		return arr, 0
	} else {
		middle := len(arr) / 2
		leftPart := arr[:middle]
		rightPart := arr[middle:]

		sortedLeftPart, leftInversionsCount := mergeSort(leftPart)
		sortedRightPart, rightInversionsCount := mergeSort(rightPart)

		result := make([]int, len(arr))
		i := 0
		j := 0
		k := 0
		inversionsCount := 0
		for i < len(sortedLeftPart) && j < len(sortedRightPart) {
			if sortedLeftPart[i] <= sortedRightPart[j] {
				result[k] = sortedLeftPart[i]
				i++
			} else {
				result[k] = sortedRightPart[j]
				inversionsCount += len(sortedLeftPart) - i
				j++
			}
			k++
		}

		for ; i < len(sortedLeftPart); i++ {
			result[k] = sortedLeftPart[i]
			k++
		}

		for ; j < len(sortedRightPart); j++ {
			result[k] = sortedRightPart[j]
			k++
		}

		return result, leftInversionsCount + rightInversionsCount + inversionsCount
	}
}

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	queriesCount, _ := strconv.Atoi(in.Text())

	for i := 0; i < queriesCount; i++ {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			in.Scan()
			value, _ := strconv.Atoi(in.Text())
			arr[i] = value
		}

		result := solve(arr)

		if i != queriesCount-1 {
			fmt.Fprintf(out, "%d\n", result)
		} else {
			fmt.Fprintf(out, "%d", result)
		}
	}
}
