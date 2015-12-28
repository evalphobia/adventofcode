package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/evalphobia/adventofcode/lib"
)

var lightGrid [1000][1000]bool
var brightGrid [1000][1000]int

func init() {
	lightGrid = [1000][1000]bool{}
	brightGrid = [1000][1000]int{}
}

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
	for _, line := range lines {
		fn := getLightAction(line)
		x1, y1 := getRangeFrom(line)
		x2, y2 := getRangeTo(line)
		fn(x1, y1, x2, y2)
	}
	return countLights()
}

func getLightAction(text string) func(int, int, int, int) {
	switch {
	case strings.HasPrefix(text, "turn on"):
		return turnOnLight
	case strings.HasPrefix(text, "turn off"):
		return turnOffLight
	case strings.HasPrefix(text, "toggle"):
		return toggleLight
	}
	log.Fatalf("cannot parse light action: %s", text)
	return nil
}

func turnOnLight(x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j] = true
		}
	}
}

func turnOffLight(x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j] = false
		}
	}
}

func toggleLight(x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			lightGrid[i][j] = !lightGrid[i][j]
		}
	}
}

func getRangeFrom(text string) (int, int) {
	return getRange(text, 0)
}

func getRangeTo(text string) (int, int) {
	return getRange(text, 1)
}

func getRange(text string, i int) (int, int) {
	if i != 0 && i != 1 {
		log.Fatalf("illegal range is set: %d", i)
	}

	lines := strings.Split(text, " through ")
	if len(lines) != 2 {
		log.Fatalf("cannot parse line: %s", text)
	}

	lines = strings.Split(lines[i], " ")
	ranges := strings.Split(lines[len(lines)-1], ",")
	if len(ranges) != 2 {
		log.Fatalf("cannot parse ranges: %s", text)
	}

	x, err := strconv.Atoi(ranges[0])
	if err != nil {
		log.Fatalf("cannot parse range to int: %v, error=%s", ranges, err.Error())
	}
	y, err := strconv.Atoi(ranges[1])
	if err != nil {
		log.Fatalf("cannot parse range to int: %v, error=%s", ranges, err.Error())
	}
	return x, y
}

func countLights() int {
	var count int
	for _, rows := range lightGrid {
		for _, grid := range rows {
			if grid == true {
				count++
			}
		}
	}
	return count
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	lines := strings.Split(val, "\n")
	for _, line := range lines {
		fn := getBrightAction(line)
		x1, y1 := getRangeFrom(line)
		x2, y2 := getRangeTo(line)
		fn(x1, y1, x2, y2)
	}
	return countBrights()
}

func getBrightAction(text string) func(int, int, int, int) {
	switch {
	case strings.HasPrefix(text, "turn on"):
		return brightUpLight
	case strings.HasPrefix(text, "turn off"):
		return brightDownLight
	case strings.HasPrefix(text, "toggle"):
		return toggleBright
	}
	log.Fatalf("cannot parse light action: %s", text)
	return nil
}

func brightUpLight(x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			brightGrid[i][j]++
		}
	}
}

func brightDownLight(x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if brightGrid[i][j] == 0 {
				continue
			}
			brightGrid[i][j]--
		}
	}
}

func toggleBright(x1, y1, x2, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			brightGrid[i][j] = brightGrid[i][j] + 2
		}
	}
}

func countBrights() int {
	var count int
	for _, rows := range brightGrid {
		for _, grid := range rows {
			count += grid
		}
	}
	return count
}
