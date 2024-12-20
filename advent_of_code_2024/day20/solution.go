package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	rowIdx int
	colIdx int
}

func (this *Point) outOfBounds(field Field) bool {
	return this.rowIdx < 0 || this.rowIdx >= field.height ||
		this.colIdx < 0 || this.colIdx >= field.width
}

type Field struct {
	startPoint Point
	endPoint   Point
	width      int
	height     int
	wallMap    map[Point]bool
}

func getInputData() Field {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := Field{wallMap: make(map[Point]bool)}

	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		line := scanner.Text()

		field.width = len(line)
		for colIdx, val := range line {
			point := Point{rowIdx, colIdx}

			switch val {
			case '#':
				field.wallMap[point] = true
			case 'S':
				field.startPoint = point
			case 'E':
				field.endPoint = point
			}
		}

		field.height = rowIdx
	}

	return field
}

var UP = Point{1, 0}
var DOWN = Point{-1, 0}
var RIGHT = Point{0, 1}
var LEFT = Point{0, -1}

var directions = []Point{DOWN, RIGHT, UP, LEFT}

type Path = map[Point]int

func findPath(field Field) Path {
	path := make(Path)

	current := field.startPoint

	step := 0
	for current != field.endPoint {
		path[current] = step

		for _, dir := range directions {
			nextPoint := Point{current.rowIdx + dir.rowIdx, current.colIdx + dir.colIdx}

			if _, ok := field.wallMap[nextPoint]; ok {
				continue
			}

			if _, ok := path[nextPoint]; ok {
				continue
			}

			current = nextPoint
			break
		}

		step++
	}
	path[current] = step

	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(startPoint, finalPoint Point) int {
	return abs(startPoint.rowIdx-finalPoint.rowIdx) + abs(startPoint.colIdx-finalPoint.colIdx)
}

func countCheats(field Field, path Path, cheatPicoseconds int) int {
	cheatCount := 0

	for startPoint := range path {
		queue := []Point{startPoint}
		candidates := map[Point]bool{startPoint: true}

		for i := 0; i < cheatPicoseconds; i++ {
			nextQueue := []Point{}

			for _, currentPoint := range queue {
				for _, dir := range directions {
					nextPoint := Point{currentPoint.rowIdx + dir.rowIdx, currentPoint.colIdx + dir.colIdx}

					if nextPoint.outOfBounds(field) {
						continue
					}

					if _, ok := candidates[nextPoint]; !ok {
						nextQueue = append(nextQueue, nextPoint)
						candidates[nextPoint] = true
					}
				}
			}

			queue = nextQueue
		}

		for finalPoint := range candidates {
			if endPathLen, ok := path[finalPoint]; ok {
				savedPicoseconds := endPathLen - path[startPoint] - manhattanDistance(startPoint, finalPoint)

				if savedPicoseconds >= 100 {
					cheatCount++
				}
			}
		}
	}

	return cheatCount
}

func main() {
	field := getInputData()
	path := findPath(field)

	fmt.Println("Part 1 solution is", countCheats(field, path, 2))
	fmt.Println("Part 2 solution is", countCheats(field, path, 20))
}
