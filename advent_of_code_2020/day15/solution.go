package main

import "fmt"

func solve(spokenNumbers []int, turns int) int {
	spokenAt := make(map[int]([]int))
	for t, num := range spokenNumbers {
		spokenAt[num] = []int{t + 1}
	}

	for turn := len(spokenNumbers) + 1; turn <= turns; turn++ {
		prevSpoken := spokenNumbers[len(spokenNumbers)-1]

		nextNumber := 0
		if prevTurns, spoken := spokenAt[prevSpoken]; spoken && len(prevTurns) > 1 {
			nextNumber = prevTurns[len(prevTurns)-1] - prevTurns[len(prevTurns)-2]
		}

		spokenNumbers = append(spokenNumbers, nextNumber)

		if _, spoken := spokenAt[nextNumber]; spoken == false {
			spokenAt[nextNumber] = make([]int, 0)
		}
		spokenAt[nextNumber] = append(spokenAt[nextNumber], turn)
	}
	return spokenNumbers[len(spokenNumbers)-1]
}

func main() {
	spokenNumbers := []int{11, 0, 1, 10, 5, 19}
	fmt.Println("Part 1 solution is", solve(spokenNumbers, 2020))
	fmt.Println("Part 2 solution is", solve(spokenNumbers, 30000000))
}
