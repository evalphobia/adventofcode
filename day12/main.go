package main

import (
	"encoding/json"
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
	obj := newJSONObject(val)
	return obj.calculateSum()
}

func runPart2(value string) {
	result := solvePart2(value)
	fmt.Printf("Answer: %d\n", result)
}

func solvePart2(val string) int {
	obj := newJSONObject(val)
	obj.ignoreRed = true
	return obj.calculateSum()
}

type jsonObject struct {
	data      interface{}
	ignoreRed bool
}

func newJSONObject(body string) *jsonObject {
	var data interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err.Error())
	}
	return &jsonObject{
		data: data,
	}
}

func (j *jsonObject) calculateSum() int {
	sum := j.calc(j.data)
	return sum
}

func (j *jsonObject) calc(value interface{}) int {
	switch typ := value.(type) {
	case float64:
		return int(typ)
	case []interface{}:
		sum := 0
		for _, v := range typ {
			sum += j.calc(v)
		}
		return sum
	case map[string]interface{}:
		sum := 0
		for _, v := range typ {
			vv, ok := v.(string)
			if j.ignoreRed && ok && vv == "red" {
				return 0
			}
			sum += j.calc(v)
		}
		return sum
	}
	return 0
}
