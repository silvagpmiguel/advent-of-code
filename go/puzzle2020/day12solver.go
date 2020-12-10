package puzzle2020

import (
	"aoc/solver"
	"strings"
)

// Day12 structure
type Day12 struct {
}

// NewDay12Solver constructs a new solver for day 12
func NewDay12Solver() solver.Solver {
	return &Day12{}
}

// ProcessInput of day 12
func (d *Day12) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	println(lines)
	return nil
}

// Part1 of day 12
func (d *Day12) Part1() (string, error) {
	return "None", nil
}

// Part2 of day 12
func (d *Day12) Part2() (string, error) {
	return "None", nil
}
