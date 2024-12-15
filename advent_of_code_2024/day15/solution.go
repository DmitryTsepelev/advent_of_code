package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	rowIdx int
	colIdx int
}

var SMALL_BOX = 0
var LEFT_SIDE = 1
var RIGHT_SIDE = 2

type Field struct {
	robot   Point
	width   int
	height  int
	boxMap  map[Point]int
	wallMap map[Point]bool
}

func getInputData() (Field, Field, []rune) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	smallField := Field{boxMap: make(map[Point]int), wallMap: make(map[Point]bool)}
	bigField := Field{boxMap: make(map[Point]int), wallMap: make(map[Point]bool)}

	movements := []rune{}

	rowIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		smallField.width = len(line)
		bigField.width = len(line) * 2

		for colIdx, val := range line {
			point := Point{rowIdx, colIdx}

			switch val {
			case '#':
				smallField.wallMap[point] = true
			case 'O':
				smallField.boxMap[point] = SMALL_BOX
			case '@':
				smallField.robot = point
			}

			lPoint := Point{rowIdx, colIdx * 2}
			rPoint := Point{rowIdx, colIdx*2 + 1}
			switch val {
			case '#':
				bigField.wallMap[lPoint] = true
				bigField.wallMap[rPoint] = true
			case 'O':
				bigField.boxMap[lPoint] = LEFT_SIDE
				bigField.boxMap[rPoint] = RIGHT_SIDE
			case '@':
				bigField.robot = lPoint
			}
		}

		rowIdx++
	}

	smallField.height = rowIdx
	bigField.height = rowIdx

	for scanner.Scan() {
		movements = append(movements, []rune(scanner.Text())...)
	}

	return smallField, bigField, movements
}

const LEFT = '<'
const RIGHT = '>'
const UP = '^'
const DOWN = 'v'

var deltas = map[rune]Point{
	LEFT:  {0, -1},
	RIGHT: {0, 1},
	UP:    {-1, 0},
	DOWN:  {1, 0},
}

