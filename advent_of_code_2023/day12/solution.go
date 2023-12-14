package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	line  string
	sizes []int
}

func getInputData() []Row {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := make([]Row, 0)

	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), " ")

		row := Row{
			line:  cmp[0],
			sizes: make([]int, 0),
		}

		for _, c := range strings.Split(cmp[1], ",") {
			size, _ := strconv.Atoi(string(c))
			row.sizes = append(row.sizes, size)
		}

		field = append(field, row)
	}

	return field
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func handleDamagedSection(line string, sizes []int, memo [][]int) int {
	sectionSize := sizes[0]
	if len(line) < sectionSize {
		return 0 // line shorter than section
	}

	for i := 0; i < sectionSize; i++ {
		if line[i] == '.' {
			return 0 // cannot build section with proper length
		}
	}

	if len(line) > sectionSize && line[sectionSize] == '#' {
		return 0 // not possible to leave a space between sections
	}

	return countVariants(line[min(len(line), sectionSize+1):], sizes[1:], memo)
}

func countVariants(line string, sizes []int, memo [][]int) int {
	// fmt.Println(len(line), len(sizes))
	// fmt.Println(memo)
	if memo[len(line)][len(sizes)] != UNDEFINED {
		return memo[len(line)][len(sizes)]
	}

	// skip initial operational
	for len(line) > 0 && line[0] == '.' {
		line = line[1:]
	}

	if len(sizes) == 0 {
		// skip initial ?
		for len(line) > 0 && (line[0] == '?' || line[0] == '.') {
			line = line[1:]
		}

		if len(line) == 0 {
			return 1 // string and section are at the end
		} else {
			return 0 // we still have unused #
		}
	}

	if len(line) == 0 {
		if len(sizes) == 0 {
			return 1 // string and section are at the end
		} else {
			return 0 // we have unhandled sections
		}
	}

	// assuming this is # or ? counted as #
	variants := handleDamagedSection(line, sizes, memo)

	if line[0] == '?' {
		// assuming this is .
		variants += countVariants(line[1:], sizes, memo)
	}

	memo[len(line)][len(sizes)] = variants

	return variants
}

const UNDEFINED = -1

func sumVariants(field []Row) int {
	count := 0
	for _, row := range field {
		memo := make([][]int, len(row.line)+1)

		for i := 0; i <= len(row.line); i++ {
			memo[i] = make([]int, len(row.sizes)+1)

			for j := 0; j <= len(row.sizes); j++ {
				memo[i][j] = UNDEFINED
			}
		}
		count += countVariants(row.line, row.sizes, memo)
	}
	return count
}

func expandField(field []Row) []Row {
	expandedField := make([]Row, len(field))

	for idx, row := range field {
		line := ""
		sizes := make([]int, 0)
		for i := 1; i <= 5; i++ {
			line += row.line
			if i < 5 {
				line += "?"
			}
			sizes = append(sizes, row.sizes...)
		}

		expandedField[idx] = Row{
			line:  line,
			sizes: sizes,
		}
	}
	return expandedField
}

func main() {
	field := getInputData()

	fmt.Println("Part 1 solution is", sumVariants(field))
	fmt.Println("Part 2 solution is", sumVariants(expandField(field)))
}
