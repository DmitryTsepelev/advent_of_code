package main

import (
	"fmt"
	"strconv"
)

func powerLevel(x, y, serialNumber int) int {
	rackId := x + 10
	baseLevel := ((rackId * y) + serialNumber) * rackId

	strLevel := strconv.Itoa(baseLevel)
	level := 0
	if len(strLevel) >= 3 {
		level, _ = strconv.Atoi(strLevel[len(strLevel)-3 : len(strLevel)-2])
	}

	return level - 5
}

func solveTask(serialNumber int, minSize int, maxSize int) (int, int, int) {
	const gridSize = 300

	sumTable := make([][]int, gridSize+1)
	for y := 0; y <= gridSize; y++ {
		sumTable[y] = make([]int, gridSize+1)
	}
	for y := 1; y <= gridSize; y++ {
		for x := 1; x <= gridSize; x++ {
			sumTable[y][x] = powerLevel(x, y, serialNumber) + sumTable[y-1][x] + sumTable[y][x-1] - sumTable[y-1][x-1]
		}
	}

	resultX := 0
	resultY := 0
	maxPower := 0
	resultSize := 0

	for squareSize := minSize; squareSize <= maxSize; squareSize++ {
		for squareY := squareSize; squareY <= gridSize; squareY++ {
			for squareX := squareSize; squareX <= gridSize; squareX++ {
				totalPower := sumTable[squareY][squareX] - sumTable[squareY-squareSize][squareX] - sumTable[squareY][squareX-squareSize] + sumTable[squareY-squareSize][squareX-squareSize]

				if totalPower > maxPower {
					maxPower = totalPower
					resultX = squareX - squareSize + 1
					resultY = squareY - squareSize + 1
					resultSize = squareSize
				}
			}
		}
	}

	return resultX, resultY, resultSize
}

func main() {
	const serialNumber = 5093

	x, y, _ := solveTask(serialNumber, 3, 3)
	fmt.Println("Task 1 solution is", x, y)

	x, y, size := solveTask(serialNumber, 1, 300)
	fmt.Println("Task 2 solution is", x, y, size)
}
