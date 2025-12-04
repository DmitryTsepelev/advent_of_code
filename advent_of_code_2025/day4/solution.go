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

	field := make([][]byte, 0)
	for scanner.Scan() {
		text := scanner.Text()
		field = append(field, []byte(text))
	}

	return field
}

var deltas = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func findAccessible(field [][]byte) [][]int {
	rows, columns := len(field), len(field[0])

	accessible := make([][]int, 0)
	for rowIdx, row := range field {
		for colIdx, val := range row {
			if val == '.' {
				continue
			}

			neighbourRolls := 0
			for _, delta := range deltas {
				nRowIdx, nColIdx := rowIdx+delta[0], colIdx+delta[1]

				if nRowIdx < 0 || nRowIdx >= rows || nColIdx < 0 || nColIdx >= columns {
					continue
				}

				if field[nRowIdx][nColIdx] == '@' {
					neighbourRolls++
				}
			}

			if neighbourRolls < 4 {
				accessible = append(accessible, []int{rowIdx, colIdx})
			}
		}
	}

	return accessible
}

func removeAll(field [][]byte) int {
	removed := 0

	for {
		accessible := findAccessible(field)
		if len(accessible) == 0 {
			return removed
		}

		removed += len(accessible)
		for _, point := range accessible {
			field[point[0]][point[1]] = '.'
		}
	}
}

func main() {
	field := getInputData()

	fmt.Println("Part 1 solution is", len(findAccessible(field)))
	fmt.Println("Part 2 solution is", removeAll(field))
}
