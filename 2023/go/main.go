package main

import (
	day1 "advent-of-code-2023/day1"
	day2 "advent-of-code-2023/day2"
	day3 "advent-of-code-2023/day3"
	"fmt"
)

type Solver interface {
	SolveTest1() string
	SolveTest2() string
	SolvePart1() string
	SolvePart2() string
}

func main() {
	days := []Solver{day1.Init(), day2.Init(), day3.Init()}

	for i, day := range days {
		fmt.Printf("----[ DAY %1d ]----\n", i+1)
		fmt.Printf("  > Part 1 (test): %s\n", day.SolveTest1())
		fmt.Printf("  > Part 2 (test): %s\n", day.SolveTest2())
		fmt.Printf("  > Part 1       : %s\n", day.SolvePart1())
		fmt.Printf("  > Part 2       : %s\n", day.SolvePart2())
		fmt.Printf("-----------------\n")
	}
}
