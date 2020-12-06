package puzzle2020

import (
	"aoc/solver"
)

// Day9 structure
type Day9 struct {
	Entries []int
	Target  int
}

// NewDay9Solver constructs a new solver for day 9
func NewDay9Solver() solver.Solver {
	return &Day9{Target: 2020}
}

// ProcessInput of day 9
func (d *Day9) ProcessInput(content string) error {
	return nil
}

// Part1 of day 9
func (d *Day9) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 9
func (d *Day9) Part2() (string, error) {
	return "None", nil
}
