package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph = [][]int

func getInputData() Graph {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := Graph{}

	for scanner.Scan() {
		row := []int{}
		for _, r := range scanner.Text() {
			h, _ := strconv.Atoi(string(r))
			row = append(row, h)
		}
		graph = append(graph, row)
	}

	return graph
}

var deltas = []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

const END_OF_TRAIL = 9

type Point struct {
	rowIdx int
	colIdx int
}

func countTrailRating(graph Graph, startRow, startCol int) (int, int) {
	var trailCounts int
	trailEndsMap := make(map[Point]bool)

	height, width := len(graph), len(graph[0])
	queue := []Point{{startRow, startCol}}

	for len(queue) > 0 {
		nextQueue := []Point{}

		for _, current := range queue {
			for _, delta := range deltas {
				nextRow := current.rowIdx + delta.rowIdx
				nextCol := current.colIdx + delta.colIdx

				if nextRow < 0 || nextRow == height ||
					nextCol < 0 || nextCol == width ||
					graph[nextRow][nextCol]-graph[current.rowIdx][current.colIdx] != 1 {
					continue
				}

				if graph[nextRow][nextCol] == END_OF_TRAIL {
					trailEndsMap[Point{nextRow, nextCol}] = true
					trailCounts++
				} else {
					nextQueue = append(nextQueue, Point{nextRow, nextCol})
				}
			}
		}

		queue = nextQueue
	}

	return len(trailEndsMap), trailCounts
}

func findTrails(graph Graph) (int, int) {
	var totalRating, totalCount int

	for rowIdx, row := range graph {
		for colIdx, val := range row {
			if val == 0 {
				rating, count := countTrailRating(graph, rowIdx, colIdx)

				totalRating += rating
				totalCount += count
			}
		}
	}

	return totalRating, totalCount
}

func main() {
	graph := getInputData()
	trailScore, trailRating := findTrails(graph)
	fmt.Println("Part 1 solution is", trailScore)
	fmt.Println("Part 2 solution is", trailRating)
}
