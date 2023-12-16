package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	// so we can store coordinate pairs outside of the context of a cell
	X, Y int
}

// Given a coord, get the coord one step in the given cardinal direction
func (co Coord) North() Coord { return Coord{co.X, co.Y - 1} }
func (co Coord) South() Coord { return Coord{co.X, co.Y + 1} }
func (co Coord) East() Coord  { return Coord{co.X + 1, co.Y} }
func (co Coord) West() Coord  { return Coord{co.X - 1, co.Y} }
func (co Coord) CardinalAdjacent() []Coord {
	return []Coord{co.North(), co.South(), co.East(), co.West()}
}

type Matrix[C Cell] struct {
	Cells   [][]C
	NumRows int // aka height; NumRows - 1 = maximum allowable y value
	NumCols int // aka width; NumCols - 1 = maximum allowable x value
}

type Cell interface {
	Coordinates() Coord
	Value() string
}

type CellConstructor[C Cell] func(x int, y int, val string) C

type SimpleCell struct {
	X   int
	Y   int
	Val string
}

var _ Cell = SimpleCell{}

func (c SimpleCell) Coordinates() Coord {
	return Coord{c.X, c.Y}
}

func (c SimpleCell) Value() string {
	return c.Val
}

func (c SimpleCell) AsInt() (val int, ok bool) {
	i, err := strconv.Atoi(c.Val)
	return i, err == nil
}

func NewSimpleCell(x int, y int, val string) SimpleCell {
	return SimpleCell{x, y, val}
}

// NewMatrix creates a Matrix object from a list of strings, where each string
// represents a row and each character in that string is an element in its own column.
func NewMatrix[C Cell](input []string, cellCtr CellConstructor[C]) (Matrix[C], error) {
	if len(input) < 1 {
		return Matrix[C]{}, fmt.Errorf("need at least one row with which to construct a matrix")
	}

	rowLen := len(input[0]) // all rows must be the same length

	var cells [][]C
	for y, s := range input {
		if len(s) != rowLen {
			return Matrix[C]{}, fmt.Errorf("row at y=%d is not of expected row length %d (row: '%s')", y, rowLen, s)
		}
		var row []C
		for x, ch := range strings.Split(s, "") {
			row = append(row, cellCtr(x, y, ch))
		}
		cells = append(cells, row)
	}
	return Matrix[C]{
		Cells:   cells,
		NumRows: len(cells),
		NumCols: rowLen,
	}, nil
}

func MustMatrix[C Cell](input []string, cellCtr CellConstructor[C]) Matrix[C] {
	matrix, err := NewMatrix[C](input, cellCtr)
	if err != nil {
		LogfErrorAndExit(err, "when making matrix that was definitely gonna be okay")
	}
	return matrix
}

func MustSimpleCellMatrix(input []string) Matrix[SimpleCell] {
	return MustMatrix[SimpleCell](input, NewSimpleCell)
}

func (m Matrix[C]) Get(x int, y int) (C, error) {
	var zeroCell C
	if x < 0 || x >= m.NumCols {
		return zeroCell, fmt.Errorf("invalid x value %d (must be 0 <= x <= %d)", x, m.NumCols-1)
	}
	if y < 0 || y >= m.NumRows {
		return zeroCell, fmt.Errorf("invalid y value %d (must be 0 <= y <= %d)", y, m.NumRows-1)
	}
	return m.Cells[y][x], nil
}

func (m Matrix[C]) GetByCoord(c Coord) (C, error) {
	return m.Get(c.X, c.Y)
}

// Flatten returns a list containing all the cells in the matrix
// such that they can be iterated over
func (m Matrix[C]) Flatten() []C {
	var cells []C
	for _, row := range m.Cells {
		cells = append(cells, row...)
	}
	return cells
}
