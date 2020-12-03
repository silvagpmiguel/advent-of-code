package puzzle

import (
	"aoc/solver"
)

// Day8 structure
type Day8 struct {
	Entries []int
	Target  int
}

// NewDay8Solver constructs a new solver for day 8
func NewDay8Solver() solver.Solver {
	return &Day8{Target: 2020}
}

// ProcessInput of day 8
func (d *Day8) ProcessInput(content string) error {
	return nil
}

// Part1 of day 8
func (d *Day8) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 8
func (d *Day8) Part2() (string, error) {
	return "None", nil
}
