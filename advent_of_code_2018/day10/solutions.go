package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Star struct {
	X  int
	Y  int
	DX int
	DY int
}

func getData() *[]*Star {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := []*Star{}

	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), "> velocity=<")

		coorditatesString := strings.Replace(cmp[0], "position=<", "", 1)
		coordinates := strings.Split(coorditatesString, ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coordinates[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coordinates[1]))

		velocityString := strings.Replace(cmp[1], ">", "", 1)
		velocity := strings.Split(velocityString, ",")
		dx, _ := strconv.Atoi(strings.TrimSpace(velocity[0]))
		dy, _ := strconv.Atoi(strings.TrimSpace(velocity[1]))

		star := Star{X: x, Y: y, DX: dx, DY: dy}
		data = append(data, &star)
	}

	return &data
}

func draw(stars *[]*Star) {
	minX := math.MaxInt32
	maxX := math.MinInt32
	minY := math.MaxInt32
	maxY := math.MinInt32

	for _, star := range *stars {
		if star.X < minX {
			minX = star.X
		}

		if star.X > maxX {
			maxX = star.X
		}

		if star.Y < minY {
			minY = star.Y
		}

		if star.Y > maxY {
			maxY = star.Y
		}
	}

	for y := minY - 5; y < maxY+5; y++ {
		line := ""

		shouldBePrinted := false

		for x := minX - 5; x < maxX+5; x++ {
			hasStar := false

			for _, star := range *stars {
				if star.X == x && star.Y == y {
					line += "#"
					hasStar = true
					shouldBePrinted = true
					break
				}
			}

			if !hasStar {
				line += "."
			}
		}

		if shouldBePrinted {
			fmt.Println(line)
		}
	}
}

func move(stars *[]*Star) {
	for _, star := range *stars {
		star.X += star.DX
		star.Y += star.DY
	}
}

func solveTask() (result *[]*Star, seconds int) {
	stars := getData()

	for i := 0; i < 1000000; i++ {
		move(stars)

		minX := 0
		maxX := 0
		minY := 0
		maxY := 0

		for _, star := range *stars {
			if star.X < minX {
				minX = star.X
			}

			if star.X > maxX {
				maxX = star.X
			}

			if star.Y < minY {
				minY = star.Y
			}

			if star.Y > maxY {
				maxY = star.Y
			}
		}

		if maxX-minX < 238 {
			return stars, i + 1
		}
	}

	return nil, 0
}

func main() {
	stars, seconds := solveTask()
	draw(stars)
	fmt.Println("Seconds", seconds)
}
