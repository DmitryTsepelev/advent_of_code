package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() [][]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		cmp := strings.Split(line, " ")
		report := make([]int, len(cmp))

		for idx, sNum := range cmp {
			num, _ := strconv.Atoi(sNum)
			report[idx] = num
		}

		reports = append(reports, report)
	}

	return reports
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func isSafe(report []int) bool {
	isSafe := true

	isIncreasing := report[1] > report[0]

	for i := 0; i < len(report)-1; i++ {
		l, r := report[i], report[i+1]
		diff := r - l

		if diff == 0 {
			isSafe = false
			break
		} else if isIncreasing {
			if diff < 0 || abs(diff) > 3 {
				isSafe = false
				break
			}
		} else if diff > 0 || abs(diff) > 3 {
			isSafe = false
			break
		}
	}

	return isSafe
}

func countSafeReports(reports [][]int) int {
	safeCount := 0

	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	return safeCount
}

func countSafeReportsFixed(reports [][]int) int {
	safeCount := 0

	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		} else {
			for i := 0; i < len(report); i++ {
				fixedReport := []int{}

				for j := 0; j < len(report); j++ {
					if i == j {
						continue
					}
					fixedReport = append(fixedReport, report[j])
				}

				if isSafe(fixedReport) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount
}

func main() {
	reports := getInputData()
	fmt.Println("Part 1 solution is", countSafeReports(reports))
	fmt.Println("Part 2 solution is", countSafeReportsFixed(reports))
}
