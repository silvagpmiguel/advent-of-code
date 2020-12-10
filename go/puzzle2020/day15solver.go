package puzzle2020

import (
	"aoc/solver"
	"strings"
)

// Day15 structure
type Day15 struct {
}

// NewDay15Solver constructs a new solver for day 15
func NewDay15Solver() solver.Solver {
	return &Day15{}
}

// ProcessInput of day 15
func (d *Day15) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	println(lines)
	return nil
}

// Part1 of day 15
func (d *Day15) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 15
func (d *Day15) Part2() (string, error) {
	return "None", nil
}
