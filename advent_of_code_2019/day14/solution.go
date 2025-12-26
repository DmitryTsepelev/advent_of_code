package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Recipe struct {
	ingredients   map[string]int
	outcomeAmount int
}

func getInputData() map[string]Recipe {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	recipes := make(map[string]Recipe, 0)
	for scanner.Scan() {
		text := scanner.Text()
		// banks = append(banks, text)

		cmp := strings.Split(text, " => ")

		outcomeCmp := strings.Split(cmp[1], " ")
		amount, _ := strconv.Atoi(outcomeCmp[0])
		outcomeType := outcomeCmp[1]

		recipe := Recipe{ingredients: make(map[string]int), outcomeAmount: amount}

		sources := strings.Split(cmp[0], ", ")
		for _, source := range sources {
			sourceCmp := strings.Split(source, " ")
			amount, _ := strconv.Atoi(sourceCmp[0])
			ingType := sourceCmp[1]

			recipe.ingredients[ingType] = amount
		}

		recipes[outcomeType] = recipe
	}

	return recipes
}

func ore(want map[string]int, reacs map[string]Recipe) int {
loop:
	for {
		for w := range want {
			if w != "ORE" && want[w] > 0 {
				amount := (want[w]-1)/reacs[w].outcomeAmount + 1
				want[w] -= reacs[w].outcomeAmount * amount

				for r := range reacs[w].ingredients {
					want[r] += reacs[w].ingredients[r] * amount
				}
				continue loop
			}
		}
		return want["ORE"]
	}
}

func main() {
	recipes := getInputData()
	fmt.Println(ore(map[string]int{"FUEL": 1}, recipes))
	fmt.Println(sort.Search(1000000000000, func(n int) bool {
		return ore(map[string]int{"FUEL": n}, recipes) > 1000000000000
	}) - 1)
}
