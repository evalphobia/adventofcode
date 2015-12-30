package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/evalphobia/adventofcode/lib"
)

var distances map[string]map[string]int
var cities map[string]bool

func init() {
	distances = make(map[string]map[string]int)
	cities = make(map[string]bool)
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
		parseRoute(line)
	}
	return calculateShortestRoutes()
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	lines := strings.Split(val, "\n")
	for _, line := range lines {
		parseRoute(line)
	}
	return calculateLongestRoutes()
}

func parseRoute(text string) {
	l := strings.Split(text, " = ")
	if len(l) != 2 {
		log.Fatalf("cannot parse route: %s", text)
	}
	distance, err := strconv.Atoi(l[1])
	if err != nil {
		log.Fatalf("cannot parse distance: %s", text)
	}

	c := strings.Split(l[0], " to ")
	if len(c) != 2 {
		log.Fatalf("cannot parse city: %s, %v", text)
	}

	if _, ok := distances[c[0]]; !ok {
		distances[c[0]] = make(map[string]int)
	}
	if _, ok := distances[c[1]]; !ok {
		distances[c[1]] = make(map[string]int)
	}
	distances[c[0]][c[1]] = distance
	distances[c[1]][c[0]] = distance
	cities[c[0]] = true
	cities[c[1]] = true
}

func calculateShortestRoutes() int {
	var c []string
	for city := range cities {
		c = append(c, city)
	}
	routes := createPermurations(c)
	minCost := math.MaxInt32
	for _, route := range routes {
		cost := getRouteCost(route)
		if cost < minCost {
			fmt.Println("cost=%d, route=%v", cost, route)
			minCost = cost
		}
	}
	return minCost
}

func calculateLongestRoutes() int {
	var c []string
	for city := range cities {
		c = append(c, city)
	}
	routes := createPermurations(c)
	maxCost := 0
	for _, route := range routes {
		cost := getRouteCost(route)
		if cost > maxCost {
			fmt.Println("cost=%d, route=%v", cost, route)
			maxCost = cost
		}
	}
	return maxCost
}

func getRouteCost(route []string) int {
	dist := 0
	for i, max := 1, len(route); i < max; i++ {
		c1, c2 := route[i-1], route[i]
		d, ok := distances[c1][c2]
		if !ok {
			log.Fatalf("cannot find route distance: %s, %s", c1, c2)
		}
		dist += d
	}
	return dist
}

func createPermurations(cities []string) [][]string {
	var result [][]string
	var list []string
	return permutations(result, cities, list)
}

func permutations(result [][]string, cities, list []string) [][]string {
	if len(list) == len(cities) {
		cp := make([]string, len(list))
		copy(cp, list)
		return append(result, cp)
	}
	for _, city := range cities {
		if isMember(list, city) {
			continue
		}
		result = permutations(result, cities, append(list, city))
	}
	return result
}

func isMember(list []string, city string) bool {
	for _, v := range list {
		if v == city {
			return true
		}
	}
	return false
}
