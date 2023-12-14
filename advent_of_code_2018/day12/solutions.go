package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func getData() (*[]string, *[][]string) {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	initialState := strings.Split(strings.Replace(scanner.Text(), "initial state: ", "", 1), "")

	patterns := [][]string{}
	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), " => ")

		if len(cmp) > 1 && cmp[1] == "#" {
			pots := strings.Split(cmp[0], "")
			patterns = append(patterns, pots)
		}
	}

	return &initialState, &patterns
}

func extendGeneration(shift int, generation []string) (int, *[]string) {
	result := []string{}
	newShift := shift

	maxLeftPadding := 4
	for i := 0; i < maxLeftPadding; i++ {
		if generation[i] == "#" {
			newShift += maxLeftPadding - i + 1
			for j := 0; j <= maxLeftPadding-i; j++ {
				result = append(result, ".")
			}
			break
		}
	}

	result = append(result, generation...)

	maxRightPadding := 4
	for i := 0; i < maxRightPadding; i++ {
		if generation[len(generation)-i-1] == "#" {
			for j := 0; j < maxRightPadding-i; j++ {
				result = append(result, ".")
			}
		}
	}

	return newShift, &result
}

func calculateGenerations(initialState *[]string, patterns *[][]string, generationCount int) int {
	initialShift, extendedGeneration := extendGeneration(0, *initialState)
	shift := initialShift
	prevGeneration := *extendedGeneration
	currentGeneration := []string{}

	for generation := 0; generation < generationCount; generation++ {
		currentGeneration = append(currentGeneration, ".", ".")

		for middlePot := 2; middlePot <= len(prevGeneration)-3; middlePot++ {
			potSlice := prevGeneration[middlePot-2 : middlePot+3]
			matching := false
			for _, pattern := range *patterns {
				if reflect.DeepEqual(pattern, potSlice) {
					matching = true
					break
				}
			}

			if matching {
				currentGeneration = append(currentGeneration, "#")
			} else {
				currentGeneration = append(currentGeneration, ".")
			}
		}

		newShift, extendedGeneration := extendGeneration(shift, currentGeneration)
		shift = newShift
		prevGeneration = *extendedGeneration

		currentGeneration = []string{}
	}

	total := 0
	for i, pot := range prevGeneration {
		if pot == "#" {
			total += i - shift
		}
	}

	return total
}

func solveTask1(initialState *[]string, patterns *[][]string) int {
	return calculateGenerations(initialState, patterns, 20)
}

func solveTask2(initialState *[]string, patterns *[][]string) int {
	var generation int
	currentDiff := 0
	repeated := 0
	for generation = 1; ; generation += 10 {
		diff := calculateGenerations(initialState, patterns, generation) - calculateGenerations(initialState, patterns, generation-1)

		if diff == currentDiff {
			repeated++
		}

		if repeated > 5 {
			break
		}

		currentDiff = diff
	}

	increment := calculateGenerations(initialState, patterns, generation) - calculateGenerations(initialState, patterns, generation-1)
	return calculateGenerations(initialState, patterns, generation) + (50000000000-(generation))*increment
}

func main() {
	initialState, patterns := getData()
	fmt.Println("Task 1 solution is", solveTask1(initialState, patterns))
	fmt.Println("Task 2 solution is", solveTask2(initialState, patterns))
}
