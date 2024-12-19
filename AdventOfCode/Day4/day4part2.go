package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func countMatchingSubMatricies(lines *[]string) int {
	accumulator := 0
	for i := 0; i < len(*lines)-2; i++ {
		for j := 0; j < len(*lines)-2; j++ {
			sm, err := getSubMatrix(lines, i, j)
			if err != nil {
				log.Fatalf("Error getting submatrix at %d,%d,%v", i, j, err)
			}

			matrixMultIdentityX(sm)
			if testAgainstPossibleMatches(sm) {
				log.Printf("Got match at %d, %d", i, j)
				accumulator++
			}
		}
	}
	return accumulator
}

// getSubMatrix returns a 3x3 submatrix with i,j being in the top left
func getSubMatrix(lines *[]string, i, j int) (*[][]uint8, error) {
	if len(*lines) < i+3 {
		return nil, errors.New(fmt.Sprintf("i (%d) is too big. i + 2 would be out of range of len (lines) (%d)", i, len(*lines)))
	}
	if len((*lines)[i]) < j+3 {
		return nil, errors.New(fmt.Sprintf("j (%d) is too big. j + 2 would be out of range of len (lines[i]) (%d)", j, len((*lines)[i])))
	}

	subMatrix := make([][]uint8, 3)

	for k := 0; k < 3; k++ {
		subMatrix[k] = make([]uint8, 3)
		for l := 0; l < 3; l++ {
			subMatrix[k][l] = (*lines)[i+k][j+l]
		}
	}
	return &subMatrix, nil
}

var identityX [][]uint8 = [][]uint8{
	[]uint8{1, 0, 1},
	[]uint8{0, 1, 0},
	[]uint8{1, 0, 1},
}

// matrixMultIdentityX mutates the submatrix, multiplying each element by
// 1 0 1
// 0 1 0
// 1 0 1
func matrixMultIdentityX(subMatrix *[][]uint8) {
	for i := 0; i < len(*subMatrix); i++ {
		for j := 0; j < len((*subMatrix)[i]); j++ {
			(*subMatrix)[i][j] = (*subMatrix)[i][j] * identityX[i][j]
		}
	}
}

var possibleMatches [][][]uint8 = [][][]uint8{
	[][]uint8{
		[]uint8{'M', 0, 'M'},
		[]uint8{0, 'A', 0},
		[]uint8{'S', 0, 'S'},
	},
	[][]uint8{
		[]uint8{'S', 0, 'S'},
		[]uint8{0, 'A', 0},
		[]uint8{'M', 0, 'M'},
	},
	[][]uint8{
		[]uint8{'M', 0, 'S'},
		[]uint8{0, 'A', 0},
		[]uint8{'M', 0, 'S'},
	},
	[][]uint8{
		[]uint8{'S', 0, 'M'},
		[]uint8{0, 'A', 0},
		[]uint8{'S', 0, 'M'},
	},
}

func testAgainstPossibleMatches(subMatrix *[][]uint8) bool {
	for _, m := range possibleMatches {
		if reflect.DeepEqual(*subMatrix, m) {
			return true
		}
	}
	return false
}
