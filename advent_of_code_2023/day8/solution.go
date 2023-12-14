package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mapping struct {
	l string
	r string
}

type Transitions = map[string]Mapping

func getInputData() (string, Transitions) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	path := scanner.Text()
	scanner.Scan()

	transitions := make(Transitions, 0)
	for scanner.Scan() {
		line := scanner.Text()

		cmp := strings.Split(line, " = ")

		lr := strings.Split(cmp[1], ", ")

		transitions[cmp[0]] = Mapping{l: lr[0][1:], r: lr[1][:3]}
	}

	return path, transitions
}

func findPath(path string, transitions Transitions, startNode string, isFinal func(string) bool) int {
	step := 0
	current := startNode

	for {
		if isFinal(current) {
			break
		}

		instruction := path[step%(len(path))]

		if instruction == 'L' {
			current = transitions[current].l
		} else {
			current = transitions[current].r
		}

		step++
	}

	return step
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func findLoops(path string, transitions Transitions) int {
	currentNodes := []string{}
	for k := range transitions {
		if k[2] == 'A' {
			currentNodes = append(currentNodes, k)
		}
	}

	lengths := make([]int, len(currentNodes))

	for idx, startNode := range currentNodes {
		step := findPath(path, transitions, startNode, func(s string) bool {
			return s[2] == 'Z'
		})

		lengths[idx] = step
	}

	return LCM(lengths[0], lengths[1], lengths[2:]...)
}

func main() {
	path, transitions := getInputData()

	solution1 := findPath(path, transitions, "AAA", func(s string) bool {
		return s == "ZZZ"
	})

	fmt.Println("Part 1 solution is", solution1)
	fmt.Println("Part 2 solution is", findLoops(path, transitions))
}
