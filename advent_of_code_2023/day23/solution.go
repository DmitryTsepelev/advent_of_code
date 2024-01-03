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

// -- part 2 --

type Graph = map[Point](map[Point]int)

func gridToGraph(input []string) Graph {
	graph := Graph{}

	for rowIdx, row := range input {
		for colIdx, col := range row {
			if col == '#' {
				continue
			}

			currentPoint := Point{rowIdx, colIdx}

			for _, delta := range deltas {
				dRow, dCol := delta[0], delta[1]

				nextPoint := Point{
					row: rowIdx + dRow,
					col: colIdx + dCol,
				}

				if nextPoint.row < 0 || nextPoint.col < 0 || nextPoint.row >= len(input) || nextPoint.col >= len(input[0]) || input[nextPoint.row][nextPoint.col] == '#' {
					// out of bounds
					continue
				}

				if _, found := graph[currentPoint]; !found {
					graph[currentPoint] = make(map[Point]int)
				}
				graph[currentPoint][nextPoint] = 1

				if _, found := graph[nextPoint]; !found {
					graph[nextPoint] = make(map[Point]int)
				}
				graph[nextPoint][currentPoint] = 1
			}
		}
	}

	return graph
}

func collapseGraph(graph Graph) {
	for {
		everythingCollapsed := true

		for currentPoint, neighbourMap := range graph {
			if len(neighbourMap) != 2 {
				continue
			}

			keys := []Point{}
			for key := range neighbourMap {
				keys = append(keys, key)
			}

			left, right := keys[0], keys[1]

			distance := graph[left][currentPoint] + graph[right][currentPoint] - 1

			delete(graph, currentPoint)
			delete(graph[left], currentPoint)
			delete(graph[right], currentPoint)

			graph[left][right] = distance + 1
			graph[right][left] = distance + 1

			everythingCollapsed = false
		}

		if everythingCollapsed {
			break
		}
	}
}

func dfsWithoutSlopes(graph Graph, destination Point, currentPath []Point, visited map[Point]bool, currentDistance int, results *[]int) {
	lastVisited := currentPath[len(currentPath)-1]

	if lastVisited == destination {
		*results = append(*results, currentDistance)
		return
	}

	for point, distance := range graph[lastVisited] {
		if visited[point] {
			continue
		}

		visited[point] = true

		dfsWithoutSlopes(
			graph,
			destination,
			append(currentPath, point),
			visited,
			currentDistance+distance,
			results,
		)

		visited[point] = false
	}
}

func findLongestPathWithoutSlopes(graph Graph, destination Point) int {
	allPaths := make([]int, 0)
	visited := make(map[Point]bool)
	dfsWithoutSlopes(graph, destination, []Point{{row: 0, col: 1}}, visited, 0, &allPaths)

	longest := 0

	for _, pathLen := range allPaths {
		if pathLen > longest {
			longest = pathLen
		}
	}

	return longest
}

func main() {
	field := getInputData()
	fmt.Println("Solution 1 is", findLongestPath(field, true))

	graph := gridToGraph(field)
	collapseGraph(graph)
	fmt.Println("Solution 2 is", findLongestPathWithoutSlopes(graph, Point{row: len(field) - 1, col: len(field[0]) - 2}))
}
