package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartOne(t *testing.T) {
	inputStr := "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"

	actual := partOne(inputStr)
	assert.Equal(t, 35, actual)
}

func TestPartTwo(t *testing.T) {
	inputStr := "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"

	actual := partTwo(inputStr)
	assert.Equal(t, 46, actual)
}

func TestParseSeedsPartOne(t *testing.T) {
	cases := map[string][]int{ // map input to expected output
		"seeds: 79 14 55 13":     {79, 14, 55, 13},
		"seeds: 79":              {79},
		"seeds: 1 1 1 1 1 1 1 1": {1, 1, 1, 1, 1, 1, 1, 1},
		"seeds: 79 14 : 55 13":   nil, // expect error
		"seeds: 71 a 14":         nil, // expect error
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			seeds, err := parseSeedsPartOne(input)
			assert.Equal(t, expected, seeds)
			if len(expected) == 0 {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestParseSeedsPartTwo(t *testing.T) {
	cases := map[string][]int{ // map input to expected output
		"seeds: 79 1 55 2":     {79, 55, 56},
		"seeds: 79 1":          {79},
		"seeds: 79 14 : 55 13": nil, // expect error
		"seeds: 71 a 14":       nil, // expect error
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			seeds, err := parseSeedsPartTwo(input)
			assert.Equal(t, expected, seeds)
			if len(expected) == 0 {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestParseMapping(t *testing.T) {
	emptyMapping := mapping{0, 0, 0}
	cases := map[string]mapping{ // map input to expected output
		"1 2 3":     mapping{1, 2, 3},
		"1 2 3 4":   {}, // expect error
		"1 2 three": mapping{},
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			m, err := parseMapping(input)
			assert.Equal(t, expected, m)
			if expected == emptyMapping {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestParseMappingSet(t *testing.T) {
	cases := map[string]mappingSet{ // map input to expected output
		"1 2 3\n4 5 6\n7 8 9": {
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9}},
		"1 2 3": {{1, 2, 3}},
		"foo-to-bar mapping:\n1 2 3\n4 5 6": {
			{1, 2, 3},
			{4, 5, 6}},
		"foo-to-bar mapping": {}, // expect error
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			ms, err := parseMappingSet(input)
			assert.Equal(t, expected, ms)
			if len(expected) == 0 { // "expected" is an empty mapping set, i.e. error case
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestParseMappingChain(t *testing.T) {
	input := `foo-to-bar mapping
1 2 3
1 2 3
1 2 3

bar-to-baz mapping:
4 4 4
8 8 8
12 12 12

baz-to-beep mapping:
1 1 1
`

	expected := mapChain{
		mappingSet{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
		mappingSet{{4, 4, 4}, {8, 8, 8}, {12, 12, 12}},
		mappingSet{{1, 1, 1}},
	}
	mc, err := parseMapChain(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, mc)

}

func TestMapping_SourceValWithinRange(t *testing.T) {
	noRangeMapping := mapping{0, 10, 1}
	largeRangeMapping := mapping{0, 10, 10}
	cases := []struct {
		name     string
		m        mapping
		val      int
		expected bool
	}{
		{"within trivial range", noRangeMapping, 10, true},
		{"outside trivial range", noRangeMapping, 11, false},
		{"start of range", largeRangeMapping, 10, true},
		{"middle of range", largeRangeMapping, 15, true},
		{"end of range", largeRangeMapping, 19, true},
		{"out of range - below", largeRangeMapping, 9, false},
		{"out of range - above", largeRangeMapping, 20, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, c.m.sourceValWithinRange(c.val))
		})
	}
}

func TestMapping_DestValForSource(t *testing.T) {
	m := mapping{100, 0, 10}
	cases := map[int]int{ // map input to expected output
		0:   100,
		5:   105,
		9:   109,
		10:  -1, // input val out of range so expect return of -1, !ok
		100: -1, // input val out of range so expect return of -1, !ok
	}

	for input, expected := range cases {
		t.Run(fmt.Sprintf("%d --> %d", input, expected), func(t *testing.T) {
			actual, ok := m.destValForSource(input)
			assert.Equal(t, expected, actual)
			if expected != -1 {
				assert.True(t, ok)
			} else {
				assert.False(t, ok)
			}
		})
	}
}

func TestMappingSet_MapVal(t *testing.T) {
	ms := mappingSet{
		mapping{50, 98, 2},
		mapping{52, 50, 48},
	}

	cases := []struct {
		name     string
		input    int
		expected int
	}{
		{"within first range", 99, 51},
		{"within second range", 60, 62},
		{"maps to self", 107, 107},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.expected, ms.mapVal(c.input))
		})
	}
}

func TestMapChain_MapVal(t *testing.T) {
	mc := mapChain{
		mappingSet{
			mapping{50, 98, 2},
			mapping{52, 50, 48},
		},
		mappingSet{
			mapping{0, 50, 10},
			mapping{90, 20, 4},
		},
		mappingSet{
			mapping{100, 72, 1},
			mapping{14, 0, 5},
		},
	}
	cases := map[int]int{ // map input to expected output
		99:    15,
		70:    100,
		4:     18,
		10000: 10000,
	}

	for input, expected := range cases {
		t.Run(fmt.Sprintf("%d --> %d", input, expected), func(t *testing.T) {
			assert.Equal(t, expected, mc.mapVal(input))
		})
	}

}
