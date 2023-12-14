package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() (uint64, int, int) {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	depth := strings.Replace(scanner.Text(), "depth: ", "", 1)
	intDepth, _ := strconv.Atoi(depth)

	scanner.Scan()
	coordinates := strings.Split(strings.Replace(scanner.Text(), "target: ", "", 1), ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])

	return uint64(intDepth), x, y
}

func main() {
	depth, targetX, targetY := getInput()

	geologicIndex := [][]uint64{}
	erosionLevel := [][]uint64{}

	totalRisk := 0
	for y := 0; y <= targetY; y++ {
		geoRow := []uint64{}
		erosionRow := []uint64{}

		for x := 0; x <= targetX; x++ {
			var index uint64
			var level uint64

			if x == 0 && y == 0 || x == targetX && y == targetY {
				index = 0
			} else if y == 0 {
				index = uint64(16807 * x)
			} else if x == 0 {
				index = uint64(48271 * y)
			} else {
				index = erosionRow[x-1] * erosionLevel[y-1][x]
			}

			level = (index + depth) % 20183

			geoRow = append(geoRow, index)
			erosionRow = append(erosionRow, level)

			totalRisk += int(level % 3)
		}

		geologicIndex = append(geologicIndex, geoRow)
		erosionLevel = append(erosionLevel, erosionRow)
	}

	fmt.Println("Task 1 solution is", totalRisk)
}
