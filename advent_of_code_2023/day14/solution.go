package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() [][]rune {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, []rune(line))
	}

	return field
}

func moveRoundedRock(row int, col int, horizontalBorder int, verticalBorder int, dRow int, dCol int, field *[][]rune) {
	if (*field)[row][col] != 'O' {
		return
	}

	currentRow, currentCol := row, col

	for {
		if currentCol == verticalBorder ||
			currentRow == horizontalBorder ||
			(*field)[currentRow+dRow][currentCol+dCol] != '.' {
			(*field)[currentRow][currentCol] = 'O'
			break
		}
		(*field)[currentRow][currentCol] = '.'

		currentRow += dRow
		currentCol += dCol
	}
}

const NO_BORDER = -1

func tiltUp(field *[][]rune) {
	leftHorizontalBorder := 0

	for row := 0; row < len(*field); row++ {
		for col := 0; col < len((*field)[0]); col++ {
			moveRoundedRock(row, col, leftHorizontalBorder, NO_BORDER, -1, 0, field)
		}
	}
}

func tiltDown(field *[][]rune) {
	rightHorizontalBorder := len(*field) - 1
	dRow := 1

	for row := len(*field) - 1; row >= 0; row-- {
		for col := len((*field)[0]) - 1; col >= 0; col-- {
			moveRoundedRock(row, col, rightHorizontalBorder, NO_BORDER, dRow, 0, field)
		}
	}
}

func tiltRight(field *[][]rune) {
	rightVerticalBorder := len((*field)[0]) - 1

	for col := len((*field)[0]) - 1; col >= 0; col-- {
		for row := len(*field) - 1; row >= 0; row-- {
			moveRoundedRock(row, col, NO_BORDER, rightVerticalBorder, 0, 1, field)
		}
	}
}

func tiltLeft(field *[][]rune) {
	leftVerticalBorder := 0

	for col := 0; col < len((*field)[0]); col++ {
		for row := 0; row < len(*field); row++ {
			moveRoundedRock(row, col, NO_BORDER, leftVerticalBorder, 0, -1, field)
		}
	}
}

func sumWeight(field [][]rune) int {
	sum := 0
	for row := 0; row < len(field); row++ {
		for col := 0; col < len((field)[0]); col++ {
			if field[row][col] == 'O' {
				sum += len(field) - row
			}
		}
	}
	return sum
}

func solve1() int {
	field := getInputData()
	tiltUp(&field)
	return sumWeight(field)
}

func findLoop(arr []int) (int, int) {
	for offset := 0; offset < len(arr)/2; offset++ {
		offseted := arr[offset:]
		for l := 3; l < len(offseted)/2+1; l++ {
			isLoop := true
			for idx := 0; idx <= l; idx++ {
				if offseted[idx] != offseted[l+idx] {
					isLoop = false
					break
				}
			}

			if isLoop {
				return offset, l
			}
		}
	}

	return 0, 0
}

func solve2() int {
	field := getInputData()

	weights := make([]int, 1000) // assuming we have that loop in first 1000 elements
	for i := 0; i < len(weights); i++ {
		tiltUp(&field)
		tiltLeft(&field)
		tiltDown(&field)
		tiltRight(&field)
		weights[i] = sumWeight(field)
	}

	offset, loopLen := findLoop(weights)
	weightIdx := offset + (1000000000-offset)%loopLen - 1

	return weights[weightIdx]
}

func main() {
	fmt.Println("Part 1 solution is", solve1())
	fmt.Println("Part 2 solution is", solve2())
}
