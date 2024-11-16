package day1

import (
	"advent-of-code-2023/helper"
	"fmt"
	"strings"
)

const day = "day4"

var (
	TestPart1File  = fmt.Sprintf("./input/%s/test1.in", day)
	TestPart2File  = fmt.Sprintf("./input/%s/test2.in", day)
	FinalPart1File = fmt.Sprintf("./input/%s/final1.in", day)
	FinalPart2File = fmt.Sprintf("./input/%s/final2.in", day)
)

type game struct {
	WinningCards []int
	PlayerCards  []int
}

func parseCardInput(line string) game {
	idx := strings.Index(line, ":")
	parts := strings.Split(line[idx+2:], "|")

	var result game

	for i := 0; i < len(parts[0])-2; i += 3 {
		v := strings.TrimSpace(parts[0][i : i+2])
		result.WinningCards = append(result.WinningCards, helper.MustParseInt(v))
	}

	for i := 1; i < len(parts[1]); i += 3 {
		v := strings.TrimSpace(parts[1][i : i+2])
		result.PlayerCards = append(result.PlayerCards, helper.MustParseInt(v))
	}

	return result
}

func part1(file string) string {
	lines := helper.ReadLines(file)

	result := 0

	// Well, I know the numbers are in range <1, 100>
	hashTable := make([]uint8, 101)
	for idx, line := range lines {
		r := parseCardInput(line)
		currentCounter := uint8(idx + 1)

		for _, winning := range r.WinningCards {
			hashTable[winning] = currentCounter
		}

		score := 1
		for _, player := range r.PlayerCards {
			if hashTable[player] == currentCounter {
				score *= 2
			}
		}

		if score > 1 {
			result += score / 2
		}
	}

	return fmt.Sprintf("%d", result)
}

func part2(file string) string {
	lines := helper.ReadLines(file)

	var result int
	copies := make([]int, len(lines))

	// Well, I know the numbers are in range <1, 100>
	hashTable := make([]int, 101)
	for idx, line := range lines {
		r := parseCardInput(line)
		currentCounter := idx + 1

		for _, winning := range r.WinningCards {
			hashTable[winning] = currentCounter
		}

		wins := 0
		for _, player := range r.PlayerCards {
			if hashTable[player] == currentCounter {
				wins++
				copies[idx+wins] += copies[idx] + 1
			}
		} //   > Part 2       : 55768, 9880422,
		result += wins * (copies[idx] + 1)
	}
	result += len(lines)

	//for i, v := range copies {
	//	fmt.Printf("%03d: %d\n", i+1, v)
	//}

	return fmt.Sprintf("%d", result)
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
	//return "TODO"
	return part2(FinalPart2File)
}
