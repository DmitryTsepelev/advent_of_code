package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() [][]string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fields := make([][]string, 0)
	field := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			fields = append(fields, field)
			field = make([]string, 0)
		} else {
			field = append(field, line)
		}
	}

	fields = append(fields, field)

	return fields
}

func transpose(matrix []string) []string {
	result := make([]string, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		result[i] = ""
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			result[j] += string(matrix[i][j])
		}
	}

	return result
}

func findMirrors(field []string) []int {
	mirrors := []int{}

	for initRow := 1; initRow < len(field); initRow++ {
		l := initRow - 1
		r := initRow
		for {
			if l == -1 || r == len(field) {
				// reached the beginning or end
				mirrors = append(mirrors, initRow)
				break
			}
			if field[l] != field[r] {
				// not matching lines
				break
			}

			l--
			r++
		}
	}

	return mirrors
}

func findAnyMirror(field []string) (int, int) {
	mirrors := findMirrors(field)

	if len(mirrors) > 0 {
		return mirrors[0], HORIZONTAL
	}

	mirrors = findMirrors(transpose(field))

	if len(mirrors) > 0 {
		return mirrors[0], VERTICAL
	}

	return -1, NONE
}

func solve1(fields [][]string) int {
	sum := 0
	for _, field := range fields {
		mirror, kind := findAnyMirror(field)

		if kind == HORIZONTAL {
			mirror *= 100
		}
		sum += mirror
	}
	return sum
}

const NONE = -1
const HORIZONTAL = 0
const VERTICAL = 1

func replaceSmuggle(field []string, row int, col int) []string {
	newField := make([]string, len(field))

	for rowIdx, line := range field {
		if rowIdx != row {
			newField[rowIdx] = line
			continue
		}

		newField[rowIdx] = ""
		for idx, c := range line {
			if idx == col {
				if c == '#' {
					newField[rowIdx] += string('.')
				} else {
					newField[rowIdx] += string('#')
				}
			} else {
				newField[rowIdx] += string(c)
			}
		}
	}

	return newField
}

func findSmuggleMirror(field []string) (int, int) {
	origMirror, origKind := findAnyMirror(field)

	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[0]); col++ {
			newField := replaceSmuggle(field, row, col)

			horizontalMirrors := findMirrors(newField)
			if origKind == HORIZONTAL {
				for _, candidate := range horizontalMirrors {
					if candidate != origMirror {
						return candidate, HORIZONTAL
					}
				}
			} else if len(horizontalMirrors) > 0 {
				return horizontalMirrors[0], HORIZONTAL
			}

			verticalMirrors := findMirrors(transpose(newField))
			if origKind == VERTICAL {
				for _, candidate := range verticalMirrors {
					if candidate != origMirror {
						return candidate, VERTICAL
					}
				}
			} else if len(verticalMirrors) > 0 {
				return verticalMirrors[0], VERTICAL
			}
		}
	}

	// this can never happen
	return -1, -1
}

func solve2(fields [][]string) int {
	sum := 0
	for _, field := range fields {
		mirror, kind := findSmuggleMirror(field)
		if kind == HORIZONTAL {
			mirror *= 100
		}
		sum += mirror
	}
	return sum
}

func main() {
	fields := getInputData()
	fmt.Println("Part 1 solution is", solve1(fields))
	fmt.Println("Part 2 solution is", solve2(fields))
}
