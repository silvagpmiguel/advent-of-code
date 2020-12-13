package puzzle2020

import (
	"aoc/solver"
	"fmt"
	"strconv"
	"strings"
)

// BusNums type
type BusNums map[int]int

// Day13 structure
type Day13 struct {
	CanDepartAt int
	BusNums
}

// NewDay13Solver constructs a new solver for day 13
func NewDay13Solver() solver.Solver {
	return &Day13{BusNums: make(map[int]int)}
}

// ProcessInput of day 13
func (d *Day13) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	i, err := strconv.Atoi(lines[0])

	if err != nil {
		return fmt.Errorf("Couldn't cast %v to integer", lines[0])
	}

	d.CanDepartAt = i

	for i, val := range strings.Split(lines[1], ",") {
		if val == "x" {
			continue
		}
		busID, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("Couldn't cast %v to integer", val)
		}
		d.BusNums[busID] = i
	}

	return nil
}

// Part1 of day 13
func (d *Day13) Part1() (string, error) {
	waitTime, earliestBus := d.getEarliestBus()
	return strconv.Itoa(computeTime(waitTime, earliestBus)), nil
}

// Part2 of day 13
func (d *Day13) Part2() (string, error) {
	return strconv.Itoa(d.findConsecutive()), nil
}

func (d *Day13) getEarliestBus() (int, int) {
	waitTime := d.CanDepartAt
	earliestBus := -1

	for lineNum := range d.BusNums {
		minToWait := lineNum - (d.CanDepartAt % lineNum)
		if minToWait < waitTime {
			waitTime = minToWait
			earliestBus = lineNum
		}
	}

	return waitTime, earliestBus
}

func (d *Day13) findConsecutive() int {
	minValue := 0
	runningProduct := 1

	for k, v := range d.BusNums {
		for (minValue+v)%k != 0 {
			minValue += runningProduct
		}
		runningProduct *= k
	}

	return minValue
}

func computeTime(waitTime, earliestBus int) int {
	return earliestBus * waitTime
}
