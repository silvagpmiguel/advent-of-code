package puzzle

import (
	"aoc/solver"
)

// Day6 structure
type Day6 struct {
	Entries []int
	Target  int
}

// NewDay6Solver constructs a new solver for day 6
func NewDay6Solver() solver.Solver {
	return &Day6{Target: 2020}
}

// ProcessInput of day 6
func (d *Day6) ProcessInput(lines []string) error {
	return nil
}

// Part1 of day 6
func (d *Day6) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 6
func (d *Day6) Part2() (string, error) {
	return "None", nil
}
