package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() [][]byte {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := [][]byte{}
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}

	return input
}

var deltas = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func validPoint(input [][]byte, nextPoint []int) bool {
	row, col := nextPoint[0], nextPoint[1]

	return row >= 0 && row < len(input) &&
		col >= 0 && col < len(input[0])

	// if row >= 0 && row < len(input) &&
	// 	col >= 0 && col < len(input[0]) {
	// 	// fmt.Println(row, col, string(input[row][col]), string(word[step]))
	// 	return input[row][col] == word[step]
	// }
	// return false
	//==

	// return row >= 0 && row < len(input) &&
	// 	col >= 0 && col < len(input[0]) &&
	// 	input[row][col] == word[step]
}

// func bfs(input [][]byte, row, col int) int {
// 	fmt.Println("BFS", row, col)
// 	word := "XMAS"

// 	queue := [][]int{{row, col}}

// 	for step := 0; step <= 3; step++ {
// 		fmt.Println("step", step)
// 		for _, p := range queue {
// 			fmt.Println(string(input[p[0]][p[1]]))
// 		}
// 		// fmt.Println(queue)
// 		nextQueue := make(map[[2]int]bool)

// 		for _, point := range queue {
// 			for _, delta := range deltas {
// 				nextPoint := []int{point[0] + delta[0], point[1] + delta[1]}

// 				if step < 3 && validPoint(input, nextPoint, word, step+1) && input[nextPoint[0]][nextPoint[1]] == word[step+1] {
// 					// nextQueue = append(nextQueue, nextPoint)
// 					nextQueue[[2]int{nextPoint[0], nextPoint[1]}] = true
// 				}
// 			}
// 		}

// 		if step == 3 {
// 			break
// 		}

// 		// queue = nextQueue

// 		queue = [][]int{}
// 		for p := range nextQueue {
// 			queue = append(queue, []int{p[0], p[1]})
// 		}
// 	}
// 	// fmt.Println("=--", queue)

// 	uniquePoints := make(map[[2]int]bool)
// 	for _, p := range queue {
// 		uniquePoints[[2]int{p[0], p[1]}] = true
// 	}
// 	fmt.Println("adding", len(uniquePoints))

// 	return len(uniquePoints)
// }

func part1(input [][]byte) int {
	var count int
	word := "XMAS"

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			// if input[i][j] == 'X' {
			// 	// count += bfs(input, i, j)

			// }

			if input[i][j] != 'X' {
				continue
			}

			for _, delta := range deltas {
				row, col := i, j
				validWord := true

				for step := 0; step <= 3; step++ {
					fmt.Println("row", row, "col", col)
					if !validPoint(input, []int{row, col}) || word[step] != input[row][col] {
						validWord = false
						fmt.Println("invalid!")
						break
					}

					row += delta[0]
					col += delta[1]
				}

				if validWord {
					fmt.Println("valid")
					count++
				}
			}
		}
	}

	return count
}

func main() {
	input := getInputData()

	fmt.Println(part1(input))
}
