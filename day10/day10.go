package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
)

var PIPE_TO_CONNECTIONS = map[string]func(c utils.Coord) utils.Set[utils.Coord]{
	//| is a vertical pipe connecting north and south.
	"|": func(c utils.Coord) utils.Set[utils.Coord] {
		return utils.NewSet[utils.Coord](c.North(), c.South())
	},
	//- is a horizontal pipe connecting east and west.
	"-": func(c utils.Coord) utils.Set[utils.Coord] {
		return utils.NewSet[utils.Coord](c.East(), c.West())
	},
	//L is a 90-degree bend connecting north and east.
	"L": func(c utils.Coord) utils.Set[utils.Coord] {
		return utils.NewSet[utils.Coord](c.North(), c.East())
	},
	//J is a 90-degree bend connecting north and west.
	"J": func(c utils.Coord) utils.Set[utils.Coord] {
		return utils.NewSet[utils.Coord](c.North(), c.West())
	},
	//7 is a 90-degree bend connecting south and west.
	"7": func(c utils.Coord) utils.Set[utils.Coord] {
		return utils.NewSet[utils.Coord](c.South(), c.West())
	},
	//F is a 90-degree bend connecting south and east.
	"F": func(c utils.Coord) utils.Set[utils.Coord] {
		return utils.NewSet[utils.Coord](c.South(), c.East())
	},
	//. is ground; there is no pipe in this tile.
	".": func(c utils.Coord) utils.Set[utils.Coord] { return utils.NewSet[utils.Coord]() },
	//S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
	"S": func(c utils.Coord) utils.Set[utils.Coord] { return utils.NewSet[utils.Coord]() },
}

func main() {
	inputLines := utils.MustReadFileAsLines("day10/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	_, start := matrixFromInput(inputLines)
	possibleDirs := []utils.Coord{
		start.coords.North(), start.coords.South(),
		start.coords.East(), start.coords.West(),
	}
	for _, dir := range possibleDirs {
		numSteps, isLoop := start.isLoop(dir)
		if isLoop {
			return numSteps/2 + 1
		}
	}
	return 0
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}

type pipeCell struct {
	coords      utils.Coord
	val         string
	connections utils.Set[utils.Coord]
	matrix      *utils.Matrix // each cell points back to its containing matrix
}

var c utils.Cell = pipeCell{}

func (c pipeCell) Coordinates() utils.Coord {
	return c.coords
}

func (c pipeCell) Value() string {
	return c.val
}

func NewPipeCell(x int, y int, val string) utils.Cell {
	coords := utils.Coord{X: x, Y: y}
	getConns, ok := PIPE_TO_CONNECTIONS[val]
	if !ok {
		utils.LogfAndExit("Invalid input character '%s'", val)
	}
	return pipeCell{
		coords:      coords,
		val:         val,
		connections: getConns(coords),
	}
}

func (c pipeCell) isStart() bool { return c.val == "S" }

func (c pipeCell) step(from utils.Coord) (pipeCell, bool) {
	if !c.connections.Contains(from) {
		// This cell isn't connected to the `from` cell so this is an illegal move.
		return pipeCell{}, false
	}
	if len(c.connections) != 2 {
		panic("malformed pipe cell -- expect exactly two connections")
	}

	// find the coords that are not the `from` coords -- these are the `to` coords
	var toCoords utils.Coord
	for coords, _ := range c.connections {
		if coords != from {
			toCoords = coords
			break
		}
	}
	toCell, err := c.matrix.GetByCoord(toCoords)
	if err != nil {
		return pipeCell{}, false
	}
	return toCell.(pipeCell), true
}

// isLoop determines whether starting from cell c and stepping to the
// cell at firstStep and continuing on from there will result in a loop
// i.e. will bring us back to cell c).
func (c pipeCell) isLoop(firstStep utils.Coord) (numSteps int, isLoop bool) {
	// todo: check if firstStep is adjacent to c
	prevCell := c
	curCellGeneric, err := c.matrix.GetByCoord(firstStep)
	if err != nil {
		utils.LogfErrorAndExit(err, "getting expected cell as first step of loop")
	}
	curCell := curCellGeneric.(pipeCell)
	var nextCell pipeCell
	ok := true
	for ok {
		nextCell, ok = curCell.step(prevCell.coords)
		numSteps += 1
		prevCell = curCell
		curCell = nextCell
		if curCell.coords == c.coords {
			return numSteps, true
		}
	}
	return 0, false
}
func matrixFromInput(input []string) (utils.Matrix, pipeCell) {
	m := utils.MustMatrix(input, NewPipeCell)
	var startCell pipeCell
	// attach reference to the parent matrix to each cell,
	// and also grab the "S" cell.
	for y, row := range m.Cells {
		for x, cell := range row {
			pc := cell.(pipeCell)
			pc.matrix = &m
			row[x] = pc
			if pc.isStart() {
				startCell = pc
			}
		}
		m.Cells[y] = row
	}

	return m, startCell
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
