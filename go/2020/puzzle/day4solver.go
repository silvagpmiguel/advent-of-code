package puzzle

import (
	"aoc/solver"
)

// Day4 structure
type Day4 struct {
	Entries []int
	Target  int
}

// NewDay4Solver constructs a new solver for day 4
func NewDay4Solver() solver.Solver {
	return &Day4{Target: 2020}
}

// ProcessInput of day 4
func (d *Day4) ProcessInput(lines []string) error {
	return nil
}

// Part1 of day 4
func (d *Day4) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 4
func (d *Day4) Part2() (string, error) {
	return "None", nil
}
