package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() [][2]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := [][2]int{}
	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), ",")

		x, _ := strconv.Atoi(cmp[0])
		y, _ := strconv.Atoi(cmp[1])

		points = append(points, [2]int{x, y})
	}

	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func areaBetween(p1, p2 [2]int) int {
	return (max(p1[0], p2[0]) - min(p1[0], p2[0]) + 1) * (max(p1[1], p2[1]) - min(p1[1], p2[1]) + 1)
}

func solve1(points [][2]int) int {
	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			maxArea = max(maxArea, areaBetween(p1, p2))
		}
	}

	return maxArea
}

type Segment struct {
	A, B [2]int
}

func (s *Segment) intersectsRect(x1, x2, y1, y2 int) bool {
	recMinX := min(x1, x2) + 1
	recMaxX := max(x1, x2) - 1
	recMinY := min(y1, y2) + 1
	recMaxY := max(y1, y2) - 1

	segMinX := min(s.A[0], s.B[0])
	segMaxX := max(s.A[0], s.B[0])
	segMinY := min(s.A[1], s.B[1])
	segMaxY := max(s.A[1], s.B[1])

	if segMaxX < recMinX || segMinX > recMaxX {
		return false
	}
	if segMaxY < recMinY || segMinY > recMaxY {
		return false
	}
	return true
}

func buildLines(points [][2]int) []Segment {
	segments := make([]Segment, 0, len(points)+1)
	for i := 0; i < len(points)-1; i++ {
		segment := Segment{A: points[i], B: points[i+1]}
		segments = append(segments, segment)
	}

	segments = append(segments, Segment{A: points[len(points)-1], B: points[0]})
	return segments
}

func solve2(points [][2]int) int {
	greenSegments := buildLines(points)

	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]

			x1 := max(p1[0], p2[0])
			x2 := min(p1[0], p2[0])

			y1 := max(p1[1], p2[1])
			y2 := min(p1[1], p2[1])

			isValid := true
			for _, greenSegment := range greenSegments {
				if greenSegment.intersectsRect(x1, x2, y1, y2) {
					isValid = false
					break
				}
			}

			if isValid {
				maxArea = max(maxArea, areaBetween(p1, p2))
			}
		}
	}

	return maxArea
}

func main() {
	points := getInputData()

	fmt.Println("Part 1 solution is", solve1(points))
	fmt.Println("Part 2 solution is", solve2(points))
}
