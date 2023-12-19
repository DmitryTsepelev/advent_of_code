package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}

func decompress(s string, isV2 bool) int {
	length := 0

	idx := 0

	for idx < len(s) {
		c := s[idx]

		if c == '(' {
			// look for marker
			idx++
			decompressedLengthS := ""
			for s[idx] != 'x' {
				decompressedLengthS += string(s[idx])
				idx++
			}

			idx++ // skip x

			multS := ""
			for s[idx] != ')' {
				multS += string(s[idx])
				idx++
			}
			idx++ // skip )

			decompressedLength, _ := strconv.Atoi(decompressedLengthS)
			mult, _ := strconv.Atoi(multS)

			if isV2 {
				length += decompress(s[idx:idx+decompressedLength], true) * mult
			} else {
				length += decompressedLength * mult
			}

			idx += decompressedLength
		} else {
			length++
			idx++
		}
	}

	return length
}

func main() {
	input := getInputData()
	fmt.Println("Solution 1 is", decompress(input, false))
	fmt.Println("Solution 2 is", decompress(input, true))
}
