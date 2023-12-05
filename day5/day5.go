package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
)

func main() {
	inputLines := utils.MustReadFileAsLines("day5/input.txt")
	fmt.Println("The answer to Part One is:", partOne(inputLines))
	fmt.Println("The answer to Part Two is:", partTwo(inputLines))
}

func partOne(inputLines []string) int {
	return len(inputLines)
}

func partTwo(inputLines []string) int {
	return len(inputLines)
}

type mapping struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLen         int
}

func (m mapping) sourceValWithin(srcVal int) bool {
	return false
}

func (m mapping) destValForSource(srcVal int) (destVal int, ok bool) {
	return
}

// Imagining that for each attr we're mapping, there's a []mapping and we go through one by one to see which applies to the seed we're checking.
// if there were a prohibitive number of mappings for each property we could bin search thru, but there are few enough that i bet we can just iterate
// To parse the input:
// - split on double newline
// - can we map the header text (e.g. "temperature-to-humidity map:") to the type constructor? or uh, the mappings will all be of the same type I guess.

// maybe each attr combination e.g. temp-to-humidity has a mappingSet
// for seed in seeds:
//   prevVal = seed
//   for conversionMaps in orderedMapSets:
//     newVal = conversationMaps.convert(prevVal)
//     (maybe log prevVal --> newVal somewhere? or just print)
//     prevVal = newVal  (i could just assign this directly i guess... but i like this readability)
//   seedsToLocation[seed] = prevVal  (this is the value we ended the iteration at)
