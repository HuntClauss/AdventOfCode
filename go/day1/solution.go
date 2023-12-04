package day1

import (
	"advent-of-code-2023/helper"
	"fmt"
	"strings"
)

const day = "day1"

var (
	TestPart1File  = fmt.Sprintf("./input/%s/test1.in", day)
	TestPart2File  = fmt.Sprintf("./input/%s/test2.in", day)
	FinalPart1File = fmt.Sprintf("./input/%s/final1.in", day)
	FinalPart2File = fmt.Sprintf("./input/%s/final2.in", day)
)

const _0 = rune('0')
const _9 = rune('9')

func isNumber(c rune) (int, bool) {
	return int(c - _0), c >= _0 && c <= _9
}

const MinimalWordLength = 3

func isTextNumber(s string) (int, bool) {
	textNumbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	if len(s) < MinimalWordLength {
		return -1, false
	}

	for i, v := range textNumbers {
		if strings.HasPrefix(s, v) {
			return i, true
		}
	}
	return -1, false
}

func getNumbers(s string) []int {
	numbers := make([]int, 0, 15)
	for pos, _ := range s {
		if num, ok := isNumber(rune(s[pos])); ok {
			numbers = append(numbers, num)
		} else if num, ok = isTextNumber(s[pos:]); ok {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func part1(file string) string {
	input := helper.ReadInput(file)
	lines := strings.Split(input, "\n")

	checksum := 0
	for _, line := range lines {
		first, last := -1, -1
		for _, c := range line {
			if c < _0 || c > _9 {
				continue
			}
			if first == -1 {
				first = int(c - _0)
			}
			last = int(c - _0)
		}
		checksum += first*10 + last
	}

	return fmt.Sprintf("%d", checksum)
}

func part2(file string) string {
	input := helper.ReadInput(file)
	lines := strings.Split(input, "\n")

	checksum := 0
	for _, line := range lines {
		numbers := getNumbers(line)

		checksum += numbers[0]*10 + numbers[len(numbers)-1]
	}

	return fmt.Sprintf("%d", checksum)
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
