package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines := preProcess()

	for i, _ := range *lines {
		for j, _ := range (*lines)[i] {
			if (*lines)[i][j] == 'X' {
				xmasMatch(lines, i, j)
			}
		}
	}

}

func xmasMatch(lines *[]string, i, j int) bool {
	return false
}

type direction struct {
	iOffset int
	jOffset int
}

var directions = []direction{
	direction{-1, 0},  // up
	direction{1, 0},   // down
	direction{0, 1},   // right
	direction{0, -1},  // left
	direction{-1, -1}, // up left
	direction{-1, 1},  // up right
	direction{1, -1},  // down left
	direction{1, 1},   // down right
}

func scanForM(lines *[]string, i, j int) []direction {
	possibleDirections := []direction{}
	for _, d := range directions {
		if !isValid(lines, i, j, i+d.iOffset, j+d.jOffset) {
			continue
		}
		if (*lines)[i+d.iOffset][j+d.jOffset] == 'M' {
			possibleDirections = append(possibleDirections, d)
		}
	}
	return possibleDirections
}

func isValid(lines *[]string, i, j, k, l int) bool {
	if k < 0 {
		return false
	}
	if k >= len(*lines) {
		return false
	}
	if l < 0 {
		return false
	}
	if l >= len((*lines)[i]) {
		return false
	}
	return true
}

func preProcess() *[]string {
	file, err := os.Open("C:\\Users\\Chris\\AppData\\Local\\JetBrains\\GoLand2024.3\\demo\\LearnGoProject\\AdventOfCode\\Day4\\input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return &lines
}
