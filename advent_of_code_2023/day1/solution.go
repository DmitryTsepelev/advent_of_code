package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func solve(lines []string, parseWords bool) int {
	sum := 0

	for _, line := range lines {
		digits := []int{}

		idx := 0
		for {
			r := line[idx]
			if r >= '0' && r <= '9' {
				ri, _ := strconv.Atoi(string(r))
				digits = append(digits, ri)
			} else if parseWords {
				for word, digit := range wordToDigit {
					if len(word) <= len(line[idx:]) && strings.HasPrefix(line[idx:], word) {
						digits = append(digits, digit)
						break
					}
				}
			}

			idx++

			if idx == len(line) {
				break
			}
		}

		calibrationValue := 10*digits[0] + digits[len(digits)-1]
		sum += calibrationValue
	}

	return sum
}

func main() {
	lines := getInputData()

	solution1 := solve(lines, false)
	fmt.Println("Part 1 solution is", solution1)

	solution2 := solve(lines, true)
	fmt.Println("Part 2 solution is", solution2)
}
