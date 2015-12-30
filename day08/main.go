package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/evalphobia/adventofcode/lib"
)

func main() {
	lib.ParseFlag()

	value := lib.GetValue()
	if value == "" {
		log.Fatal("cannot get value")
	}

	switch {
	case lib.GetPart() == 2:
		runPart2(value)
	default:
		runPart1(value)
	}
}

func runPart1(value string) {
	result := solvePart1(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart1(val string) int {
	lines := strings.Split(val, "\n")
	var literalCount, memoryCount int
	for _, line := range lines {
		literalCount += len(line)
		memoryCount += getInMemoryCount(line)
	}
	return literalCount - memoryCount
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	lines := strings.Split(val, "\n")
	var literalCount, escapedCount int
	for _, line := range lines {
		literalCount += len(line)
		escapedCount += getEscapedCount(line)
	}
	return escapedCount - literalCount
}

func getInMemoryCount(text string) int {
	converted, err := strconv.Unquote(text)
	if err != nil {
		log.Fatal("cannot convert text: %s, error=%s", text, err.Error())
	}
	return len(converted)
}

func getEscapedCount(text string) int {
	converted := strconv.Quote(text)
	return len(converted)
}
