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
	fields := strings.Fields(str)
	n, _ := strconv.Atoi(fields[0])
	destination, _ := strconv.Atoi(fields[1])
	alice := destination - 1

	in.Scan()
	str = in.Text()
	line1 := strings.Fields(str)

	in.Scan()
	str = in.Text()
	line2 := strings.Fields(str)

	found := false
	if line1[0] == "1" {
		if line1[alice] == "1" {
			found = true
		} else if line2[alice] == "1" {
			for i := alice + 1; i < n; i++ {
				if line1[i] == "1" && line2[i] == "1" {
					found = true
					break
				}
			}
		}
	}

	if found {
		fmt.Fprintf(out, "YES")
	} else {
		fmt.Fprintf(out, "NO")
	}
	out.Flush()

}
