package solver

// Solver is an AoC puzzle solver interface
type Solver interface {

	// ProcessInput of a puzzle.
	ProcessInput(content string) error

	// Part1 is the solution for the part 1 of the puzzle
	Part1() (string, error)

	// Part2 is the solution for the part 2 of the puzzle
	Part2() (string, error)
}
