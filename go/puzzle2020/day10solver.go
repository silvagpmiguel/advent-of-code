package puzzle2020

import (
	"aoc/solver"
)

// Day10 structure
type Day10 struct {
	Entries []int
	Target  int
}

// NewDay10Solver constructs a new solver for day 10
func NewDay10Solver() solver.Solver {
	return &Day10{Target: 2020}
}

// ProcessInput of day 10
func (d *Day10) ProcessInput(content string) error {
	return nil
}

// Part1 of day 10
func (d *Day10) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 10
func (d *Day10) Part2() (string, error) {
	return "None", nil
}
