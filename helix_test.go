package main

import (
	"fmt"
	"os"
	"testing"
)

func TestHelixFunc(t *testing.T) {
	tests := []struct {
		in  [][]int
		out []int
	}{
		{in: [][]int{}, out: []int{}},
		{in: [][]int{{1}}, out: []int{1}},
		{in: [][]int{{1, 2}}, out: []int{}},
		{in: [][]int{{1}, {2}}, out: []int{}},
		{in: [][]int{{1, 2}, {3, 4}}, out: []int{1, 2, 4, 3}},
		{in: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, out: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{in: [][]int{{1, 2, 3, 1}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}}, out: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8}},
		{in: [][]int{{1, 2, 3, 1, 4}, {4, 5, 6, 4}, {7, 8, 9, 7}, {7, 8, 9, 7}, {4, 5, 6, 4}}, out: []int{}},
		{in: [][]int{{1, 2, 3, 1, 4}, {4, 5, 6, 4, 7}, {7, 8, 9, 7, 9}, {7, 8, 9, 7, 4}, {4, 5, 6, 4, 8}}, out: []int{1, 2, 3, 1, 4, 7, 9, 4, 8, 4, 6, 5, 4, 7, 7, 4, 5, 6, 4, 7, 7, 9, 8, 8, 9}},
		{in: [][]int{{1, 2, 3, 1, 4, 2}, {4, 5, 6, 1, 6, 4}, {7, 2, 8, 1, 9, 7}, {7, 4, 9, 8, 9, 7}, {1, 2, 5, 1, 3, 7}, {7, 8, 1, 1, 2, 2}}, out: []int{1, 2, 3, 1, 4, 2, 4, 7, 7, 7, 2, 2, 1, 1, 8, 7, 1, 7, 7, 4, 5, 6, 1, 6, 9, 9, 3, 1, 5, 2, 4, 2, 8, 1, 8, 9}},
	}

	for caseNum, test := range tests {
		assert(RunHelix(test.in), test.out, caseNum)
	}
}

func assert(actual []int, expected []int, caseNum int) {
	if fmt.Sprint(actual) != fmt.Sprint(expected) {
		fmt.Printf("\n%c[35m%s%d\n%v != %v\n%c[0m\n", 27, "Case number: ", caseNum, actual, expected, 27)
		os.Exit(1)
	}
}
