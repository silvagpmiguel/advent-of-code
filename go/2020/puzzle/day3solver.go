package puzzle

import (
	"aoc/solver"
)

// Day3 structure
type Day3 struct {
	Entries []int
	Target  int
}

// NewDay3Solver constructs a new solver for day 3
func NewDay3Solver() solver.Solver {
	return &Day3{Target: 2020}
}

// ProcessInput of day 3
func (d *Day3) ProcessInput(lines []string) error {
	return nil
}

// Part1 of day 3
func (d *Day3) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 3
func (d *Day3) Part2() (string, error) {
	return "None", nil
}
