package puzzle2020

import (
	"aoc/solver"
	"strings"
)

// Day17 structure
type Day17 struct {
}

// NewDay17Solver constructs a new solver for day 17
func NewDay17Solver() solver.Solver {
	return &Day17{}
}

// ProcessInput of day 17
func (d *Day17) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), ",")
	println(lines)
	return nil
}

// Part1 of day 17
func (d *Day17) Part1() (string, error) {
	return "none", nil
}

// Part2 of day 17
func (d *Day17) Part2() (string, error) {
	return "none", nil
}
