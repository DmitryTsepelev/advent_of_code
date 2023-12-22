package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Position struct {
	x int
	y int
	z int
}

type Brick struct {
	position1 Position
	position2 Position
}

func getInputData() []Brick {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	re, _ := regexp.Compile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)

	scanner := bufio.NewScanner(file)

	bricks := []Brick{}

	for scanner.Scan() {
		line := scanner.Text()

		match := re.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			x1S, y1S, z1S, x2S, y2S, z2S := match[0][1], match[0][2], match[0][3], match[0][4], match[0][5], match[0][6]

			x1, _ := strconv.Atoi(x1S)
			y1, _ := strconv.Atoi(y1S)
			z1, _ := strconv.Atoi(z1S)
			x2, _ := strconv.Atoi(x2S)
			y2, _ := strconv.Atoi(y2S)
			z2, _ := strconv.Atoi(z2S)

			bricks = append(bricks, Brick{
				position1: Position{x: x1, y: y1, z: z1},
				position2: Position{x: x2, y: y2, z: z2},
			})

			continue
		}

		panic(line)
	}

	return bricks
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

type Occupations = map[Position]int

func getBrickPositions(brick Brick) []Position {
	positions := make([]Position, 0)

	minX, maxX := minmax(brick.position1.x, brick.position2.x)
	minY, maxY := minmax(brick.position1.y, brick.position2.y)
	minZ, maxZ := minmax(brick.position1.z, brick.position2.z)

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			for z := minZ; z <= maxZ; z++ {
				positions = append(positions, Position{x: x, y: y, z: z})
			}
		}
	}

	return positions
}

func bricksToMap(bricks []Brick) Occupations {
	occupied := make(Occupations, 0)

	for id, brick := range bricks {
		for _, position := range getBrickPositions(brick) {
			occupied[position] = id
		}
	}

	return occupied
}

func fallAndSettle(occupations *Occupations, bricks *[]Brick) {
	for {
		brickMoved := false

		for brickId := 0; brickId < len(*bricks); brickId++ {
			brick := (*bricks)[brickId]

			positions := getBrickPositions(brick)

			canMove := true
			for _, position := range positions {
				positionBelow := Position{x: position.x, y: position.y, z: position.z - 1}

				if positionBelow.z < 1 {
					// ground level
					canMove = false
					break
				}

				if id, found := (*occupations)[positionBelow]; found && id != brickId {
					// another brick
					canMove = false
					break
				}
			}

			if !canMove {
				continue
			}

			brickMoved = true

			// move down brick itself
			(*bricks)[brickId].position1.z--
			(*bricks)[brickId].position2.z--

			// move down brick on map
			for _, position := range positions {
				delete(*occupations, position)
			}

			for _, position := range positions {
				positionBelow := Position{x: position.x, y: position.y, z: position.z - 1}
				(*occupations)[positionBelow] = brickId
			}
		}

		if !brickMoved {
			break
		}
	}
}

func supportedBy(bricks *[]Brick, occupations *Occupations) [][]bool {
	graph := make([][]bool, len(*bricks))
	for i := 0; i < len(*bricks); i++ {
		graph[i] = make([]bool, len(graph))
	}

	for id, brick := range *bricks {
		positions := getBrickPositions(brick)

		for _, position := range positions {
			positionBelow := Position{x: position.x, y: position.y, z: position.z - 1}

			if foundId, found := (*occupations)[positionBelow]; found && foundId != id {
				graph[id][foundId] = true
			}
		}
	}

	return graph
}

func checkSafety(occupations *Occupations, bricks *[]Brick) int {
	graph := supportedBy(bricks, occupations)

	safeCount := 0

	for brickId := range *bricks {
		notSupportsAnything := true
		allRowsHaveOtherSupporters := true

		for row := 0; row < len(*bricks); row++ {
			if row == brickId {
				// skip itself
				continue
			}

			if graph[row][brickId] {
				// supports
				notSupportsAnything = false
				supportersCount := 0
				for col := 0; col < len(*bricks); col++ {
					if graph[row][col] {
						supportersCount++
					}
				}

				if supportersCount <= 1 {
					allRowsHaveOtherSupporters = false
				}
			}
		}

		if notSupportsAnything || allRowsHaveOtherSupporters {
			safeCount++
		}
	}

	return safeCount
}

func chainReaction(rootBrickId int, graph [][]bool) []int {
	fallen := make(map[int]bool, 0)
	queue := []int{rootBrickId}

	for len(queue) > 0 {
		nextQueue := []int{}

		for _, currentBrickId := range queue {
			fallen[currentBrickId] = true

			for row := 0; row < len(graph); row++ {
				if row == currentBrickId {
					// skip itself
					continue
				}

				if graph[row][currentBrickId] {
					// supports
					supportersCount := 0
					for col := 0; col < len(graph); col++ {
						if graph[row][col] && !fallen[col] {
							supportersCount++
						}
					}

					if supportersCount == 0 {
						nextQueue = append(nextQueue, row)
					}
				}
			}
		}

		queue = nextQueue
	}

	list := []int{}
	for k := range fallen {
		if k == rootBrickId {
			continue
		}
		list = append(list, k)
	}

	return list
}

func findMaxFailing(occupations *Occupations, bricks *[]Brick) int {
	graph := supportedBy(bricks, occupations)

	sum := 0
	for brickId := range *bricks {
		sum += len(chainReaction(brickId, graph))
	}

	return sum
}

func main() {
	bricks := getInputData()
	occupations := bricksToMap(bricks)
	fallAndSettle(&occupations, &bricks)

	fmt.Println("Solution 1 is", checkSafety(&occupations, &bricks))
	fmt.Println("Solution 2 is", findMaxFailing(&occupations, &bricks))
}
