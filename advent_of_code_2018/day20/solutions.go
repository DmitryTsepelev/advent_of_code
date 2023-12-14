package main

import (
	"bufio"
	"fmt"
	"os"
)

func getRegex() string {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	regex := scanner.Text()

	return regex
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Room struct {
	X int
	Y int
}

func main() {
	regex := getRegex()

	positions := []Room{}
	currentRoom := Room{X: 0, Y: 0}
	prevRoom := Room{currentRoom.X, currentRoom.Y}
	connected := map[Room][]Room{}
	distances := map[Room]int{}

	for i := 1; i < len(regex)-1; i++ {
		char := regex[i : i+1]

		if char == "(" {
			positions = append(positions, Room{X: currentRoom.X, Y: currentRoom.Y})
		} else if char == ")" {
			currentRoom, positions = positions[len(positions)-1], positions[:len(positions)-1]
		} else if char == "|" {
			currentRoom = positions[len(positions)-1]
		} else {
			switch char {
			case "N":
				currentRoom.Y--
			case "E":
				currentRoom.X++
			case "S":
				currentRoom.Y++
			case "W":
				currentRoom.X--
			}

			contains := false
			for _, el := range connected[currentRoom] {
				if el == prevRoom {
					contains = true
					break
				}
			}
			if !contains {
				connected[currentRoom] = append(connected[currentRoom], prevRoom)
			}

			if distances[currentRoom] != 0 {
				distances[currentRoom] = min(distances[currentRoom], distances[prevRoom]+1)
			} else {
				distances[currentRoom] = distances[prevRoom] + 1
			}
		}

		prevRoom = currentRoom
	}

	maxDistance := 0
	for _, distance := range distances {
		if distance > maxDistance {
			maxDistance = distance
		}
	}
	fmt.Println("Task 1 solution is", maxDistance)

	roomCount := 0
	for _, distance := range distances {
		if distance >= 1000 {
			roomCount++
		}
	}
	fmt.Println("Task 2 solution is", roomCount)
}
