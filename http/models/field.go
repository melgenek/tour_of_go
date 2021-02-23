package models

import "math/rand"

const width = 3500
const length = 3500
const depth = 10

type Field = [width][length][depth]bool

func CreateField() Field {
	var field Field
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			for k := 0; k < depth; k++ {
				field[i][j][k] = rand.Intn(2) == 1
			}
		}
	}
	return field
}
