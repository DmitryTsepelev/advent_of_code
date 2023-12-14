package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToIntList(s string) []int {
	sTrimmed := strings.Trim(s, " ")

	cmp := strings.Split(sTrimmed, " ")
	result := make([]int, 0)
	for _, sNum := range cmp {
		trimmedNum := strings.Trim(sNum, " ")
		if len(trimmedNum) == 0 {
			continue
		}
		num, _ := strconv.Atoi(trimmedNum)
		result = append(result, num)
	}
	return result
}

func getInputData() [][]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sequences := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		sequence := stringToIntList(line)

		sequences = append(sequences, sequence)
	}

	return sequences
}

func buildDiffs(sequence []int) [][]int {
	sequences := [][]int{sequence}

	for {
		currentSequence := sequences[len(sequences)-1]
		nextSequence := make([]int, len(currentSequence)-1)

		allZeroes := true
		for i := 0; i < len(currentSequence)-1; i++ {
			value := currentSequence[i+1] - currentSequence[i]
			if value != 0 {
				allZeroes = false
			}
			nextSequence[i] = value
		}

		sequences = append(sequences, nextSequence)

		if allZeroes {
			break
		}
	}

	return sequences
}

func predictNext(sequence []int) int {
	sequences := buildDiffs(sequence)

	prevSequenceVal := 0
	for sid := len(sequences) - 2; sid >= 0; sid-- {
		current := sequences[sid]
		leftVal := current[len(current)-1]

		prevSequenceVal += leftVal
	}

	return prevSequenceVal
}

func predictPrev(sequence []int) int {
	sequences := buildDiffs(sequence)

	prevSequenceVal := 0
	for sid := len(sequences) - 2; sid >= 0; sid-- {
		current := sequences[sid]
		rightVal := current[0]

		prevSequenceVal = rightVal - prevSequenceVal
	}

	return prevSequenceVal
}

func predictAndSum(sequences [][]int, predict func([]int) int) int {
	sum := 0
	for _, sequence := range sequences {
		sum += predict(sequence)
	}
	return sum
}

func main() {
	sequences := getInputData()
	fmt.Println("Part 1 solution is", predictAndSum(sequences, predictNext))
	fmt.Println("Part 2 solution is", predictAndSum(sequences, predictPrev))
}
