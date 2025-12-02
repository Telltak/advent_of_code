# Advent of Code 2025

## Usage

Run a specific day and part:
```bash
go run . -day 1 -part 1
```

## Adding a New Day

1. Create `dayXX.go` with the structure:
```go
package main

type DayXX struct{}

func (d *DayXX) Part1(input string) int {
    // Your solution
    return 0
}

func (d *DayXX) Part2(input string) int {
    // Your solution
    return 0
}
```

2. Add your day to the `getSolver()` map in `main.go`:
```go
solvers := map[int]Solver{
    1: &Day01{},
    XX: &DayXX{},  // Add this line
}
```

3. Add your input to `inputs/dayXX.txt`

## Structure

- `main.go` - Entry point, handles CLI args and routing
- `solver.go` - Interface that all days implement
- `dayXX.go` - Individual day solutions
- `utils.go` - Common helper functions
- `inputs/` - Puzzle inputs
