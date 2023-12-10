package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

	actual := partOne(inputStr)
	assert.Equal(t, 6440, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

	actual := partTwo(inputStr)
	assert.Equal(t, 123, actual)
}

func TestHand_TypeRank(t *testing.T) {
	cases := map[string]typeRank{ // map input to expected output
		"86524": HIGH_CARD,
		"AA452": ONE_PAIR,
		"77QQ8": TWO_PAIR,
		"444JK": THREE_OF_A_KIND,
		"22333": FULL_HOUSE,
		"AA8AA": FOUR_OF_A_KIND,
		"AAAAA": FIVE_OF_A_KIND,
		"7A7A7": FULL_HOUSE,
		"83333": FOUR_OF_A_KIND,
		"87778": FULL_HOUSE,
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			h := hand{cards: parseCards(input)} // trivial hand without a bid
			actual := h.typeRank()
			assert.Equalf(t, expected, actual, "expected %s / got %s", expected.toString())
		})
	}
}

func TestHand_Cmp(t *testing.T) {
	cases := []struct {
		h1          string
		h2          string
		expectedCmp int
	}{
		{"86524", "AAAAA", -1},
		{"77QQ8", "AA452", 1},
		{"444JK", "555JK", -1},
		{"444JK", "23555", 1},
		{"AA8AA", "AA7AA", 1},
		{"34567", "35467", -1},
		{"JJJJJ", "JJJJJ", 0},
		{"83333", "87778", 1},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			hand1 := hand{cards: parseCards(c.h1)} // trivial hand without a bid
			hand2 := hand{cards: parseCards(c.h2)} // trivial hand without a bid
			assert.Equal(t, c.expectedCmp, hand1._cmp(hand2))
		})
	}
}
