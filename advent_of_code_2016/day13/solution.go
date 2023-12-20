package main

import (
	"fmt"
	"strconv"
)

type Point struct {
	row int
	col int
}

type Field = map[Point]bool

func isOpen(field Field, point Point, num int) bool {
	if _, ok := field[point]; !ok {
		value := point.col*point.col + 3*point.col + 2*point.col*point.row + point.row + point.row*point.row + num
		binary := strconv.FormatInt(int64(value), 2)

		ones := 0
		for _, r := range binary {
			if r == '1' {
				ones++
			}
		}

		field[point] = ones%2 == 0
	}

	return field[point]
}

var deltas = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findShortestPath(num int, destination Point, maxSteps int) (int, int) {
	field := make(Field)
	start := Point{row: 1, col: 1}
	visited := map[Point]bool{start: true}

	queue := []Point{start}

	steps := 0
	reachedAt50 := 0

	for {
		nextQueue := make([]Point, 0)

		if steps == 50 {
			reachedAt50 = len(visited)
		}

		for _, point := range queue {
			if point == destination {
				return steps, reachedAt50
			}

			for _, delta := range deltas {
				nextPoint := Point{row: point.row + delta[0], col: point.col + delta[1]}

				if nextPoint.row < 0 || nextPoint.col < 0 || visited[nextPoint] || !isOpen(field, nextPoint, num) {
					continue
				}

				visited[nextPoint] = true

				nextQueue = append(nextQueue, nextPoint)
			}
		}

		steps++
		queue = nextQueue
	}
}

func main() {
	fmt.Println(findShortestPath(1358, Point{row: 39, col: 31}, 50))
}
