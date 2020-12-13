package puzzle2020

import (
	"aoc/solver"
	"fmt"
	"strconv"
	"strings"
)

// Action type
type Action byte

// Action enum
const (
	North   = 'N'
	South   = 'S'
	East    = 'E'
	West    = 'W'
	Left    = 'L'
	Right   = 'R'
	Forward = 'F'
)

// NavInstruction structure
type NavInstruction struct {
	Action
	Value int
}

// Coordinates structure
type Coordinates struct {
	X NavInstruction
	Y NavInstruction
}

// NavInstructions type
type NavInstructions []NavInstruction

// Day12 structure
type Day12 struct {
	NavInstructions
	Coordinates *Coordinates
	Facing      Action
	Waypoint    *Coordinates
}

// NewDay12Solver constructs a new solver for day 12
func NewDay12Solver() solver.Solver {
	initialCoords := Coordinates{
		X: NavInstruction{Action: East, Value: 0},
		Y: NavInstruction{Action: North, Value: 0},
	}
	waypointCoords := Coordinates{
		X: NavInstruction{Action: East, Value: 10},
		Y: NavInstruction{Action: North, Value: 1},
	}
	return &Day12{Coordinates: &initialCoords, Facing: East, Waypoint: &waypointCoords}
}

func (d *Day12) resetDay() {
	initialCoords := Coordinates{
		X: NavInstruction{Action: East, Value: 0},
		Y: NavInstruction{Action: North, Value: 0},
	}
	d.Coordinates = &initialCoords
}

// ProcessInput of day 12
func (d *Day12) ProcessInput(content string) error {
	lines := strings.Split(strings.TrimSpace(content), "\n")

	for _, line := range lines {
		value, err := strconv.Atoi(line[1:len(line)])
		if err != nil {
			return err
		}
		inst := NavInstruction{Action: Action(line[0]), Value: value}
		d.NavInstructions = append(d.NavInstructions, inst)
	}
	return nil
}

// Part1 of day 12
func (d *Day12) Part1() (string, error) {
	for _, inst := range d.NavInstructions {
		d.refreshCoordinates(&inst)
	}

	return strconv.Itoa(d.computeManhattanDistance()), nil
}

// Part2 of day 12
func (d *Day12) Part2() (string, error) {
	d.resetDay()
	for _, inst := range d.NavInstructions {
		d.refreshCoordinates2(&inst)
	}
	fmt.Printf("%v\n", d)
	return strconv.Itoa(d.computeManhattanDistance()), nil
}

func (d *Day12) refreshCoordinates2(inst *NavInstruction) {
	action := inst.Action

	if action == North || action == South || action == East || action == West {
		d.Waypoint.goTo(inst, '0')
	} else if action == Left || action == Right {
		d.Waypoint.Y.Action = d.Waypoint.Y.turn(inst)
		d.Waypoint.X.Action = d.Waypoint.X.turn(inst)

		if inst.Value != 180 {
			aux := d.Waypoint.Y
			d.Waypoint.Y = d.Waypoint.X
			d.Waypoint.X = aux
		}
	} else {
		d.goForward(inst)
	}
}

func (d *Day12) goForward(inst *NavInstruction) {
	times := inst.Value
	waypoint := d.Waypoint
	waypointX := waypoint.X.Action
	waypointY := waypoint.Y.Action
	valX := waypoint.X.Value
	valY := waypoint.Y.Value

	if waypointY == d.Coordinates.Y.Action {
		d.Coordinates.Y.Value += times * valY
	} else {
		d.Coordinates.Y = computeWhenNegative(waypointY, d.Coordinates.Y, valY, times)
	}
	if waypointX == d.Coordinates.X.Action {
		d.Coordinates.X.Value += times * valX
	} else {
		d.Coordinates.X = computeWhenNegative(waypointX, d.Coordinates.X, valX, times)
	}
}

func (d *Day12) refreshCoordinates(inst *NavInstruction) {
	action := inst.Action

	if action == North || action == South || action == East || action == West {
		d.Coordinates.goTo(inst, d.Facing)
	} else if action == Left || action == Right {
		d.Facing = d.Facing.turn(inst)
	} else {
		d.Coordinates.goTo(inst, d.Facing)
	}
}

func (d *Day12) computeManhattanDistance() int {
	return d.Coordinates.X.Value + d.Coordinates.Y.Value
}

func (d *Day12) String() string {
	return fmt.Sprintf("%v\n%v", d.Coordinates.String(), d.NavInstructions)
}

/** Action Methods */

func (a Action) turn(inst *NavInstruction) Action {
	arr := []byte{'N', 'E', 'S', 'W'}

	if inst.Action == Left {
		arr = []byte{'N', 'W', 'S', 'E'}
	}

	pos := indexOf(arr, a)
	inc := inst.Value / 90

	return Action(arr[(pos+inc)%4])
}

func (a Action) String() string {
	return fmt.Sprintf("%c", a)
}

/** Coordinates Methods */

func (c *Coordinates) goTo(inst *NavInstruction, part1Facing Action) {
	val := inst.Value
	action := inst.Action
	facing := action

	if action == 'F' {
		facing = part1Facing
	}

	if isFacingRightWay(c, facing) {
		if facing == North || facing == South {
			c.Y.Value += val
		} else {
			c.X.Value += val
		}
	} else {
		if facing == North || facing == South {
			c.Y = computeWhenNegative(facing, c.Y, val, 1)
		} else {
			c.X = computeWhenNegative(facing, c.X, val, 1)
		}
	}
}

func (c *Coordinates) String() string {
	return fmt.Sprintf("Coordinates: (%c %v,%c %v)", c.X.Action, c.X.Value, c.Y.Action, c.Y.Value)
}

/** NavInstruction Methods */

func (n *NavInstruction) String() string {
	return fmt.Sprintf("(%c %v)", n.Action, n.Value)
}

/** NavInstructions Methods */

func (n NavInstructions) String() string {
	if len(n) == 0 {
		return "()"
	}

	str := n[0].String()

	for _, inst := range n[1:] {
		str += "\n" + inst.String()
	}

	return str
}

/** Aux Functions */

func isFacingRightWay(coords *Coordinates, facing Action) bool {
	return coords.X.Action == facing || coords.Y.Action == facing
}

func computeWhenNegative(action Action, inst NavInstruction, val int, times int) NavInstruction {
	realVal := val * times
	if inst.Value > realVal {
		inst.Value = inst.Value - realVal
		return inst
	}
	inst.Action = action
	inst.Value = realVal - inst.Value
	return inst
}

func indexOf(arr []byte, action2 Action) int {
	for ind, aux := range arr {
		action := Action(aux)
		if action2 == action {
			return ind
		}
	}
	return -1
}
