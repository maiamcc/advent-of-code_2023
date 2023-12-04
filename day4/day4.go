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
	id          int
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
