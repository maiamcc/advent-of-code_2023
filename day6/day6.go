package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"strings"
)

func main() {
	fullInput := utils.MustReadFileAsString("day6/input.txt")
	fmt.Println("The answer to Part One is:", partOne(fullInput))
	//fmt.Println("The answer to Part Two is:", partTwo(fullInput))
}

func partOne(fullInput string) int {
	races, err := parseRaces(fullInput)
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing races")
	}
	total := 1
	for _, r := range races {
		total *= r.numWinningOptions()
	}
	return total
}

func partTwo(fullInput string) int {
	return len(fullInput)
}

type race struct {
	dur    int // in ms
	record int // in ms
}

func (r race) numWinningOptions() int {
	count := 0
	fmt.Printf("== ANALYZING RACE: dur. %d ms, record %d mm ==\n", r.dur, r.record)
	for i := 0; i <= r.dur; i++ {
		// i ms holding button, dur - i ms traveling at a speed of i mm/ms
		actualDist := (r.dur - i) * i
		fmt.Printf("%d ms holding button --> %d ms of travel at %d mm/ms --> dist of %d\n", i, r.dur-i, i, actualDist)
		if actualDist > r.record {
			fmt.Println("-- incrementing win count")
			count += 1
		}
	}
	fmt.Printf("--> %d winning options\n\n", count)
	return count
}

func parseRaces(input string) ([]race, error) {
	lns, err := utils.SplitIntoExpectedParts(input, "\n", 2)
	if err != nil {
		return nil, err
	}

	durPts, err := utils.SplitIntoExpectedParts(lns[0], ":", 2)
	if err != nil {
		return nil, err
	}
	durVals, err := utils.StringsToInts(strings.Split(strings.TrimSpace(durPts[1]), " "))

	distPts, err := utils.SplitIntoExpectedParts(lns[1], ":", 2)
	if err != nil {
		return nil, err
	}
	distVals, err := utils.StringsToInts(strings.Split(strings.TrimSpace(distPts[1]), " "))

	if len(durVals) != len(distVals) {
		return nil, fmt.Errorf("different length duration and distance arrays")
	}
	races := make([]race, len(durVals))
	for i := range durVals {
		races[i] = race{dur: durVals[i], record: distVals[i]}
	}
	return races, nil
}
