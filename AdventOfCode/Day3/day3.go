package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	r, err := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)")
	if err != nil {
		log.Fatal(err)
	}

	matches := preProcess(r)
	accumulator := 0
	enabled := true
	for _, m := range matches {
		switch m {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				accumulator += doMul(m)
			}

		}
	}

	fmt.Println(accumulator)

}

func doMul(cmd string) int {
	m1, m2 := findArguments(cmd)
	return m1 * m2
}

func findArguments(b string) (int, int) {
	commaIndex := strings.Index(b, ",")
	if commaIndex == -1 {
		return 0, 0
	}
	sFirstNum := b[4:commaIndex]
	sSecondNum := b[commaIndex+1 : len(b)-1]

	log.Printf("First %s, second %s", sFirstNum, sSecondNum)

	firstNum, _ := strconv.Atoi(sFirstNum)
	secondNum, _ := strconv.Atoi(sSecondNum)
	return firstNum, secondNum
}

func preProcess(r *regexp.Regexp) []string {
	data, err := os.ReadFile("AdventOfCode/Day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strData := string(data)

	return r.FindAllString(strData, -1)
}
