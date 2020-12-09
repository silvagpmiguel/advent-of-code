package puzzle2020

import (
	"aoc/solver"
	"strconv"
	"strings"
)

// XMAS struct
type XMAS struct {
	Preamble int
	Numbers  []int
}

// Day9 structure
type Day9 struct {
	XMAS
}

// NewDay9Solver constructs a new solver for day 9
func NewDay9Solver() solver.Solver {
	return &Day9{XMAS{Preamble: 25}}
}

// ProcessInput of day 9
func (d *Day9) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	for _, line := range lines {
		num, err := strconv.Atoi(line)

		if err != nil {
			return err
		}

		d.Numbers = append(d.Numbers, num)
	}
	return nil
}

// Part1 of day 9
func (d *Day9) Part1() (string, error) {
	wrong := 0
	before := []int{}

	for ind, num := range d.XMAS.Numbers {
		before = append(before, num)
		if (ind + 1) > d.XMAS.Preamble {
			if !checkPreamble(num, before[ind-d.XMAS.Preamble:]) {
				wrong = num
				break
			}
		}
	}
	return strconv.Itoa(wrong), nil
}

// Part2 of day 9
func (d *Day9) Part2() (string, error) {
	wrong := 0
	before := []int{}
	smallest, largest, found := 0, 0, false

	for ind, num := range d.XMAS.Numbers {
		before = append(before, num)
		if (ind + 1) > d.XMAS.Preamble {
			if !checkPreamble(num, before[ind-d.XMAS.Preamble:]) {
				wrong = num
				break
			}
		}
	}

	for i := 2; i < len(d.XMAS.Numbers); i++ {
		smallest, largest, found = findWrongSumSet(d.XMAS.Numbers, wrong, i)
		if found {
			break
		}
	}

	return strconv.Itoa(smallest + largest), nil
}

func checkPreamble(num int, val []int) bool {
	for i := 0; i < len(val)-1; i++ {
		for j := i + 1; j < len(val); j++ {
			if val[i]+val[j] == num {
				return true
			}
		}
	}

	return false
}

func findWrongSumSet(numbers []int, wrong int, k int) (int, int, bool) {
	length := len(numbers)
	j := 0

	for i := 0; i < length-k; i++ {
		smallest, largest := numbers[i], numbers[i]
		sum := numbers[i]
		for j = i + 1; j < k+i; j++ {
			num := numbers[j]
			sum += num
			if smallest > num {
				smallest = num
			}
			if largest < num {
				largest = num
			}
			if sum == wrong {
				return smallest, largest, true
			}
		}
	}

	return 0, 0, false
}
