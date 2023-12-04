package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"strconv"
	"unicode"
)

func main() {
	inputLines := utils.MustReadFileAsLines("day3/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	total := 0
	matrix := utils.MustMatrix(inputLines)
	numCells := numbersForRow(matrix.Flatten()) // we can actually analyze the whole matrix at once, neat!
	for _, num := range numCells {
		isPartNum, err := isPartNumber(matrix, num)
		if err != nil {
			utils.LogfErrorAndExit(err, "checking for part num")
		}
		if isPartNum {
			i, err := num.asInt()
			if err != nil {
				utils.LogfErrorAndExit(err, "this should be a valid int")
			}
			total += i
		}
	}
	return total
}

// getAdjacentCoords returns all theoretically adjacent coordinates for the coordinates given (i.e. all
// squares 1 away on either axis). This includes diagonally adjacent squares. Note that not all
// adjacent coordinates may be valid on the actual matrix in question. Note also that the given
// coord pair will not appear in the returned list, as it is not adjacent to itself.
func getAdjacentCoords(c utils.Coord) []utils.Coord {
	var adjacent []utils.Coord
	for x := c.X - 1; x <= c.X+1; x++ {
		for y := c.Y - 1; y <= c.Y+1; y++ {
			if !(x == c.X && y == c.Y) { // input coordinate shouldn't show up in its own adjacency list
				adjacent = append(adjacent, utils.Coord{X: x, Y: y})
			}
		}
	}
	return adjacent
}

// getAllAdjacentCoords returns all theoretically adjacent coordinates for ALL given coordinates.
func getAllAdjacentCoords(coords []utils.Coord) []utils.Coord {
	var adjacent []utils.Coord
	for _, coord := range coords {
		adjacent = append(adjacent, getAdjacentCoords(coord)...)
	}
	return dedupeCoords(adjacent)
}

func dedupeCoords(coords []utils.Coord) []utils.Coord {
	var unique []utils.Coord
	seen := make(map[utils.Coord]struct{}) // Go, why don't you have built-in sets?!?
	for _, coord := range coords {
		if _, ok := seen[coord]; ok {
			continue
		}
		unique = append(unique, coord)
		seen[coord] = struct{}{}
	}
	return unique
}

// A list of adjacent cells representing a number (with each cell represents a digit)
type numberCells []utils.Cell

func (num numberCells) coords() []utils.Coord {
	var coords []utils.Coord
	for _, cell := range num {
		coords = append(coords, cell.Coordinates())
	}
	return coords
}

func (num numberCells) asInt() (int, error) {
	// nb: don't verify they're adjacent, assume we did it right lol
	var s string
	for _, cell := range num {
		s += cell.Val
	}
	return strconv.Atoi(s)
}

// numbersForRow returns an array of Cell arrays (each of the latter
// representing a series of horizontally adjacent cells containing digits,
// which together can be taken to represent a number)
func numbersForRow(row []utils.Cell) []numberCells {
	var allNumberCells []utils.Cell
	for _, cell := range row {
		if _, ok := cell.AsInt(); ok {
			allNumberCells = append(allNumberCells, cell)
		}
	}
	if len(allNumberCells) == 0 {
		return []numberCells{}
	}

	var numbers []numberCells
	var curNum []utils.Cell
	prevXCoord := allNumberCells[0].X - 1 // make sure first loop below adds first number cell to the array
	for _, cell := range allNumberCells {
		if cell.X == prevXCoord+1 {
			// This cell is adjacent to the previous number cell, i.e. it's part of the same number
			curNum = append(curNum, cell)
		} else {
			// otherwise, it's the start of a new number; add the number we've
			// been accumulating to the return array and start a new one
			numbers = append(numbers, curNum)
			curNum = []utils.Cell{cell}
		}
		prevXCoord = cell.X
	}
	// add whatever number we were working on before to the return array
	numbers = append(numbers, curNum)
	return numbers
}

func isSymbol(matrix utils.Matrix, coord utils.Coord) bool {
	cell, err := matrix.Get(coord.X, coord.Y)
	if err != nil {
		// ehh, implementation is such that we might get passed bunk coordinates from our
		// overeager adjacent coordinate finding. If it's not valid, just return false
		// instead of erroring out.
		return false
	}
	// Just checking if it's a digit or a period. If letters are possible in the grid,
	// will have to change this.
	return !unicode.IsNumber(utils.MustRune(cell.Val)) && cell.Val != "."
}

func anyIsSymbol(matrix utils.Matrix, coords []utils.Coord) (bool, error) {
	for _, coord := range coords {
		if isSymbol(matrix, coord) {
			return true, nil
		}
	}
	return false, nil
}

func isPartNumber(matrix utils.Matrix, numCells numberCells) (bool, error) {
	adjacentCoords := getAllAdjacentCoords(numCells.coords())
	return anyIsSymbol(matrix, adjacentCoords)
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}
