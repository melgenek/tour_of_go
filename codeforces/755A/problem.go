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
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	str := in.Text()
	n, _ := strconv.Atoi(str)

	if n == 1 {
		fmt.Fprintf(out, "%d", 3)
	} else if n == 2 {
		fmt.Fprintf(out, "%d", 4)
	} else {
		fmt.Fprintf(out, "%d", n-2)
	}

}
