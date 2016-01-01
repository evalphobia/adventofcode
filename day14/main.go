package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/evalphobia/adventofcode/lib"
)

var reindeers []*reindeer

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
		parseLine(line)
	}

	const sec = 2503
	maxDistance := 0
	for _, reindeer := range reindeers {
		distance := reindeer.run(sec)
		if distance > maxDistance {
			maxDistance = distance
			fmt.Printf("distance:%d, reindeer:%v\n", distance, reindeer)
		}
	}
	return maxDistance
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	lines := strings.Split(val, "\n")
	for _, line := range lines {
		parseLine(line)
	}

	const sec = 2503
	for i := 1; i <= sec; i++ {
		maxDistance := 0
		for _, reindeer := range reindeers {
			distance := reindeer.run(i)
			if distance > maxDistance {
				maxDistance = distance
			}
		}

		for _, reindeer := range reindeers {
			if reindeer.latestDistance == maxDistance {
				reindeer.WinPoint()
			}
		}
	}

	maxPoint := 0
	for _, reindeer := range reindeers {
		if reindeer.point > maxPoint {
			maxPoint = reindeer.point
		}
	}
	return maxPoint
}

func parseLine(text string) {
	words := strings.Split(text, " ")
	if len(words) != 15 {
		log.Fatalf("cannot parse line: %s", text)
	}

	speed, err := strconv.Atoi(words[3])
	if err != nil {
		log.Fatalf("cannot parse speed: %s", text)
	}

	stamina, err := strconv.Atoi(words[6])
	if err != nil {
		log.Fatalf("cannot parse stamina: %s", text)
	}

	resting, err := strconv.Atoi(words[13])
	if err != nil {
		log.Fatalf("cannot parse resting: %s", text)
	}

	r := &reindeer{
		name:    words[0],
		speed:   speed,
		stamina: stamina,
		resting: resting,
	}

	reindeers = append(reindeers, r)
}

type reindeer struct {
	name           string
	speed          int
	stamina        int
	resting        int
	point          int
	latestDistance int
}

func (r *reindeer) run(sec int) int {
	distance := 0
	distance += sec / (r.stamina + r.resting) * r.stamina * r.speed
	extra := sec % (r.stamina + r.resting)
	if extra < r.stamina {
		distance += extra * r.speed
	} else {
		distance += r.stamina * r.speed
	}
	r.latestDistance = distance
	return distance
}

func (r *reindeer) WinPoint() {
	r.point++
}
