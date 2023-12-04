package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"math"
	"regexp"
	"strings"
)

var gameIdRe = regexp.MustCompile("Card +(\\d+)")

func main() {
	inputLines := utils.MustReadFileAsLines("day4/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	total := 0
	for _, ln := range inputLines {
		g := parseGame(ln)
		total += g.score()
	}
	return total
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}

type game struct {
	id          int // this is now denormalized data but whatever we'll just leave it
	count       int // number of occurences of this card (incl. original and any copies) in the stack
	winningNums utils.IntSet
	numsOnCard  []int
}

func (g game) countWinners() int {
	count := 0
	for _, num := range g.numsOnCard {
		if g.winningNums.Contains(num) {
			count += 1
		}
	}
	return count
}

func (g game) score() int {
	return int(math.Pow(2, float64(g.countWinners()-1)))
}

func parseGame(input string) game {
	gameIdAndNums, err := utils.SplitIntoExpectedParts(input, ":", 2)
	if err != nil {
		utils.LogfErrorAndExit(err, "splitting input (into game id and rest)")
	}

	gameId := utils.NumResultFromRe(gameIdAndNums[0], gameIdRe)
	if gameId == 0 {
		utils.LogfErrorAndExit(nil, "couldn't parse game id from input string '%s", input)
	}

	parts, err := utils.SplitIntoExpectedParts(strings.TrimSpace(gameIdAndNums[1]), " | ", 2)
	if err != nil {
		utils.LogfErrorAndExit(err, "splitting input (into winning numbers / numbers on card)")
	}

	return game{
		id:          gameId,
		winningNums: utils.NewIntSet(utils.MustStringsToInts(strings.Split(parts[0], " "))),
		numsOnCard:  utils.MustStringsToInts(strings.Split(parts[1], " ")),
	}
}

type cardStack []game // index games by id (leave a dummy game at cardStack[0] so it's one-indexed

func newCardStack(games []game) cardStack {
	// assume that we don't have holes in our list of games: that a list of 5 games
	// implies that we have games with ids 1 --> 5 (and need a list of len 6 to hold
	// them, b/c we're 1-indexed
	cs := make(cardStack, len(games)+1)
	for _, g := range games {
		cs[g.id] = g
	}
	return cs
}

func (cs cardStack) scoreGameWithId(id int) {
	if id >= len(cs) {
		// nonexistent game id, idk if this might happen when scoring the very last game?
		// just treat it as a no-op.
		return
	}
	curGame := cs[id]
	winCount := curGame.countWinners()
	for i := id + 1; i < id+winCount+1; i++ {
		if i >= len(cs) {
			// can't get extra copies of a game you don't have I guess
			break
		}
		cs[i].count += curGame.count
	}
}

func (cs cardStack) totalCards() int {
	score := 0
	for _, g := range cs {
		score += g.count
	}
	return score
}
