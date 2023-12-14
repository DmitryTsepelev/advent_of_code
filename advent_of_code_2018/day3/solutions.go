package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Area struct {
	ID       int
	X        int
	Y        int
	Width    int
	Height   int
	Overlaps bool
}

func getInputData() *[]string {
	data := []string{}

	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return &data
}

func lineToArea(line string) *Area {
	base := strings.Split(line, "@")

	id, _ := strconv.Atoi(strings.TrimSpace(strings.Replace(base[0], "#", "", 1)))

	cmp := strings.Split(base[1], ":")

	coordinates := strings.Split(cmp[0], ",")
	x, _ := strconv.Atoi(strings.TrimSpace(coordinates[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(coordinates[1]))

	size := strings.Split(cmp[1], "x")
	width, _ := strconv.Atoi(strings.TrimSpace(size[0]))
	height, _ := strconv.Atoi(strings.TrimSpace(size[1]))

	return &Area{
		ID:       id,
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
		Overlaps: false,
	}
}

func solveTask1(areas *[]*Area) int {
	field := [1000][1000]int{}

	for _, area := range *areas {
		x := area.X
		y := area.Y

		for i := 0; i < area.Height; i++ {
			for j := 0; j < area.Width; j++ {
				if field[x+j][y+i] == 0 {
					field[x+j][y+i] = area.ID
				} else {
					field[x+j][y+i] = -1
				}
			}
		}
	}

	overlaps := 0

	for i, row := range field {
		for j := range row {
			if field[i][j] == -1 {
				overlaps += 1
			}
		}
	}

	return overlaps
}

func solveTask2(areas *[]*Area) int {
	areasIndex := make(map[int]*Area)

	field := [1000][1000]int{}

	for _, area := range *areas {
		areasIndex[area.ID] = area

		x := area.X
		y := area.Y

		for i := 0; i < area.Height; i++ {
			for j := 0; j < area.Width; j++ {
				areaId := field[x+j][y+i]

				if areaId == 0 {
					field[x+j][y+i] = area.ID
				} else {
					areasIndex[area.ID].Overlaps = true
					areasIndex[areaId].Overlaps = true
				}
			}
		}
	}

	for _, area := range *areas {
		if !area.Overlaps {
			return area.ID
		}
	}

	return 0
}

func main() {
	areas := []*Area{}
	data := *getInputData()

	for _, line := range data {
		areas = append(
			areas,
			lineToArea(line),
		)
	}

	fmt.Println("Task 1 solution is", solveTask1(&areas))
	fmt.Println("Task 2 solution is", solveTask2(&areas))
}
