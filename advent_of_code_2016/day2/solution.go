package main

import (
	"bufio"
	"fmt"
	"os"
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

type Point struct {
	row int
	col int
}

func move(dir rune, point Point) Point {
	newPoint := Point{row: point.row, col: point.col}

	switch dir {
	case 'L':
		newPoint.col--
	case 'R':
		newPoint.col++
	case 'U':
		newPoint.row--
	case 'D':
		newPoint.row++
	}

	if newPoint.col < 0 || newPoint.col > 2 || newPoint.row < 0 || newPoint.row > 2 {
		return point
	}

	return newPoint
}

func move2(dir rune, point Point) Point {
	newPoint := Point{row: point.row, col: point.col}

	switch dir {
	case 'L':
		newPoint.col--
	case 'R':
		newPoint.col++
	case 'U':
		newPoint.row--
	case 'D':
		newPoint.row++
	}

	if ((newPoint.row == 0 || newPoint.row == 4) && newPoint.col != 2) ||
		((newPoint.row == 1 || newPoint.row == 3) && (newPoint.col <= 0 || newPoint.col >= 4)) ||
		(newPoint.row == 2 && (newPoint.col < 0 || newPoint.col > 4)) ||
		(newPoint.row < 0 || newPoint.row > 4) {
		return point
	}

	return newPoint
}

func solve(lines []string, pad [][]rune, start Point, moveF func(rune, Point) Point) string {
	result := ""
	point := start

	for _, line := range lines {
		for _, dir := range line {
			point = moveF(dir, point)
		}

		result += string(pad[point.row][point.col])
	}

	return result
}

func solve1(lines []string) string {
	pad := [][]rune{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}

	point := Point{row: 1, col: 1}

	return solve(lines, pad, point, move)
}

func solve2(lines []string) string {
	point := Point{row: 2, col: 0}
	pad := [][]rune{
		{'0', '0', '1', '0', '0'},
		{'0', '2', '3', '4', '0'},
		{'5', '6', '7', '8', '9'},
		{'0', 'A', 'B', 'C', '0'},
		{'0', '0', 'D', '0', '0'},
	}

	return solve(lines, pad, point, move2)
}

func main() {
	lines := getInputData()

	fmt.Println("Solution 1 is", solve1(lines))
	fmt.Println("Solution 2 is", solve2(lines))
}
