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
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	total := 0
	for _, ln := range inputLines {
		g, err := parseGame(ln)
		if err != nil {
			utils.LogfErrorAndExit(err, "parsing line '%s'", ln)
		}
		if g.possibleWith(12, 13, 14) {
			total += g.id
		}
	}
	return total
}

func partTwo(inputLines []string) int {
	total := 0
	for _, ln := range inputLines {
		g, err := parseGame(ln)
		if err != nil {
			utils.LogfErrorAndExit(err, "parsing line '%s'", ln)
		}
		minSet := g.minSet()
		total += minSet.power()
	}
	return total
}

type round struct {
	numRed   int
	numGreen int
	numBlue  int
}

func (rnd round) isEmpty() bool {
	return rnd.numRed == 0 && rnd.numBlue == 0 && rnd.numGreen == 0
}

// possibleWith indicates whether the receiver round would be possible with
// the given number of red, green, and blue blocks.
func (rnd round) possibleWith(red int, green int, blue int) bool {
	return rnd.numRed <= red && rnd.numGreen <= green && rnd.numBlue <= blue
}

func (rnd round) power() int {
	return rnd.numRed * rnd.numGreen * rnd.numBlue
}

type game struct {
	id     int
	rounds []round
}

// possibleWith indicates whether the receiver game would be possible with
// the given number of red, green, and blue blocks.
func (g game) possibleWith(red int, green int, blue int) bool {
	for _, rnd := range g.rounds {
		if !rnd.possibleWith(red, green, blue) {
			return false
		}
	}
	return true
}

// minSet returns the minimum number of red/green/blue cubes needed to play the receiver game.
// (We return this info as a round b/c it's the easiest way of storing info about r/g/b cubes lol.)
func (g game) minSet() round {
	var minRed, minGreen, minBlue int
	for _, rnd := range g.rounds {
		if rnd.numRed > minRed {
			minRed = rnd.numRed
		}
		if rnd.numGreen > minGreen {
			minGreen = rnd.numGreen
		}

		if rnd.numBlue > minBlue {
			minBlue = rnd.numBlue
		}
	}
	return round{minRed, minGreen, minBlue}
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
