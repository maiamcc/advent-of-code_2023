package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"os"
	"sort"
	"strings"
)

var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}

type typeRank int

const (
	HIGH_CARD typeRank = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func (tr typeRank) toString() string {
	switch tr {
	case HIGH_CARD:
		return "high card"
	case ONE_PAIR:
		return "one pair"
	case TWO_PAIR:
		return "two pair"
	case THREE_OF_A_KIND:
		return "three of a kind"
	case FULL_HOUSE:
		return "full house"
	case FOUR_OF_A_KIND:
		return "four of a kind"
	case FIVE_OF_A_KIND:
		return "five of a kind"
	default:
		return "[invalid type]"
	}
}

func main() {
	fullInput := utils.MustReadFileAsString("day7/input.txt")
	fmt.Println("The answer to Part One is:", partOne(fullInput))
	//fmt.Println("The answer to Part Two is:", partTwo(fullInput))
}

func partOne(fullInput string) int {
	total := 0
	hands := parseHandsWithBids(fullInput)
	sort.Sort(sortable(hands))
	for i, h := range hands {
		rank := i + 1 // 1-index us
		total += rank * h.bid
	}
	return total
}

func partTwo(fullInput string) int {
	return len(fullInput)
}

func newCard(s string) card {
	_, ok := cardValues[s]
	if !ok {
		fmt.Printf("invalid label '%s', will not create card\n", s)
		os.Exit(1)
	}
	return card(s)
}

type card string

func (c card) val() int {
	v, ok := cardValues[string(c)]
	if !ok {
		// this is probs unnecessary now that we have validation on card creation but whatevs
		fmt.Printf("couldn't find numeric value for label '%s'\n", c)
		os.Exit(1)
	}
	return v
}

func (c card) _cmp(c2 card) int {
	if c.val() < c2.val() {
		return -1
	}
	if c.val() > c2.val() {
		return 1
	}
	return 0
}

func (c card) lessThan(c2 card) bool {
	return c._cmp(c2) == -1
}

func (c card) greaterThan(c2 card) bool {
	return c._cmp(c2) == 1
}

type hand struct {
	cards []card
	bid   int
}

func (h hand) _cmp(h2 hand) int {
	tr1 := h.typeRank()
	tr2 := h2.typeRank()

	if tr1 < tr2 {
		return -1
	} else if tr1 > tr2 {
		return 1
	}

	// type ranks are equal, compare card by card
	for i, c := range h.cards {
		if c == h2.cards[i] {
			continue
		}
		if c.lessThan(h2.cards[i]) {
			return -1
		} else {
			return 1
		}
	}
	// hands are identical
	return 0
}

func parseCards(s string) []card {
	cardStrs := strings.Split(s, "")
	cards := make([]card, len(cardStrs))
	for i, c := range cardStrs {
		cards[i] = newCard(c)
	}
	return cards
}

func parseHandWithBid(s string) hand {
	cardsAndBid, err := utils.SplitIntoExpectedParts(s, " ", 2)
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing hand")
	}
	cards := parseCards(cardsAndBid[0])
	bid := utils.MustAtoI(cardsAndBid[1])
	return hand{cards, bid}
}

func parseHandsWithBids(s string) []hand {
	lns := strings.Split(s, "\n")
	hands := make([]hand, len(lns))
	for i, ln := range lns {
		hands[i] = parseHandWithBid(ln)
	}
	return hands
}

//func parseHands(s string) []hand {
//	lns := strings.Split(s, "\n")
//	var hands []hand
//	for _, ln := range lns {
//		hands = append(hands, parseCards(ln))
//	}
//	return hands
//}

func (h hand) typeRank() typeRank {
	cardFreqs := make(map[card]int)
	for _, c := range h.cards {
		cardFreqs[c] += 1
	}
	if len(cardFreqs) == 5 {
		return HIGH_CARD
	} else if len(cardFreqs) == 4 {
		return ONE_PAIR
	} else if len(cardFreqs) == 3 {
		for _, count := range cardFreqs {
			if count == 3 {
				return THREE_OF_A_KIND
			}
		}
		return TWO_PAIR
	} else if len(cardFreqs) == 2 {
		for _, count := range cardFreqs {
			if count == 4 {
				return FOUR_OF_A_KIND
			}
			return FULL_HOUSE
		}
	}
	return FIVE_OF_A_KIND
}

// sortable implements sort.Interface for []hand based on the strength of the hands
type sortable []hand

func (s sortable) Len() int           { return len(s) }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortable) Less(i, j int) bool { return s[i]._cmp(s[j]) == -1 }
