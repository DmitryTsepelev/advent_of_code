package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	args      []int
	operation byte
}

func getTasksPart1() map[int]*Task {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tasks := map[int]*Task{}
	for scanner.Scan() {
		components := strings.Fields(scanner.Text())
		for colIdx, component := range components {
			if _, ok := tasks[colIdx]; !ok {
				tasks[colIdx] = &Task{args: []int{}}
			}

			if component == "+" || component == "*" {
				tasks[colIdx].operation = component[0]
			} else {
				val, _ := strconv.Atoi(component)
				tasks[colIdx].args = append(tasks[colIdx].args, val)
			}
		}
	}

	return tasks
}

func getTasksPart2() map[int]*Task {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := []string{}
	width, height := 0, 0
	for scanner.Scan() {
		line := scanner.Text()
		height++
		width = max(width, len(line))
		matrix = append(matrix, line)
	}

	taskIdx := 0
	tasks := map[int]*Task{}

	for colIdx := 0; colIdx < width; colIdx++ {
		allSpaces := true
		sArg := []byte{}

		if _, ok := tasks[taskIdx]; !ok {
			tasks[taskIdx] = &Task{args: []int{}}
		}

		for rowIdx := 0; rowIdx < height; rowIdx++ {
			if colIdx >= len(matrix[rowIdx]) {
				continue
			}
			val := matrix[rowIdx][colIdx]
			if val == ' ' {
				continue
			}

			allSpaces = false

			if val == '*' || val == '+' {
				tasks[taskIdx].operation = val
			} else {
				sArg = append(sArg, val)
			}
		}

		if allSpaces {
			taskIdx++
		} else {
			arg, _ := strconv.Atoi(string(sArg))
			tasks[taskIdx].args = append(tasks[taskIdx].args, arg)
		}
	}

	return tasks
}

func solve(tasks map[int]*Task) int {
	sum := 0
	for _, task := range tasks {
		acc := task.args[0]
		for i := 1; i < len(task.args); i++ {
			if task.operation == '+' {
				acc += task.args[i]
			} else {
				acc *= task.args[i]
			}
		}

		sum += acc
	}
	return sum
}

func main() {
	fmt.Println("Part 1 solution is", solve(getTasksPart1()))
	fmt.Println("Part 2 solution is", solve(getTasksPart2()))
}
