package main

import (
	"crypto/md5"
	"encoding/hex"
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
	var i int
	for {
		if i%100000 == 0 {
			fmt.Printf("progress... %d\n", i)
		}

		key := fmt.Sprintf("%s%d", val, i)
		if strings.HasPrefix(getMD5Hash(key), "00000") {
			break
		}

		i++
	}
	return i
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	var i int
	for {
		if i%100000 == 0 {
			fmt.Printf("progress... %d\n", i)
		}

		key := fmt.Sprintf("%s%d", val, i)
		if strings.HasPrefix(getMD5Hash(key), "000000") {
			break
		}

		i++
	}
	return i
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
