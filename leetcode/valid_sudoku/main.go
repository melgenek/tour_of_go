package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	columns := make([][]bool, 10)
	for i := range columns {
		columns[i] = make([]bool, 10)
	}

	quadrants := make([][][]bool, 3)
	for i := range quadrants {
		quadrants[i] = make([][]bool, 3)
		for j := range quadrants[i] {
			quadrants[i][j] = make([]bool, 10)
		}
	}

	for i, row := range board {
		rowSeen := make([]bool, 10)
		for j, v := range row {
			if v == '.' {
				continue
			}
			n := v - '0'

			if rowSeen[n] {
				return false
			}
			rowSeen[n] = true

			seenColumn := columns[j][n]
			if seenColumn {
				return false
			}
			columns[j][n] = true

			quadrantI := i / 3
			quadrantJ := j / 3
			seenQuadrant := quadrants[quadrantI][quadrantJ][n]
			if seenQuadrant {
				return false
			}
			quadrants[quadrantI][quadrantJ][n] = true
		}
	}

	return true
}

// [
//["7",".",".",".","4",".",".",".","."],
//[".",".",".","8","6","5",".",".","."],
//[".","1",".","2",".",".",".",".","."],
//[".",".",".",".",".","9",".",".","."],
//[".",".",".",".","5",".","5",".","."],
//[".",".",".",".",".",".",".",".","."],
//[".",".",".",".",".",".","2",".","."],
//[".",".",".",".",".",".",".",".","."],
//[".",".",".",".",".",".",".",".","."]
//]
