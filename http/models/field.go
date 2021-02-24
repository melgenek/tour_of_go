package models

import "math/rand"

const width = 3500
const length = 3500
const depth = 10

type Cell struct {
	HasGold bool
	goldAt  int
	depth   int
}

type Field struct {
	Cells [width][length]*Cell
}

func CreateField() *Field {
	var field Field
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			field.Cells[i][j] = &Cell{HasGold: rand.Intn(5) == 1, goldAt: rand.Intn(depth), depth: 0}
		}
	}
	return &field
}

func (field *Field) Dig(posX int, posY int) bool {
	cell := field.Cells[posX][posY]
	cell.depth += 1
	return cell.depth == cell.goldAt
}
