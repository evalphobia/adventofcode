package main

import (
	"fmt"
	"log"
	"sort"
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

	result := 0
	for _, line := range lines {
		result += calculatePaper(getDimentions(line))
	}
	return result
}

func getDimentions(line string) (int, int, int) {
	strD := strings.Split(line, "x")
	if len(strD) != 3 {
		log.Fatalf("cannot parse dimention: %s", line)
	}

	intD := make([]int, 3)
	for i, str := range strD {
		var err error
		intD[i], err = strconv.Atoi(str)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	switch {
	case intD[0] == 0, intD[1] == 0, intD[2] == 0:
		log.Fatalf("cannot have 0 feet length: %s", line)
	}

	return intD[0], intD[1], intD[2]
}

func calculatePaper(l, w, h int) int {
	d1 := l * w
	d2 := w * h
	d3 := h * l
	sum := (d1 + d2 + d3) * 2
	min := getMin(getMin(d1, d2), d3)
	return sum + min
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	lines := strings.Split(val, "\n")

	result := 0
	for _, line := range lines {
		result += calculateRibon(getDimentions(line))
	}
	return result
}

func calculateRibon(l, w, h int) int {
	d1, d2 := getMinTwo(l, w, h)
	sum := (d1 + d2) * 2
	bow := (l * w * h)
	return sum + bow
}

func getMinTwo(list ...int) (int, int) {
	sort.Ints(list)
	return list[0], list[1]
}
