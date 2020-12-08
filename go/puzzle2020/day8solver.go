package puzzle2020

import (
	"aoc/solver"
	"fmt"
	"strconv"
	"strings"
)

// Instruction structure
type Instruction struct {
	IP    int
	Key   string
	Value int
}

// Day8 structure
type Day8 struct {
	IP           int
	Stacktrace   map[int]int // IP - Map[IP]
	Instructions []Instruction
	Acc          int
	Swap         []Instruction
	IsFixed      bool
}

// NewDay8Solver constructs a new solver for day 8
func NewDay8Solver() solver.Solver {
	return &Day8{IsFixed: false}
}

// ProcessInput of day 8
func (d *Day8) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")
	for i, line := range lines {
		splitted := strings.Split(line, " ")
		key := splitted[0]
		val, err := strconv.Atoi(splitted[1])

		if err != nil {
			return fmt.Errorf("Couldn't cast to int the instruction %v with value %v: %v", key, val, err)
		}

		inst := Instruction{
			IP:    i,
			Key:   key,
			Value: val,
		}
		d.Instructions = append(d.Instructions, inst)

		if key == "nop" || key == "jmp" {
			d.Swap = append(d.Instructions, inst)
		}

	}
	return nil
}

// Part1 of day 8
func (d *Day8) Part1() (string, error) {
	err := execProgram(d)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return strconv.Itoa(d.Acc), nil
	}

	return strconv.Itoa(d.Acc), nil
}

// Part2 of day 8
func (d *Day8) Part2() (string, error) {
	for i := 0; !d.IsFixed; i++ {
		bruteForce(d, i)
	}

	err := execProgram(d)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return strconv.Itoa(d.Acc), nil
}

// Reset Day8
func (d *Day8) Reset() {
	d.Acc = 0
	d.IP = 0
	d.Stacktrace = make(map[int]int)
}

func execProgram(d *Day8) error {
	d.Reset()
	insts := d.Instructions
	stacktrace := d.Stacktrace

	for d.IP < len(d.Instructions) {
		key := insts[d.IP].Key
		val := insts[d.IP].Value

		if stacktrace[d.IP] > 1 {
			return fmt.Errorf("Infinite Loop : %v", insts[d.IP])
		}

		switch key {
		case "acc":
			d.Acc += val
			d.IP++
		case "jmp":
			d.IP += val
			if d.IP < 0 {
				return fmt.Errorf("seg fault : accessing position %v in the memory", d.IP)
			}
		default:
			d.IP++
		}

		stacktrace[d.IP]++
	}

	return nil
}

func bruteForce(d *Day8, ind int) {
	d.Reset()
	insts := d.Instructions
	stacktrace := d.Stacktrace
	toSwap := d.Swap[ind]
	toSwapIP := toSwap.IP

	if insts[toSwapIP].Key == "nop" {
		insts[toSwapIP].Key = "jmp"
	} else {
		insts[toSwapIP].Key = "nop"
	}

	for d.IP < len(d.Instructions) {
		key := insts[d.IP].Key
		val := insts[d.IP].Value

		if stacktrace[d.IP] > 1 {
			insts[toSwap.IP] = toSwap
			return
		}

		switch key {
		case "acc":
			d.IP++
		case "jmp":
			d.IP += val
		default:
			d.IP++
		}

		stacktrace[d.IP]++
	}

	d.IsFixed = true
}

func contains(ip int, ips []int) bool {
	for _, val := range ips {
		if val == ip {
			return true
		}
	}
	return false
}
