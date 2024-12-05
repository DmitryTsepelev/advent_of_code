package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rules = map[int]([]int)

func getInputData() (Rules, [][]int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := Rules{}
	for scanner.Scan() && len(scanner.Text()) > 0 {
		cmp := strings.Split(scanner.Text(), "|")
		l, _ := strconv.Atoi(cmp[0])
		r, _ := strconv.Atoi(cmp[1])

		if _, ok := rules[r]; !ok {
			rules[r] = []int{}
		}

		rules[r] = append(rules[r], l)
	}

	manuals := [][]int{}
	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), ",")

		manual := []int{}

		for _, c := range cmp {
			page, _ := strconv.Atoi(c)
			manual = append(manual, page)
		}

		manuals = append(manuals, manual)
	}

	return rules, manuals
}

func prepareApplicableRules(rules Rules, manual []int) Rules {
	pageMap := map[int]bool{}
	for _, page := range manual {
		pageMap[page] = true
	}

	applicableRules := Rules{}
	for from, to := range rules {
		if _, ok := pageMap[from]; !ok {
			continue
		}

		applicableRules[from] = []int{}
		for _, page := range to {
			if _, ok := pageMap[page]; !ok {
				continue
			}

			applicableRules[from] = append(applicableRules[from], page)
		}
	}

	return applicableRules
}

func reorderManual(rules Rules, manual []int) []int {
	inDegree := make(map[int]int)
	for _, page := range manual {
		if _, ok := inDegree[page]; !ok {
			inDegree[page] = 0
		}
	}
	for _, to := range rules {
		for _, page := range to {
			inDegree[page]++
		}
	}

	fixed := make([]int, len(manual))
	for i := 0; i < len(manual); i++ {
		for page, degree := range inDegree {
			if degree == 0 {
				delete(inDegree, page)
				fixed[i] = page

				for _, blockedPage := range rules[page] {
					inDegree[blockedPage]--
				}
				break
			}
		}
	}
	return fixed
}

func findMiddlesOfCorrectManuals(rules Rules, manuals [][]int, fixOrder bool) int {
	sumOfMiddles := 0

	for _, manual := range manuals {
		applicableRules := prepareApplicableRules(rules, manual)

		validManual := true
		for idx, page := range manual {
			prereq := map[int]bool{}
			for _, reqPage := range applicableRules[page] {
				prereq[reqPage] = false
			}

			for prevIdx := 0; prevIdx < idx; prevIdx++ {
				printedPage := manual[prevIdx]

				if _, ok := prereq[printedPage]; ok {
					prereq[printedPage] = true
				}
			}

			allPrerequisitesMet := true
			for _, isMet := range prereq {
				if isMet == false {
					allPrerequisitesMet = false
					break
				}
			}

			if allPrerequisitesMet == false {
				validManual = false
				break
			}
		}

		if fixOrder {
			if validManual == false {
				fixed := reorderManual(applicableRules, manual)
				sumOfMiddles += fixed[len(fixed)/2]
			}
		} else if validManual {
			sumOfMiddles += manual[len(manual)/2]
		}
	}

	return sumOfMiddles
}

func main() {
	rules, manuals := getInputData()

	fmt.Println("Part 1 solution is", findMiddlesOfCorrectManuals(rules, manuals, false))
	fmt.Println("Part 2 solution is", findMiddlesOfCorrectManuals(rules, manuals, true))
}
