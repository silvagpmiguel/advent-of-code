package solver

import (
	"fmt"
	"io"
	"time"
)

// Solution is the structure that stores the details of a puzzle solution
type Solution struct {
	ProcessingTime time.Duration
	Part1          string
	Part1Time      time.Duration
	Part2Time      time.Duration
	Part2          string
}

// PrintSolution to System.Out
func (s *Solution) PrintSolution(w io.Writer) error {
	writer := NewErrorTolerantWriter(w)

	fmt.Fprintf(w, "Processing: %v\n", s.ProcessingTime)

	fmt.Fprintf(w, "Part 1: %s (in %v)\n", s.Part1, s.Part1Time)

	fmt.Fprintf(w, "Part 2: %s (in %v)\n", s.Part2, s.Part2Time)

	return writer.Error()
}
