package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func getInputData() []Hand {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()

		cmp := strings.Split(line, " ")
		bid, _ := strconv.Atoi(cmp[1])

		hand := Hand{
			cards: cmp[0],
			bid:   bid,
		}

		hands = append(hands, hand)
	}

	return hands
}

const (
	FIVE_OF_A_KIND  int = 7
	FOUR_OF_A_KIND      = 6
	FULL_HOUSE          = 5
	THREE_OF_A_KIND     = 4
	TWO_PAIR            = 3
	ONE_PAIR            = 2
	HIGH_CARD           = 1
)

func strenghOf(cards string) int {
	counts := make(map[rune]int, 0)

	for _, card := range cards {
		if _, ok := counts[card]; ok {
			counts[card]++
		} else {
			counts[card] = 1
		}
	}

	for _, v := range counts {
		if v == 5 {
			return FIVE_OF_A_KIND
		}
		if v == 4 {
			return FOUR_OF_A_KIND
		}
	}
	if len(counts) == 2 {
		return FULL_HOUSE
	}
	for _, v := range counts {
		if v == 3 {
			return THREE_OF_A_KIND
		}
	}

	c := 0
	for _, v := range counts {
		if v == 2 {
			c++
		}
	}
	if c == 2 {
		return TWO_PAIR
	} else if c == 1 {
		return ONE_PAIR
	}

	return HIGH_CARD
}

func strenghOfWithJoker(cards string) int {
	counts := make(map[rune]int, 0)

	for _, card := range cards {
		if _, ok := counts[card]; ok {
			counts[card]++
		} else {
			counts[card] = 1
		}
	}

	if _, ok := counts['J']; !ok {
		return strenghOf(cards)
	}

	var maxCount int
	var maxRune rune

	for r, c := range counts {
		if r == 'J' {
			continue
		}

		if maxCount < c {
			maxCount = c
			maxRune = r
		}
	}

	replaced := strings.ReplaceAll(cards, "J", string(maxRune))
	return strenghOf(replaced)
}

func cardToIdx(card rune, cardOrder []rune) int {
	for i := 0; i < len(cardOrder); i++ {
		if cardOrder[i] == card {
			return i
		}
	}

	return 0
}

func compareCards(cardOrder []rune) func(string, string) bool {
	return func(cards1, cards2 string) bool {
		for idx, card1 := range cards1 {
			card2 := rune(cards2[idx])

			if card1 != card2 {
				return cardToIdx(card2, cardOrder) < cardToIdx(card1, cardOrder)
			}
		}

		return true
	}
}

func solve(hands []Hand, strenghOf func(string) int, compareCards func(string, string) bool) int {
	sort.Slice(hands, func(i, j int) bool {
		iCards := hands[i].cards
		jCards := hands[j].cards

		si := strenghOf(iCards)
		sj := strenghOf(jCards)

		return sj == si && compareCards(iCards, jCards) || si < sj
	})

	amount := 0
	for idx, hand := range hands {
		amount += hand.bid * (idx + 1)
	}

	return amount
}

func main() {
	hands := getInputData()

	solution1 := solve(hands, strenghOf, compareCards([]rune{
		'A',
		'K',
		'Q',
		'J',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	}))
	fmt.Println("Part 1 solution is", solution1)

	solution2 := solve(hands, strenghOfWithJoker, compareCards([]rune{
		'A',
		'K',
		'Q',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
		'J',
	}))
	fmt.Println("Part 2 solution is", solution2)
}
