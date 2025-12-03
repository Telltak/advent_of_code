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

func init() {
    RegisterSolver(XX, &DayXX{})
}

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

2. Add your input to `inputs/dayXX.txt`

That's it! The day auto-registers via `init()` - no need to modify `main.go`.

## Structure

- `main.go` - Entry point, handles CLI args and routing
- `solver.go` - Interface that all days implement
- `dayXX.go` - Individual day solutions
- `utils.go` - Common helper functions
- `inputs/` - Puzzle inputs
