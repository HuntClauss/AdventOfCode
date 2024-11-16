package day1

import (
	"advent-of-code-2023/helper"
	"fmt"
	"log"
	"strconv"
)

const day = "day3"

var (
	TestPart1File  = fmt.Sprintf("./input/%s/test1.in", day)
	TestPart2File  = fmt.Sprintf("./input/%s/test2.in", day)
	FinalPart1File = fmt.Sprintf("./input/%s/final1.in", day)
	FinalPart2File = fmt.Sprintf("./input/%s/final2.in", day)
)

const (
	_0, _9 = rune('0'), rune('9')
)

func isSymbol(c rune) bool {
	return !isNumber(c) && c != '.'
}

func isNumber(c rune) bool {
	return c >= _0 && c <= _9
}

func getElem(matrix []string, row, col int) (rune, bool) {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) {
		return ' ', false
	}
	return rune(matrix[row][col]), true
}

func hasNeighbor(table []string, row, col int, checkFunc func(rune) bool) bool {
	neights := []helper.Pair[int]{{row, col - 1}, {row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1}, {row, col + 1}, {row + 1, col + 1}, {row + 1, col}, {row + 1, col - 1}}
	for _, n := range neights {
		char, ok := getElem(table, n.A, n.B)
		if ok && checkFunc(char) {
			return true
		}
	}
	return false
}

func getNumber(table []string, row, col int) int {
	start, stop := col, col
	for i := col; i >= 0; i-- {
		if !isNumber(rune(table[row][i])) {
			break
		}

		start = i
	}

	for i := col; i < len(table[row]); i++ {
		if !isNumber(rune(table[row][i])) {
			break
		}

		stop = i
	}

	num, err := strconv.Atoi(table[row][start : stop+1])
	if err != nil {
		log.Fatalln("cannot convert number:", err)
	}

	return num
}

func part1(file string) string {
	table := make([]string, 0, 100)

	for _, line := range helper.ReadLines(file) {
		table = append(table, line)
	}

	result := 0
	for i, row := range table {
		start, end, serial := -1, 0, false
		for j, char := range row {
			if !isNumber(char) {
				end = j
				if serial && start != -1 {
					num, err := strconv.Atoi(row[start:end])
					if err != nil {
						log.Fatalln("cannot convert number:", err)
					}
					result += num
				}
				serial = false
				continue
			}

			if start < end {
				start = j
			}

			if !serial && hasNeighbor(table, i, j, isSymbol) {
				serial = true
			}
		}
	}
	return fmt.Sprintf("%d", result)
}

func part2(file string) string {
	table := make([]string, 0, 100)

	for _, line := range helper.ReadLines(file) {
		table = append(table, line)
	}

	//for i, row := range table {
	//	for j, char := range row {
	//		if char != '*' {
	//			continue
	//		}
	//
	//
	//	}
	//}

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
