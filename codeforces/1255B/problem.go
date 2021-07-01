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
	tests, _ := strconv.Atoi(in.Text())

	for i := 0; i < tests; i++ {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		in.Scan()
		m, _ := strconv.Atoi(in.Text())

		weights := make([]int, n)
		sum := 0
		for j := 0; j < n; j++ {
			in.Scan()
			w, _ := strconv.Atoi(in.Text())
			sum += w
			weights[j] = w
		}

		if n <= 2 {
			fmt.Fprintf(out, "-1\n")
		} else if m < n {
			fmt.Fprintf(out, "-1\n")
		} else if m == n {
			fmt.Fprintf(out, "%d\n", sum*2)
			fmt.Fprintf(out, "%d %d\n", 1, n)
			for j := 1; j < n; j++ {
				fmt.Fprintf(out, "%d %d", j, j+1)
				if !(i == tests-1 && j == n-1) {
					fmt.Fprintf(out, "\n")
				}
			}
		}
	}

	out.Flush()

}
