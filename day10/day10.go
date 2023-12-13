package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
)

func main() {
	inputLines := utils.MustReadFileAsLines("day10/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	return len(inputLines)
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}

type pipeCell struct {
	coords      utils.Coord
	val         string
	connections utils.Set[utils.Coord]
}

var c utils.Cell = pipeCell{}

func (c pipeCell) Coordinates() utils.Coord {
	return c.coords
}

func (c pipeCell) Value() string {
	return c.val
}

func (c pipeCell) step(from utils.Coord) (pipeCell, bool) {
	//if !c.connections.Contains(from) {
	//
	//}
	c.connections
}

/*
- refactor Matrix to take a cell creator and make cell an interface i guess?? or just make a new matrix class
- every cell has 0 or two dest cells
- cell.step(from coord) -- `from` must be one of the cell's two connection coords. if so, return the cell at the other dest coord
	- e.g. the cell at 1,1 connects to 1,0 and 2,1 (i.e. connects north and east): cell.step(from 1,1) returns the cell at 2,1
	(you entered at 1,1 so you exit at 2,1)
- given a [][]string, loop thru to make a Matrix --> [][]cell where every cell knows its connection coords, and also save S cell
- from the S cell, check N S E and W for valid connecting cells (i.e. a cell that includes S as one of its connections)
- for each valid cell connecting to S, check for a loop: step to next until you hit S again, or a dead end (i.e. period) (and count the steps it takes you)
- once you find your loop, which takes n steps, go n/2 steps through the loop and return that cell.
	- will it ever be an odd number of steps?? shrug.
*/
