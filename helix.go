package main

const (
	RIGHT int = iota
	DOWN
	LEFT
	UP
)

func RunHelix(input [][]int) (output []int) {
	// check if it is empty or not square
	if len(input) == 0 || len(input[0]) == 0 || !isSquareMatrix(input) {
		return output
	}

	n := len(input)
	expectedLen := n * n // amount of all elements
	output = make([]int, expectedLen)

	orient := []bool{true, false, false, false} // RIGHT, DOWN, LEFT, UP

	var row, col, round int

	for i := 0; i < expectedLen; i++ {
		if orient[LEFT] && col == 0+round { // LEFT -> UP
			orient[LEFT] = false
			orient[UP] = true
		}
		if orient[RIGHT] && col == n-1-round { // RIGHT -> DOWN
			orient[RIGHT] = false
			orient[DOWN] = true
		}
		if orient[DOWN] && row == n-1-round { // DOWN -> LEFT
			orient[DOWN] = false
			orient[LEFT] = true
		}
		if orient[UP] && row == 0+round { // UP -> RIGHT
			orient[UP] = false
			orient[RIGHT] = true
			round++
			row++
			col++
		}

		output[i] = input[row][col]

		if orient[LEFT] { // move LEFT
			col = col - 1
		} else if orient[DOWN] { // move DOWN
			row = row + 1
		} else if orient[RIGHT] { // move RIGHT
			col = col + 1
		} else { // move UP
			row = row - 1
		}
	}
	return output
}

func isSquareMatrix(input [][]int) bool {
	length := len(input)
	for _, el := range input {
		if len(el) != length {
			return false
		}
	}
	return true
}
