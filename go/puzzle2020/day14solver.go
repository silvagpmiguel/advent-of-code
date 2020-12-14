package puzzle2020

import (
	"aoc/solver"
	"fmt"
	"strconv"
	"strings"
)

// Item type
type Item struct {
	MaskOn  uint64
	MaskOff uint64
	Mask    string
	Key     string
	Value   uint64
}

// Day14 structure
type Day14 struct {
	Nums []Item
}

// NewDay14Solver constructs a new solver for day 14
func NewDay14Solver() solver.Solver {
	return &Day14{}
}

// ProcessInput of day 14
func (d *Day14) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	maskon, maskoff := uint64(0), uint64(0)
	var mask string
	var err error

	for _, line := range lines {
		splitted := strings.Split(line, " = ")
		switch splitted[0] {
		case "mask":
			maskon, err = strconv.ParseUint(strings.ReplaceAll(splitted[1], "X", "1"), 2, 64)
			if err != nil {
				return err
			}
			maskoff, err = strconv.ParseUint(strings.ReplaceAll(splitted[1], "X", "0"), 2, 64)
			if err != nil {
				return err
			}
			mask = splitted[1]
		default:
			ind := splitted[0][4 : len(splitted[0])-1]
			val, err := strconv.ParseUint(splitted[1], 10, 64)

			if err != nil {
				return err
			}

			d.Nums = append(d.Nums, Item{Key: ind, Value: val, MaskOn: maskon, MaskOff: maskoff, Mask: mask})
		}
	}
	return nil
}

// Part1 of day 14
func (d *Day14) Part1() (string, error) {
	sum := uint64(0)
	heap := make(map[string]uint64)

	for _, val := range d.Nums {
		heap[val.Key] = applyMask(val.Value, val.MaskOn, val.MaskOff)
	}

	for _, num := range heap {
		sum += num
	}

	return strconv.FormatUint(sum, 10), nil
}

// Part2 of day 14
func (d *Day14) Part2() (string, error) {
	sum := uint64(0)
	heap := make(map[uint64]uint64)

	for _, val := range d.Nums {
		addr, err := strconv.ParseUint(val.Key, 10, 64)

		if err != nil {
			return "", err
		}

		computedValue := applyMask2(addr, val.Mask)
		combs := maskCombs(computedValue)

		for _, maskedAddrString := range combs {
			maskedAddr, _ := strconv.ParseUint(maskedAddrString, 2, 64)
			heap[maskedAddr] = val.Value
		}
	}

	for _, num := range heap {
		sum += num
	}
	return strconv.FormatUint(sum, 10), nil
}

func applyMask(val uint64, maskon uint64, maskoff uint64) uint64 {
	fmt.Println(val, (val|maskoff)&maskon)
	return (val | maskoff) & maskon
}

func applyMask2(val uint64, mask string) string {
	aux := fmt.Sprintf("%036b", val)

	var numBuilder strings.Builder
	for i := range aux {
		switch mask[i] {
		case '1':
			numBuilder.WriteRune('1')
		case 'X':
			numBuilder.WriteRune('X')
		case '0':
			numBuilder.WriteRune(rune(aux[i]))
		}
	}

	return numBuilder.String()
}

func maskCombs(mask string) []string {
	if len(mask) == 0 {
		return []string{""}
	}

	var res []string
	toMerge := maskCombs(mask[1:])

	switch mask[0] {
	case 'X':
		for _, m := range toMerge {
			res = append(res, "1"+m)
			res = append(res, "0"+m)
		}
	case '1':
		for _, m := range toMerge {
			res = append(res, "1"+m)
		}
	case '0':
		for _, m := range toMerge {
			res = append(res, "0"+m)
		}
	default:
	}

	return res
}
