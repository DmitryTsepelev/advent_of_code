package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() ([2]int, map[[2]int]bool, int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	start := [2]int{}
	splitters := map[[2]int]bool{}

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		for colIdx := 0; colIdx < len(line); colIdx++ {
			point := [2]int{rowIdx, colIdx}
			if line[colIdx] == 'S' {
				start = point
			} else if line[colIdx] == '^' {
				splitters[point] = true
			}
		}

		rowIdx++
	}

	return start, splitters, rowIdx
}

func main() {
	start, splitters, height := getInputData()

	beams := map[[2]int]int{start: 1}
	splits := 0

	for rowIdx := 1; rowIdx < height; rowIdx++ {
		newBeams := map[[2]int]int{}
		for beam, variants := range beams {
			beam[0]++

			if _, ok := splitters[beam]; ok {
				splits++
				newBeams[[2]int{beam[0], beam[1] - 1}] += variants
				newBeams[[2]int{beam[0], beam[1] + 1}] += variants
			} else {
				newBeams[beam] += variants
			}
		}

		beams = newBeams
	}

	timelines := 0
	for _, v := range beams {
		timelines += v
	}

	fmt.Println("Part 1 solution is", splits)
	fmt.Println("Part 2 solution is", timelines)
}
