package puzzle

import (
	"aoc/solver"
	"fmt"
	"strconv"
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

// ReadInput
func (d *Day1) ProcessInput(lines []string) error {
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("couldn't convert %s to int : %v", line, err)
		}

		d.Entries = append(d.Entries, i)
	}

	return nil
}

func (d *Day1) Part1() (string, error) {
	var arr []int
	entry1, entry2 := -1, -1
	for _, fst := range d.Entries {
		for _, snd := range arr {
			if fst+snd == d.Target {
				entry1 = fst
				entry2 = snd
				break
			}
		}
		if entry1 != -1 {
			break
		}
		arr = append(arr, fst)
	}
	return fmt.Sprintf("%d", entry1*entry2), nil
}

func (d *Day1) Part2() (string, error) {
	entry1, entry2, entry3 := -1, -1, -1
	leng := len(d.Entries)
	for i := 0; i < leng-2 && entry1+entry2+entry3 != d.Target; i++ {
		for j := i + 1; j < leng-1 && entry1+entry2+entry3 != d.Target; j++ {
			for k := i + 2; k < leng && entry1+entry2+entry3 != d.Target; k++ {
				entry1 = d.Entries[i]
				entry2 = d.Entries[j]
				entry3 = d.Entries[k]
			}
		}
	}

	return fmt.Sprintf("%d", entry1*entry2*entry3), nil
}
