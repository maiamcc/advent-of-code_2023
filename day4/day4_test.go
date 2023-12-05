package main

import (
	"github.com/maiamcc/advent-of-code_2023/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	inputLns := strings.Split(inputStr, "\n")

	actual := partOne(inputLns)
	assert.Equal(t, 13, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 30, actual)
}

func TestParseCard(t *testing.T) {
	input := "Card 52: 87 83  26 28 3 |  88 3  777 12 93    22 82 36"
	expected := card{
		id:          52,
		count:       1,
		winningNums: utils.NewIntSet([]int{87, 83, 26, 28, 3}),
		numsOnCard:  []int{88, 3, 777, 12, 93, 22, 82, 36},
	}
	assert.Equal(t, expected, parseCard(input))
}

func TestCard_CountWinners(t *testing.T) {
	c := card{
		id:          111,
		winningNums: utils.NewIntSet([]int{13, 32, 20, 16, 61}),
		numsOnCard:  []int{16, 3, 777, 12, 93, 61, 82, 36},
	}
	assert.Equal(t, 2, c.countWinners())
}

func TestCard_Score(t *testing.T) {
	c := card{
		id:          1,
		winningNums: utils.NewIntSet([]int{41, 48, 83, 86, 17}),
		numsOnCard:  []int{83, 86, 6, 31, 17, 9, 48, 53},
	}
	assert.Equal(t, 8, c.score())
}

func TestCardStackFromLines(t *testing.T) {
	input := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	}
	expected := cardStack{
		// is this a meaningful test? Cnclear!
		parseCard("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"),
		parseCard("Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"),
		parseCard("Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"),
	}
	assert.Equal(t, expected, cardStackFromLines(input))
}

func TestCardStack_ScoreCardWithId(t *testing.T) {
	cs := cardStack{
		card{id: 1, count: 3},
		card{id: 2, count: 3,
			winningNums: utils.NewIntSet([]int{13, 32, 20, 16, 61}),
			numsOnCard:  []int{16, 3, 777, 12, 93, 61, 82, 36}}, // 2 winners
		card{id: 3, count: 6},
		card{id: 4, count: 2},
		card{id: 5, count: 1},
	}
	expected := cardStack{
		card{id: 1, count: 3},
		card{id: 2, count: 3,
			winningNums: utils.NewIntSet([]int{13, 32, 20, 16, 61}),
			numsOnCard:  []int{16, 3, 777, 12, 93, 61, 82, 36}},
		card{id: 3, count: 9},
		card{id: 4, count: 5},
		card{id: 5, count: 1},
	}
	cs.scoreCardWithId(2)
	assert.Equal(t, expected, cs)
}

func TestCardStack_TotalCards(t *testing.T) {
	cs := cardStack{
		card{id: 0, count: 0},
		card{id: 1, count: 3},
		card{id: 2, count: 3},
		card{id: 3, count: 6},
		card{id: 4, count: 2},
		card{id: 5, count: 1},
	}
	assert.Equal(t, 15, cs.totalCards())
}
