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

	res := make([]int, 100)
	strs := make([][]string, 100)

	for i := 0; i < n; i++ {
		in.Scan()
		v, _ := strconv.Atoi(in.Text())
		res[v]++

		in.Scan()
		if i < n/2 {
			strs[v] = append(strs[v], "-")
		} else {
			strs[v] = append(strs[v], in.Text())
		}
	}

	total := 0
	for idx, count := range res {
		for i := 0; i < count; i++ {
			fmt.Fprint(out, strs[idx][i])
			if total != n-1 {
				fmt.Fprint(out, " ")
			}
			total++
		}
	}

}
