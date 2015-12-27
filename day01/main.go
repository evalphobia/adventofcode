package main

import (
	"fmt"
	"log"
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
	up := strings.Count(val, "(")
	down := strings.Count(val, ")")
	return up - down
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	floor := 0
	for i, v := range val {
		floor += getUpOrDown(string(v))
		if floor < 0 {
			return i + 1
		}
	}
	log.Fatalf("does not reach basement")
	return 0
}

func getUpOrDown(v string) int {
	switch v {
	case "(":
		return 1
	case ")":
		return -1
	}
	log.Fatalf("cannot parse value: %s", v)
	return 0
}
