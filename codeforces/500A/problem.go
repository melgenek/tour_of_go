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
	in.Scan()
	to, _ := strconv.Atoi(in.Text())

	portals := make([]int, n+1)
	for i := 1; i <= n-1; i++ {
		in.Scan()
		v, _ := strconv.Atoi(in.Text())
		portals[i] = v
	}

	now := 1
	for now < to {
		now = now + portals[now]
	}

	if now == to {
		fmt.Fprintf(out, "YES")
	} else {
		fmt.Fprintf(out, "NO")
	}
	out.Flush()

}
