package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	banks := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		banks = append(banks, text)
	}

	return banks
}

func sMax(x, y string) string {
	xi, _ := strconv.Atoi(x)
	yi, _ := strconv.Atoi(y)

	if xi > yi {
		return x
	}

	return y
}

func helper(bank string, length int, idx int, memo map[[2]int]string) string {
	if idx == len(bank) || length == 0 {
		return ""
	}

	key := [2]int{length, idx}

	if _, ok := memo[key]; !ok {
		take := string(bank[idx]) + helper(bank, length-1, idx+1, memo)
		skip := helper(bank, length, idx+1, memo)
		memo[key] = sMax(take, skip)
	}

	return memo[key]
}

func maxJoltage(bank string, length int) int {
	sMax := helper(bank, length, 0, make(map[[2]int]string))
	max, _ := strconv.Atoi(sMax)
	return max
}

func solve(banks []string, length int) int {
	total := 0
	for _, bank := range banks {
		total += maxJoltage(bank, length)
	}
	return total
}

func main() {
	banks := getInputData()

	fmt.Println("Part 1 solution is", solve(banks, 2))
	fmt.Println("Part 2 solution is", solve(banks, 12))
}
