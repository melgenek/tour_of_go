package matrix_multiplication

func Mul(a [][]int, b [][]int) [][]int {
	height := len(a)
	innerLen := len(b)
	width := len(b[0])

	c := make([][]int, len(a))
	for i, _ := range c {
		c[i] = make([]int, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for k := 0; k < innerLen; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}

func MulRow(a [][]int, b [][]int) [][]int {
	height := len(a)
	innerLen := len(b)
	width := len(b[0])

	c := make([][]int, len(a))
	for i, _ := range c {
		c[i] = make([]int, width)
	}

	for i := 0; i < height; i++ {
		for k := 0; k < innerLen; k++ {
			for j := 0; j < width; j++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return c
}
