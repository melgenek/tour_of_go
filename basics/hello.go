package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := 0; i < n; i++ {
		if unicode.IsLower(rune(b[i])) {
			b[i] = 'a' + ((b[i] - 'a' + 13) % 26)
		}
		if unicode.IsUpper(rune(b[i])) {
			b[i] = 'A' + ((b[i] - 'A' + 13) % 26)
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
