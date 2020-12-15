package puzzle2020

import (
	"aoc/solver"
	"strings"
)

// Day16 structure
type Day16 struct {
}

// NewDay16Solver constructs a new solver for day 16
func NewDay16Solver() solver.Solver {
	return &Day16{}
}

// ProcessInput of day 16
func (d *Day16) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), ",")
	println(lines)
	return nil
}

// Part1 of day 16
func (d *Day16) Part1() (string, error) {
	return "none", nil
}

// Part2 of day 16
func (d *Day16) Part2() (string, error) {
	return "none", nil
}
