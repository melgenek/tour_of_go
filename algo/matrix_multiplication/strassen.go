package matrix_multiplication

func Strassen(m1 [][]int, m2 [][]int) [][]int {
	if len(m1) < 128 {
		return MulRow(m1, m2)
	} else {
		a, b, c, d := split(m1)
		e, f, g, h := split(m2)

		p1 := Strassen(a, Subtract(f, h))
		p2 := Strassen(Sum(a, b), h)
		p3 := Strassen(Sum(c, d), e)
		p4 := Strassen(d, Subtract(g, e))
		p5 := Strassen(Sum(a, d), Sum(e, h))
		p6 := Strassen(Subtract(b, d), Sum(g, h))
		p7 := Strassen(Subtract(a, c), Sum(e, f))

		c11 := Sum(Subtract(Sum(p5, p4), p2), p6)
		c12 := Sum(p1, p2)
		c21 := Sum(p3, p4)
		c22 := Subtract(Subtract(Sum(p1, p5), p3), p7)

		width := len(m1)
		result := make([][]int, width)
		for i, _ := range result {
			result[i] = make([]int, width)
			if i < len(c11) {
				copy(result[i], c11[i])
				copy(result[i][len(c11[i]):], c12[i])
			} else {
				bottomI := i - len(c11)
				copy(result[i], c21[bottomI])
				copy(result[i][len(c21[bottomI]):], c22[bottomI])
			}
		}

		return result
	}
}

func split(a [][]int) ([][]int, [][]int, [][]int, [][]int) {
	height := len(a)
	half := height / 2

	topLeft := make([][]int, half)
	topRight := make([][]int, half)

	for i, row := range a[:half] {
		topLeft[i] = row[:half]
		topRight[i] = row[half:]
	}

	bottomLeft := make([][]int, height-half)
	bottomRight := make([][]int, height-half)
	for i, row := range a[half:] {
		bottomLeft[i] = row[:half]
		bottomRight[i] = row[half:]
	}

	return topLeft, topRight, bottomLeft, bottomRight
}

func Sum(a [][]int, b [][]int) [][]int {
	return sum(a, b, func(a1 int, a2 int) int {
		return a1 + a2
	})
}

func Subtract(a [][]int, b [][]int) [][]int {
	return sum(a, b, func(a1 int, a2 int) int {
		return a1 - a2
	})
}

func sum(a [][]int, b [][]int, op func(int, int) int) [][]int {
	width := len(a)

	c := make([][]int, width)
	for i, _ := range c {
		c[i] = make([]int, width)
		for j, _ := range c[i] {
			c[i][j] = op(a[i][j], b[i][j])
		}
	}

	return c
}
