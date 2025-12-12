package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Present = [3][3]bool

type Requirement struct {
	height, width int
	presentCounts map[int]int
}

func getInputData() (map[int]Present, []Requirement) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	presents := map[int]Present{}
	requirements := []Requirement{}

	for scanner.Scan() {
		line := scanner.Text()

		if line[len(line)-1] == ':' {
			cmp := strings.Split(line, ":")
			presentId, _ := strconv.Atoi(cmp[0])

			present := [3][3]bool{}

			for i := 0; i < 3; i++ {
				scanner.Scan()
				line := scanner.Text()

				for j := 0; j < 3; j++ {
					present[i][j] = line[j] == '#'
				}
			}

			presents[presentId] = present
			scanner.Scan()
		} else {
			cmp := strings.Split(line, ": ")

			dimCmp := strings.Split(cmp[0], "x")
			width, _ := strconv.Atoi(dimCmp[0])
			height, _ := strconv.Atoi(dimCmp[1])

			reqCmp := strings.Split(cmp[1], " ")

			presentCounts := map[int]int{}
			for idx, sCount := range reqCmp {
				count, _ := strconv.Atoi(sCount)
				presentCounts[idx] = count
			}

			requirements = append(requirements, Requirement{
				height:        height,
				width:         width,
				presentCounts: presentCounts,
			})
		}
	}

	return presents, requirements
}

func canFit(presents map[int]Present, requirements map[int]int, height, width int) bool {
	expectedArea := 0

	for _, count := range requirements {
		expectedArea += count * 9
	}

	return height*width >= expectedArea
}

func main() {
	presents, requirements := getInputData()

	ans := 0
	for _, r := range requirements {
		if canFit(presents, r.presentCounts, r.height, r.width) {
			ans++
		}
	}
	fmt.Println("solution", ans)
}
