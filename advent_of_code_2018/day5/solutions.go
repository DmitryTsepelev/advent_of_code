package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput() []byte {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Bytes()
}

func react(polymer *[]byte) *[]byte {
	reacted := make([]byte, len(*polymer))
	copy(reacted, *polymer)

	i := 0

	for {
		if i > len(reacted)-2 {
			break
		}

		currentChar := string(reacted[i])
		nextChar := string(reacted[i+1])

		if currentChar != nextChar && strings.ToUpper(nextChar) == strings.ToUpper(currentChar) {
			copy(reacted[i:], reacted[i+2:])
			reacted = reacted[:len(reacted)-2]

			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}

	return &reacted
}

func solveTask1() int {
	polymer := getInput()
	reacted := react(&polymer)
	return len(*reacted)
}

func removeUnit(polymer *[]byte, unit byte) *[]byte {
	withoutUnit := []byte{}

	for i := 0; i < len(*polymer); i++ {
		currentByte := (*polymer)[i]

		if currentByte != unit && currentByte != unit-32 {
			withoutUnit = append(withoutUnit, currentByte)
		}
	}

	return &withoutUnit
}

func solveTask2() int {
	polymer := getInput()

	minLength := len(polymer)

	for unit := 'a'; unit <= 'z'; unit++ {
		withoutUnit := removeUnit(&polymer, byte(unit))
		reacted := react(withoutUnit)
		length := len(*reacted)
		if length < minLength {
			minLength = length
		}
	}

	return minLength
}

func main() {
	fmt.Println("Task 1 solution is", solveTask1())
	fmt.Println("Task 2 solution is", solveTask2())
}
