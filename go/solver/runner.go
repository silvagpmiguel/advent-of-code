package solver

import (
	"fmt"
	"io/ioutil"
	"time"
)

// Runner is the structure of a puzzle runner
type Runner struct {
	Input  []byte
	Solver Solver
}

// NewRunners creates new puzzle runners
func NewRunners(filepath string, solvers ...Solver) ([]Runner, error) {
	bytes, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("could not read from input file '%s': %w", filepath, err)
	}

	var runners []Runner

	for _, el := range solvers {
		runners = append(runners, Runner{
			Input:  bytes,
			Solver: el,
		})
	}

	return runners, nil
}

// Run a puzzle runner
func Run(r *Runner) (*Solution, error) {
	var s Solution

	err := r.Solver.ProcessInput(string(r.Input))

	if err != nil {
		return nil, err
	}

	p1Start := time.Now()
	p1, err := r.Solver.Part1()

	if err != nil {
		return nil, fmt.Errorf("solver could not solve part 1: %w", err)
	}

	s.Part1Time = time.Since(p1Start)
	s.Part1 = p1

	p2Start := time.Now()
	p2, err := r.Solver.Part2()

	if err != nil {
		return nil, fmt.Errorf("solver could not solve part 2: %w", err)
	}

	s.Part2Time = time.Since(p2Start)
	s.Part2 = p2

	return &s, nil
}
