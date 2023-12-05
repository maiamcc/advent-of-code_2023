package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"math"
	"regexp"
	"strings"
)

var cardIdRe = regexp.MustCompile("Card +(\\d+)")

func main() {
	inputLines := utils.MustReadFileAsLines("day4/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	total := 0
	for _, ln := range inputLines {
		g := parseCard(ln)
		total += g.score()
	}
	return total
}

func partTwo(inputLines []string) int {
	cards := cardStackFromLines(inputLines)
	for i, _ := range cards {
		cards.scoreCardWithId(i + 1)
	}
	return cards.totalCards()
}

type card struct {
	id          int // this is now denormalized data but whatever we'll just leave it
	count       int // number of occurences of this card (incl. original and any copies) in the stack
	winningNums utils.IntSet
	numsOnCard  []int
}

func (g card) countWinners() int {
	count := 0
	for _, num := range g.numsOnCard {
		if g.winningNums.Contains(num) {
			count += 1
		}
	}
	return count
}

func (g card) score() int {
	return int(math.Pow(2, float64(g.countWinners()-1)))
}

func parseCard(input string) card {
	cardIdAndNums, err := utils.SplitIntoExpectedParts(input, ":", 2)
	if err != nil {
		utils.LogfErrorAndExit(err, "splitting input (into card id and rest)")
	}

	cardId := utils.NumResultFromRe(cardIdAndNums[0], cardIdRe)
	if cardId == 0 {
		utils.LogfErrorAndExit(nil, "couldn't parse card id from input string '%s", input)
	}

	parts, err := utils.SplitIntoExpectedParts(strings.TrimSpace(cardIdAndNums[1]), " | ", 2)
	if err != nil {
		utils.LogfErrorAndExit(err, "splitting input (into winning numbers / numbers on card)")
	}

	return card{
		id:          cardId,
		count:       1, // initally, there's one of each card
		winningNums: utils.NewIntSet(utils.MustStringsToInts(strings.Split(parts[0], " "))),
		numsOnCard:  utils.MustStringsToInts(strings.Split(parts[1], " ")),
	}
}

// assume cards are in order and there are no gaps; thus, card 1 is the first
// elem (index 0), card 20 is the 20th element (index 19), etc.
type cardStack []card

func cardStackFromLines(lns []string) cardStack {
	cs := make(cardStack, len(lns))
	for i, ln := range lns {
		cs[i] = parseCard(ln)
	}
	return cs
}

func (cs cardStack) scoreCardWithId(id int) {
	if id > len(cs) {
		// nonexistent card id, idk if this might happen when scoring the very last card?
		// just treat it as a no-op.
		return
	}
	curCard := cs[id-1]
	winCount := curCard.countWinners()
	for i := id + 1; i < id+winCount+1; i++ {
		if i > len(cs) {
			// can't get extra copies of a card you don't have I guess
			break
		}
		cs[i-1].count += curCard.count
	}
}

func (cs cardStack) totalCards() int {
	score := 0
	for _, g := range cs {
		score += g.count
	}
	return score
}
