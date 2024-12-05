package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() [][]byte {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := [][]byte{}
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}

	return input
}

var deltas1 = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func validPoint(input [][]byte, nextPoint []int) bool {
	row, col := nextPoint[0], nextPoint[1]

	return row >= 0 && row < len(input) &&
		col >= 0 && col < len(input[0])
}

func part1(input [][]byte) int {
	var count int
	word := "XMAS"

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			for _, delta := range deltas1 {
				row, col := i, j
				validWord := true

				for step := 0; step <= 3; step++ {
					if !validPoint(input, []int{row, col}) || word[step] != input[row][col] {
						validWord = false
						break
					}

					row += delta[0]
					col += delta[1]
				}

				if validWord {
					count++
				}
			}
		}
	}

	return count
}

var deltas2 = [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func part2(input [][]byte) int {
	var count int

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[0])-1; j++ {
			if input[i][j] == 'A' {
				xRunes := []byte{}

				for _, delta := range deltas2 {
					xRunes = append(xRunes, input[i+delta[0]][j+delta[1]])
				}

				xCorners := string(xRunes)

				if xCorners == "MSMS" ||
					xCorners == "MMSS" ||
					xCorners == "SMSM" ||
					xCorners == "SSMM" {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	input := getInputData()

	fmt.Println("Part 1 solution is", part1(input))
	fmt.Println("Part 2 solution is", part2(input))
}
