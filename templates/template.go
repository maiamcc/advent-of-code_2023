package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
)

func main() {
	inputLines := utils.MustReadFileAsLines("DAY_N/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	return len(inputLines)
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}
