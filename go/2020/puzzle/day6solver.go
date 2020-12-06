package puzzle

import (
	"aoc/solver"
	"strconv"
	"strings"
)

// Group structure
type Group struct {
	Questions map[rune]int
	Persons   int
}

// Day6 structure
type Day6 struct {
	Groups []Group
}

// NewDay6Solver constructs a new solver for day 6
func NewDay6Solver() solver.Solver {
	return &Day6{}
}

// ProcessInput of day 6
func (d *Day6) ProcessInput(content string) error {
	lines := strings.Split(content, "\n")
	group := Group{Questions: make(map[rune]int)}

	for _, line := range lines {
		if line == "" {
			d.Groups = append(d.Groups, group)
			group = Group{Questions: make(map[rune]int)}
			continue
		}
		for _, ch := range line {
			group.Questions[ch]++
		}
		group.Persons++
	}

	return nil
}

// Part1 of day 6
func (d *Day6) Part1() (string, error) {
	count := 0

	for _, group := range d.Groups {
		count += len(group.Questions)
	}

	return strconv.Itoa(count), nil
}

// Part2 of day 6
func (d *Day6) Part2() (string, error) {
	count := 0

	for _, group := range d.Groups {
		persons := group.Persons
		for _, c := range group.Questions {
			if c == persons {
				count++
			}
		}
	}

	return strconv.Itoa(count), nil
}
