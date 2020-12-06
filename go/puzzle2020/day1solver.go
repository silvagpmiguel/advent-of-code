package puzzle2020

import (
	"aoc/solver"
	"fmt"
	"strconv"
	"strings"
)

// Day1 structure
type Day1 struct {
	Entries []int
	Target  int
}

// NewDay1Solver constructs a new solver for day 1
func NewDay1Solver() solver.Solver {
	return &Day1{Target: 2020}
}

// ProcessInput of day 1
func (d *Day1) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")

	for _, line := range lines {
		i, err := strconv.Atoi(line)

		if err != nil {
			return fmt.Errorf("couldn't convert %s to int : %v", line, err)
		}

		d.Entries = append(d.Entries, i)
	}

	return nil
}

// Part1 of day 1
func (d *Day1) Part1() (string, error) {
	var entry1, entry2 int
	leng := len(d.Entries)

	for i := 0; i < leng-1; i++ {
		for j := i + 1; j < leng; j++ {
			entry1 = d.Entries[i]
			entry2 = d.Entries[j]

			if entry1+entry2 == d.Target {
				return strconv.Itoa(entry1 * entry2), nil
			}
		}
	}

	return "None", nil
}

// Part2 of day 1
func (d *Day1) Part2() (string, error) {
	var entry1, entry2, entry3 int
	leng := len(d.Entries)

	for i := 0; i < leng-2; i++ {
		for j := i + 1; j < leng-1; j++ {
			for k := i + 2; k < leng; k++ {
				entry1 = d.Entries[i]
				entry2 = d.Entries[j]
				entry3 = d.Entries[k]

				if entry1+entry2+entry3 == d.Target {
					return strconv.Itoa(entry1 * entry2 * entry3), nil
				}
			}
		}
	}

	return "None", nil
}
