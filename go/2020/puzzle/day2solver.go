package puzzle

import (
	"aoc/solver"
	"strconv"
	"strings"
)

/**
For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
*/

// Info for day 2
type Info struct {
	Policy1  int
	Policy2  int
	Letter   byte
	Password string
}

// Day2 Array of Info
type Day2 []Info

// NewDay2Solver constructs a new solver for day 2
func NewDay2Solver() solver.Solver {
	return &Day2{}
}

// ProcessInput of day 2
func (d *Day2) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	for _, line := range lines {
		splitted := strings.Split(line, " ")
		policy := strings.Split(splitted[0], "-")
		policy1, err := strconv.Atoi(policy[0])

		if err != nil {
			return err
		}

		policy2, err := strconv.Atoi(policy[1])

		if err != nil {
			return err
		}

		*d = append(*d, Info{
			Policy1:  policy1,
			Policy2:  policy2,
			Letter:   splitted[1][0],
			Password: splitted[2],
		})
	}

	return nil
}

// Part1 of day 2
func (d *Day2) Part1() (string, error) {
	counter := 0

	for _, info := range *d {
		count := strings.Count(info.Password, string(info.Letter))

		if count >= info.Policy1 && count <= info.Policy2 {
			counter++
		}
	}

	return strconv.Itoa(counter), nil
}

// Part2 of day 2
func (d *Day2) Part2() (string, error) {
	counter := 0

	for _, info := range *d {
		password := info.Password
		letter := info.Letter
		check := 0
		if password[info.Policy1-1] == letter {
			check++
		}
		if password[info.Policy2-1] == letter {
			check++
		}
		if check == 1 {
			counter++
		}
	}

	return strconv.Itoa(counter), nil
}
