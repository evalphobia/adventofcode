package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/evalphobia/adventofcode/lib"
)

var bitMap = make(map[string]uint16)

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
	q := &Queue{}
	for _, line := range lines {
		q.Enqueue(line)
	}

	for q.HasNext() {
		line := q.Dequeue().(string)
		if line == "" {
			continue
		}

		b := NewBitwise(line)
		b.parseText()
		has := b.parseOperand()
		if !has {
			q.Enqueue(line)
			continue
		}
		b.emulateOperator()
	}
	if _, ok := bitMap["a"]; !ok {
		fmt.Printf("%v\n", bitMap)
		log.Fatal("cannot get `a`")
	}

	return int(bitMap["a"])
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	a := solvePart1(val)
	fmt.Printf("old a = %d\n", a)
	bitMap = make(map[string]uint16)
	bitMap["b"] = uint16(a)
	return solvePart1(val)
}

type Bitwise struct {
	text       string
	operatorFn func()
	opText1    string
	opText2    string

	operand1 uint16
	operand2 uint16
	result   string
}

func NewBitwise(text string) *Bitwise {
	return &Bitwise{
		text: text,
	}
}

func (b *Bitwise) parseText() {
	lines := strings.Split(b.text, " -> ")
	if len(lines) != 2 {
		log.Fatalf("cannot parse connector: %s", b.text)
	}

	b.result = lines[1]
	words := strings.Split(lines[0], " ")
	switch {
	case len(words) == 1:
		b.setAsAssign(words)
	case words[0] == "NOT":
		b.setAsNOT(words)
	case len(words) != 3:
		log.Fatalf("cannot parse words: %s", b.text)
	case words[1] == "AND":
		b.setAsAND(words)
	case words[1] == "OR":
		b.setAsOR(words)
	case words[1] == "LSHIFT":
		b.setAsLSHIFT(words)
	case words[1] == "RSHIFT":
		b.setAsRSHIFT(words)
	default:
		log.Fatalf("cannot parse words: %s", b.text)
	}
}

func (b *Bitwise) setAsAssign(words []string) {
	b.operatorFn = func() {
		if _, ok := bitMap[b.result]; !ok {
			bitMap[b.result] = b.operand1
		}
	}
	b.opText1 = words[0]
}

func (b *Bitwise) setAsNOT(words []string) {
	b.operatorFn = func() {
		bitMap[b.result] = ^b.operand1
	}
	b.opText1 = words[1]
}

func (b *Bitwise) setAsAND(words []string) {
	b.operatorFn = func() {
		bitMap[b.result] = b.operand1 & b.operand2
	}
	b.opText1 = words[0]
	b.opText2 = words[2]
}

func (b *Bitwise) setAsOR(words []string) {
	b.operatorFn = func() {
		bitMap[b.result] = b.operand1 | b.operand2
	}
	b.opText1 = words[0]
	b.opText2 = words[2]
}

func (b *Bitwise) setAsLSHIFT(words []string) {
	b.operatorFn = func() {
		bitMap[b.result] = b.operand1 << b.operand2
	}
	b.opText1 = words[0]
	b.opText2 = words[2]
}

func (b *Bitwise) setAsRSHIFT(words []string) {
	b.operatorFn = func() {
		bitMap[b.result] = b.operand1 >> b.operand2
	}
	b.opText1 = words[0]
	b.opText2 = words[2]
}

func (b *Bitwise) parseOperand() bool {
	var has bool
	b.operand1, has = getOrParseOperand(b.opText1)
	if !has {
		return false
	}
	b.operand2, has = getOrParseOperand(b.opText2)
	if !has {
		return false
	}
	return true
}

func getOrParseOperand(v string) (uint16, bool) {
	if v == "" {
		return 0, true
	}

	// int value
	i, err := strconv.Atoi(v)
	if err == nil {
		return uint16(i), true
	}

	// string variable
	if i, ok := bitMap[v]; ok {
		return i, true
	}

	// missing variable
	return 0, false
}

func (b *Bitwise) emulateOperator() {
	if b.operatorFn == nil {
		log.Fatalf("operator is missing: %s", b.text)
	}
	b.operatorFn()
}

type Queue struct {
	queue []interface{}
}

func (q *Queue) HasNext() bool {
	return len(q.queue) != 0
}

func (q *Queue) Enqueue(item interface{}) {
	q.queue = append(q.queue, item)
}

func (q *Queue) Dequeue() interface{} {
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}
