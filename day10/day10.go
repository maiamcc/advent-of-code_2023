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

type cellMark int

const (
	UNKNOWN cellMark = iota
	LOOP
	INTERNAL
	EXTERNAL
)

func (cm cellMark) toString() string {
	switch cm {
	case UNKNOWN:
		return "UNKNOWN"
	case LOOP:
		return "LOOP"
	case INTERNAL:
		return "INTERNAL"
	case EXTERNAL:
		return "EXTERNAL"
	default:
		return "[invalid type]"
	}
}

func main() {
	inputLines := utils.MustReadFileAsLines("day10/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))

	ans, _ := partTwo(inputLines)
	fmt.Println("The answer to Part Two is:", ans)
}

func partOne(inputLines []string) int {
	_, start := matrixFromInput(inputLines)
	loopCoords := loopCoordsFromStart(start)
	return len(loopCoords) / 2
}

func partTwo(inputLines []string) (int, utils.Matrix[*pipeCell]) {
	matrix, start := matrixFromInput(inputLines)
	findAndMarkLoop(start)
	for _, row := range matrix.Cells {
		for _, cell := range row {
			if cell.mark == UNKNOWN {
				cell.radiateAndMark()
			}
		}
	}
	count := 0
	for _, cell := range matrix.Flatten() {
		if cell.mark == INTERNAL {
			count += 1
		}
	}
	return count, matrix
}

type pipeCell struct {
	coords      utils.Coord
	val         string
	connections utils.Set[utils.Coord]
	matrix      *utils.Matrix[*pipeCell] // each cell points back to its containing matrix
	mark        cellMark
}

var _ utils.Cell = &pipeCell{}

func (c *pipeCell) Coordinates() utils.Coord {
	return c.coords
}

func (c *pipeCell) Value() string {
	return c.val
}

func (c *pipeCell) debugStr() string {
	if c.mark == UNKNOWN {
		return "?"
	} else if c.mark == INTERNAL {
		return "I"
	} else if c.mark == EXTERNAL {
		return "0"
	}
	return c.val
}

func NewPipeCell(x int, y int, val string) *pipeCell {
	coords := utils.Coord{X: x, Y: y}
	getConns, ok := PIPE_TO_CONNECTIONS[val]
	if !ok {
		utils.LogfAndExit("Invalid input character '%s'", val)
	}
	return &pipeCell{
		coords:      coords,
		val:         val,
		connections: getConns(coords),
	}
}

func (c *pipeCell) isStart() bool { return c.val == "S" }

func (c *pipeCell) step(from utils.Coord) (*pipeCell, bool) {
	if !c.connections.Contains(from) {
		// This cell isn't connected to the `from` cell so this is an illegal move.
		return nil, false
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
		return nil, false
	}
	return toCell, true
}

// isLoop determines whether starting from cell c and stepping to the
// cell at firstStep and continuing on from there will result in a loop
// i.e. will bring us back to cell c).
func (c *pipeCell) isLoop(firstStep utils.Coord) (loopCoords utils.Set[utils.Coord], isLoop bool) {
	// Assume that firstStep is adjacent to the current cell I guess

	loopCoords = utils.NewSet[utils.Coord](c.coords)
	prevCell := c
	curCell, err := c.matrix.GetByCoord(firstStep)
	if err != nil {
		// we fell off the edge when attempting to start the loop in this direction,
		// so this sure isn't a valid loop
		return utils.NewSet[utils.Coord](), false
	}
	var nextCell *pipeCell
	ok := true
	for ok {
		nextCell, ok = curCell.step(prevCell.coords)
		if !ok {
			break
		}
		prevCell = curCell
		loopCoords.Add(curCell.coords)
		curCell = nextCell
		if curCell.coords == c.coords {
			return loopCoords, true
		}
	}
	return utils.NewSet[utils.Coord](), false
}

func loopCoordsFromStart(start *pipeCell) utils.Set[utils.Coord] {
	for _, dir := range start.coords.CardinalAdjacent() {
		loopCoords, isLoop := start.isLoop(dir)
		if isLoop {
			return loopCoords
		}
	}
	return nil
}

// radiateAndMark radiates out adjacent cells, finding all adjacent cells that are not
// the edge of the board or a part of the main loop; once all adjacent cells have been found,
// mark them as either internal (if we didn't hit the edge of the board, just loop cells)
// or external (if this group of cells hit the edge of the board).
func (c *pipeCell) radiateAndMark() {
	visited, anyIsEdge := c.radiateAdjacent(utils.NewSet[utils.Coord]())
	mark := INTERNAL
	if anyIsEdge {
		mark = EXTERNAL
	}
	markCoords(c.matrix, visited, mark)
}

func (c *pipeCell) radiateAdjacent(visitedSoFar utils.Set[utils.Coord]) (visited utils.Set[utils.Coord], anyIsEdge bool) {
	visitedSoFar.Add(c.coords)
	var foundEdge bool
	for _, coord := range c.coords.CardinalAdjacent() {
		if visitedSoFar.Contains(coord) {
			continue
		}
		cell, err := c.matrix.GetByCoord(coord)
		if err != nil {
			// index out of range error = we just attempted to "get" a cell outside the grid;
			// i.e. the requested coord was off the grid, i.e. current cell is on the edge.
			anyIsEdge = true
			continue
		}
		if cell.mark == LOOP {
			// don't have to radiate from this cell as it's part of the loop
			continue
		}
		visitedSoFar, foundEdge = cell.radiateAdjacent(visitedSoFar)
		anyIsEdge = anyIsEdge || foundEdge
	}
	return visitedSoFar, anyIsEdge
}

func matrixFromInput(input []string) (utils.Matrix[*pipeCell], *pipeCell) {
	m := utils.MustMatrix[*pipeCell](input, NewPipeCell)
	var startCell *pipeCell
	// attach reference to the parent matrix to each cell,
	// and also grab the "S" cell.
	for y, row := range m.Cells {
		for x, cell := range row {
			cell.matrix = &m
			row[x] = cell
			if cell.isStart() {
				startCell = cell
			}
		}
		m.Cells[y] = row
	}

	return m, startCell
}

// findAndMarkLoop finds the cells that are part of the loop starting at the given
// start location and marks them with mark = LOOP.
func findAndMarkLoop(start *pipeCell) {
	loopCoords := loopCoordsFromStart(start)
	markCoords(start.matrix, loopCoords, LOOP)
}

func markCoords(matrix *utils.Matrix[*pipeCell], coords utils.Set[utils.Coord], mark cellMark) {
	for coord, _ := range coords {
		cell, err := matrix.GetByCoord(coord)
		if err != nil {
			utils.LogfErrorAndExit(err, "didn't expect an invalid coord in call to `markCoords`")
		}
		cell.mark = mark
	}
}

func debugPrint(matrix utils.Matrix[*pipeCell]) {
	for _, row := range matrix.Cells {
		var printRow []string
		for _, cell := range row {
			printRow = append(printRow, cell.debugStr())
		}
		fmt.Println(printRow)
	}
}
