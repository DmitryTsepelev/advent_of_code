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
	scanner.Scan()
	line := scanner.Text()

	cmp := strings.Split(line, ": ")

	ints := make([]int, 0)
	for _, num := range stringToIntList(cmp[1]) {
		ints = append(ints, num)
	}

	return ints
}

func getInputData() ([]int, []int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	return parseIntLine(scanner), parseIntLine(scanner)
}

func parseIntFullLine(scanner *bufio.Scanner) int {
	scanner.Scan()
	line := scanner.Text()

	cmp := strings.Split(line, ": ")

	trimmed := strings.ReplaceAll(cmp[1], " ", "")
	trimmedNum, _ := strconv.Atoi(trimmed)
	return trimmedNum
}

func getInputData2() (int, int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	return parseIntFullLine(scanner), parseIntFullLine(scanner)
}

func findBeats(time, record int) int {
	beats := 0
	for holdT := 1; holdT <= time; holdT++ {
		distance := holdT * (time - holdT)

		if distance > record {
			beats++
		}
	}
	return beats
}

func solve1(times []int, distances []int) int {
	result := 1

	for i := 0; i < len(distances); i++ {
		result *= findBeats(times[i], distances[i])
	}

	return result
}

func main() {
	times, distances := getInputData()
	fmt.Println("Part 1 solution is", solve1(times, distances))

	time, distance := getInputData2()
	fmt.Println("Part 2 solution is", findBeats(time, distance))
}
