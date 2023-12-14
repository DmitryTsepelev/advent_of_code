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

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	point := Point{}
	field := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		for col, c := range line {
			if c == 'S' {
				point.row = len(field)
				point.col = col
			}
		}

		field = append(field, line)
	}

	return field
}

const (
	up    = 0
	down  = 1
	left  = 2
	right = 3
)

var currentToDirections = map[rune]([]int){
	'S': {up, down, left, right},
	'|': {up, down},
	'-': {left, right},
	'L': {up, right},
	'J': {left, up},
	'7': {left, down},
	'F': {right, down},
}

var deltas = map[int]([]int){
	up:    []int{-1, 0},
	down:  []int{1, 0},
	left:  []int{0, -1},
	right: []int{0, 1},
}

var validNextPipes = map[rune]string{
	up:    "|7FS",
	down:  "|LJS",
	left:  "-LFS",
	right: "-J7S",
}

func adjacentFrom(field []string, current Point) []Point {
	result := make([]Point, 0)

	currentPipe := rune(field[current.row][current.col])
	for _, direction := range currentToDirections[currentPipe] {
		delta := deltas[direction]
		nextPoint := Point{row: current.row + delta[0], col: current.col + delta[1]}
		if nextPoint.row < 0 || nextPoint.row == len(field) || nextPoint.col < 0 || nextPoint.col == len(field[0]) {
			// out of borders
			continue
		}

		nextPipe := rune(field[nextPoint.row][nextPoint.col])
		if validPipes, ok := validNextPipes[rune(direction)]; ok {
			for _, valid := range validPipes {
				if valid == nextPipe {
					result = append(result, nextPoint)
					break
				}
			}
		}
	}

	return result
}

func dfs(field []string, current Point, parent Point, start Point, visited [][]bool, currentPath []Point) []Point {
	visited[current.row][current.col] = true

	for _, nextPoint := range adjacentFrom(field, current) {
		if nextPoint.row == parent.row && nextPoint.col == parent.col {
			// do not go back to parent
			continue
		}

		if visited[nextPoint.row][nextPoint.col] {
			if nextPoint.row == start.row && nextPoint.col == start.col {
				// loop found
				return append(currentPath, current)
			}
		} else {
			path := dfs(field, nextPoint, current, start, visited, append(currentPath, current))
			if len(path) > 0 {
				return path
			}
		}
	}

	return []Point{}
}

func findStart(field []string) Point {
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			if field[i][j] == 'S' {
				return Point{row: i, col: j}
			}
		}
	}

	return Point{}
}

func findLoop(field []string) []Point {
	start := findStart(field)

	visited := make([][]bool, len(field))
	for i := 0; i < len(field); i++ {
		visited[i] = make([]bool, len(field[0]))
	}

	impossibleParent := Point{row: -1, col: -1}
	path := dfs(field, start, impossibleParent, start, visited, []Point{})

	return path
}

// ---

func countInside(field []string) int {
	clearField := make([][]rune, 0)
	for i := 0; i < len(field); i++ {
		line := ""
		for j := 0; j < len(field[0]); j++ {
			line += "."
		}
		clearField = append(clearField, []rune(line))
	}

	path := findLoop(field)

	for _, point := range path {
		clearField[point.row][point.col] = rune(field[point.row][point.col])
	}

	sum := 0
	for i := 0; i < len(clearField); i++ {
		count := 0
		fBefore := false
		lBefore := false
		for j := 0; j < len(clearField); j++ {
			cell := rune(clearField[i][j])

			if cell == '|' || cell == 'S' {
				count++
				lBefore = false
				fBefore = false
			} else if cell == 'F' {
				fBefore = true
				lBefore = false
			} else if cell == 'L' {
				lBefore = true
				fBefore = false
			} else if cell == 'J' {
				if fBefore {
					count++
					fBefore = false
					lBefore = false
				}
			} else if cell == '7' {
				if lBefore {
					count++
					fBefore = false
					lBefore = false
				}
			} else if cell == '-' {
				// nothing
			} else if cell == '.' {
				fBefore = false
				lBefore = false

				if count%2 == 1 {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	field := getInputData()
	fmt.Println("Part 1 solution is", len(findLoop(field))/2)
	fmt.Println("Part 2 solution is", countInside(field))
}
