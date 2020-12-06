package puzzle2020

import (
	"aoc/solver"
	"strconv"
	"strings"
)

// Day3 structure
type Day3 struct {
	Map    []string
	Length int
	DirX   []int
	DirY   []int
	Result []int
}

// NewDay3Solver constructs a new solver for day 3
func NewDay3Solver() solver.Solver {
	x := []int{1, 3, 5, 7, 1}
	y := []int{1, 1, 1, 1, 2}
	res := []int{0, 0, 0, 0, 0}
	return &Day3{Length: 5, DirX: x, DirY: y, Result: res}
}

// ProcessInput of day 3
func (d *Day3) ProcessInput(content string) error {
	d.Map = strings.Split(strings.TrimSpace(content), "\n")
	return nil
}

// Part1 of day 3
func (d *Day3) Part1() (string, error) {
	trees := 0
	l := len(d.Map[0])

	for i, j := 1, 3; i < len(d.Map); i, j = i+1, j+3 {
		if d.Map[i][j%l] == '#' {
			trees++
		}
	}

	return strconv.Itoa(trees), nil
}

// Part2 of day 3
func (d *Day3) Part2() (string, error) {
	l := len(d.Map[0])

	for k := 0; k < d.Length; k++ {
		stepX := d.DirX[k]
		stepY := d.DirY[k]
		for x, y := 0, 0; y < len(d.Map); x, y = x+stepX, y+stepY {
			if d.Map[y][x%l] == '#' {
				d.Result[k]++
			}
		}
	}

	return strconv.Itoa(mult(d.Result)), nil
}

func mult(arr []int) int {
	res := 1

	for _, num := range arr {
		res *= num
	}

	return res
}
