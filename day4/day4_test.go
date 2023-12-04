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
	inputStr := "input\ngoes\nhere"
	inputLns := strings.Split(inputStr, "\n")

	actual := partTwo(inputLns)
	assert.Equal(t, 123, actual)
}

func TestParseGame(t *testing.T) {
	input := "Card 52: 87 83  26 28 3 |  88 3  777 12 93    22 82 36"
	expected := game{
		id:          52,
		winningNums: utils.NewIntSet([]int{87, 83, 26, 28, 3}),
		numsOnCard:  []int{88, 3, 777, 12, 93, 22, 82, 36},
	}
	assert.Equal(t, expected, parseGame(input))
}

func TestGame_CountWinners(t *testing.T) {
	g := game{
		id:          111,
		winningNums: utils.NewIntSet([]int{13, 32, 20, 16, 61}),
		numsOnCard:  []int{16, 3, 777, 12, 93, 61, 82, 36},
	}
	assert.Equal(t, 2, g.countWinners())
}

func TestGame_Score(t *testing.T) {
	g := game{
		id:          1,
		winningNums: utils.NewIntSet([]int{41, 48, 83, 86, 17}),
		numsOnCard:  []int{83, 86, 6, 31, 17, 9, 48, 53},
	}
	assert.Equal(t, 8, g.score())
}

func TestNewCardStack(t *testing.T) {

}
func TestCardStack_ScoreGameWithId(t *testing.T) {
	cs := cardStack{
		game{id: 0, count: 0},
		game{id: 1, count: 3},
		game{id: 2, count: 3,
			winningNums: utils.NewIntSet([]int{13, 32, 20, 16, 61}),
			numsOnCard:  []int{16, 3, 777, 12, 93, 61, 82, 36}}, // 2 winners
		game{id: 3, count: 6},
		game{id: 4, count: 2},
		game{id: 5, count: 1},
	}
	expected := cardStack{
		game{id: 0, count: 0},
		game{id: 1, count: 3},
		game{id: 2, count: 3,
			winningNums: utils.NewIntSet([]int{13, 32, 20, 16, 61}),
			numsOnCard:  []int{16, 3, 777, 12, 93, 61, 82, 36}},
		game{id: 3, count: 9},
		game{id: 4, count: 5},
		game{id: 5, count: 1},
	}
	cs.scoreGameWithId(2)
	assert.Equal(t, expected, cs)
}

func TestCardStack_TotalCards(t *testing.T) {
	cs := cardStack{
		game{id: 0, count: 0},
		game{id: 1, count: 3},
		game{id: 2, count: 3},
		game{id: 3, count: 6},
		game{id: 4, count: 2},
		game{id: 5, count: 1},
	}
	assert.Equal(t, 15, cs.totalCards())
}
