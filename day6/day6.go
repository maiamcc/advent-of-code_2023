package main

import (
	"fmt"
	"github.com/maiamcc/advent-of-code_2023/utils"
	"strconv"
	"strings"
)

func main() {
	fullInput := utils.MustReadFileAsString("day6/input.txt")
	fmt.Println("The answer to Part One is:", partOne(fullInput))
	fmt.Println("The answer to Part Two is:", partTwo(fullInput))
}

func partOne(fullInput string) int {
	races, err := parseRaces(fullInput)
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing races")
	}
	total := 1
	for _, r := range races {
		fmt.Printf("== ANALYZING RACE: dur. %d ms, record %d mm ==\n", r.dur, r.record)
		total *= r.numWinningOptionsV2()
	}
	return total
}

func partTwo(fullInput string) int {
	r, err := parseRacePartTwo(fullInput)
	if err != nil {
		utils.LogfErrorAndExit(err, "parsing the one big race")
	}
	return r.numWinningOptionsV2()
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
		} else {
			if count > 0 {
				// we've found a loser after finding at least one winner;
				//everything from here will be a loser and we can just bail
				break
			}
		}
	}
	fmt.Printf("--> %d winning options\n\n", count)
	return count
}

func (r race) distForMsButtonPress(ms int) int {
	return (r.dur - ms) * ms
}

func (r race) numWinningOptionsV2() int {
	fmt.Println("Searching for first winning button")
	median := r.dur / 2 // uhh will the middle number always be a winner? i hope so!
	firstWinningButtonMs := binSearchInflection(0, median, func(i int) bool {
		return r.distForMsButtonPress(i) > r.record
	})
	firstLosingButtonMs := binSearchInflection(median, r.dur, func(i int) bool {
		return r.distForMsButtonPress(i) <= r.record
	})
	return firstLosingButtonMs - firstWinningButtonMs
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

func parseRacePartTwo(input string) (race, error) {
	lns, err := utils.SplitIntoExpectedParts(input, "\n", 2)
	if err != nil {
		return race{}, err
	}

	dur, err := parseMegaString(lns[0])
	if err != nil {
		return race{}, err
	}

	dist, err := parseMegaString(lns[1])
	if err != nil {
		return race{}, err
	}

	return race{dur, dist}, nil
}

// look idk how naming works. Take input `foobar:  1    3    5   7` --> int 1357
func parseMegaString(s string) (int, error) {
	pts, err := utils.SplitIntoExpectedParts(s, ":", 2)
	if err != nil {
		return 0, err
	}
	s = strings.Replace(pts[1], " ", "", -1)
	return strconv.Atoi(s)
}

// binSearchInflection finds the inflection point — the first value where checker(i)
// returns true having previously returned false
func binSearchInflection(start int, end int, checker func(i int) bool) int {
	if checker(start) {
		panic("expected input where left side returned false")
	} else if !checker(end) {
		panic("expected input where right side returned true")
	}

	if end-start == 1 {
		// they're one apart, we've found the inflection!
		return end
	}

	median := start + (end-start)/2 // b/c of int implementation, this rounds down
	fmt.Printf("= bin search %d -> %d (median: %d)\n", start, end, median)
	if !checker(median) {
		fmt.Printf("== recursing on %d -> %d\n", median, end)
		return binSearchInflection(median, end, checker)
	} else {
		fmt.Printf("== recursing on %d -> %d\n", start, median)
		return binSearchInflection(start, median, checker)
	}
}
