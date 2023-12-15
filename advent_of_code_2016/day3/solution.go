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

func parseIntLine(scanner *bufio.Scanner) []int {
	line := scanner.Text()

	ints := make([]int, 0)
	for _, num := range stringToIntList(line) {
		ints = append(ints, num)
	}

	return ints
}

func getInputData() [][]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := make([][]int, 0)
	for scanner.Scan() {
		cmp := parseIntLine(scanner)
		lines = append(lines, cmp)
	}

	return lines
}

func isValid(x, y, z int) bool {
	return x+y > z && x+z > y && y+z > x
}

func countValidTriangles(lines [][]int) int {
	count := 0
	for _, line := range lines {
		if isValid(line[0], line[1], line[2]) {
			count++
		}
	}
	return count
}

func countValidGroups(lines [][]int) int {
	count := 0

	for i := 0; i < len(lines)/3; i++ {
		line1, line2, line3 := lines[i*3], lines[i*3+1], lines[i*3+2]

		for i := 0; i < 3; i++ {
			if isValid(line1[i], line2[i], line3[i]) {
				count++
			}
		}
	}
	return count
}

func main() {
	lines := getInputData()
	fmt.Println("Solution 1 is", countValidTriangles(lines))
	fmt.Println("Solution 2 is", countValidGroups(lines))
}
