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

	students := make([]int, n+1)
	for i := 1; i <= n; i++ {
		in.Scan()
		v, _ := strconv.Atoi(in.Text())
		students[i] = v
	}

	for i := 1; i <= n; i++ {
		seen := make([]bool, n+1)
		cur := i
		for !seen[cur] {
			seen[cur] = true
			cur = students[cur]
		}
		fmt.Fprintf(out, "%d", cur)
		if i != n {
			fmt.Fprintf(out, " ")
		}
	}
	out.Flush()

}
