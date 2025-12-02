package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() [][2]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	cmp := strings.Split(scanner.Text(), ",")

	ranges := [][2]int{}
	for _, c := range cmp {
		spl := strings.Split(c, "-")
		lCmp, rCmp := spl[0], spl[1]

		l, _ := strconv.Atoi(lCmp)
		r, _ := strconv.Atoi(rCmp)
		ranges = append(ranges, [2]int{l, r})
	}

	return ranges
}

func countInvalid(ranges [][2]int, comparator func(string) bool) int64 {
	var sum int64
	for _, rng := range ranges {
		for i := rng[0]; i <= rng[1]; i++ {
			s := strconv.Itoa(i)

			if comparator(s) {
				sum += int64(i)
			}
		}
	}
	return sum
}

func main() {
	ranges := getInputData()

	fmt.Println("Part 1 solution is", countInvalid(ranges, func(s string) bool {
		return len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:]
	}))

	fmt.Println("Part 2 solution is", countInvalid(ranges, func(s string) bool {
		for patternEnd := 1; patternEnd <= len(s)/2; patternEnd++ {
			pattern := s[0:patternEnd]

			candidate := pattern + pattern
			for len(candidate) <= len(s) {
				if s == candidate {
					return true
				}

				candidate += pattern
			}
		}

		return false
	}))
}
