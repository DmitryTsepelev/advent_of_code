package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bags = map[string](map[string]int)

func getInputData() Bags {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	bags := Bags{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		bagName := ""

		for !strings.HasPrefix(line, " bags contain") {
			bagName += string(line[0])
			line = line[1:]
		}

		if _, found := bags[bagName]; !found {
			bags[bagName] = make(map[string]int)
		}

		containedLine := line[14 : len(line)-1]

		if containedLine != "no other bags" {
			cmp := strings.Split(containedLine, ", ")

			for _, spec := range cmp {
				bagCount, _ := strconv.Atoi(string(spec[0]))
				bagNameCmp := strings.Split(spec[2:], " ")
				containedBagName := strings.Join(bagNameCmp[:len(bagNameCmp)-1], " ")

				bags[bagName][containedBagName] = bagCount
			}
		}
	}

	return bags
}

const SHINY_GOLD = "shiny gold"

// ------------------------------------------- part 1 -------------------------------------------

func checkBagCanContainShinyGold(currentBag string, bags Bags, memo map[string]bool) bool {
	if _, found := memo[currentBag]; !found {
		if currentBag == SHINY_GOLD {
			memo[currentBag] = true
		} else if len(bags[currentBag]) == 0 {
			memo[currentBag] = false
		} else {
			isContaining := false
			for containedBag := range bags[currentBag] {
				if checkBagCanContainShinyGold(containedBag, bags, memo) {
					memo[currentBag] = true
					isContaining = true
				}
			}

			if !isContaining {
				memo[currentBag] = false
			}
		}
	}

	return memo[currentBag]
}

func countBagsContainingGoldenShiny(bags Bags) int {
	memo := map[string]bool{}
	for bagName := range bags {
		checkBagCanContainShinyGold(bagName, bags, memo)
	}

	count := -1 // exclude golden back itself
	for _, contains := range memo {
		if contains {
			count++
		}
	}

	return count
}

// ------------------------------------------- part 2 -------------------------------------------

func countNestedFrom(currentBag string, bags Bags) int {
	totalCount := 0

	if currentBag != SHINY_GOLD {
		totalCount++ // add itself
	}

	for nestedBagName, bagCount := range bags[currentBag] {
		nestedCount := countNestedFrom(nestedBagName, bags)

		if nestedCount > 0 {
			totalCount += bagCount * nestedCount
		} else {
			totalCount += bagCount
		}
	}

	return totalCount
}

func main() {
	bags := getInputData()

	fmt.Println("Solution 1 is", countBagsContainingGoldenShiny(bags))
	fmt.Println("Solution 2 is", countNestedFrom(SHINY_GOLD, bags))
}
