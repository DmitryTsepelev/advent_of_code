package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func getInputData() []Point {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := []Point{}

	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(cmp[0])
		y, _ := strconv.Atoi(cmp[1])

		points = append(points, Point{x, y})
	}

	return points
}

var directions = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var NO_PATH = -1

func bfs(width, height int, fallingBytes []Point) int {
	currentPositions := []Point{{0, 0}}

	obstacles := map[Point]bool{}
	for _, b := range fallingBytes {
		obstacles[b] = true
	}

	visited := make(map[Point]bool)
	steps := 0

	for len(currentPositions) > 0 {
		nextPositions := []Point{}

		for _, currentPosition := range currentPositions {
			if currentPosition.x == width-1 && currentPosition.y == height-1 {
				return steps
			}

			for _, dir := range directions {
				nextPoint := Point{x: currentPosition.x + dir.x, y: currentPosition.y + dir.y}

				if nextPoint.x < 0 || nextPoint.x >= width || nextPoint.y < 0 || nextPoint.y >= height {
					continue
				}

				if _, ok := obstacles[nextPoint]; ok {
					continue
				}

				if _, ok := visited[nextPoint]; ok {
					continue
				}

				visited[nextPoint] = true

				nextPositions = append(nextPositions, nextPoint)
			}
		}

		currentPositions = nextPositions
		steps++
	}

	return NO_PATH
}

func findBlockingBit(width, height int, fallingBytes []Point) string {
	l, r := 0, len(fallingBytes)-1

	for l < r {
		mid := l + (r-l)/2

		if bfs(width, height, fallingBytes[:mid]) == NO_PATH {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	resByte := fallingBytes[l]
	return strconv.Itoa(resByte.x) + "," + strconv.Itoa(resByte.y)
}

func main() {
	points := getInputData()

	fmt.Println("Part 1 solution is", bfs(71, 71, points[:1024]))
	fmt.Println("Part 2 solution is", findBlockingBit(71, 71, points))
}
