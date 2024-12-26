package main

import (
	"bufio"
	"fmt"
	"os"
)

type Key = [5]int
type Lock = [5]int

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getInputData() ([]Key, []Lock) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	keys := []Key{}
	locks := []Lock{}

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			scanner.Scan()
			continue
		}

		mx := [7]string{}
		for i := 0; i < 7; i++ {
			mx[i] = scanner.Text()
			scanner.Scan()
		}

		obj := [5]int{}
		for j := 0; j < 5; j++ {
			for i := 0; i < 7; i++ {
				if mx[i][j] == '#' {
					obj[j]++
				}
			}
		}

		isLock := true
		for j := 0; j < 5; j++ {
			if mx[0][j] == '.' {
				isLock = false
				break
			}
		}

		if isLock {
			locks = append(locks, obj)
		} else {
			keys = append(keys, obj)
		}
	}

	return keys, locks
}

func canOpen(key Key, lock Lock) bool {
	for idx, kHeight := range key {
		if lock[idx]+kHeight > 7 {
			return false
		}
	}

	return true
}

func remove(s []Key, i int) []Key {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	keys, locks := getInputData()

	var pairs int
	for _, lock := range locks {
		for _, key := range keys {
			if canOpen(key, lock) {
				pairs++
			}
		}
	}
	fmt.Println("Part 1 solution is", pairs)
}
