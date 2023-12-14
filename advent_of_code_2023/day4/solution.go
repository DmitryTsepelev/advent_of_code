package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	winning map[int]bool
	have    []int
}

func stringToIntList(s string) []int {
	sTrimmed := strings.Trim(s, " ")

	cmp := strings.Split(sTrimmed, " ")
	result := make([]int, 0)
	for _, sNum := range cmp {
		trimmedNum := strings.Trim(sNum, " ")
		if len(trimmedNum) == 0 {
			continue
		}
		num, _ := strconv.Atoi(trimmedNum)
		result = append(result, num)
	}
	return result
}

func getInputData() []Card {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := []Card{}
	for scanner.Scan() {
		line := scanner.Text()
		cmp := strings.Split(line, ": ")
		cmp = strings.Split(cmp[1], " | ")

		winning := make(map[int]bool)

		for _, num := range stringToIntList(cmp[0]) {
			if _, ok := winning[num]; !ok {
				winning[num] = true
			}
		}

		card := Card{
			winning: winning,
			have:    stringToIntList(cmp[1]),
		}

		cards = append(cards, card)
	}

	return cards
}

func pow2(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			result = 1
		} else {
			result *= 2
		}
	}

	return result
}

func solve1(cards []Card) int {
	sum := 0

	for _, card := range cards {
		pow := 0
		for _, candidate := range card.have {
			if _, ok := card.winning[candidate]; ok {
				pow++
			}
		}

		sum += pow2(pow)
	}

	return sum
}

func solve2(cards []Card) int {
	wonCards := make(map[int]([]int), 0) // which cards we win for this card
	counts := make(map[int]int, 0)

	for idx, card := range cards {
		counts[idx] = 1 // we have one card of each type initially

		matchCount := 0
		for _, candidate := range card.have {
			if _, ok := card.winning[candidate]; ok {
				matchCount++
			}
		}

		wins := make([]int, 0)
		for i := idx + 1; i <= idx+matchCount; i++ {
			wins = append(wins, i)
		}

		wonCards[idx] = wins
	}

	// for each initial card we add won cards to counts
	for i := 0; i < len(cards); i++ {
		for _, wonIdx := range wonCards[i] {
			counts[wonIdx] += counts[i]
		}
	}

	totalCount := 0
	for _, count := range counts {
		totalCount += count
	}

	return totalCount
}

func main() {
	cards := getInputData()

	solution1 := solve1(cards)
	fmt.Println("Part 1 solution is", solution1)

	solution2 := solve2(cards)
	fmt.Println("Part 2 solution is", solution2)
}
