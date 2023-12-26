package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() *[][]byte {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := [][]byte{}
	for scanner.Scan() {
		value := []byte(scanner.Text())
		numbers = append(numbers, value)
	}

	return &numbers
}

func getMostCommonValue(index int, numbers *[][]byte) byte {
	count := 0

	for _, row := range *numbers {
		if row[index] == '1' {
			count += 1
		} else {
			count -= 1
		}
	}

	if count >= 0 {
		return '1'
	}

	return '0'
}

func getRate(numbers *[][]byte, controlByte byte) int {
	rate := []byte{}

	for i := 0; i < len((*numbers)[0]); i++ {
		mostCommon := getMostCommonValue(i, numbers)

		if mostCommon == controlByte {
			rate = append(rate, '1')
		} else {
			rate = append(rate, '0')
		}
	}

	decimal, _ := strconv.ParseInt(string(rate), 2, 64)

	return int(decimal)
}

func powerConsumption(numbers [][]byte) int {
	gammaRate := getRate(&numbers, '1')
	epsilonRate := getRate(&numbers, '0')

	return gammaRate * epsilonRate
}

func getRatingByCriteria(numbers [][]byte, bitCriteria func(byte, byte) bool) int {
	for bitPosition := 0; bitPosition < len(numbers[0]); bitPosition++ {
		if len(numbers) == 1 {
			break
		}
		mostCommon := getMostCommonValue(bitPosition, &numbers)

		newNumbers := [][]byte{}

		for _, number := range numbers {
			if bitCriteria(number[bitPosition], mostCommon) {
				newNumbers = append(newNumbers, number)
			}
		}

		numbers = newNumbers
	}

	decimal, _ := strconv.ParseInt(string(numbers[0]), 2, 64)

	return int(decimal)
}

func lifeSupportRating(numbers [][]byte) int {
	oxygenGeneratorRating := getRatingByCriteria(numbers, func(b1, b2 byte) bool {
		return b1 == b2
	})

	co2ScrubberRating := getRatingByCriteria(numbers, func(b1, b2 byte) bool {
		return b1 != b2
	})

	return oxygenGeneratorRating * co2ScrubberRating
}

func main() {
	numbers := *getInputData()
	fmt.Println("Part 1 solution:", powerConsumption(numbers))
	fmt.Println("Part 2 solution:", lifeSupportRating(numbers))
}
