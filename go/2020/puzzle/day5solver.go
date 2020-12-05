package puzzle

import (
	"aoc/solver"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Range structure
type Range struct {
	From float64
	To   float64
}

// Seat structure
type Seat struct {
	Rep string
	Row int
	Col int
	ID  int
}

// Day5 structure
type Day5 struct {
	Map     map[int][]Seat // Map: Row -> Seat List
	Biggest int
	NumRows int
	NumCols int
}

// NewDay5Solver constructs a new solver for day 5
func NewDay5Solver() solver.Solver {
	return &Day5{Map: make(map[int][]Seat), Biggest: 0, NumRows: 128, NumCols: 8}
}

// ProcessInput of day 5
func (d *Day5) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	for _, line := range lines {
		seat := computeSeat(line, float64(d.NumRows-1), float64(d.NumCols-1), 7, 3)
		id := seat.ID
		row := seat.Row
		d.Map[row] = append(d.Map[row], seat)
		if d.Biggest < id {
			d.Biggest = id
		}
	}

	return nil
}

// Part1 of day 5
func (d *Day5) Part1() (string, error) {
	return strconv.Itoa(d.Biggest), nil
}

// Part2 of day 5
func (d *Day5) Part2() (string, error) {
	for _, row := range d.Map {
		if len(row) == d.NumCols-1 {
			return strconv.Itoa(findMySeat(row)), nil
		}
	}
	return "Seat not found", nil
}

func findMySeat(seats []Seat) int {
	sort.Slice(seats, func(p, q int) bool {
		return seats[p].ID < seats[q].ID
	})

	for i := 0; i+1 < len(seats); i++ {
		actual := seats[i].ID
		next := seats[i+1].ID

		if actual+1 != next {
			return actual + 1
		}
	}
	return 0
}

func computeSeat(rep string, maxRow float64, maxCol float64, rowChars int, colChars int) Seat {
	var i int
	isLower := map[byte]bool{
		'F': true,
		'B': false,
		'L': true,
		'R': false,
	}
	rowRange := &Range{
		From: 0,
		To:   maxRow,
	}
	colRange := &Range{
		From: 0,
		To:   maxCol,
	}
	for i = 0; i < rowChars; i++ {
		rowRange = computeHalf(rowRange, isLower[rep[i]])
	}
	for j := i; j < rowChars+colChars; j++ {
		colRange = computeHalf(colRange, isLower[rep[j]])
	}

	row := int(rowRange.From)
	col := int(colRange.From)
	id := computeSeatID(row, col)

	return Seat{
		Rep: rep,
		Row: row,
		Col: col,
		ID:  id,
	}
}

func computeSeatID(row int, col int) int {
	return row*8 + col
}

func computeHalf(r *Range, isLower bool) *Range {
	var from float64
	var to float64
	if isLower {
		from = r.From
		to = math.Floor((from + r.To) / 2)
	} else {
		to = r.To
		from = math.Ceil((to + r.From) / 2)
	}
	return &Range{From: from, To: to}
}
