package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction byte
	distance  int64
	color     string
}

const (
	RIGHT = 'R'
	LEFT  = 'L'
	UP    = 'U'
	DOWN  = 'D'
)

func getInputData(fromColor bool) []Instruction {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := []Instruction{}

	for scanner.Scan() {
		line := scanner.Text()

		cmp := strings.Split(line, " ")

		instruction := Instruction{}

		if fromColor {
			hex := cmp[2][1 : len(cmp[2])-1]

			switch hex[len(hex)-1] {
			case '0':
				instruction.direction = RIGHT
			case '1':
				instruction.direction = DOWN
			case '2':
				instruction.direction = LEFT
			case '3':
				instruction.direction = UP
			}

			convertedDistance, _ := strconv.ParseInt(hex[1:len(hex)-1], 16, 64)
			instruction.distance = convertedDistance
		} else {
			instruction.direction = cmp[0][0]
			convertedDistance, _ := strconv.Atoi(cmp[1])
			instruction.distance = int64(convertedDistance)
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

type Point struct {
	y int
	x int
}

var deltas = map[byte]([]int){
	RIGHT: {0, 1},
	LEFT:  {0, -1},
	UP:    {-1, 0},
	DOWN:  {1, 0},
}

func buildPoints(instructions []Instruction) []Point {
	points := []Point{}

	currentRow := 0
	currentCol := 0

	for _, instruction := range instructions {
		delta := deltas[instruction.direction]
		dRow, dCol := delta[0], delta[1]

		points = append(points, Point{y: currentRow, x: currentCol})

		currentRow += dRow * int(instruction.distance)
		currentCol += dCol * int(instruction.distance)
	}
	return points
}

func perimeter(points []Point) int {
	result := 0

	for idx := 1; idx <= len(points); idx++ {
		prev := points[idx-1]

		current := points[0]
		if idx < len(points) {
			current = points[idx]
		}

		length := (current.y - prev.y) + (current.x - prev.x)
		if length < 0 {
			length = -length
		}
		result += length
	}

	return result
}

func shoelace(points []Point) int {
	sum := 0

	for currentIdx := 0; currentIdx < len(points); currentIdx++ {
		nextIdx := currentIdx + 1
		if currentIdx == len(points)-1 {
			nextIdx = 0
		}
		sum += points[currentIdx].x*points[nextIdx].y - points[nextIdx].x*points[currentIdx].y
	}

	return sum / 2
}

func findArea(instructions []Instruction) int {
	points := buildPoints(instructions)
	return shoelace(points) + perimeter(points)/2 + 1
}

func main() {
	fmt.Println("Solution 1 is", findArea(getInputData(false)))
	fmt.Println("Solution 2 is", findArea(getInputData(true)))
}
