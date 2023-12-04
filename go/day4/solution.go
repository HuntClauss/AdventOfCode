package day1

import (
	"fmt"
)

const day = "day1"

var (
	TestPart1File  = fmt.Sprintf("./input/%s/test1.in", day)
	TestPart2File  = fmt.Sprintf("./input/%s/test2.in", day)
	FinalPart1File = fmt.Sprintf("./input/%s/final1.in", day)
	FinalPart2File = fmt.Sprintf("./input/%s/final2.in", day)
)

func part1(file string) string {
	return "TODO"
}

func part2(file string) string {
	return "TODO"
}

type ProblemSolution struct{}

func Init() *ProblemSolution {
	return &ProblemSolution{}
}

func (_ *ProblemSolution) SolveTest1() string {
	return part1(TestPart1File)
}

func (_ *ProblemSolution) SolvePart1() string {
	return part1(FinalPart1File)
}

func (_ *ProblemSolution) SolveTest2() string {
	return part2(TestPart2File)
}

func (_ *ProblemSolution) SolvePart2() string {
	return part2(FinalPart2File)
}
