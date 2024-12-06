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

type ObstacleMap = map[Point]bool

func getInputData() (int, int, Point, ObstacleMap) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	width, height := 0, 0
	start := Point{}
	obstacles := ObstacleMap{}

	for scanner.Scan() {
		width = len(scanner.Text())

		for col, val := range scanner.Text() {
			switch val {
			case '#':
				obstacles[Point{height, col}] = true
			case '^':
				start = Point{height, col}
			}
		}

		height++
	}

	return height, width, start, obstacles
}

var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func simulate(height, width int, currentPosition Point, obstacles ObstacleMap) *map[Point]([]int) {
	dirIdx := 0
	visited := map[Point]([]int){}

	for {
		if previousDirs, ok := visited[currentPosition]; ok {
			for _, prevDirIdx := range previousDirs {
				if prevDirIdx == dirIdx {
					return nil
				}
			}
		} else {
			visited[currentPosition] = append(visited[currentPosition], dirIdx)
		}

		direction := directions[dirIdx]

		nextPosition := Point{
			currentPosition.row + direction[0],
			currentPosition.col + direction[1],
		}

		if nextPosition.row < 0 || nextPosition.row == height ||
			nextPosition.col < 0 || nextPosition.col == width {
			return &visited
		}

		if _, ok := obstacles[nextPosition]; ok {
			// rotate right
			dirIdx = (dirIdx + 1) % len(directions)
		} else {
			currentPosition = nextPosition
		}
	}
}

func countObstaclesMakingLoop(height, width int, startPosition Point, obstacles ObstacleMap) (count int) {
	possibleObstacles := simulate(height, width, startPosition, obstacles)

	for obstacle := range *possibleObstacles {
		if startPosition == obstacle {
			continue
		}

		obstacles[obstacle] = true
		if simulate(height, width, startPosition, obstacles) == nil {
			count++
		}
		delete(obstacles, obstacle)
	}

	return count
}

func main() {
	height, width, start, obstacles := getInputData()

	fmt.Println("Part 1 solution is", len(*simulate(height, width, start, obstacles)))
	fmt.Println("Part 2 solution is", countObstaclesMakingLoop(height, width, start, obstacles))
}
