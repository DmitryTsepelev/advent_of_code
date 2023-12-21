package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	row int
	col int
}

type Field = []string

func getInputData() (Field, Point) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := Field{}
	var point Point

	for scanner.Scan() {
		line := scanner.Text()

		field = append(field, line)

		for col, cell := range line {
			if cell == 'S' {
				point = Point{row: len(field) - 1, col: col}
			}
		}
	}

	return field, point
}

func isRock(field Field, row, col int) bool {
	row %= len(field)
	if row < 0 {
		row = len(field) + row
	}

	col %= len(field[0])
	if col < 0 {
		col = len(field[0]) + col
	}
	return field[row][col] == '#'
}

var deltas = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func bfs(field Field, start Point, maxSteps int) int {
	queue := []Point{start}
	steps := 0

	for {
		nextQueueSet := make(map[Point]bool)
		if steps == maxSteps {
			return len(queue)
		}

		for _, point := range queue {
			for _, delta := range deltas {
				nextRow := (point.row + delta[0])
				nextCol := (point.col + delta[1])
				nextPoint := Point{row: nextRow, col: nextCol}

				if isRock(field, nextRow, nextCol) {
					// rock
					continue
				}

				nextQueueSet[nextPoint] = true
			}
		}

		steps++

		queue = make([]Point, 0)
		for point := range nextQueueSet {
			queue = append(queue, point)
		}
	}
}

func interpolate(field Field, start Point, n int) int {
	if len(field) != len(field[0]) {
		panic("area is not a square")
	}

	for i := 0; i < len(field); i++ {
		if field[i][start.col] == '#' {
			panic("obstacle at " + string(i) + " " + string(start.col))
		}

		if field[start.row][i] == '#' {
			panic("obstacle at " + string(start.row) + " " + string(i))
		}
	}

	a0 := bfs(field, start, 65)
	a1 := bfs(field, start, 65+131)
	a2 := bfs(field, start, 65+2*131)

	b0 := a0
	b1 := a1 - a0
	b2 := a2 - a1

	return b0 + b1*n + (n*(n-1)/2)*(b2-b1)
}

func main() {
	field, start := getInputData()
	fmt.Println("Solution 1 is", bfs(field, start, 64))
	fmt.Println("Solution 2 is", interpolate(field, start, 26501365/len(field)))
}
