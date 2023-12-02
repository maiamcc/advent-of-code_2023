package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"regexp"
	"strings"
)

var redRe = regexp.MustCompile("(\\d+) red")
var greenRe = regexp.MustCompile("(\\d+) green")
var blueRe = regexp.MustCompile("(\\d+) blue")
var gameIdRe = regexp.MustCompile("Game (\\d+)")

func main() {
	inputLines := utils.MustReadFileAsLines("day2/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	//fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	return len(inputLines)
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}

type round struct {
	numRed   int
	numGreen int
	numBlue  int
}

func (rnd round) isEmpty() bool {
	return rnd.numRed == 0 && rnd.numBlue == 0 && rnd.numGreen == 0
}

type game struct {
	id     int
	rounds []round
}

func numResultFromRe(input string, re *regexp.Regexp) int {
	res := re.FindStringSubmatch(input)
	if len(res) == 0 {
		// no match
		return 0
	}
	return utils.MustAtoI(res[1])
}
func parseRound(input string) round {
	return round{
		numRed:   numResultFromRe(input, redRe),
		numGreen: numResultFromRe(input, greenRe),
		numBlue:  numResultFromRe(input, blueRe),
	}
}

func parseGame(input string) (game, error) {
	parts := strings.Split(input, ":")
	if len(parts) != 2 {
		return game{}, fmt.Errorf("unexpected num parts when splitting input string '%s' on colon", input)
	}
	gameId := numResultFromRe(parts[0], gameIdRe)
	if gameId == 0 {
		return game{}, fmt.Errorf("couldn't parse game id from input string '%s", input)
	}
	roundStrs := strings.Split(parts[1], ";")
	var rounds []round
	for _, rnd := range roundStrs {
		round := parseRound(rnd)
		if round.isEmpty() {
			return game{}, fmt.Errorf("couldn't parse round from input string '%s", rnd)
		}
		rounds = append(rounds, round)

	}
	return game{
		id:     gameId,
		rounds: rounds,
	}, nil
}
