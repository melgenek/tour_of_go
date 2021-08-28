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

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		in.Scan()
		arr[i], _ = strconv.Atoi(in.Text())
	}

	for j := 1; j < n; j++ {
		cur := arr[j]
		i := j
		for ; i > 0 && arr[i-1] > cur; i-- {
			arr[i] = arr[i-1]
		}
		arr[i] = cur
		print(out, arr)
		if j != n-1 {
			fmt.Fprint(out, "\n")
		}
	}

}

func print(out *bufio.Writer, arr []int) {
	for idx, v := range arr {
		if idx != len(arr)-1 {
			fmt.Fprintf(out, "%d ", v)
		} else {
			fmt.Fprintf(out, "%d", v)
		}
	}
}
