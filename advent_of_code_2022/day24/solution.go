package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	UP    = '^'
	DOWN  = 'v'
	LEFT  = '<'
	RIGHT = '>'
)

type Blizzard struct {
	row int
	col int
	dir rune
}

func getInputData() ([]*Blizzard, map[Point]bool, int, int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	blizzards := make([]*Blizzard, 0)
	wallMap := map[Point]bool{}

	row := 0
	col := 0
	for scanner.Scan() {
		line := scanner.Text()

		col = len(line)

		for col, dir := range line {
			switch dir {
			case '.':
				continue
			case '#':
				wallMap[Point{row, col}] = true
			default:
				blizzard := Blizzard{row: row, col: col, dir: dir}

				blizzards = append(blizzards, &blizzard)
			}
		}

		row++
	}

	return blizzards, wallMap, row - 1, col - 1
}

func moveBlizzards(blizzards []*Blizzard, wallMap map[Point]bool, fieldHeight int, fieldWidth int) {
	for _, blizzard := range blizzards {
		switch blizzard.dir {
		case RIGHT:
			blizzard.col++
			if _, ok := wallMap[Point{blizzard.row, blizzard.col}]; ok {
				blizzard.col = 1
			}
		case LEFT:
			blizzard.col--
			if _, ok := wallMap[Point{blizzard.row, blizzard.col}]; ok {
				blizzard.col = fieldWidth - 1
			}
		case DOWN:
			blizzard.row++
			if _, ok := wallMap[Point{blizzard.row, blizzard.col}]; ok {
				blizzard.row = 1
			}
		case UP:
			blizzard.row--
			if _, ok := wallMap[Point{blizzard.row, blizzard.col}]; ok {
				blizzard.row = fieldHeight - 1
			}
		}
	}
}

type Point struct {
	row int
	col int
}

type Elf struct {
	point Point
	stage Stage
}

var elfDeltas = []Point{
	{0, 0},
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func validElf(elfPoint Point, fieldHeight, fieldWidth int, blizzardMap map[Point]bool, wallMap map[Point]bool) bool {
	if elfPoint.row < 0 || elfPoint.row > fieldHeight || elfPoint.col < 0 || elfPoint.col > fieldWidth {
		return false
	}

	if _, ok := blizzardMap[elfPoint]; ok {
		return false
	}

	if _, ok := wallMap[elfPoint]; ok {
		return false
	}

	return true
}

type Stage = int

const FIRST_FORWARD = 0
const BACK = 1
const SECOND_FORWARD = 2

func reachedStartPoint(elf Point) bool {
	return elf.row == 0 && elf.col == 1
}

func reachedFinishPoint(elf Point, fieldHeight, fieldWidth int) bool {
	return elf.row == fieldHeight && elf.col == fieldWidth-1
}

func findPath(blizzards []*Blizzard, wallMap map[Point]bool, fieldHeight int, fieldWidth int, withReturn bool) int {
	elves := map[Elf]bool{
		{Point{0, 1}, FIRST_FORWARD}: true,
	}

	minute := 1

	for {
		moveBlizzards(blizzards, wallMap, fieldHeight, fieldWidth)

		nextElves := map[Elf]bool{}

		blizzardMap := map[Point]bool{}
		for _, blizzard := range blizzards {
			blizzardMap[Point{blizzard.row, blizzard.col}] = true
		}

		for elf := range elves {
			for _, delta := range elfDeltas {
				nextElf := Elf{Point{elf.point.row + delta.row, elf.point.col + delta.col}, elf.stage}

				if validElf(nextElf.point, fieldHeight, fieldWidth, blizzardMap, wallMap) {
					if reachedFinishPoint(nextElf.point, fieldHeight, fieldWidth) {
						if withReturn {
							if elf.stage == SECOND_FORWARD {
								return minute
							} else {
								nextElf.stage = BACK
							}
						} else {
							if elf.stage == FIRST_FORWARD {
								return minute
							}
						}
					} else if reachedStartPoint(nextElf.point) && elf.stage == BACK {
						nextElf.stage = SECOND_FORWARD
					}

					nextElves[nextElf] = true
				}
			}
		}

		minute++
		elves = nextElves
	}
}

func main() {
	blizzards, wallMap, fieldHeight, fieldWidth := getInputData()

	fmt.Println("Part 1 solution is", findPath(blizzards, wallMap, fieldHeight, fieldWidth, false))
	fmt.Println("Part 2 solution is", findPath(blizzards, wallMap, fieldHeight, fieldWidth, true))
}
