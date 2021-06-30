package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	Solution(os.Stdin, os.Stdout)
}

type Coord struct {
	X int
	Y int
}

var (
	directions = []Coord{
		{1, 0},
		{0, 1},
		{0, -1},
		{-1, 0},
	}
)

func Solution(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	out := bufio.NewWriter(output)
	defer out.Flush()

	in.Scan()
	str := in.Text()
	fields := strings.Fields(str)
	rows, _ := strconv.Atoi(fields[0])
	columns, _ := strconv.Atoi(fields[1])

	field := make([][]rune, rows)
	queue := list.New()

	for i := 0; i < rows; i++ {
		in.Scan()
		row := []rune(in.Text())
		field[i] = make([]rune, columns)
		for j := 0; j < columns; j++ {
			field[i][j] = row[j]

			if field[i][j] == 'W' {
				queue.PushFront(Coord{i, j})
			}
		}
	}

	seen := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		seen[i] = make([]bool, columns)
	}

	success := true
	for queue.Len() > 0 {
		element := queue.Front()
		coord := element.Value.(Coord)
		seen[coord.X][coord.Y] = true
		for _, direction := range directions {
			if coord.X+direction.X < 0 || coord.Y+direction.Y < 0 ||
				coord.X+direction.X >= rows || coord.Y+direction.Y >= columns {
				continue
			}
			adjacentCell := field[coord.X+direction.X][coord.Y+direction.Y]
			if adjacentCell == 'S' {
				if field[coord.X][coord.Y] == 'W' {
					success = false
					queue.Init()
					break
				} else {
					field[coord.X][coord.Y] = 'D'
				}
			}
			if adjacentCell != 'S' && adjacentCell != 'D' && !seen[coord.X+direction.X][coord.Y+direction.Y] {
				queue.PushFront(Coord{coord.X + direction.X, coord.Y + direction.Y})
			}
		}
		queue.Remove(element)
	}

	if success {
		fmt.Fprintf(out, "Yes\n")
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				fmt.Fprintf(out, "%c", field[i][j])
			}
			if i != rows-1 {
				fmt.Fprintf(out, "\n")
			}
		}
	} else {
		fmt.Fprintf(out, "No")
	}
	out.Flush()

}
