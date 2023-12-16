package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, line)
	}

	return field
}

func countChars(lines []string, pos int) map[byte]int {
	counts := make(map[byte]int, 0)

	for _, line := range lines {
		c := line[pos]

		if _, ok := counts[c]; ok {
			counts[c]++
		} else {
			counts[c] = 1
		}
	}

	return counts
}

func getMostFrequentChar(lines []string, pos int) rune {
	counts := make(map[byte]int, 0)
	bestCount := 0
	var freq rune

	for _, line := range lines {
		c := line[pos]

		if _, ok := counts[c]; ok {
			counts[c]++
		} else {
			counts[c] = 1
		}

		if counts[c] > bestCount {
			bestCount = counts[c]
			freq = rune(c)
		}
	}

	return freq
}

func getLeastFrequentChar(lines []string, pos int) rune {
	charCounts := countChars(lines, pos)

	counts := make([]int, 0)
	for _, count := range charCounts {
		counts = append(counts, count)
	}
	sort.Ints(counts)

	for char, count := range charCounts {
		if count == counts[0] {
			return rune(char)
		}
	}

	return 'a'
}

func findMessage(lines []string, getChar func([]string, int) rune) string {
	messageLen := len(lines[0])
	message := make([]rune, messageLen)
	for pos := 0; pos < messageLen; pos++ {
		message[pos] = getChar(lines, pos)
	}

	return string(message)
}

func main() {
	lines := getInputData()

	fmt.Println("Solution 1 is", findMessage(lines, getMostFrequentChar))
	fmt.Println("Solution 2 is", findMessage(lines, getLeastFrequentChar))
}
