package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInputData() ([]int, []int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left, right := make([]int, 0), make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		cmp := strings.Split(line, "   ")

		lNum, _ := strconv.Atoi(cmp[0])
		left = append(left, lNum)
		rNum, _ := strconv.Atoi(cmp[1])
		right = append(right, rNum)
	}

	return left, right
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func findDiff(left, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	totalDiff := 0
	for idx, lNum := range left {
		rNum := right[idx]
		totalDiff += abs(rNum - lNum)
	}

	return totalDiff
}

func findSimilarity(left, right []int) int {
	fq := make(map[int]int)

	for _, num := range right {
		if _, ok := fq[num]; !ok {
			fq[num] = 0
		}

		fq[num]++
	}

	similarity := 0
	for _, num := range left {
		similarity += num * fq[num]
	}
	return similarity
}

func main() {
	left, right := getInputData()

	fmt.Println("Part 1 solution is", findDiff(left, right))
	fmt.Println("Part 2 solution is", findSimilarity(left, right))
}
