package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	levels := preProcess()
	accumulator := 0
	accumulatorWithRetry := 0
	for _, level := range levels {
		problem := checkLevel(level)
		if problem == -1 {
			accumulator++
		}

		err := checkLevelsWithRetry(level)
		if err == nil {
			accumulatorWithRetry++
		}
	}
	fmt.Printf("Without buffer: %d\n", accumulator)
	fmt.Printf("With buffer: %d\n", accumulatorWithRetry)
}

func checkLevelsWithRetry(levels []int) error {
	problem := checkLevel(levels)

	if problem == -1 {
		return nil
	}

	dest := make([]int, len(levels)-1)

	for i, _ := range dest {
		rebuildLevel(&levels, &dest, i)
		problem = checkLevel(dest)
		if problem == -1 {
			return nil
		}
	}

	return errors.New("unable to solve this level with retry")
}

// rebuildLevels assumes destination is a []int with size one less than source
func rebuildLevel(source, dest *[]int, toSkip int) {
	i := 0
	for ; i < toSkip; i++ {
		(*dest)[i] = (*source)[i]
	}
	for ; i < len(*dest); i++ {
		(*dest)[i] = (*source)[i+1]
	}
}

// checkLevel returns the index of the problem
func checkLevel(levels []int) int {
	isIncreasing := true
	isInitialized := false
	minDiff := 1
	maxDiff := 3

	prevVal := levels[0]
	for i := 1; i < len(levels); i++ {
		thisVal := levels[i]
		if !isInitialized {
			isInitialized = true
			isIncreasing = thisVal > prevVal
		}
		//check for increasing / decreasing
		if (thisVal > prevVal) != isIncreasing {
			return i
		}
		//check for minDiff / maxDiff
		diff := thisVal - prevVal
		if diff < 0 {
			diff = -diff
		}
		if diff < minDiff {
			return i
		} else if diff > maxDiff {
			return i
		}

		prevVal = thisVal
	}
	return -1
}

func preProcess() [][]int {
	levelsInput := make([][]int, 0)

	file, err := os.Open("AdventOfCode/Day2/input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return levelsInput
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		thisLine := scanner.Text()
		levels := strings.Fields(thisLine)
		levelsInput = append(levelsInput, convertLevel(levels))
	}
	return levelsInput
}

func convertLevel(strLevel []string) []int {
	intLevel := make([]int, len(strLevel))
	fmt.Println(strLevel)
	fmt.Println(intLevel)
	for i, s := range strLevel {
		intVal, _ := strconv.Atoi(s)
		intLevel[i] = intVal
	}
	return intLevel
}
