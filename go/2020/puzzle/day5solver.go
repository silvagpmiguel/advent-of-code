package puzzle

import (
	"aoc/solver"
)

// Day5 structure
type Day5 struct {
	Entries []int
	Target  int
}

// NewDay5Solver constructs a new solver for day 5
func NewDay5Solver() solver.Solver {
	return &Day6{Target: 2020}
}

// ProcessInput of day 5
func (d *Day5) ProcessInput(lines []string) error {
	return nil
}

// Part1 of day 5
func (d *Day5) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 5
func (d *Day5) Part2() (string, error) {
	return "None", nil
}
