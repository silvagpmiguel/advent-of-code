package puzzle2020

import (
	"aoc/solver"
	"strings"
)

// Day14 structure
type Day14 struct {
}

// NewDay14Solver constructs a new solver for day 14
func NewDay14Solver() solver.Solver {
	return &Day14{}
}

// ProcessInput of day 14
func (d *Day14) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	println(lines)
	return nil
}

// Part1 of day 14
func (d *Day14) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 14
func (d *Day14) Part2() (string, error) {
	return "None", nil
}
