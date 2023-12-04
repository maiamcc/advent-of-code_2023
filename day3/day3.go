package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
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
	symbolCoords := coordsForSymbols(matrix)
	symbolAdjacentCoords := getAllAdjacentCoords(symbolCoords)
	for _, coord := range symbolAdjacentCoords {
		cell, err := matrix.Get(coord.X, coord.Y)
		if err != nil {
			utils.LogfErrorAndExit(err, "getting cell at (%d, %d)", coord.X, coord.Y)
		}
		if i, ok := cell.AsInt(); ok {
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

// coordsForSymbols returns all coordinates representing (non-period) symbols.
// Current implementation is: non-numeric, non-period values. If alphabetical values show up we're SOL.
func coordsForSymbols(matrix utils.Matrix) []utils.Coord {
	var coords []utils.Coord
	for _, cell := range matrix.Flatten() {
		if !unicode.IsNumber(utils.MustRune(cell.Val)) && cell.Val != "." {
			coords = append(coords, utils.Coord{X: cell.X, Y: cell.Y})
		}
	}
	return coords
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}
