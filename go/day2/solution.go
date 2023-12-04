package day1

import (
	"advent-of-code-2023/helper"
	"fmt"
	"strconv"
	"strings"
)

const day = "day2"

var (
	TestPart1File  = fmt.Sprintf("./input/%s/test1.in", day)
	TestPart2File  = fmt.Sprintf("./input/%s/test2.in", day)
	FinalPart1File = fmt.Sprintf("./input/%s/final1.in", day)
	FinalPart2File = fmt.Sprintf("./input/%s/final2.in", day)
)

type Boxes struct {
	Red, Green, Blue int
}

func ParseBoxInfo(s string) (string, int) {
	idx := strings.Index(s, " ")
	count, _ := strconv.Atoi(s[:idx])
	color := s[idx+1:]
	return color, count
}

func ParseGame(s string) (int, []Boxes) {
	result := make([]Boxes, 0, 5)

	ii := strings.Index(s, ":")
	id, _ := strconv.Atoi(s[5:ii])

	sets := strings.Split(s[ii+2:], "; ")
	for _, set := range sets {
		boxes := strings.Split(set, ", ")

		var parsedSet Boxes
		for _, box := range boxes {
			color, count := ParseBoxInfo(box)
			switch color {
			case "red":
				parsedSet.Red += count
			case "blue":
				parsedSet.Blue += count
			case "green":
				parsedSet.Green += count
			}
		}
		result = append(result, parsedSet)
	}
	return id, result
}

func IsGameValid(game []Boxes, redLimit, greenLimit, blueLimit int) bool {
	for _, set := range game {
		if set.Red > redLimit || set.Green > greenLimit || set.Blue > blueLimit {
			return false
		}
	}

	return true
}

func part1(file string) string {
	lines := helper.ReadLines(file)

	checksum := 0
	for idx, line := range lines {
		id, game := ParseGame(line)
		_ = idx
		if IsGameValid(game, 12, 13, 14) {
			checksum += id
		}
	}

	return fmt.Sprintf("%d", checksum)
}

func part2(file string) string {
	lines := helper.ReadLines(file)

	checksum := 0
	for _, line := range lines {
		_, game := ParseGame(line)

		var maxBox Boxes
		for _, boxes := range game {
			maxBox.Red = max(maxBox.Red, boxes.Red)
			maxBox.Green = max(maxBox.Green, boxes.Green)
			maxBox.Blue = max(maxBox.Blue, boxes.Blue)
		}

		checksum += maxBox.Red * maxBox.Green * maxBox.Blue
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
