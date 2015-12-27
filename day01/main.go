package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	url   string
	file  string
	value string
)

func init() {
	flag.StringVar(&url, "url", "", "set url")
	flag.StringVar(&file, "file", "", "set file")
	flag.StringVar(&value, "value", "", "set value")
}

func main() {
	flag.Parse()

	value := getValue()
	if value == "" {
		log.Fatal("cannot get value")
	}

	result := solve(value)
	fmt.Printf("Answer: %d\n", result)
}

func getValue() string {
	switch {
	case url != "":
		return getValueFromURL(url)
	case file != "":
		return getValueFromFile(file)
	case value != "":
		return value
	}
	log.Fatalf("set flag parameter \n")
	return ""
}

func getValueFromURL(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer resp.Body.Close()
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(byt)
}

func getValueFromFile(file string) string {
	fp, err := os.Open(file)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer fp.Close()

	var result []string
	s := bufio.NewScanner(fp)
	for s.Scan() {
		result = append(result, s.Text())
	}
	return strings.Join(result, "")
}

func solve(val string) int {
	up := strings.Count(val, "(")
	down := strings.Count(val, ")")
	return up - down
}
