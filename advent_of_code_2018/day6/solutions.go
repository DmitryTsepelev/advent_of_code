package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	ID       string
	X        int
	Y        int
	Infinite bool
	Count    int
}

func loadPoints() *[]*Point {
	points := []*Point{}

	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	id := 'A'
	for scanner.Scan() {
		coordinates := strings.Split(scanner.Text(), ", ")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		points = append(points, &Point{ID: string(id), X: x, Y: y, Infinite: false, Count: 0})
		id++
	}

	return &points
}

func distanceBetween(point *Point, x int, y int) int {
	return int(math.Abs(float64(point.X-x)) + math.Abs(float64(point.Y-y)))
}

func solveTask1() int {
	points := *loadPoints()

	const fieldSize = 360

	for x := 0; x < fieldSize; x++ {
		for y := 0; y < fieldSize; y++ {
			minDistance := fieldSize + fieldSize

			distances := make(map[*Point]int)

			for _, point := range points {
				distance := distanceBetween(point, x, y)

				distances[point] = distance
				if distance < minDistance {
					minDistance = distance
				}
			}

			var closestPoint *Point

			for point, distance := range distances {
				if distance == minDistance {
					if closestPoint == nil {
						closestPoint = point
					} else {
						closestPoint = nil
					}
				}
			}

			if closestPoint != nil {
				closestPoint.Count++

				if x == 0 || x == fieldSize-1 || y == 0 || y == fieldSize-1 {
					closestPoint.Infinite = true
				}
			}
		}
	}

	maxCount := 0
	for _, point := range points {
		if !point.Infinite && maxCount < point.Count {
			maxCount = point.Count
		}
	}

	return maxCount
}

func solveTask2() int {
	points := *loadPoints()

	const fieldSize = 360
	const maxDistance = 10000

	count := 0

	for x := 0; x < fieldSize; x++ {
		for y := 0; y < fieldSize; y++ {
			distance := 0
			for _, point := range points {
				distance += distanceBetween(point, x, y)
			}

			if distance < maxDistance {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println("Task 1 solution is", solveTask1())
	fmt.Println("Task 2 solution is", solveTask2())
}
