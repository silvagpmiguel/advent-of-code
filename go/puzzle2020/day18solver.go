package puzzle2020

import (
	"aoc/solver"
	"strings"
)

// Day18 structure
type Day18 struct {
}

// NewDay18Solver constructs a new solver for day 18
func NewDay18Solver() solver.Solver {
	return &Day18{}
}

// ProcessInput of day 18
func (d *Day18) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), ",")
	println(lines)
	return nil
}

// Part1 of day 18
func (d *Day18) Part1() (string, error) {
	return "none", nil
}

// Part2 of day 18
func (d *Day18) Part2() (string, error) {
	return "none", nil
}
