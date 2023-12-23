package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

type Point struct {
	row int
	col int
}

func isDestination(field []string, point Point) bool {
	return point.row == len(field)-1 && point.col == len(field[0])-2
}

var deltas = map[byte]([]int){
	'>': {0, 1},
	'<': {0, -1},
	'^': {-1, 0},
	'v': {1, 0},
}

func adjacentPoints(currentPath []Point, field []string, visited map[Point]bool, slipSlopes bool) []Point {
	lastVisited := currentPath[len(currentPath)-1]

	validPoints := []Point{}

	for validSlope, delta := range deltas {
		dRow, dCol := delta[0], delta[1]

		row := lastVisited.row + dRow
		col := lastVisited.col + dCol

		if row < 0 || col < 0 || row >= len(field) || col >= len(field[0]) {
			// out of bounds
			continue
		}

		if slipSlopes {
			if !(field[row][col] == '.' || field[row][col] == validSlope) {
				// illegal move
				continue
			}
		} else {
			if field[row][col] == '#' {
				// illegal move
				continue
			}
		}

		candidate := Point{row: row, col: col}
		if v, found := visited[candidate]; !found || !v {
			validPoints = append(validPoints, candidate)
		}
	}

	return validPoints
}

func dfs(field []string, currentPath []Point, visited map[Point]bool, slipSlopes bool, results *[][]Point) {
	lastVisited := currentPath[len(currentPath)-1]

	if isDestination(field, lastVisited) {
		*results = append(*results, currentPath)
		return
	}

	for _, point := range adjacentPoints(currentPath, field, visited, slipSlopes) {
		visited[point] = true

		dfs(
			field,
			append(currentPath, point),
			visited,
			slipSlopes,
			results,
		)

		visited[point] = false
	}
}

func findLongestPath(field []string, slipSlopes bool) int {
	allPaths := make([][]Point, 0)
	visited := make(map[Point]bool)
	dfs(field, []Point{{row: 0, col: 1}}, visited, slipSlopes, &allPaths)

	longest := 0

	for _, path := range allPaths {
		if len(path)-1 > longest {
			longest = len(path) - 1
		}
	}

	return longest
}

func main() {
	field := getInputData()
	fmt.Println("Solution 1 is", findLongestPath(field, true))
	fmt.Println("Solution 2 is", findLongestPath(field, false))
}
