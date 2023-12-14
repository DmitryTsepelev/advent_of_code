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

type Starmap = [][]rune

func getInputData() Starmap {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	point := Point{}
	field := Starmap{}

	for scanner.Scan() {
		line := scanner.Text()

		for col, c := range line {
			if c == 'S' {
				point.row = len(field)
				point.col = col
			}
		}

		field = append(field, []rune(line))
	}

	return field
}

const GALAXY = '#'

type ExpansionMap = map[int]bool

func findExpansionMap(field Starmap) (ExpansionMap, ExpansionMap) {
	emptyRows := make(ExpansionMap, len(field))
	for i := 0; i < len(field); i++ {
		emptyRows[i] = true
	}

	emptyCols := make(ExpansionMap, len(field[0]))
	for i := 0; i < len(field[0]); i++ {
		emptyCols[i] = true
	}

	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[0]); col++ {
			if field[row][col] == GALAXY {
				emptyRows[row] = false
				emptyCols[col] = false
			}
		}
	}

	return emptyRows, emptyCols
}

func findGalaxies(field Starmap) []Point {
	galaxies := make([]Point, 0)
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[0]); col++ {
			if field[row][col] == GALAXY {
				galaxies = append(galaxies, Point{row: row, col: col})
			}
		}
	}
	return galaxies
}

func minmax(x, y int) (int, int) {
	if x < y {
		return x, y
	}
	return y, x
}

func calcDistance(p1 Point, p2 Point, expansionDistance int, emptyRows, emptyCols ExpansionMap) int {
	minRow, maxRow := minmax(p1.row, p2.row)
	minCol, maxCol := minmax(p1.col, p2.col)

	distance := 0

	for row := minRow + 1; row <= maxRow; row++ {
		if emptyRows[row] {
			distance += expansionDistance
		} else {
			distance += 1
		}
	}

	for col := minCol + 1; col <= maxCol; col++ {
		if emptyCols[col] {
			distance += expansionDistance
		} else {
			distance += 1
		}
	}

	return distance
}

func calcDistances(galaxies []Point, expansionDistance int, emptyRows, emptyCols ExpansionMap) []int {
	distances := make([]int, 0)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := calcDistance(galaxies[i], galaxies[j], expansionDistance, emptyRows, emptyCols)
			distances = append(distances, distance)
		}
	}

	return distances
}

func sumDistances(field Starmap, expansionDistance int) int {
	emptyRows, emptyCols := findExpansionMap(field)
	galaxies := findGalaxies(field)
	distances := calcDistances(galaxies, expansionDistance, emptyRows, emptyCols)

	sum := 0
	for _, distance := range distances {
		sum += distance
	}

	return sum
}

func main() {
	starfield := getInputData()
	fmt.Println("Part 1 solution is", sumDistances(starfield, 2))
	fmt.Println("Part 2 solution is", sumDistances(starfield, 1000000))
}
