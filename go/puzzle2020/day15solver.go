package puzzle2020

import (
	"aoc/solver"
	"strconv"
	"strings"
)

// NumInfo struct
type NumInfo struct {
	Pos   int
	Count int
}

// Day15 structure
type Day15 struct {
	Numbers map[int]NumInfo // number -> pos found
	LastNum int
}

// NewDay15Solver constructs a new solver for day 15
func NewDay15Solver() solver.Solver {
	return &Day15{Numbers: make(map[int]NumInfo)}
}

// ProcessInput of day 15
func (d *Day15) ProcessInput(content string) error {
	i := 0
	lines := strings.Split(strings.TrimSpace(content), ",")
	for ; i < len(lines)-1; i++ {
		num, err := strconv.Atoi(lines[i])

		if err != nil {
			return err
		}

		d.Numbers[num] = NumInfo{Pos: i, Count: 1}
	}
	last, err := strconv.Atoi(lines[i])
	if err != nil {
		return err
	}
	d.LastNum = last
	return nil
}

// Part1 of day 15
func (d *Day15) Part1() (string, error) {
	guess := guessRound(d.cloneMap(), 2020, d.LastNum)
	return strconv.Itoa(guess), nil
}

// Part2 of day 15
func (d *Day15) Part2() (string, error) {
	guess := guessRound(d.Numbers, 30000000, d.LastNum)
	return strconv.Itoa(guess), nil
}

func guessRound(numbers map[int]NumInfo, turn int, last int) int {
	for i := len(numbers); i < turn-1; i++ {
		info, ok := numbers[last]

		if ok && info.Count+1 > 1 {
			numbers[last] = NumInfo{Pos: i, Count: info.Count + 1}
			last = i - info.Pos
		} else if info.Count == 1 {
			numbers[last] = NumInfo{Pos: info.Pos, Count: info.Count + 1}
			last = 0
		} else {
			numbers[last] = NumInfo{Pos: i, Count: 1}
			last = 0
		}
	}

	return last
}

func (d *Day15) cloneMap() map[int]NumInfo {
	newMap := make(map[int]NumInfo)
	for k, v := range d.Numbers {
		newMap[k] = v
	}
	return newMap
}
