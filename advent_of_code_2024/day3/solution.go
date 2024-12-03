package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func getInputData() string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	program := ""
	for scanner.Scan() {
		program += scanner.Text()
	}

	return program
}

var KIND_DO = 1
var KIND_DONT = 2
var KIND_VAL = 3

var do_regex = regexp.MustCompile(`do\(\)`)
var dont_regex = regexp.MustCompile(`don't\(\)`)
var mul_regex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func part1(input string) int64 {
	matches := mul_regex.FindAllStringSubmatch(input, -1)

	var sum int64
	for _, match := range matches {
		l, _ := strconv.Atoi(match[1])
		r, _ := strconv.Atoi(match[2])

		sum += int64(l * r)
	}

	return sum
}

type Instruction struct {
	kind int
	pos  int
	val  int
}

func part2(input string) int64 {
	byte_input := []byte(input)

	instructions := []Instruction{}

	do_matches := do_regex.FindAllIndex(byte_input, -1)
	for _, m := range do_matches {
		instructions = append(instructions, Instruction{kind: KIND_DO, pos: m[0]})
	}

	dont_matches := dont_regex.FindAllIndex(byte_input, -1)
	for _, m := range dont_matches {
		instructions = append(instructions, Instruction{kind: KIND_DONT, pos: m[0]})
	}

	matches := mul_regex.FindAllStringSubmatchIndex(input, -1)
	for _, m := range matches {
		l, _ := strconv.Atoi(input[m[2]:m[3]])
		r, _ := strconv.Atoi(input[m[4]:m[5]])
		val := l * r
		instructions = append(instructions, Instruction{kind: KIND_VAL, pos: m[0], val: val})
	}

	sort.Slice(instructions, func(i, j int) bool {
		return instructions[i].pos < instructions[j].pos
	})

	var sum int64
	should_do := true

	for _, instruction := range instructions {
		switch instruction.kind {
		case KIND_DO:
			should_do = true
		case KIND_DONT:
			should_do = false
		case KIND_VAL:
			if should_do {
				sum += int64(instruction.val)
			}
		}
	}

	return sum
}

func main() {
	input := getInputData()

	fmt.Println("Part 1 solution is", part1(input))
	fmt.Println("Part 2 solution is", part2(input))
}
