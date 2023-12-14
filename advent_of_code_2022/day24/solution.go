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

func getInputData() ([]Blizzard, int, int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	blizzards := make([]Blizzard, 0)

	row := 0
	col := 0
	for scanner.Scan() {
		line := scanner.Text()

		col = len(line)

		for col, dir := range line {
			if dir == '.' || dir == '#' {
				continue
			}

			blizzard := Blizzard{row: row, col: col}

			blizzard.dir = dir

			blizzards = append(blizzards, blizzard)
		}

		row++
	}

	return blizzards, row - 1, col - 1
}

func moveBlizzards(blizzards []Blizzard, fieldHeight int, fieldWidth int) []Blizzard {
	newBlizzards := make([]Blizzard, len(blizzards))

	for idx, blizzard := range blizzards {
		newBlizzard := Blizzard{
			row: blizzard.row,
			col: blizzard.col,
			dir: blizzard.dir,
		}

		switch newBlizzard.dir {
		case RIGHT:
			newBlizzard.col++
			if newBlizzard.col == fieldWidth {
				newBlizzard.col = 1
			}
		case LEFT:
			newBlizzard.col--
			if newBlizzard.col == 0 {
				newBlizzard.col = fieldWidth - 1
			}
		case DOWN:
			newBlizzard.row++
			if newBlizzard.row == fieldHeight {
				newBlizzard.row = 1
			}
		case UP:
			newBlizzard.row--
			if newBlizzard.row == 0 {
				newBlizzard.row = fieldHeight - 1
			}
		}

		newBlizzards[idx] = newBlizzard
	}

	return newBlizzards
}

func printField(blizzards []Blizzard, fieldHeight int, fieldWidth int) {
	for row := 0; row <= fieldHeight; row++ {
		for col := 0; col <= fieldWidth; col++ {
			if (row == 0 && col != 1) || (row == fieldHeight && col != fieldWidth-1) || col == 0 || col == fieldWidth {
				fmt.Print("#")
			} else {
				found := make([]rune, 0)
				for _, blizzard := range blizzards {
					if blizzard.row == row && blizzard.col == col {
						found = append(found, blizzard.dir)
					}
				}

				if len(found) == 0 {
					fmt.Print(".")
				} else if len(found) == 1 {
					fmt.Print(string(found[0]))
				} else {
					fmt.Print(len(found))
				}
			}
		}

		fmt.Println()
	}
}

// func findLoop(blizzards []Blizzard, fieldHeight int, fieldWidth int) int {
// 	current := blizzards

// 	idx := 0
// 	for {
// 		current = moveBlizzards(current, fieldHeight, fieldWidth)
// 		if reflect.DeepEqual(blizzards, current) {
// 			return idx
// 		}
// 		idx++
// 	}
// }

type Point struct {
	row int
	col int
}

func bfs(blizzards []Blizzard, fieldHeight int, fieldWidth int) int {
	queue := map[Point]bool{
		{row: 0, col: 1}: true,
	}
	// printField(blizzards, fieldHeight, fieldWidth)
	// fmt.Println()

	minute := 0
	for {
		minute++

		// fmt.Println("minute", minute)

		nextQueue := make(map[Point]bool, 0)
		// fmt.Println("Before", len(blizzards))
		blizzards = moveBlizzards(blizzards, fieldHeight, fieldWidth)
		// fmt.Println("after", len(blizzards))
		// printField(blizzards, fieldHeight, fieldWidth)

		for point := range queue {
			// fmt.Println(fieldHeight, fieldWidth)
			if point.row == fieldHeight && point.col == fieldWidth-1 {
				return minute
			}
			canStay := true
			canGoLeft := point.row > 0 && point.col > 1
			canGoRight := point.row > 0 && point.col < fieldWidth-1
			canGoUp := point.row > 1
			canGoDown := point.row < fieldHeight-1 || point.col == fieldWidth-1
			// fmt.Println("fieldHeight", fieldHeight)
			// fmt.Println(point, canStay, canGoLeft, canGoRight, canGoUp, canGoDown)
			for _, nextBlizzard := range blizzards {
				if point.row == nextBlizzard.row && point.col == nextBlizzard.col {
					canStay = false
				}

				if canGoUp && point.row-1 == nextBlizzard.row && point.col == nextBlizzard.col {
					canGoUp = false
				}

				if canGoDown && point.row+1 == nextBlizzard.row && point.col == nextBlizzard.col {
					canGoDown = false
				}

				if canGoLeft && point.row == nextBlizzard.row && point.col-1 == nextBlizzard.col {
					canGoLeft = false
				}

				if canGoRight && point.row == nextBlizzard.row && point.col+1 == nextBlizzard.col {
					canGoRight = false
				}
			}
			// fmt.Println(point, "canStay =", canStay, "canGoLeft =", canGoLeft, "canGoRight =", canGoRight, "canGoUp =", canGoUp, "canGoDown =", canGoDown)

			// fmt.Println(point)
			if canStay {
				// fmt.Println("canStay")
				nextQueue[point] = true
			}

			if canGoLeft {
				// fmt.Println("canGoLeft")
				nextQueue[Point{row: point.row, col: point.col - 1}] = true
			}

			if canGoRight {
				// fmt.Println("canGoRight")
				nextQueue[Point{row: point.row, col: point.col + 1}] = true
			}

			if canGoUp {
				// fmt.Println("canGoUp")
				nextQueue[Point{row: point.row - 1, col: point.col}] = true
			}

			if canGoDown {
				// fmt.Println("canGoDown")
				nextQueue[Point{row: point.row + 1, col: point.col}] = true
			}
			// fmt.Println()
		}

		queue = nextQueue
	}
}

func main() {
	blizzards, fieldHeight, fieldWidth := getInputData()

	// for i := 0; i < 6; i++ {
	// 	printField(blizzards, fieldHeight, fieldWidth)
	// 	blizzards = moveBlizzards(blizzards, fieldHeight, fieldWidth)
	// 	fmt.Println()
	// }

	fmt.Println(bfs(blizzards, fieldHeight, fieldWidth))
}
