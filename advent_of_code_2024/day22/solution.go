package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getInputData() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nums := []int{}

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	return nums
}

func findNext(secretNumber int) int {
	secretNumber = prune(mix((secretNumber * 64), secretNumber))
	secretNumber = prune(mix((secretNumber / 32), secretNumber))
	secretNumber = prune(mix((secretNumber * 2048), secretNumber))
	return secretNumber
}

func mix(x, y int) int {
	return x ^ y
}

func prune(x int) int {
	return x % 16777216
}

func part1(secretNumbers []int) int {
	var sum int
	for _, secretNumber := range secretNumbers {
		for i := 0; i < 2000; i++ {
			secretNumber = findNext(secretNumber)
		}
		sum += secretNumber
	}
	return sum
}

func secretToPrice(num int) int {
	return num % 10
}

type Sequence = [4]int

func part2(secretNumbers []int) int {
	limit := 2000

	sequences := map[Sequence](map[int]int){}

	for monkeyId, secretNumber := range secretNumbers {
		data := [][3]int{}
		prevPrice := secretToPrice(secretNumber)
		for i := 0; i <= limit; i++ {
			secretNumber = findNext(secretNumber)
			currentPrice := secretToPrice(secretNumber)
			data = append(data, [3]int{secretNumber, currentPrice, currentPrice - prevPrice})
			prevPrice = currentPrice
		}

		for i := 0; i < len(data)-4; i++ {
			sequence := Sequence{}
			for j := 0; j < 4; j++ {
				sequence[j] = data[i+j][2]
			}

			if _, ok := sequences[sequence]; !ok {
				sequences[sequence] = make(map[int]int)
			}

			if _, ok := sequences[sequence][monkeyId]; !ok {
				sequences[sequence][monkeyId] = data[i+3][1]
			}
		}
	}

	best := math.MinInt
	for _, monkeyPrices := range sequences {
		candidate := 0
		for _, price := range monkeyPrices {
			candidate += price
		}

		if candidate > best {
			best = candidate
		}
	}

	return best
}

func main() {
	secretNumbers := getInputData()

	fmt.Println("Part 1 solution is", part1(secretNumbers))
	fmt.Println("Part 2 solution is", part2(secretNumbers))
}
