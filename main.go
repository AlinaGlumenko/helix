package main

import "fmt"

func main() {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	output := RunHelix(input)

	fmt.Println("Input: ", input)
	fmt.Println("Output: ", output)
	fmt.Println("Expected: ", expected)
}
