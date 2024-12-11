package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumCounts = map[int]int

func getInputData() NumCounts {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numCounts := make(NumCounts)

	for scanner.Scan() {
		for _, sNum := range strings.Split(scanner.Text(), " ") {
			num, _ := strconv.Atoi(sNum)

			if _, ok := numCounts[num]; !ok {
				numCounts[num] = 0
			}
			numCounts[num]++
		}
	}

	return numCounts
}

func simulate(numCounts NumCounts, steps int) NumCounts {
	for i := 0; i < steps; i++ {
		nextNumCounts := make(NumCounts)

		for num, count := range numCounts {
			if num == 0 {
				nextNumCounts[1] += count
			} else {
				sVal := strconv.Itoa(num)

				if len(sVal)%2 == 0 {
					mid := len(sVal) / 2
					lVal, _ := strconv.Atoi(sVal[:mid])
					rVal, _ := strconv.Atoi(sVal[mid:])

					nextNumCounts[lVal] += count
					nextNumCounts[rVal] += count
				} else {
					nextNumCounts[num*2024] += count
				}
			}
		}

		numCounts = nextNumCounts
	}

	return numCounts
}

func countNums(numCounts NumCounts) int {
	var result int

	for _, count := range numCounts {
		result += count
	}

	return result
}

func main() {
	numMap := getInputData()
	numMap = simulate(numMap, 25)
	fmt.Println("Part 1 solution is", countNums(numMap))

	numMap = getInputData()
	numMap = simulate(numMap, 75)
	fmt.Println("Part 2 solution is", countNums(numMap))
}
