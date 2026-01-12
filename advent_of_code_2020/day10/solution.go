package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getInputData() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	adapters := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		adapter, _ := strconv.Atoi(text)
		adapters = append(adapters, adapter)
	}

	sort.Ints(adapters)

	adapters = append(adapters, adapters[len(adapters)-1]+3)
	adapters = append([]int{0}, adapters...)

	return adapters
}

func part1() int {
	adapters := getInputData()

	diffs := make(map[int]int)
	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		diffs[diff]++
	}

	return diffs[1] * diffs[3]
}

func part2() int {
	adapters := getInputData()

	memo := make(map[int]int)

	var helper func(idx int) int
	helper = func(idx int) int {
		if idx == len(adapters)-1 {
			return 1
		}

		if _, ok := memo[idx]; !ok {
			variants := 0
			for i := idx + 1; i < len(adapters) && adapters[i]-adapters[idx] <= 3; i++ {
				variants += helper(i)
			}

			memo[idx] = variants
		}

		return memo[idx]
	}

	return helper(0)
}

func main() {
	fmt.Println("Part 1 solution is", part1())
	fmt.Println("Part 2 solution is", part2())
}
