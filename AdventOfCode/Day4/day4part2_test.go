package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getSubMatrix(t *testing.T) {
	testLines := preProcess("testInput.txt")

	testSub, err := getSubMatrix(testLines, 0, 0)
	if err != nil {
		t.Errorf("Error in getSubMatrix %v", err)
	}

	expected := [][]uint8{
		[]uint8{'M', 'M', 'M'},
		[]uint8{'M', 'S', 'A'},
		[]uint8{'A', 'M', 'X'},
	}

	if !reflect.DeepEqual(*testSub, expected) {
		t.Errorf("Test failed, expected %v, got %v", expected, *testSub)
	}

	_, err2 := getSubMatrix(testLines, 0, 8)
	if err2 == nil {
		t.Errorf("Expected error when trying to GetSubMatrix(0, 8), got nil instead")
	}

	_, err3 := getSubMatrix(testLines, 8, 0)
	if err3 == nil {
		t.Errorf("Expected error when trying to GetSubMatrix(8, 0), got nil instead")
	}

}

func Test_matrixMultIdentityX(t *testing.T) {
	testLines := preProcess("testInput.txt")
	testSub, _ := getSubMatrix(testLines, 0, 0)

	matrixMultIdentityX(testSub)

	expected := [][]uint8{
		[]uint8{'M', 0, 'M'},
		[]uint8{0, 'S', 0},
		[]uint8{'A', 0, 'X'},
	}

	if !reflect.DeepEqual(*testSub, expected) {
		t.Errorf("Test failed, expected %v, got %v", expected, *testSub)
	}
}

func Test_testAgainstPossilbeMatches(t *testing.T) {
	type args struct {
		subMatrix *[][]uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"basic match",
			args{&[][]uint8{
				[]uint8{'M', 0, 'M'},
				[]uint8{0, 'A', 0},
				[]uint8{'S', 0, 'S'},
			}},
			true,
		},
		{
			"basic fail",
			args{&[][]uint8{
				[]uint8{'A', 0, 'M'},
				[]uint8{0, 'B', 0},
				[]uint8{'S', 0, 'S'},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testAgainstPossibleMatches(tt.args.subMatrix); got != tt.want {
				t.Errorf("testAgainstPossibleMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMatchingSubMatricies(t *testing.T) {
	lines := preProcess("testInput.txt")
	matches := countMatchingSubMatricies(lines)
	if matches != 9 {
		t.Errorf("Test failed, expected 9, got %d", matches)
	}
}

func Test_specificSubmatrix(t *testing.T) {
	lines := preProcess("testInput.txt")
	testSub, _ := getSubMatrix(lines, 2, 1)
	matrixMultIdentityX(testSub)
	fmt.Println(*testSub)

	isMatch := testAgainstPossibleMatches(testSub)
	if !isMatch {
		t.Errorf("Test failed, expected %v, got %v", true, isMatch)
	}

}
