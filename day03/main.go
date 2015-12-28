package main

import (
	"fmt"
	"log"

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
	var x, y int
	totalHouse := 1
	locMap := make(map[int]map[int]int)
	locMap[0] = make(map[int]int)
	locMap[0][0] = 1
	for _, v := range val {
		switch string(v) {
		case ">":
			x++
		case "<":
			x--
		case "^":
			y++
		case "v":
			y--
		default:
			log.Fatalf("cannot parse: %s", v)
		}
		if locMap[x] == nil {
			locMap[x] = make(map[int]int)
		}
		if locMap[x][y] == 0 {
			totalHouse++
		}
		locMap[x][y]++
	}
	return totalHouse
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	x := make([]int, 2)
	y := make([]int, 2)
	totalHouse := 1
	locMap := make(map[int]map[int]int)
	locMap[0] = make(map[int]int)
	locMap[0][0] = 1
	for i, v := range val {
		santaOrRobo := i % 2
		switch string(v) {
		case ">":
			x[santaOrRobo]++
		case "<":
			x[santaOrRobo]--
		case "^":
			y[santaOrRobo]++
		case "v":
			y[santaOrRobo]--
		default:
			log.Fatalf("cannot parse: %s", v)
		}
		if locMap[x[santaOrRobo]] == nil {
			locMap[x[santaOrRobo]] = make(map[int]int)
		}
		if locMap[x[santaOrRobo]][y[santaOrRobo]] == 0 {
			totalHouse++
		}
		locMap[x[santaOrRobo]][y[santaOrRobo]]++
	}
	return totalHouse
}
