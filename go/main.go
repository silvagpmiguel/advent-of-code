package main

import (
	"aoc/puzzle2020"
	"aoc/service"
	"aoc/solver"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var solverMap = map[string]map[string]solver.Solver{
	"2020": map[string]solver.Solver{
		"1":  puzzle2020.NewDay1Solver(),
		"2":  puzzle2020.NewDay2Solver(),
		"3":  puzzle2020.NewDay3Solver(),
		"4":  puzzle2020.NewDay4Solver(),
		"5":  puzzle2020.NewDay5Solver(),
		"6":  puzzle2020.NewDay6Solver(),
		"7":  puzzle2020.NewDay7Solver(),
		"8":  puzzle2020.NewDay8Solver(),
		"9":  puzzle2020.NewDay9Solver(),
		"10": puzzle2020.NewDay10Solver(),
	},
}

func getSolver(year string, day string) (solver.Solver, error) {
	var yearSolvers map[string]solver.Solver
	var s solver.Solver
	var ok bool

	if yearSolvers, ok = solverMap[year]; !ok {
		return nil, fmt.Errorf("there is no solver for the year %v", year)
	}

	if s, ok = yearSolvers[day]; !ok {
		return nil, fmt.Errorf("there is no solver for day %v", day)
	}

	return s, nil
}

func main() {
	var input []byte
	args, err := service.NewArgs(len(os.Args), os.Args)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	endpoint := fmt.Sprintf("https://adventofcode.com/%s/day/%s/%s", args.Year, args.Day, args.Action)

	switch args.Action {
	case "solve":
		solv, err := getSolver(args.Year, args.Day)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		runners, err := solver.NewRunners(args.Filepath, solv)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		for _, runner := range runners {
			sol, err := solver.Run(&runner)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

			err = sol.PrintSolution(os.Stdout)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	case "answer":
		result, err := service.Post(endpoint, args.Cookie, args.Level, args.Answer)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		fmt.Println(result)
	case "input":
		input, err = service.Get(endpoint, args.Cookie)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		err = os.MkdirAll(filepath.Dir(args.Filepath), 0755)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		err = ioutil.WriteFile(args.Filepath, input, 0644)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		fmt.Println("Created AoC input file in " + args.Filepath)
	default:
		fmt.Println("Error: Wrong action.\nUsage: ./aoc \"input|answer|solve\" \"year\" \"day\" [level] [answer]")
		os.Exit(1)
	}

}
