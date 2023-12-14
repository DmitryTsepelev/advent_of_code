package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Recipe struct {
	Value int
	Next  *Recipe
	Prev  *Recipe
}

type Scoreboard struct {
	First       *Recipe
	Last        *Recipe
	RecipeCount int
}

func (scoreboard *Scoreboard) Move(recipe *Recipe) *Recipe {
	result := recipe
	for i := 0; i < recipe.Value+1; i++ {
		result = result.Next
		if result == nil {
			result = scoreboard.First
		}
	}
	return result
}

func (scoreboard *Scoreboard) AddRecipe(recipe *Recipe) {
	scoreboard.Last.Next = recipe
	recipe.Prev = scoreboard.Last
	scoreboard.Last = recipe
	scoreboard.RecipeCount++
}

func (scoreboard *Scoreboard) LastSequence(length int) string {
	recipe := scoreboard.Last

	result := ""
	for i := 0; i < length; i++ {
		if recipe == nil {
			break
		}
		result = strconv.Itoa(recipe.Value) + result
		recipe = recipe.Prev
	}

	return result
}

func initScoreboard() *Scoreboard {
	SecondRecipe := Recipe{Value: 7}
	FirstRecipe := Recipe{Value: 3, Next: &SecondRecipe}
	SecondRecipe.Prev = &FirstRecipe
	return &Scoreboard{First: &FirstRecipe, Last: &SecondRecipe, RecipeCount: 2}
}

func makeNewRecipes(firstElf *Recipe, secondElf *Recipe) *[]string {
	values := strings.Split(strconv.Itoa(firstElf.Value+secondElf.Value), "")
	return &values
}

func solveTask1(input int) string {
	scoreboard := initScoreboard()

	firstElf := scoreboard.First
	secondElf := scoreboard.Last

	resultRecipes := []string{}

	for i := 0; ; i++ {
		values := makeNewRecipes(firstElf, secondElf)

		for _, value := range *values {
			stringValue, _ := strconv.Atoi(value)
			scoreboard.AddRecipe(&Recipe{Value: stringValue})

			if scoreboard.RecipeCount > input {
				resultRecipes = append(resultRecipes, value)
			}
		}

		if len(resultRecipes) >= 10 {
			return strings.Join(resultRecipes[0:10], "")
		}

		firstElf = scoreboard.Move(firstElf)
		secondElf = scoreboard.Move(secondElf)
	}
}

func solveTask2(sequence string) int {
	scoreboard := initScoreboard()

	firstElf := scoreboard.First
	secondElf := scoreboard.Last

	for i := 0; ; i++ {
		values := makeNewRecipes(firstElf, secondElf)

		for _, value := range *values {
			stringValue, _ := strconv.Atoi(value)
			scoreboard.AddRecipe(&Recipe{Value: stringValue})

			if scoreboard.LastSequence(len(sequence)) == sequence {
				return scoreboard.RecipeCount - len(sequence)
			}
		}

		firstElf = scoreboard.Move(firstElf)
		secondElf = scoreboard.Move(secondElf)
	}
}

func main() {
	const input = "147061"

	recipeCount, _ := strconv.Atoi(input)
	fmt.Println("Task 1 solution is", solveTask1(recipeCount))
	fmt.Println("Task 2 solution is", solveTask2(input))
}
