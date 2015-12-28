package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/evalphobia/adventofcode/lib"
)

var (
	badStrings = []string{"ab", "cd", "pq", "xy"}
	vowels     = []string{"a", "e", "i", "o", "u"}
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
	var niceCount int

	lines := strings.Split(val, "\n")
	for _, line := range lines {
		switch {
		case !hasThreeVowels(line):
			continue
		case !hasTwice(line):
			continue
		case hasBadStrings(line):
			continue
		}
		niceCount++
	}
	return niceCount
}

func hasThreeVowels(text string) bool {
	var count int
	for _, v := range vowels {
		count += strings.Count(text, v)
		if count >= 3 {
			return true
		}
	}
	return false
}

func hasTwice(text string) bool {
	var prev rune
	for _, v := range text {
		if prev == v {
			return true
		}
		prev = v
	}
	return false
}

func hasBadStrings(text string) bool {
	for _, v := range badStrings {
		if strings.Contains(text, v) {
			return true
		}
	}
	return false
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	var niceCount int

	lines := strings.Split(val, "\n")
	for _, line := range lines {
		switch {
		case !hasTwicePair(line):
			continue
		case !hasSandwitchLetter(line):
			continue
		}
		niceCount++
	}
	return niceCount
}

func hasTwicePair(text string) bool {
	for i := range text {
		if i == 0 {
			continue
		}
		pair := string(text[i-1]) + string(text[i])
		if strings.Count(text, pair) >= 2 {
			return true
		}
	}
	return false
}

func hasSandwitchLetter(text string) bool {
	for i := range text {
		if i == 0 || i == 1 {
			continue
		}
		if text[i-2] == text[i] {
			return true
		}
	}
	return false
}
