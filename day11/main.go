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
	fmt.Printf("Answer: %s\n", result)
}

func solvePart1(val string) string {
	pass := newPassword(val)
	return pass.generate()
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %s\n", result)
}

func solvePart2(val string) string {
	p1 := newPassword(val)
	gen := p1.generate()
	p2 := newPassword(gen)
	return p2.generate()
}

type password struct {
	phrase string
	code   []int
}

func newPassword(text string) *password {
	p := &password{
		phrase: text,
		code:   make([]int, len(text)),
	}
	p.encode()
	return p
}

func (p *password) encode() {
	for i, s := range p.phrase {
		p.code[i] = convertToInt(string(s))
	}
}

func (p *password) decode() string {
	phrases := make([]string, len(p.code))
	for i, num := range p.code {
		phrases[i] = convertToString(num)
	}
	return strings.Join(phrases, "")
}

func (p *password) generate() string {
	lastIdx := len(p.code) - 1
	for {
		p.increment(lastIdx)
		if p.isValid() {
			return p.decode()
		}
	}
}

func (p *password) increment(idx int) {
	if idx >= len(p.code) || idx < 0 {
		log.Fatalf("invalid index: %d, code=%v", idx, p.code)
	}

	p.code[idx]++
	if p.code[idx] > 26 {
		p.code[idx] = 1
		p.increment(idx - 1)
	}
}

func (p *password) isValid() bool {
	var prev1, prev2 int
	var pairCount int
	var maxStreak int
	currentStreak := 1
	for _, code := range p.code {
		switch {
		case p.isBadLetter(code):
			return false
		case prev1 == 0:
			prev1 = code
			continue
		case prev1 == code && prev2 != prev1:
			pairCount++
		case prev1+1 == code:
			currentStreak++
			if maxStreak < currentStreak {
				maxStreak = currentStreak
			}
			prev2 = prev1
			prev1 = code
			continue
		}
		prev2 = prev1
		prev1 = code
		currentStreak = 1
	}

	switch {
	case pairCount < 2:
		return false
	case maxStreak < 3:
		return false
	}

	return true
}

func (p *password) isBadLetter(i int) bool {
	switch i {
	case 9, 12, 15:
		return true
	}
	return false
}

func convertToInt(s string) int {
	switch s {
	case "a":
		return 1
	case "b":
		return 2
	case "c":
		return 3
	case "d":
		return 4
	case "e":
		return 5
	case "f":
		return 6
	case "g":
		return 7
	case "h":
		return 8
	case "i":
		return 9
	case "j":
		return 10
	case "k":
		return 11
	case "l":
		return 12
	case "m":
		return 13
	case "n":
		return 14
	case "o":
		return 15
	case "p":
		return 16
	case "q":
		return 17
	case "r":
		return 18
	case "s":
		return 19
	case "t":
		return 20
	case "u":
		return 21
	case "v":
		return 22
	case "w":
		return 23
	case "x":
		return 24
	case "y":
		return 25
	case "z":
		return 26
	}
	log.Fatalf("cannot convert to int: %s", s)
	return 0
}

func convertToString(i int) string {
	switch i {
	case 1:
		return "a"
	case 2:
		return "b"
	case 3:
		return "c"
	case 4:
		return "d"
	case 5:
		return "e"
	case 6:
		return "f"
	case 7:
		return "g"
	case 8:
		return "h"
	case 9:
		return "i"
	case 10:
		return "j"
	case 11:
		return "k"
	case 12:
		return "l"
	case 13:
		return "m"
	case 14:
		return "n"
	case 15:
		return "o"
	case 16:
		return "p"
	case 17:
		return "q"
	case 18:
		return "r"
	case 19:
		return "s"
	case 20:
		return "t"
	case 21:
		return "u"
	case 22:
		return "v"
	case 23:
		return "w"
	case 24:
		return "x"
	case 25:
		return "y"
	case 26:
		return "z"
	}
	log.Fatalf("cannot convert to string: %d", i)
	return ""
}
