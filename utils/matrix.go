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
type Matrix struct {
	Cells   [][]Cell
	NumRows int // aka height; NumRows - 1 = maximum allowable y value
	NumCols int // aka width; NumCols - 1 = maximum allowable x value
}

// NewMatrix creates a Matrix object from a list of strings, where each string
// represents a row and each character in that string is an element in its own column.
func NewMatrix(input []string) (Matrix, error) {
	if len(input) < 1 {
		return Matrix{}, fmt.Errorf("need at least one row with which to construct a matrix")
	}

	rowLen := len(input[0]) // all rows must be the same length

	var cells [][]Cell
	for y, s := range input {
		if len(s) != rowLen {
			return Matrix{}, fmt.Errorf("row at y=%d is not of expected row length %d (row: '%s')", y, rowLen, s)
		}
		var row []Cell
		for x, ch := range strings.Split(s, "") {
			row = append(row, Cell{
				X:   x,
				Y:   y,
				Val: ch,
			})
		}
		cells = append(cells, row)
	}
	return Matrix{
		Cells:   cells,
		NumRows: len(cells),
		NumCols: rowLen,
	}, nil
}

func MustMatrix(input []string) Matrix {
	matrix, err := NewMatrix(input)
	if err != nil {
		LogfErrorAndExit(err, "when making matrix that was definitely gonna be okay")
	}
	return matrix
}

func (m Matrix) Get(x int, y int) (Cell, error) {
	if x < 0 || x >= m.NumCols {
		return Cell{}, fmt.Errorf("invalid x value %d (must be 0 <= x <= %d)", x, m.NumCols-1)
	}
	if y < 0 || y >= m.NumRows {
		return Cell{}, fmt.Errorf("invalid y value %d (must be 0 <= y <= %d)", y, m.NumRows-1)
	}
	return m.Cells[y][x], nil
}

func (m Matrix) GetByCoord(c Coord) (Cell, error) {
	return m.Get(c.X, c.Y)
}

// Flatten returns a list containing all the cells in the matrix
// such that they can be iterated over
func (m Matrix) Flatten() []Cell {
	var cells []Cell
	for _, row := range m.Cells {
		cells = append(cells, row...)
	}
	return cells
}

type Cell struct {
	X   int
	Y   int
	Val string
}

func (c Cell) Coordinates() Coord {
	return Coord{c.X, c.Y}
}
func (c Cell) AsInt() (val int, ok bool) {
	i, err := strconv.Atoi(c.Val)
	return i, err == nil
}
