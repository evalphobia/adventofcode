package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/evalphobia/adventofcode/lib"
)

var happinessTable map[string]map[string]int
var attendees map[string]bool

func init() {
	happinessTable = make(map[string]map[string]int)
	attendees = make(map[string]bool)
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
		parseLine(line)
	}

	var list []string
	for key := range attendees {
		list = append(list, key)
	}
	sort.Strings(list)

	perms := createNecklacePermurations(list)
	maxHapiness := 0
	for _, perm := range perms {
		h := calculateHapiness(perm)
		if maxHapiness < h {
			maxHapiness = h
			fmt.Printf("%d : %#v\n", h, perm)
		}
	}
	return maxHapiness
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

	var list []string
	for key := range attendees {
		list = append(list, key)
	}
	sort.Strings(list)

	perms := createNecklacePermurations(list)
	maxHapiness := 0
	for _, perm := range perms {
		for i := range perm {
			p := append(perm[:i+1], perm[i:]...)
			p[i] = ""
			h := calculateHapiness(p)
			if maxHapiness < h {
				maxHapiness = h
				fmt.Printf("%d : %#v\n", h, p)
			}
		}
	}
	return maxHapiness
}

func parseLine(text string) {
	text = strings.Replace(text, "gain ", "", 1)
	text = strings.Replace(text, "lose ", "-", 1)

	s1 := strings.Split(text, " happiness units by sitting next to ")
	if len(s1) != 2 {
		log.Fatalf("cannot parse line: %s", text)
	}

	s2 := strings.Split(s1[0], " would ")
	if len(s2) != 2 {
		log.Fatalf("cannot parse line: %s", text)
	}

	happiness, err := strconv.Atoi(s2[1])
	if err != nil {
		log.Fatalf("cannot parse hapiness: %v", s2)
	}

	person1 := s2[0]
	person2 := strings.TrimRight(s1[1], ".")
	attendees[person1] = true
	attendees[person2] = true

	if _, ok := happinessTable[person1]; !ok {
		happinessTable[person1] = make(map[string]int)
	}
	happinessTable[person1][person2] = happiness
}

func createNecklacePermurations(list []string) [][]string {
	max := len(list)
	var result [][]int
	for i := 1; i < max; i++ {
		initialList := []int{i, 0}
		necklacePermutations(&result, initialList, 1, max)
	}

	perms := make([][]string, len(result))
	for i, idxList := range result {
		perm := make([]string, len(idxList))
		for j, idx := range idxList {
			perm[j] = list[idx]
		}
		perms[i] = perm
	}
	return perms
}

func necklacePermutations(result *[][]int, list []int, n, max int) []int {
	if len(list) == max {
		cp := make([]int, max)
		copy(cp, list)
		return cp
	}
	for i := 1; i < max; i++ {
		switch {
		case i == 1 && list[0] <= i:
			continue
		case isMember(list, i):
			continue
		}
		l := necklacePermutations(result, append(list, i), n+1, max)
		if l != nil {
			*result = append(*result, l)
		}
	}
	return nil
}

func isMember(list []int, city int) bool {
	for _, v := range list {
		if v == city {
			return true
		}
	}
	return false
}

func calculateHapiness(list []string) int {
	sum := 0
	max := len(list)
	for i, person := range list {
		if isMe(person) {
			continue
		}
		var prev, next string
		if i == 0 {
			prev = list[max-1]
		} else {
			prev = list[i-1]
		}
		if i == max-1 {
			next = list[0]
		} else {
			next = list[i+1]
		}

		if !isMe(prev) {
			sum += happinessTable[person][prev]
		}
		if !isMe(next) {
			sum += happinessTable[person][next]
		}
	}
	return sum
}

func isMe(person string) bool {
	return person == ""
}