func drawField(field Field) {
	fmt.Println("===================================")

	for rowIdx := 0; rowIdx < field.height; rowIdx++ {
		for colIdx := 0; colIdx < field.width; colIdx++ {
			point := Point{rowIdx, colIdx}

			if field.robot == point {
				fmt.Print("@")
			} else if boxKind, ok := field.boxMap[point]; ok {
				switch boxKind {
				case SMALL_BOX:
					fmt.Print("O")
				case LEFT_SIDE:
					fmt.Print("[")
				case RIGHT_SIDE:
					fmt.Print("]")
				}

			} else if _, ok := field.wallMap[point]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}
}

func simulate(field Field, movements []rune, drawMode bool) {
	for _, movDir := range movements {
		delta := deltas[movDir]

		canMove := true
		var movedBox *Point
		currentPoint := Point{field.robot.rowIdx, field.robot.colIdx}

		for {
			currentPoint.rowIdx += delta.rowIdx
			currentPoint.colIdx += delta.colIdx

			if _, ok := field.boxMap[currentPoint]; ok {
				// found box to move
				if movedBox == nil {
					movedBox = &Point{currentPoint.rowIdx, currentPoint.colIdx}
				}
			} else if _, ok := field.wallMap[currentPoint]; ok {
				// found wall
				canMove = false
				break
			} else {
				// found empty space
				break
			}
		}

		if canMove {
			if movedBox != nil {
				delete(field.boxMap, *movedBox)
				field.boxMap[currentPoint] = SMALL_BOX
			}

			field.robot.rowIdx += delta.rowIdx
			field.robot.colIdx += delta.colIdx
		}

		if drawMode {
			drawField(field)
		}
	}
}

func addBoxToList(queue *[]Point, boxKind int, boxPoint Point) {
	*queue = append(*queue, boxPoint)

	if boxKind == LEFT_SIDE {
		*queue = append(*queue, Point{boxPoint.rowIdx, boxPoint.colIdx + 1})
	} else {
		*queue = append(*queue, Point{boxPoint.rowIdx, boxPoint.colIdx - 1})
	}
}

func (field *Field) getAnotherBoxSide(boxPoint Point) Point {
	boxKind := field.boxMap[boxPoint]
	if boxKind == LEFT_SIDE {
		return Point{boxPoint.rowIdx, boxPoint.colIdx + 1}
	} else {
		return Point{boxPoint.rowIdx, boxPoint.colIdx - 1}
	}
}

func (field *Field) canMoveBoxVertically(delta Point, boxPoint Point, boxMovements map[Point]Point) bool {
	boxPoint2 := field.getAnotherBoxSide(boxPoint)

	nextPoint := Point{boxPoint.rowIdx + delta.rowIdx, boxPoint.colIdx + delta.colIdx}
	nextPoint2 := Point{boxPoint2.rowIdx + delta.rowIdx, boxPoint2.colIdx + delta.colIdx}

	boxMovements[boxPoint] = nextPoint
	boxMovements[boxPoint2] = nextPoint2

	if _, ok := field.wallMap[nextPoint]; ok {
		return false
	}

	if _, ok := field.wallMap[nextPoint2]; ok {
		return false
	}

	canMove := true
	if _, ok := field.boxMap[nextPoint]; ok {
		canMove = canMove && field.canMoveBoxVertically(delta, nextPoint, boxMovements)
	}

	if _, ok := field.boxMap[nextPoint2]; ok {
		canMove = canMove && field.canMoveBoxVertically(delta, nextPoint2, boxMovements)
	}

	return canMove
}

func (field *Field) tryMoveBoxHorizontally(delta Point, boxPoint Point) bool {
	boxPoint2 := field.getAnotherBoxSide(boxPoint)

	nextPoint := Point{boxPoint2.rowIdx + delta.rowIdx, boxPoint2.colIdx + delta.colIdx}

	if _, ok := field.wallMap[nextPoint]; ok {
		return false
	}

	canMove := true
	if _, ok := field.boxMap[nextPoint]; ok {
		// next box can be moved
		canMove = field.tryMoveBox(delta, nextPoint)
	}

	// can move box
	if canMove {
		field.boxMap[nextPoint] = field.boxMap[boxPoint2]
		field.boxMap[boxPoint2] = field.boxMap[boxPoint]
		delete(field.boxMap, boxPoint)
	}

	return canMove
}

func (field *Field) tryMoveBox(delta Point, boxPoint Point) bool {
	if delta.rowIdx == 0 {
		return field.tryMoveBoxHorizontally(delta, boxPoint)
	}

	// vertical movement
	boxMovements := map[Point]Point{}
	if field.canMoveBoxVertically(delta, boxPoint, boxMovements) {
		boxMapCopy := make(map[Point]int)
		for point, side := range field.boxMap {
			if target, ok := boxMovements[point]; ok {
				boxMapCopy[target] = side
			} else {
				boxMapCopy[point] = side
			}
		}
		field.boxMap = boxMapCopy

		return true
	}

	return false
}

func (field *Field) tryMoveRobot(delta Point) bool {
	nextPoint := Point{field.robot.rowIdx + delta.rowIdx, field.robot.colIdx + delta.colIdx}
	if _, ok := field.wallMap[nextPoint]; ok {
		// wall found, no movement
		return false
	}

	if _, ok := field.boxMap[nextPoint]; ok {
		return field.tryMoveBox(delta, nextPoint)
	}

	// empty space
	return true
}

func simulateBig(field *Field, movements []rune, drawMode bool) {
	for _, movDir := range movements {
		delta := deltas[movDir]

		if field.tryMoveRobot(delta) {
			field.robot.rowIdx += delta.rowIdx
			field.robot.colIdx += delta.colIdx
		}

		if drawMode {
			drawField(*field)
		}
	}
}

func sumBoxGPS(field Field) int64 {
	var result int64
	for box, kind := range field.boxMap {
		if kind != RIGHT_SIDE {
			result += int64(100*box.rowIdx + box.colIdx)
		}
	}
	return result
}

func main() {
	field, bigField, movements := getInputData()

	simulate(field, movements, false)
	fmt.Println("Part 1 solution is", sumBoxGPS(field))

	simulateBig(&bigField, movements, false)
	fmt.Println("Part 2 solution is", sumBoxGPS(bigField))
}
