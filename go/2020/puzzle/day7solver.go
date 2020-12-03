package puzzle

import (
	"aoc/solver"
)

// Day7 structure
type Day7 struct {
	Entries []int
	Target  int
}

// NewDay7Solver constructs a new solver for day 7
func NewDay7Solver() solver.Solver {
	return &Day7{Target: 2020}
}

// ProcessInput of day 7
func (d *Day7) ProcessInput(content string) error {
	return nil
}

// Part1 of day 7
func (d *Day7) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 7
func (d *Day7) Part2() (string, error) {
	return "None", nil
}
