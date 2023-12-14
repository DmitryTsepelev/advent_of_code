package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(commands []string) int {
	up := 0
	right := 0
	direction := 'U'

	rotationsL := map[rune]rune{
		'U': 'L',
		'L': 'D',
		'D': 'R',
		'R': 'U',
	}

	rotationsR := map[rune]rune{
		'U': 'R',
		'R': 'D',
		'D': 'L',
		'L': 'U',
	}

	for _, cmd := range commands {
		cmd = strings.TrimSuffix(cmd, "\n")
		rotation := rune(cmd[0])

		if rotation == 'L' {
			direction = rotationsL[direction]
		} else {
			direction = rotationsR[direction]
		}

		length, _ := strconv.Atoi(string(cmd[1:]))

		switch direction {
		case 'U':
			up += length
		case 'D':
			up -= length
		case 'L':
			right -= length
		case 'R':
			right += length
		}
	}

	if up < 0 {
		up = -up
	}

	if right < 0 {
		right = -right
	}

	return up + right
}

type Point struct {
	up    int
	right int
}

func part2(commands []string) int {
	up := 0
	right := 0
	direction := 'U'

	rotationsL := map[rune]rune{
		'U': 'L',
		'L': 'D',
		'D': 'R',
		'R': 'U',
	}

	rotationsR := map[rune]rune{
		'U': 'R',
		'R': 'D',
		'D': 'L',
		'L': 'U',
	}

	locations := make([]Point, 0)

	for _, cmd := range commands {
		cmd = strings.TrimSuffix(cmd, "\n")

		// locations = append(locations, Point{up: up, right: right})

		rotation := rune(cmd[0])

		if rotation == 'L' {
			direction = rotationsL[direction]
		} else {
			direction = rotationsR[direction]
		}

		length, _ := strconv.Atoi(string(cmd[1:]))

		switch direction {
		case 'U':
			for i := 0; i < length; i++ {
				newPoint := Point{up: up + i, right: right}
				locations = append(locations, newPoint)
				fmt.Println(up, right)
				for _, point := range locations {
					if newPoint.up == up && point.right == right {
						fmt.Println("!!!", point)
						if up < 0 {
							up = -up
						}

						if right < 0 {
							right = -right
						}

						return up + right
					}
				}
			}
			up += length
		case 'D':
			for i := 0; i < length; i++ {
				newPoint := Point{up: up - i, right: right}
				locations = append(locations, newPoint)
				fmt.Println(up, right)
				for _, point := range locations {
					if newPoint.up == up && point.right == right {
						fmt.Println("!!!", point)
						if up < 0 {
							up = -up
						}

						if right < 0 {
							right = -right
						}

						return up + right
					}
				}
			}
			up -= length
		case 'L':
			for i := 0; i < length; i++ {
				newPoint := Point{up: up, right: right - i}
				locations = append(locations, newPoint)
				fmt.Println(up, right)
				for _, point := range locations {
					if newPoint.up == up && point.right == right {
						fmt.Println("!!!", point)
						if up < 0 {
							up = -up
						}

						if right < 0 {
							right = -right
						}

						return up + right
					}
				}
			}
			right -= length
		case 'R':
			for i := 0; i < length; i++ {
				newPoint := Point{up: up, right: right + i}
				locations = append(locations, newPoint)
				fmt.Println(up, right)
				for _, point := range locations {
					if newPoint.up == up && newPoint.right == right {
						fmt.Println("!!!", point)
						if up < 0 {
							up = -up
						}

						if right < 0 {
							right = -right
						}

						return up + right
					}
				}
			}
			right += length
		}

		fmt.Println(locations)
	}

	return 0
}

func main() {
	dat, _ := os.ReadFile("input.txt")
	commands := strings.Split(string(dat), ", ")

	fmt.Println(part1(commands))
	fmt.Println(part2(commands))
}
