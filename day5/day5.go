package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"os"
	"strings"
	"unicode"
)

func main() {
	input := utils.MustReadFileAsString("day5/input.txt")
	//fmt.Println("The answer to Part One is:", partOne(input))
	fmt.Println("The answer to Part Two is:", partTwo(input))
}

func partOne(fullInput string) int {
	seedsAndRest := strings.SplitN(fullInput, "\n\n", 2)
	if len(seedsAndRest) != 2 {
		fmt.Printf("unexpected number of parts for input")
		os.Exit(1)
	}
	seeds, err := parseSeedsPartOne(seedsAndRest[0])
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing seeds")
	}
	allMaps, err := parseMapChain(seedsAndRest[1])
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing input")
	}

	seedsToLocation := make(map[int]int)
	for _, seed := range seeds {
		seedsToLocation[seed] = allMaps.mapVal(seed)
	}
	minLocation := -1
	for _, location := range seedsToLocation {
		if minLocation == -1 {
			minLocation = location // set initial value
			continue
		}

		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

func partTwo(fullInput string) int {
	seedsAndRest := strings.SplitN(fullInput, "\n\n", 2)
	if len(seedsAndRest) != 2 {
		fmt.Printf("unexpected number of parts for input")
		os.Exit(1)
	}
	seeds, err := parseSeedsPartTwo(seedsAndRest[0])
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing seeds")
	}

	fmt.Printf("=Parsed %d seeds\n\n", len(seeds))

	allMaps, err := parseMapChain(seedsAndRest[1])
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing input")
	}

	minLocation := -1
	for _, seed := range seeds {
		location := allMaps.mapVal(seed)
		fmt.Printf("=Seed %d --> location %d\n", seed, location)
		if minLocation == -1 {
			minLocation = location // set initial value
			continue
		}

		if location < minLocation {
			minLocation = location
		}
	}
	return minLocation
}

// parseSeedsPartOne parses a "seeds: a b c d" line as a list of ints
func parseSeedsPartOne(s string) ([]int, error) {
	parts, err := utils.SplitIntoExpectedParts(s, ": ", 2)
	if err != nil {
		return nil, err
	}
	seedVals := strings.Split(parts[1], " ")
	return utils.StringsToInts(seedVals)
}

// parseSeedsPartTwo parses a "seeds: a b c d" line as a series of int ranges and returns a list of all ints contained therein
func parseSeedsPartTwo(s string) ([]int, error) {
	parts, err := utils.SplitIntoExpectedParts(s, ": ", 2)
	if err != nil {
		return nil, err
	}
	nums, err := utils.StringsToInts(strings.Split(parts[1], " "))
	if err != nil {
		return nil, err
	}

	var res []int
	rangeStart := -1
	for _, n := range nums {
		if rangeStart == -1 {
			// beginning of a range
			rangeStart = n
		} else {
			// n represents the range length
			res = append(res, utils.Rng(rangeStart, rangeStart+n)...)
			rangeStart = -1
		}
	}
	return res, nil
}

type seedRange struct {
	startVal  int
	rangeSize int
}

func parseMapping(s string) (mapping, error) {
	parts, err := utils.SplitIntoExpectedParts(s, " ", 3)
	if err != nil {
		return mapping{}, err
	}

	ints, err := utils.StringsToInts(parts)
	if err != nil {
		return mapping{}, err
	}
	return mapping{
		destRangeStart:   ints[0],
		sourceRangeStart: ints[1],
		rangeLen:         ints[2],
	}, nil
}

// parseMappingSet takes a block beginning with a header (e.g. "seed-to-soil mapping:") and
// some number of lines, each representing a mapping, and parses it into a mappingSet
func parseMappingSet(input string) (mappingSet, error) {
	return parseMappingSetHelper(strings.Split(input, "\n"))
}

func parseMappingSetHelper(lns []string) (mappingSet, error) {
	if len(lns) < 1 {
		return mappingSet{}, fmt.Errorf("no lines to parse")
	}
	if !(unicode.IsNumber(utils.MustRune(lns[0][0:1]))) {
		// First line is the block header, chop it off and parse the remaining
		if len(lns) == 1 {
			// wait what, there are no actual mapping lines to parse, this line is the only one!
			return mappingSet{}, fmt.Errorf("no lines to parse")
		}
		return parseMappingSetHelper(lns[1:])
	}
	result := make(mappingSet, len(lns))
	for i, ln := range lns {
		m, err := parseMapping(ln)
		if err != nil {
			return mappingSet{}, err
		}
		result[i] = m
	}
	return result, nil
}

func parseMapChain(input string) (mapChain, error) {
	blocks := strings.Split(input, "\n\n")
	var result mapChain
	for _, block := range blocks {
		block = strings.TrimSpace(block)
		if block == "" { // in case of imprecise splitting
			continue
		}
		ms, err := parseMappingSet(block)
		if err != nil {
			return mapChain{}, err
		}
		result = append(result, ms)
	}
	return result, nil
}

type mapping struct {
	destRangeStart   int
	sourceRangeStart int
	rangeLen         int
}

func (m mapping) sourceValWithinRange(srcVal int) bool {
	return m.sourceRangeStart <= srcVal &&
		srcVal < m.sourceRangeStart+m.rangeLen
}

func (m mapping) destValForSource(srcVal int) (destVal int, ok bool) {
	if !m.sourceValWithinRange(srcVal) {
		// source value isn't in the source range of this mapping, so can't return a valid result
		return -1, ok
	}
	return m.destRangeStart + (srcVal - m.sourceRangeStart), true
}

type mappingSet []mapping

func (ms mappingSet) mapVal(srcVal int) int {
	if srcVal < 0 {
		panic(fmt.Sprintf("can't process a negative source value (got: %d)", srcVal))
	}
	for _, m := range ms {
		destVal, ok := m.destValForSource(srcVal)
		if ok {
			return destVal
		}
	}
	// if we haven't successfully mapped the value yet, it doesn't fall within any of
	// the source ranges of the mappings we have; so it just corresponds to itself
	return srcVal
}

// list of mappingSets; a value should be passed through this chain in order, with
// the result from mappingSet[n].mapVal being passed to mappingSet[n+1].mapVal, etc.
type mapChain []mappingSet

func (mc mapChain) mapVal(srcVal int) int {
	destVal := srcVal
	for _, ms := range mc {
		destVal = ms.mapVal(destVal)
	}
	return destVal
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
//     prevVal = newVal  (i could just assign this directly i guess... but i like this readability) (actually wait no we need this b/c if the value never got mapped it just goes to itself, need to know if it changed)
//   seedsToLocation[seed] = prevVal  (this is the value we ended the iteration at)
