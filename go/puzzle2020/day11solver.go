package puzzle2020

import (
	"aoc/solver"
	"fmt"
	"strconv"
	"strings"
)

// Area where the seats are located
type Area [][]rune

// Day11 structure
type Day11 struct {
	Area
	Occupied int
}

// NewDay11Solver constructs a new solver for day 11
func NewDay11Solver() solver.Solver {
	return &Day11{Area: [][]rune{}, Occupied: 0}
}

// ProcessInput of day 11
func (d *Day11) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	occupied := 0

	for _, line := range lines {
		d.Area = append(d.Area, []rune(line))
		occupied += strings.Count(line, "#")
	}

	d.Occupied = occupied

	return nil
}

// Part1 of day 11
func (d *Day11) Part1() (string, error) {

	return strconv.Itoa(d.Occupied), nil
}

// Part2 of day 11
func (d *Day11) Part2() (string, error) {
	changed := true
	for ; changed; changed = d.nextRoundV2() {
	}

	return strconv.Itoa(d.Occupied), nil
}

func (d *Day11) nextRound() bool {
	length := len(d.Area)
	changed := false
	newArea := make([][]rune, length)

	for i := 0; i < length; i++ {
		row := d.Area[i]
		rowLen := len(row)
		newArea[i] = make([]rune, rowLen)
		before, after := []rune{}, []rune{}

		if i-1 >= 0 {
			before = d.Area[i-1]
		}
		if i+1 < length {
			after = d.Area[i+1]
		}

		for j := 0; j < rowLen; j++ {
			newArea[i][j] = row[j]

			switch row[j] {
			case 'L':
				if checkAdjacents(j, before, row, after) == 0 {
					d.Occupied++
					newArea[i][j] = '#'
					changed = true
				}
			case '#':
				if checkAdjacents(j, before, row, after) > 3 {
					d.Occupied--
					newArea[i][j] = 'L'
					changed = true
				}
			default:
			}
		}
	}

	d.Area = newArea
	return changed
}

func checkAdjacents(pos int, before []rune, actual []rune, after []rune) int {
	count := 0

	for i := pos - 1; i <= pos+1; i++ {
		if i < 0 || i >= len(actual) {
			continue
		}
		if len(after) > 0 && after[i] == '#' {
			count++
		}
		if len(before) > 0 && before[i] == '#' {
			count++
		}
		if i != pos && actual[i] == '#' {
			count++
		}
	}

	return count
}

func (d *Day11) nextRoundV2() bool {
	length := len(d.Area)
	changed := false
	newArea := make([][]rune, length)

	for i := 0; i < length; i++ {
		row := d.Area[i]
		rowLen := len(row)
		newArea[i] = make([]rune, rowLen)

		for j := 0; j < rowLen; j++ {
			newArea[i][j] = row[j]

			switch row[j] {
			case 'L':
				if checkAdjacentsV2(i, j, d.Area) == 0 {
					d.Occupied++
					newArea[i][j] = '#'
					changed = true
				}
			case '#':
				if checkAdjacentsV2(i, j, d.Area) >= 5 {
					d.Occupied--
					newArea[i][j] = 'L'
					changed = true
				}
			default:
			}
		}
	}
	d.Area = newArea
	return changed
}

func checkAdjacentsV2(posY int, posX int, area Area) int {
	colLen := len(area)
	rowLen := len(area[posY])
	var val rune
	count := 0

	// Up
	for i := posY + 1; i < colLen; i++ {
		val = area[i][posX]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}
	// Down
	for i := posY - 1; i >= 0; i-- {
		val = area[i][posX]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}
	// Right
	for i := posX + 1; i < rowLen; i++ {
		val = area[posY][i]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}
	// Left
	for i := posX - 1; i >= 0; i-- {
		val = area[posY][i]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}

	// Diag bottom right
	for i, j := posY+1, posX+1; ; i, j = i+1, j+1 {
		if i >= colLen || j >= rowLen {
			break
		}

		val = area[i][j]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}
	// Ddiag bottom left
	for i, j := posY+1, posX-1; ; i, j = i+1, j-1 {
		if i >= colLen || j < 0 {
			break
		}

		val = area[i][j]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}
	// Diag top left
	for i, j := posY-1, posX-1; ; i, j = i-1, j-1 {
		if i < 0 || j < 0 {
			break
		}

		val = area[i][j]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}
	// Diag top right
	for i, j := posY-1, posX+1; ; i, j = i-1, j+1 {
		if i < 0 || j >= rowLen {
			break
		}

		val = area[i][j]
		if val == '#' {
			count++
			break
		} else if val == 'L' {
			break
		}
	}

	return count
}

func printMat(mat [][]rune) {
	for _, row := range mat {
		fmt.Printf("%v\n", string(row))
	}
	fmt.Println()
}
