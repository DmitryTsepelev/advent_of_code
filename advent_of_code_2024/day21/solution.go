package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	codes := []string{}

	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	return codes
}

type Point struct {
	rowIdx int
	colIdx int
}

type Graph = [][](rune)

var directions = map[Point]rune{
	{0, -1}: '<',
	{-1, 0}: '^',
	{0, 1}:  '>',
	{1, 0}:  'v',
}

type Paths = map[[2]rune]([]string)

var INVALID_VALUE = 'X'

func adj(currentPath []Point, graph Graph, visited map[Point]bool) map[Point]rune {
	currentPoint := currentPath[len(currentPath)-1]

	res := map[Point]rune{}

	for dir, dirSym := range directions {
		nextPoint := Point{currentPoint.rowIdx + dir.rowIdx, currentPoint.colIdx + dir.colIdx}

		if nextPoint.rowIdx < 0 || nextPoint.rowIdx >= len(graph) ||
			nextPoint.colIdx < 0 || nextPoint.colIdx >= len(graph[0]) ||
			graph[nextPoint.rowIdx][nextPoint.colIdx] == INVALID_VALUE {
			continue
		}

		if visited, ok := visited[nextPoint]; !ok || visited == false {
			res[nextPoint] = dirSym
		}
	}

	return res
}

func dfs(graph Graph, destination Point, currentPath []Point, currentSequence string, visited map[Point]bool) []string {
	lastVisited := currentPath[len(currentPath)-1]
	if lastVisited == destination {
		return []string{currentSequence}
	}

	results := []string{}
	for nextPoint, dirSym := range adj(currentPath, graph, visited) {
		visited[nextPoint] = true

		currentPath = append(currentPath, nextPoint)

		results = append(
			results,
			dfs(
				graph,
				destination,
				currentPath,
				currentSequence+string(dirSym),
				visited,
			)...,
		)

		currentPath = currentPath[:len(currentPath)-1]

		visited[nextPoint] = false
	}

	return results
}

func filterShortestPaths(results []string) []string {
	filtered := []string{}

	sort.Slice(results, func(i, j int) bool {
		return len(results[i]) < len(results[j])
	})

	for i := 0; i < len(results) && len(results[i]) == len(results[0]); i++ {
		filtered = append(filtered, results[i])
	}

	return filtered
}

var numGraph = Graph{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{'X', '0', 'A'},
}

var digitGraph = Graph{
	{'X', '^', 'A'},
	{'<', 'v', '>'},
}

func codeToNum(code string) int {
	digits := ""

	for _, r := range code {
		if r >= '0' && r <= '9' {
			digits += string(r)
		}
	}

	num, _ := strconv.Atoi(digits)
	return num
}

func findShortestPathsFor(graph Graph) map[[2]rune]([]string) {
	shortestPaths := make(map[[2]rune]([]string))

	for startRowIdx := 0; startRowIdx < len(graph); startRowIdx++ {
		for startColIdx := 0; startColIdx < len(graph[startRowIdx]); startColIdx++ {
			for finishRowIdx := 0; finishRowIdx < len(graph); finishRowIdx++ {
				for finishColIdx := 0; finishColIdx < len(graph[finishRowIdx]); finishColIdx++ {
					if graph[startRowIdx][startColIdx] == INVALID_VALUE ||
						graph[finishRowIdx][finishColIdx] == INVALID_VALUE {
						continue
					}
					results := dfs(graph, Point{finishRowIdx, finishColIdx}, []Point{{startRowIdx, startColIdx}}, "", map[Point]bool{{startRowIdx, startColIdx}: true})

					shortest := filterShortestPaths(results)
					shortestPaths[[2]rune{graph[startRowIdx][startColIdx], graph[finishRowIdx][finishColIdx]}] = shortest
				}
			}
		}
	}

	return shortestPaths
}

func sequenceForCode(code string, symPaths map[[2]rune]([]string)) []string {
	prevNumSym := 'A'
	paths := []string{""}

	for _, currentNumSym := range code {
		nextPaths := []string{}

		for _, symPath := range symPaths[[2]rune{prevNumSym, currentNumSym}] {
			for _, p := range paths {
				nextPaths = append(nextPaths, p+symPath+"A")
			}
		}
		prevNumSym = currentNumSym

		paths = nextPaths
	}

	return paths
}

func codeShortestSequence(code string, padCount int, digitPaths map[[2]rune]([]string), numPaths map[[2]rune]([]string)) int {
	numSequences := sequenceForCode(code, numPaths)
	fmt.Println(numSequences)

	prevSequences := numSequences

	for i := 0; i < padCount; i++ {
		nextSequences := []string{}

		for _, sq := range prevSequences {
			nextSequences = append(nextSequences, sequenceForCode(sq, digitPaths)...)
		}

		prevSequences = nextSequences
	}

	// for _, numSequence := range numSequences {
	// 	digitSequences1 := sequenceForCode(numSequence, digitPaths)

	// 	for _, digitSequence1 := range digitSequences1 {
	// 		digitSequences2 := sequenceForCode(digitSequence1, digitPaths)

	// 		for _, digitSequence2 := range digitSequences2 {
	// 			if len(digitSequence2) < shortest {
	// 				shortest = len(digitSequence2)
	// 			}
	// 		}
	// 	}
	// }

	shortest := math.MaxInt
	for _, sq := range prevSequences {
		if len(sq) < shortest {
			shortest = len(sq)
		}
	}

	return shortest
}

func solve(codes []string, padCount int) int {
	digitPaths := findShortestPathsFor(digitGraph)

	numPaths := findShortestPathsFor(numGraph)

	var sum int
	for _, code := range codes {
		sum += codeToNum(code) * codeShortestSequence(code, padCount, digitPaths, numPaths)
	}

	return sum
}

func main() {
	codes := getInputData()
	fmt.Println("Part 1 solution is", solve(codes, 2))
}
