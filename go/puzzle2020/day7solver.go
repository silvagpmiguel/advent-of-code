package puzzle2020

import (
	"aoc/solver"
	"fmt"
	"strconv"
	"strings"
)

// BagContent structure
type BagContent struct {
	Color string
	Count int
}

// Bag structure
type Bag struct {
	Content []BagContent
}

// Day7 structure
type Day7 struct {
	Bags   map[string]Bag // Color Bag
	Target string
}

// NewDay7Solver constructs a new solver for day 7
func NewDay7Solver() solver.Solver {
	return &Day7{Bags: make(map[string]Bag), Target: "shiny gold"}
}

// ProcessInput of day 7
func (d *Day7) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	for _, line := range lines {
		splitted := strings.Split(line, " bags contain ")
		color := splitted[0]
		splitted = strings.Split(splitted[1], ", ")
		bag := Bag{}

		for _, item := range splitted {
			var count int
			var err error
			if item == "no other bags." {
				continue
			} else {
				count, err = strconv.Atoi(string(item[0]))
				if err != nil {
					return fmt.Errorf("converting %v to integer: %v", string(item[0]), err)
				}
			}

			color := strings.TrimSuffix(item[2:], ".")
			color = strings.TrimSuffix(color, "bag")
			color = strings.TrimSpace(strings.TrimSuffix(color, "bags"))
			bag.Content = append(bag.Content, BagContent{
				Color: color,
				Count: count,
			})
		}

		d.Bags[color] = bag
	}

	return nil
}

// Part1 of day 7
func (d *Day7) Part1() (string, error) {
	count := 0

	for _, content := range d.Bags {
		for _, c := range content.Content {
			if checkForBag(d.Bags, c.Color, d.Target) {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count), nil
}

// Part2 of day 7
func (d *Day7) Part2() (string, error) {
	return strconv.Itoa(countInsideBags(d.Bags, d.Target)), nil
}

func countInsideBags(bags map[string]Bag, color string) int {
	res := 0

	for _, content := range bags[color].Content {
		res += content.Count + content.Count*countInsideBags(bags, content.Color)
	}

	return res
}

func checkForBag(bags map[string]Bag, color string, target string) bool {
	if strings.Contains(color, target) {
		return true
	}

	for _, content := range bags[color].Content {
		if strings.Contains(content.Color, target) || checkForBag(bags, content.Color, target) {
			return true
		}
	}

	return false
}
