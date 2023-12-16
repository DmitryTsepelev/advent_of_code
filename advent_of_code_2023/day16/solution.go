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

const (
	RIGHT = 0
	LEFT  = 1
	UP    = 2
	DOWN  = 3
)

type Beam struct {
	row int
	col int
	dir int
}

type Point struct {
	row int
	col int
}

const (
	EMPTY_SPACE         = '.'
	MIRROR_LBTR         = '/'
	MIRROR_TLBR         = '\\'
	SPLITTER_VERTICAL   = '|'
	SPLITTER_HORIZONTAL = '-'
)

func moveBeam(beam Beam) []Beam {
	newBeam := Beam{row: beam.row, col: beam.col, dir: beam.dir}

	switch beam.dir {
	case RIGHT:
		newBeam.col++
	case LEFT:
		newBeam.col--
	case DOWN:
		newBeam.row++
	case UP:
		newBeam.row--
	}

	return []Beam{newBeam}
}

func mirrorLBTR(beam Beam) []Beam {
	newBeam := Beam{row: beam.row, col: beam.col, dir: beam.dir}

	switch beam.dir {
	case RIGHT:
		newBeam.dir = UP
		newBeam.row--
	case LEFT:
		newBeam.dir = DOWN
		newBeam.row++
	case DOWN:
		newBeam.dir = LEFT
		newBeam.col--
	case UP:
		newBeam.dir = RIGHT
		newBeam.col++
	}

	return []Beam{newBeam}
}

func mirrorTLBR(beam Beam) []Beam {
	newBeam := Beam{row: beam.row, col: beam.col, dir: beam.dir}

	switch beam.dir {
	case RIGHT:
		newBeam.dir = DOWN
		newBeam.row++
	case LEFT:
		newBeam.dir = UP
		newBeam.row--
	case DOWN:
		newBeam.dir = RIGHT
		newBeam.col++
	case UP:
		newBeam.dir = LEFT
		newBeam.col--
	}

	return []Beam{newBeam}
}

func splitterVertical(beam Beam) []Beam {
	if beam.dir == LEFT || beam.dir == RIGHT {
		return []Beam{
			{row: beam.row - 1, col: beam.col, dir: UP},
			{row: beam.row + 1, col: beam.col, dir: DOWN},
		}
	}

	return moveBeam(beam)
}

func splitterHorizontal(beam Beam) []Beam {
	if beam.dir == UP || beam.dir == DOWN {
		return []Beam{
			{row: beam.row, col: beam.col - 1, dir: LEFT},
			{row: beam.row, col: beam.col + 1, dir: RIGHT},
		}
	}

	return moveBeam(beam)
}

func validateBeams(beams []Beam, height, width int) []Beam {
	valid := make([]Beam, 0)

	for _, beam := range beams {
		if beam.row >= 0 && beam.row < height && beam.col >= 0 && beam.col < width {
			valid = append(valid, beam)
		}
	}

	return valid
}

func countEnergized(field []string, startBeam Beam) int {
	height, width := len(field), len(field[0])

	beams := []Beam{startBeam}
	energized := make(map[Point]bool, 0)

	prevEnergizedCount := 0
	stabilizedCount := 0
	for {
		newBeamsSet := make(map[Beam]bool, 0)

		for _, beam := range beams {
			energized[Point{row: beam.row, col: beam.col}] = true

			var changedBeams []Beam
			switch field[beam.row][beam.col] {
			case EMPTY_SPACE:
				changedBeams = moveBeam(beam)
			case MIRROR_LBTR:
				changedBeams = mirrorLBTR(beam)
			case MIRROR_TLBR:
				changedBeams = mirrorTLBR(beam)
			case SPLITTER_VERTICAL:
				changedBeams = splitterVertical(beam)
			case SPLITTER_HORIZONTAL:
				changedBeams = splitterHorizontal(beam)
			}

			for _, validBeam := range validateBeams(changedBeams, height, width) {
				newBeamsSet[validBeam] = true
			}
		}

		energizedCount := len(energized)
		if energizedCount == prevEnergizedCount {
			stabilizedCount++
			if stabilizedCount > 3 {
				return energizedCount
			}
		} else {
			stabilizedCount = 0
		}

		beams = make([]Beam, 0)
		for beam := range newBeamsSet {
			beams = append(beams, beam)
		}

		prevEnergizedCount = energizedCount
	}
}

func findBestConfig(field []string) int {
	height, width := len(field), len(field[0])

	candidates := []int{
		// left top
		countEnergized(field, Beam{row: 0, col: 0, dir: RIGHT}),
		countEnergized(field, Beam{row: 0, col: 0, dir: DOWN}),
		// right top
		countEnergized(field, Beam{row: 0, col: width - 1, dir: LEFT}),
		countEnergized(field, Beam{row: 0, col: width - 1, dir: DOWN}),
		// left bottom
		countEnergized(field, Beam{row: height - 1, col: 0, dir: RIGHT}),
		countEnergized(field, Beam{row: height - 1, col: 0, dir: UP}),
		// right bottom
		countEnergized(field, Beam{row: height - 1, col: width - 1, dir: LEFT}),
		countEnergized(field, Beam{row: height - 1, col: width - 1, dir: UP}),
	}

	for col := 1; col < width-2; col++ {
		candidates = append(candidates, countEnergized(field, Beam{row: 0, col: col, dir: DOWN}))
		candidates = append(candidates, countEnergized(field, Beam{row: height - 1, col: col, dir: UP}))
	}

	for row := 1; row < height-2; row++ {
		candidates = append(candidates, countEnergized(field, Beam{row: row, col: 0, dir: RIGHT}))
		candidates = append(candidates, countEnergized(field, Beam{row: row, col: width - 1, dir: LEFT}))
	}

	best := 0
	for _, candidate := range candidates {
		if candidate > best {
			best = candidate
		}
	}

	return best
}

func main() {
	field := getInputData()

	fmt.Println("Solution 1 is", countEnergized(field, Beam{row: 0, col: 0, dir: RIGHT}))
	fmt.Println("Solution 2 is", findBestConfig(field))
}
