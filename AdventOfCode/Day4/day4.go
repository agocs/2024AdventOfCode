package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines := preProcess("AdventOfCode/Day4/input.txt")
	count := 0
	for i, _ := range *lines {
		for j, _ := range (*lines)[i] {
			if (*lines)[i][j] == 'X' {
				count += xmasMatch(lines, i, j)
			}
		}
	}
	fmt.Println(count)

	fmt.Println(countMatchingSubMatricies(lines))

}

func xmasMatch(lines *[]string, i, j int) int {
	matches := 0
	possibleDirections := scanForM(lines, i, j)
	for _, dir := range possibleDirections {
		if xmasMatchInDirection(lines, i, j, dir) {
			matches++
		}
	}
	return matches
}

func xmasMatchInDirection(lines *[]string, i, j int, dir direction) bool {
	xmas := "XMAS"
	for n := range xmas {
		iOffset := n * dir.iOffset
		jOffset := n * dir.jOffset
		if !isValid(lines, i, j, i+iOffset, j+jOffset) {
			return false
		}
		charAtCoords := (*lines)[i+iOffset][j+jOffset]
		if charAtCoords != xmas[n] {
			return false
		}
	}
	return true
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

func preProcess(path string) *[]string {
	file, err := os.Open(path)
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
