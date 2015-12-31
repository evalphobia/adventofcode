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
	for i := 0; i < 40; i++ {
		val = lookAndSay(val)
	}
	return len(val)
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	for i := 0; i < 50; i++ {
		val = lookAndSay(val)
	}
	return len(val)
}

func lookAndSay(text string) string {
	var result []string
	streak := 1
	for i, max := 1, len(text); i <= max; i++ {
		prev := text[i-1]
		if i == max {
			result = append(result, fmt.Sprintf("%d%s", streak, string(prev)))
			break
		}

		cur := text[i]
		if prev == cur {
			streak++
			continue
		}

		result = append(result, fmt.Sprintf("%d%s", streak, string(prev)))
		streak = 1
	}
	return strings.Join(result, "")
}
