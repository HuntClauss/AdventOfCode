package helper

import (
	"log"
	"os"
	"strings"
)

func ReadInput(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalln("cannot read input file:", err)
	}

	return string(data)
}

func ReadLines(filepath string) []string {
	return strings.Split(ReadInput(filepath), "\n")
}

type Pair[T any] struct {
	A, B T
}
