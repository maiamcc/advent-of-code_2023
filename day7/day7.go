package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"os"
	"sort"
	"strings"
)

var cardValuesPartOne = map[string]int{
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

var cardValuesPartTwo = map[string]int{
	"J": 1, // joker now the weakest card, make it worth the least
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	// "J": 11,
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
	fmt.Println("The answer to Part Two is:", partTwo(fullInput))
}

func partOne(fullInput string) int {
	total := 0
	hands := parseHandsWithBids(fullInput)
	sort.Sort(sortable(hands))
	for i, h := range hands {
		rank := i + 1 // 1-index us
		total += rank * h.bid
		//fmt.Printf("%d: %v\t[%s]\t(bid: %d)\n", rank, h.cards, h.typeRank(false).toString(), h.bid)
	}
	return total
}

func partTwo(fullInput string) int {
	total := 0
	hands := parseHandsWithBids(fullInput)
	sort.Sort(sortableWithJokers(hands))
	for i, h := range hands {
		rank := i + 1 // 1-index us
		total += rank * h.bid
		//fmt.Printf("%d: %v\t[%s]\t(bid: %d)\n", rank, h.cards, h.typeRank().toString(), h.bid)
	}
	return total
}

func newCard(s string) card {
	_, ok := cardValuesPartOne[s] // eh this should be paramaterized but it's just validation
	if !ok {
		fmt.Printf("invalid label '%s', will not create card\n", s)
		os.Exit(1)
	}
	return card(s)
}

type card string

func (c card) val(cardVals map[string]int) int {
	v, ok := cardVals[string(c)]
	if !ok {
		// this is probs unnecessary now that we have validation on card creation but whatevs
		fmt.Printf("couldn't find numeric value for label '%s'\n", c)
		os.Exit(1)
	}
	return v
}

func (c card) _cmp(c2 card, cardVals map[string]int) int {
	if c.val(cardVals) < c2.val(cardVals) {
		return -1
	}
	if c.val(cardVals) > c2.val(cardVals) {
		return 1
	}
	return 0
}

func (c card) lessThan(c2 card, cardVals map[string]int) bool {
	return c._cmp(c2, cardVals) == -1
}

func (c card) greaterThan(c2 card, cardVals map[string]int) bool {
	return c._cmp(c2, cardVals) == 1
}

type hand struct {
	cards []card
	bid   int
}

func (h hand) _cmpPartOne(h2 hand) int {
	return h._cmp(h2, cardValuesPartOne, false)
}

func (h hand) _cmpPartTwo(h2 hand) int {
	return h._cmp(h2, cardValuesPartTwo, true)
}

func (h hand) _cmp(h2 hand, cardVals map[string]int, withJokers bool) int {
	tr1 := h.typeRank(withJokers)
	tr2 := h2.typeRank(withJokers)

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
		if c.lessThan(h2.cards[i], cardVals) {
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

func (h hand) typeRankPartOne() typeRank {
	return h.typeRank(false)
}

func (h hand) typeRankPartTwo() typeRank {
	return h.typeRank(true)
}

func (h hand) typeRank(withJokers bool) typeRank {
	cardFreqs := make(map[card]int)
	for _, c := range h.cards {
		cardFreqs[c] += 1
	}
	if withJokers {
		if jCount, ok := cardFreqs["J"]; ok && jCount != 5 {
			// jokers present, so modify the card count accordingly.
			// Hunch: the best possible use of a joker in any given situation is to
			// act as another of the card you already have the most of.

			// NB: this conditional filters out a special case, 5 jokers. If that's the hand,
			// no need to modify what the jokers stand for, it'll be parsed as 5 of a kind below.

			delete(cardFreqs, "J")   // pop jokers from freq map so they don't get double counted as part of the hand
			cardWithMost := card("") // find card we already have hte most of
			for c, count := range cardFreqs {
				if cardWithMost == "" {
					cardWithMost = c
					continue
				}
				if count > cardFreqs[cardWithMost] {
					cardWithMost = c
				}
			}
			cardFreqs[cardWithMost] += jCount // pretend we have count(J) more of that card
		}
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
		}
		return FULL_HOUSE
	}
	return FIVE_OF_A_KIND
}

// sortable implements sort.Interface for []hand based on the strength of the hands
type sortable []hand

func (s sortable) Len() int           { return len(s) }
func (s sortable) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortable) Less(i, j int) bool { return s[i]._cmpPartOne(s[j]) == -1 }

// sortable implements sort.Interface for []hand based on the strength of the hands
type sortableWithJokers []hand

func (sj sortableWithJokers) Len() int           { return len(sj) }
func (sj sortableWithJokers) Swap(i, j int)      { sj[i], sj[j] = sj[j], sj[i] }
func (sj sortableWithJokers) Less(i, j int) bool { return sj[i]._cmpPartTwo(sj[j]) == -1 }
