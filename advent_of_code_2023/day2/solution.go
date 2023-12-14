package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id     int
	rounds []Round
}

func getInputData() []Game {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	games := []Game{}
	for scanner.Scan() {
		line := scanner.Text()

		gameCmp := strings.Split(line, ": ")
		id, _ := strconv.Atoi(gameCmp[0][5:])

		roundCmp := strings.Split(gameCmp[1], "; ")
		rounds := make([]Round, len(roundCmp))
		for idx, roundSpec := range roundCmp {
			colorCmp := strings.Split(roundSpec, ", ")

			for _, colorSpec := range colorCmp {
				cmp := strings.Split(colorSpec, " ")
				count, _ := strconv.Atoi(cmp[0])

				switch cmp[1] {
				case "red":
					rounds[idx].red = count
				case "green":
					rounds[idx].green = count
				case "blue":
					rounds[idx].blue = count
				}
			}
		}

		game := Game{
			id:     id,
			rounds: rounds,
		}
		games = append(games, game)
	}

	return games
}

func solve1(games []Game) int {
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	result := 0
	for _, game := range games {
		validGame := true

		for _, round := range game.rounds {
			if round.red > maxRed || round.green > maxGreen || round.blue > maxBlue {
				validGame = false
				break
			}
		}

		if validGame {
			result += game.id
		}
	}

	return result
}

func solve2(games []Game) int {
	powerSum := 0

	for _, game := range games {
		minRed, minGreen, minBlue := 0, 0, 0

		for _, round := range game.rounds {
			if round.red > minRed {
				minRed = round.red
			}
			if round.green > minGreen {
				minGreen = round.green
			}
			if round.blue > minBlue {
				minBlue = round.blue
			}
		}

		power := minRed * minGreen * minBlue
		powerSum += power
	}

	return powerSum
}

func main() {
	games := getInputData()

	solution1 := solve1(games)
	fmt.Println("Part 1 solution is", solution1)

	solution2 := solve2(games)
	fmt.Println("Part 2 solution is", solution2)
}
