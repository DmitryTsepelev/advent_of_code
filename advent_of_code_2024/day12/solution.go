package main

import (
	"bufio"
	"fmt"
	"os"
)

type Field = [][]byte

func getInputData() Field {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := Field{}

	for scanner.Scan() {
		row := []byte{}
		for _, r := range scanner.Text() {
			row = append(row, byte(r))
		}
		field = append(field, row)
	}

	return field
}

func validPoint(rowIdx, colIdx int, field Field) bool {
	return rowIdx >= 0 && rowIdx < len(field) &&
		colIdx >= 0 && colIdx < len(field[0])
}

var deltas = []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

const NOT_VISITED = 0
const QUEUED = 1
const VISITED = 2

type Point struct {
	rowIdx int
	colIdx int
}

// based on https://github.com/shraddhaag/aoc/blob/main/2024/day12/main.go
func checkCorners(field [][]byte, current Point) int {
	count := 0
	gardenType := field[current.rowIdx][current.colIdx]
	x, y := current.colIdx, current.rowIdx

	if x == 0 && y == 0 {
		count += 1
	}

	if x == 0 && y == len(field)-1 {
		count += 1
	}

	if x == len(field[0])-1 && y == len(field)-1 {
		count += 1
	}

	if x == len(field[0])-1 && y == 0 {
		count += 1
	}

	// top left outside corner
	// ##   __   |#
	// #O   #O   |O
	if (x > 0 && y > 0 && field[y][x-1] != gardenType && field[y-1][x] != gardenType) ||
		(x > 0 && y == 0 && field[y][x-1] != gardenType) || (x == 0 && y > 0 && field[y-1][x] != gardenType) {
		count += 1
	}

	// top left inside corner
	// OO
	// O#
	if x < len(field[0])-1 && y < len(field)-1 && field[y][x+1] == gardenType && field[y+1][x] == gardenType && field[y+1][x+1] != gardenType {
		count += 1
	}

	// top right outside corner
	// ##   __    #|
	// O#   O#    O|
	if (x < len(field[0])-1 && y > 0 && field[y][x+1] != gardenType && field[y-1][x] != gardenType) ||
		(x < len(field[0])-1 && y == 0 && field[y][x+1] != gardenType) || (x == len(field[0])-1 && y > 0 && field[y-1][x] != gardenType) {
		count += 1
	}

	// top right inside corner
	// OO
	// #O
	if x > 0 && y < len(field)-1 && field[y][x-1] == gardenType && field[y+1][x] == gardenType && field[y+1][x-1] != gardenType {
		count += 1
	}

	// bottom left outside corner
	// #O   #O    |O
	// ##   --    |#
	if (x > 0 && y < len(field)-1 && field[y][x-1] != gardenType && field[y+1][x] != gardenType) ||
		(x > 0 && y == len(field)-1 && field[y][x-1] != gardenType) || (x == 0 && y < len(field)-1 && field[y+1][x] != gardenType) {
		count += 1
	}

	// bottom left inside corner
	// O#
	// OO
	if x < len(field[0])-1 && y > 0 && field[y][x+1] == gardenType && field[y-1][x] == gardenType && field[y-1][x+1] != gardenType {
		count += 1
	}

	// bottom right outside corner
	// O#   O#    O|
	// ##   --    #|
	if (x < len(field[0])-1 && y < len(field)-1 && field[y][x+1] != gardenType && field[y+1][x] != gardenType) ||
		(x < len(field[0])-1 && y == len(field)-1 && field[y][x+1] != gardenType) || (x == len(field[0])-1 && y < len(field)-1 && field[y+1][x] != gardenType) {
		count += 1
	}

	// bottom right inside corner
	// #O
	// OO
	if x > 0 && y > 0 && field[y][x-1] == gardenType && field[y-1][x] == gardenType && field[y-1][x-1] != gardenType {
		count += 1
	}
	return count
}

func fencePlant(startRowIdx, startColIdx int, field Field, visited [][]byte) (int, int) {
	plant := field[startRowIdx][startColIdx]

	queue := []Point{{startRowIdx, startColIdx}}
	plantPoints := make(map[Point]bool)
	visited[startRowIdx][startColIdx] = NOT_VISITED

	area, perimeter := 0, 0
	for len(queue) > 0 {
		currentPoint := queue[0]
		currentRowIdx, currentColIdx := currentPoint.rowIdx, currentPoint.colIdx
		queue = queue[1:]

		visited[currentRowIdx][currentColIdx] = VISITED
		plantPoints[currentPoint] = true
		area++
		perimeter += 4

		for _, delta := range deltas {
			nextRowIdx := currentRowIdx + delta.rowIdx
			nextColIdx := currentColIdx + delta.colIdx

			if validPoint(nextRowIdx, nextColIdx, field) && field[nextRowIdx][nextColIdx] == plant {
				if visited[nextRowIdx][nextColIdx] == VISITED {
					continue
				}

				if visited[nextRowIdx][nextColIdx] == NOT_VISITED {
					visited[nextRowIdx][nextColIdx] = QUEUED
					queue = append(queue, Point{nextRowIdx, nextColIdx})
				}

				perimeter -= 2
			}
		}
	}

	numSides := 0
	for point := range plantPoints {
		numSides += checkCorners(field, point)
	}

	regularPrice := area * perimeter
	discountedPrice := area * numSides

	return regularPrice, discountedPrice
}

func fence(field Field) (int, int) {
	fenceCost, discountedFenceCost := 0, 0

	visited := make([][]byte, len(field))
	for i := 0; i < len(field); i++ {
		visited[i] = make([]byte, len(field[0]))
	}

	for rowIdx, row := range field {
		for colIdx := range row {
			if visited[rowIdx][colIdx] == VISITED {
				continue
			}

			regularPrice, discountedPrice := fencePlant(rowIdx, colIdx, field, visited)

			fenceCost += regularPrice
			discountedFenceCost += discountedPrice
		}
	}

	return fenceCost, discountedFenceCost
}

func main() {
	field := getInputData()

	fenceCost, discountedFenceCost := fence(field)
	fmt.Println("Part 1 solution is", fenceCost)
	fmt.Println("Part 2 solution is", discountedFenceCost)
}
