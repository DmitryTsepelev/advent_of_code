package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInputData() ([]string, []string) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := []string{}
	towels := []string{}

	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		if rowIdx == 0 {
			towels = strings.Split(scanner.Text(), ", ")
		} else if rowIdx > 1 {
			patterns = append(patterns, scanner.Text())
		}
	}

	return patterns, towels
}

func canUseTowel(pattern string, towel string, idx int) bool {
	return len(towel) <= len(pattern)-idx && pattern[idx:len(towel)+idx] == towel
}

func dp(pattern string, towels []string, idx int, memo map[int]int) int {
	if idx == len(pattern) {
		return 1
	}

	if _, ok := memo[idx]; !ok {
		for _, towel := range towels {
			if canUseTowel(pattern, towel, idx) {
				memo[idx] += dp(pattern, towels, idx+len(towel), memo)
			}
		}
	}

	return memo[idx]
}

func solve(pattern string, towels []string) int {
	return dp(pattern, towels, 0, make(map[int]int))
}

func main() {
	patterns, towels := getInputData()

	var possbleCount int
	for _, pattern := range patterns {
		if solve(pattern, towels) > 0 {
			possbleCount++
		}
	}
	fmt.Println("Part 1 solution is", possbleCount)

	var variantsSum int
	for _, pattern := range patterns {
		variantsSum += solve(pattern, towels)
	}
	fmt.Println("Part 2 solution is", variantsSum)
}
