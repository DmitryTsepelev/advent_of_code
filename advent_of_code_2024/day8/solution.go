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

type AntennaMap = map[rune](*PointMap)
type PointMap = map[Point]bool

func getInputData() (AntennaMap, int, int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	antennas := AntennaMap{}

	row := 0
	width := 0
	for scanner.Scan() {
		width = len(scanner.Text())

		for col, v := range scanner.Text() {
			if v == '.' {
				continue
			}

			if _, ok := antennas[v]; !ok {
				pointMap := make(PointMap)
				antennas[v] = &pointMap
			}

			(*antennas[v])[Point{row, col}] = true
		}
		row++
	}

	return antennas, row, width
}

func antinodesFor(antenna1, antenna2 Point, height, width int) []Point {
	dRow := antenna2.row - antenna1.row
	dCol := antenna2.col - antenna1.col

	antinodes := []Point{}

	candidate1 := Point{antenna1.row - dRow, antenna1.col - dCol}
	if outOfBounds(candidate1, height, width) == false {
		antinodes = append(antinodes, candidate1)
	}

	candidate2 := Point{antenna1.row + 2*dRow, antenna1.col + 2*dCol}
	if outOfBounds(candidate2, height, width) == false {
		antinodes = append(antinodes, candidate2)
	}

	return antinodes
}

func resonatedAntinodesFor(antenna1, antenna2 Point, height, width int) []Point {
	dRow := antenna2.row - antenna1.row
	dCol := antenna2.col - antenna1.col

	antinodes := []Point{}

	prev := antenna1
	for {
		candidate := Point{prev.row - dRow, prev.col - dCol}

		if outOfBounds(candidate, height, width) {
			break
		}

		antinodes = append(antinodes, candidate)

		prev = candidate
	}

	prev = antenna2
	for {
		candidate := Point{prev.row + dRow, prev.col + dCol}

		if outOfBounds(candidate, height, width) {
			break
		}

		antinodes = append(antinodes, candidate)

		prev = candidate
	}

	return antinodes
}

func outOfBounds(p Point, height, width int) bool {
	return p.row < 0 || p.row >= height ||
		p.col < 0 || p.col >= width
}

func countAntinodes(antennas AntennaMap, height, width int, withResonance bool) int {
	antinodeMap := PointMap{}

	extraAntinodes := 0
	for _, points := range antennas {
		for p1 := range *points {
			for p2 := range *points {
				if p1 == p2 {
					continue
				}

				fn := antinodesFor

				if withResonance {
					fn = resonatedAntinodesFor
					// antennas are in line so they are antinodes
					antinodeMap[p1] = true
					antinodeMap[p2] = true
				}

				for _, candidate := range fn(p1, p2, height, width) {
					antinodeMap[candidate] = true
				}
			}
		}
	}

	return len(antinodeMap) + extraAntinodes
}

func main() {
	antennas, height, width := getInputData()

	fmt.Println("Part 1 solution is", countAntinodes(antennas, height, width, false))
	fmt.Println("Part 2 solution is", countAntinodes(antennas, height, width, true))
}
