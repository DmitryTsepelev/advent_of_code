package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func solve1(schema []string) int {
	currentNumber := ""
	validNumber := false
	sum := 0

	for rowIdx, row := range schema {
		for colIdx, cell := range row {
			if cell >= '0' && cell <= '9' {
				currentNumber += string(cell)

				deltas := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

				for _, d := range deltas {
					checkRow := rowIdx + d[0]
					checkCol := colIdx + d[1]

					if checkRow < 0 || checkRow >= len(schema) || checkCol < 0 || checkCol >= len(schema[0]) {
						continue
					}

					candidate := schema[checkRow][checkCol]
					if candidate != '.' && (candidate < '0' || candidate > '9') {
						validNumber = true
						break
					}
				}

			}

			lastCell := rowIdx == len(schema)-1 && colIdx == len(schema[0])-1
			if (cell < '0' || cell > '9' || lastCell) && len(currentNumber) > 0 {
				if validNumber {
					intNumber, _ := strconv.Atoi(currentNumber)
					sum += intNumber
				}

				currentNumber = ""
				validNumber = false
			}
		}
	}

	return sum
}

type Gear struct {
	row int
	col int
}

func solve2(schema []string) int {
	currentNumber := ""

	gears := make(map[Gear](*[]int), 0)
	var gearNums *[]int

	for rowIdx, row := range schema {
		for colIdx, cell := range row {
			if cell >= '0' && cell <= '9' {
				currentNumber += string(cell)

				deltas := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

				for _, d := range deltas {
					checkRow := rowIdx + d[0]
					checkCol := colIdx + d[1]

					if checkRow < 0 || checkRow >= len(schema) || checkCol < 0 || checkCol >= len(schema[0]) {
						continue
					}

					candidate := schema[checkRow][checkCol]
					if candidate == '*' {
						gear := Gear{row: checkRow, col: checkCol}
						if _, ok := gears[gear]; !ok {
							nums := make([]int, 0)
							gears[gear] = &nums
						}

						gearNums = gears[gear]
						break
					}
				}

			}

			lastCell := rowIdx == len(schema)-1 && colIdx == len(schema[0])-1
			if (cell < '0' || cell > '9' || lastCell) && len(currentNumber) > 0 {
				intNumber, _ := strconv.Atoi(currentNumber)
				if gearNums != nil {
					*gearNums = append(*gearNums, intNumber)
				}

				currentNumber = ""
				gearNums = nil
			}
		}
	}

	sum := 0
	for _, nums := range gears {
		if len(*nums) == 2 {
			sum += (*nums)[0] * (*nums)[1]
		}
	}

	return sum
}

func main() {
	schema := getInputData()

	solution1 := solve1(schema)
	fmt.Println("Part 1 solution is", solution1)

	solution2 := solve2(schema)
	fmt.Println("Part 2 solution is", solution2)
}
