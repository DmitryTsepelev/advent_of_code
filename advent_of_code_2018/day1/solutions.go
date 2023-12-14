package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() *[]int {
	data := []int{}

	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intValue, _ := strconv.Atoi(scanner.Text())
		data = append(data, intValue)
	}

	return &data
}

func solveTask1() int {
	total := 0
	data := *getInputData()

	for _, item := range data {
		total += item
	}

	return total
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func solveTask2() int {
	total := 0
	data := *getInputData()
	frequencies := []int{0}

	for {
		for _, item := range data {
			total += item
			if contains(frequencies, total) {
				return total
			}
			frequencies = append(frequencies, total)
		}
	}
}

func main() {
	fmt.Println("Task 1 solution is", solveTask1())
	fmt.Println("Task 2 solution is", solveTask2())
}
