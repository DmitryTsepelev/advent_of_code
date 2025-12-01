package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	distances := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		distance, _ := strconv.Atoi(text[1:])
		if text[0] == 'L' {
			distance = -distance
		}
		distances = append(distances, distance)
	}

	return distances
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(distances []int) (int, int) {
	zeroStops, zeroPasses := 0, 0
	currentPosition := 50

	for _, distance := range distances {
		zeroPasses += abs(distance) / 100
		distance %= 100

		currentPositionWas := currentPosition
		currentPosition += distance

		turningRight := distance > 0

		if currentPosition > 0 {
			currentPosition %= 100
			if currentPositionWas != 0 && currentPositionWas > currentPosition && turningRight {
				zeroPasses++
			}
		} else if currentPosition < 0 {
			currentPosition = 100 + currentPosition
			if currentPositionWas != 0 && currentPositionWas < currentPosition && turningRight == false {
				zeroPasses++
			}
		} else {
			zeroPasses++
		}

		if currentPosition == 0 {
			zeroStops++
		}
	}

	return zeroStops, zeroPasses
}

func main() {
	distances := getInputData()

	zeroStops, zeroPasses := solve(distances)

	fmt.Println("Part 1 solution is", zeroStops)
	fmt.Println("Part 1 solution is", zeroPasses)
}
