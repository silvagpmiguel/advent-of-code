package puzzle2020

import (
	"aoc/solver"
	"strconv"
	"strings"

	"github.com/silvagpmiguel/go-utils/utils"
)

// Chain of adaptors
type Chain struct {
	Adapters *utils.PQueue
	Exists   map[int]int
}

// Day10 structure
type Day10 struct {
	Chain
}

// NewDay10Solver constructs a new solver for day 10
func NewDay10Solver() solver.Solver {
	return &Day10{Chain{Adapters: utils.NewIntPQueue(true), Exists: make(map[int]int)}}
}

// ProcessInput of day 10
func (d *Day10) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	for _, line := range lines {
		adapter, err := strconv.Atoi(line)

		if err != nil {
			return err
		}

		d.Chain.Adapters = d.Chain.Adapters.Enqueue(adapter)
	}

	return nil
}

// Part1 of day 10
func (d *Day10) Part1() (string, error) {
	oneJoin, threeJoin := buildDiffs(d.Chain.Adapters.Clone(), 1, 3)
	return strconv.Itoa(oneJoin * threeJoin), nil
}

// Part2 of day 10
func (d *Day10) Part2() (string, error) {
	sorted := d.sort()
	arrangements := d.computeArrangements(sorted)
	return strconv.Itoa(arrangements), nil
}

func (d *Day10) computeArrangements(arr []int) int {
	length := len(arr)
	ways := make([]int, length)
	ways[length-1] = 1
	ret := 0

	for i := length - 2; i >= 0; i-- {
		sum := 0
		for diff := 1; diff <= 3; diff++ {
			if pos, ok := d.Chain.Exists[arr[i]+diff]; ok {
				sum += ways[pos]
			}
		}
		ways[i] = sum
	}

	for v := 1; v <= 3; v++ {
		if pos, ok := d.Chain.Exists[v]; ok {
			ret += ways[pos]
		}
	}

	return ret
}

func (d *Day10) sort() []int {
	sorted := []int{}

	for ind := 0; !d.Chain.Adapters.IsEmpty(); ind++ {
		interf, err := d.Chain.Adapters.Dequeue()
		jolt := interf.(int)

		if err != nil {
			return nil
		}

		d.Chain.Exists[jolt] = ind
		sorted = append(sorted, jolt)
	}

	return sorted
}

func buildDiffs(p *utils.PQueue, diff1 int, diff2 int) (int, int) {
	firstDiff, secondDiff, first := 0, 0, 0

	for i := 1; !p.IsEmpty(); i++ {
		interf, err := p.Dequeue()
		second := interf.(int)

		if err != nil {
			return 0, 0
		}

		diff := second - first

		if diff == diff1 {
			firstDiff++
		} else if diff == diff2 {
			secondDiff++
		}

		first = second
	}

	secondDiff++
	return firstDiff, secondDiff
}
