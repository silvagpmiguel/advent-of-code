package puzzle2020

import (
	"aoc/solver"
	"strings"
)

// Day13 structure
type Day13 struct {
}

// NewDay13Solver constructs a new solver for day 13
func NewDay13Solver() solver.Solver {
	return &Day13{}
}

// ProcessInput of day 13
func (d *Day13) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	println(lines)
	return nil
}

// Part1 of day 13
func (d *Day13) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 13
func (d *Day13) Part2() (string, error) {
	return "None", nil
}
