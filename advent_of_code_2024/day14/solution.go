package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Robot struct {
	location *Point
	velocity Point
}

type RobotMap = map[Robot]bool

func getInputData() RobotMap {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	robotMap := RobotMap{}

	for scanner.Scan() {
		line := scanner.Text()
		cmp := strings.Split(line, " ")

		pcmp := strings.Split(cmp[0][2:], ",")
		px, _ := strconv.Atoi(pcmp[0])
		py, _ := strconv.Atoi(pcmp[1])

		vcmp := strings.Split(cmp[1][2:], ",")
		vx, _ := strconv.Atoi(vcmp[0])
		vy, _ := strconv.Atoi(vcmp[1])

		robot := Robot{location: &Point{x: px, y: py}, velocity: Point{x: vx, y: vy}}
		robotMap[robot] = true
	}

	return robotMap
}

const EASTER_EGG_MODE = -1

func simulate(robotMap *RobotMap, width, height, steps int) int {
	for i := 1; ; i++ {
		for r := range *robotMap {
			r.location.x = (r.location.x + r.velocity.x) % width
			if r.location.x < 0 {
				r.location.x = width + r.location.x
			}

			r.location.y = (r.location.y + r.velocity.y) % height
			if r.location.y < 0 {
				r.location.y = height + r.location.y
			}
		}

		if steps == EASTER_EGG_MODE {
			if isEasterEgg(robotMap) {
				drawField(robotMap, width, height)
				return i
			}
		} else if i == steps {
			break
		}
	}

	return steps
}

func drawField(robotMap *RobotMap, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			foundCount := 0
			for robot := range *robotMap {
				if robot.location.x == x && robot.location.y == y {
					foundCount++
				}
			}

			if foundCount == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(foundCount)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func isEasterEgg(robotMap *RobotMap) bool {
	uniqPoints := map[Point]bool{}

	for robot := range *robotMap {
		if _, ok := uniqPoints[*robot.location]; ok {
			return false
		}
		uniqPoints[*robot.location] = true
	}

	return true
}

const TOP_LEFT = 0
const TOP_RIGHT = 1
const BOTTOM_LEFT = 2
const BOTTOM_RIGHT = 3

func calculateQuadrants(robotMap RobotMap, width, height int) []int {
	midY := height / 2
	midX := width / 2

	counts := make([]int, 4)

	for r := range robotMap {
		if r.location.x < midX {
			if r.location.y < midY {
				counts[TOP_LEFT]++
			} else if r.location.y > midY {
				counts[TOP_RIGHT]++
			}
		} else if r.location.x > midX {
			if r.location.y < midY {
				counts[BOTTOM_LEFT]++
			} else if r.location.y > midY {
				counts[BOTTOM_RIGHT]++
			}
		}
	}

	return counts
}

func safetyFactor(robotMap RobotMap, width, height int) int {
	counts := calculateQuadrants(robotMap, width, height)

	mult := 1
	for i := 0; i < 4; i++ {
		mult *= counts[i]
	}
	return mult
}

func main() {
	width, height := 101, 103
	robotMap := getInputData()
	simulate(&robotMap, width, height, 100)
	fmt.Println("Part 1 solution is", safetyFactor(robotMap, width, height))

	robotMap = getInputData()
	fmt.Println("Part 2 solution is", simulate(&robotMap, width, height, EASTER_EGG_MODE))
}
