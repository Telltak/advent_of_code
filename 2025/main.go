package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var solvers = make(map[int]Solver)

func RegisterSolver(day int, solver Solver) {
	solvers[day] = solver
}

func main() {
	day := flag.Int("day", 1, "Day number (1-25)")
	part := flag.Int("part", 1, "Part number (1-2)")
	flag.Parse()

	input, err := os.ReadFile(fmt.Sprintf("inputs/day%02d.txt", *day))
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	solver := solvers[*day]
	if solver == nil {
		log.Fatalf("Day %d not implemented", *day)
	}

	var result int
	if *part == 1 {
		result = solver.Part1(string(input))
	} else {
		result = solver.Part2(string(input))
	}

	fmt.Printf("Day %d Part %d: %d\n", *day, *part, result)
}
