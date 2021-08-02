package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestProblem(t *testing.T) {
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		if file.IsDir() {
			t.Run(file.Name(), func(t *testing.T) {
				input, _ := os.Open(strings.Join([]string{file.Name(), "input.txt"}, "/"))
				output, _ := os.Open(strings.Join([]string{file.Name(), "output.txt"}, "/"))
				expectedResult, _ := ioutil.ReadAll(output)
				var result bytes.Buffer
				Solution(input, bufio.NewWriter(&result))
				if string(result.Bytes()) != string(expectedResult) {
					t.Fatalf(`Expected: %s. Actual: %s`, string(expectedResult), string(result.Bytes()))
				}
			})
		}
	}
}
