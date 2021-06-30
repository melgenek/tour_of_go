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

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	str := in.Text()
	n, _ := strconv.Atoi(str)

	in.Scan()
	likesStr := in.Text()
	likes := strings.Fields(likesStr)
	values := make([]int, n)

	for i, likeStr := range likes {
		like, _ := strconv.Atoi(likeStr)
		values[i] = like - 1
	}

	for a, like := range values {
		b := values[like]
		c := values[b]
		if c == a {
			fmt.Fprintf(out, "YES")
			return
		}
	}

	fmt.Fprintf(out, "NO")
}
