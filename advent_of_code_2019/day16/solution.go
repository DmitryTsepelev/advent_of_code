package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	cmp := strings.Split(input, "")
	list := make([]int, len(cmp))
	for idx, c := range cmp {
		num, _ := strconv.Atoi(c)
		list[idx] = num
	}
	return list
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func applyPattern(list []int) []int {
	basePattern := []int{0, 1, 0, -1}

	patterns := make([][]int, len(list))

	for i := 0; i < len(list); i++ {
		patterns[i] = make([]int, 0)

		for _, val := range basePattern {
			for j := 0; j < i+1; j++ {
				patterns[i] = append(patterns[i], val)
			}
		}
	}

	modified := make([]int, len(list))

	for idx := range list {
		pattern := patterns[idx]

		var sum int
		for idx, value := range list {
			patternVal := pattern[(idx+1)%len(pattern)]
			sum += patternVal * value
		}

		modified[idx] = abs(sum % 10)
	}

	return modified
}

func repeat(list []int, times int) []int {
	repeated := make([]int, 0)
	for i := 0; i < times; i++ {
		repeated = append(repeated, list...)
	}
	return repeated
}

func takeN(list []int, n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result = result*10 + list[i]
	}
	return result
}

func part1(input []int) int {
	for i := 0; i < 100; i++ {
		input = applyPattern(input)
	}

	result := 0
	for i := 0; i < 8; i++ {
		result = result*10 + input[i]
	}
	return result
}

func part2(input []int) int {
	offset := takeN(input, 7)
	input = repeat(input, 10000)
	input_length := len(input)

	for i := 0; i < 100; i++ {
		var partial_sum int
		for j := offset; j < input_length; j++ {
			partial_sum += input[j]
		}

		for j := offset; j < input_length; j++ {
			t := partial_sum
			partial_sum -= input[j]
			if t >= 0 {
				input[j] = t % 10
			} else {
				input[j] = (-t) % 10
			}
		}
	}

	result := 0
	for i := offset; i < offset+8; i++ {
		result = result*10 + input[i]
	}
	return result
}

func main() {
	input := getInputData()
	fmt.Println("Part 1 solution is", part1(input))
	fmt.Println("Part 2 solution is", part2(input))
}
