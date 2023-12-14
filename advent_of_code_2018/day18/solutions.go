package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getField() *[][]string {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := [][]string{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")

		field = append(field, row)
	}

	return &field
}

const Open = "."
const Tree = "|"
const Lumberyard = "#"

func acresAround(field *[][]string, rowIndex int, acreIndex int) map[string]int {
	around := map[string]int{}

	for rowMod := -1; rowMod <= 1; rowMod++ {
		for acreMod := -1; acreMod <= 1; acreMod++ {
			if rowMod == 0 && acreMod == 0 {
				continue
			}
			rowNumber := rowIndex + rowMod
			if rowNumber < 0 || rowNumber >= len(*field) {
				continue
			}

			row := (*field)[rowNumber]
			acreNumber := acreIndex + acreMod

			if acreNumber < 0 || acreNumber >= len(row) {
				continue
			}

			content := row[acreNumber]
			count, found := around[content]
			if found {
				around[content] = count + 1
			} else {
				around[content] = 1
			}
		}
	}

	return around
}

func newAcre(field *[][]string, rowIndex int, acreIndex int) string {
	around := acresAround(field, rowIndex, acreIndex)

	switch (*field)[rowIndex][acreIndex] {
	case Open:
		if around[Tree] >= 3 {
			return Tree
		}
		return Open
	case Tree:
		if around[Lumberyard] >= 3 {
			return Lumberyard
		}
		return Tree
	case Lumberyard:
		if around[Lumberyard] >= 1 && around[Tree] >= 1 {
			return Lumberyard
		}
		return Open
	}

	return ""
}

func showField(field *[][]string) {
	for _, row := range *field {
		for _, acre := range row {
			fmt.Print(acre)
		}
		fmt.Println()
	}
	fmt.Println()
}

func scoreFor(field *[][]string) int {
	trees := 0
	lumberyards := 0
	for _, row := range *field {
		for _, acre := range row {
			if acre == Tree {
				trees++
			} else if acre == Lumberyard {
				lumberyards++
			}
		}
	}

	return trees * lumberyards
}

func solveTask1() int {
	field := getField()

	for minute := 1; minute <= 10; minute++ {
		newField := [][]string{}

		for rowIndex, row := range *field {
			newRow := []string{}

			for acreIndex := range row {
				newRow = append(newRow, newAcre(field, rowIndex, acreIndex))
			}

			newField = append(newField, newRow)
		}

		field = &newField
	}

	return scoreFor(field)
}

func solveTask2() int {
	field := getField()

	scores := []int{}

	for minute := 1; ; minute++ {
		newField := [][]string{}

		for rowIndex, row := range *field {
			newRow := []string{}

			for acreIndex := range row {
				newRow = append(newRow, newAcre(field, rowIndex, acreIndex))
			}

			newField = append(newField, newRow)
		}

		field = &newField

		newScore := scoreFor(field)

		candidateIndex := -1
		for index, score := range scores {
			if score == newScore {
				candidateIndex = index
				break
			}
		}

		scores = append(scores, newScore)

		period := minute - candidateIndex
		if period > 3 && candidateIndex >= 0 {

			isPattern := true
			for currentIndex := minute - 1; candidateIndex < currentIndex; currentIndex-- {
				prevIndex := currentIndex - period + 1

				if prevIndex < 0 || scores[currentIndex] != scores[prevIndex] {
					isPattern = false
					break
				}
			}

			if !isPattern {
				continue
			}

			periodScores := scores[len(scores)-period-1 : len(scores)-2]
			idx := period - ((1000000000 - minute) % period) - 1
			result := periodScores[idx]
			return result
		}
	}
}

func main() {
	fmt.Println("Task 1 solution is", solveTask1())
	fmt.Println("Task 2 solution is", solveTask2())
}
