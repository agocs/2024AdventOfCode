package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	columnA, columnB := preprocess()

	diffScore := calcDifferenceScore(columnA, columnB)
	fmt.Println(diffScore)
	simScore := calcSimScore(columnA, columnB)
	fmt.Println(simScore)
}

func calcDifferenceScore(columnA []int, columnB []int) int {
	accumulator := 0
	for i := range len(columnA) {
		a := columnA[i]
		b := columnB[i]
		diff := a - b
		if diff < 0 {
			diff = diff * -1
		}
		accumulator += diff
	}
	return accumulator
}
func calcSimScore(colA []int, colB []int) int {
	accumulator := 0
	for _, v := range colA {
		count := countTarget(colB, v)
		accumulator += v * count
	}
	return accumulator
}

func countTarget(haystack []int, needle int) int {
	accumulator := 0
	for _, v := range haystack {
		if v == needle {
			accumulator++
		}
	}
	return accumulator
}

func preprocess() ([]int, []int) {
	columnA := make([]int, 0)
	columnB := make([]int, 0)

	file, err := os.Open("AdventOfCode/Day1/input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return columnA, columnB
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		thisLine := scanner.Text()
		columns := strings.Fields(thisLine)
		if len(columns) != 2 {
			break
		}
		colA, _ := strconv.Atoi(columns[0])
		colB, _ := strconv.Atoi(columns[1])
		columnA = append(columnA, colA)
		columnB = append(columnB, colB)
	}

	sort.Ints(columnA)
	sort.Ints(columnB)

	return columnA, columnB
}
