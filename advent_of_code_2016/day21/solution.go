package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	SWAP_POSITION            = "SWAP_POSITION"
	SWAP_LETTER              = "SWAP_LETTER"
	ROTATE_LEFT              = "ROTATE_LEFT"
	ROTATE_RIGHT             = "ROTATE_RIGHT"
	ROTATE_BASED_ON_POSITION = "ROTATE_BASED_ON_POSITION"
	REVERSE_POSITIONS        = "REVERSE_POSITIONS"
	MOVE_POSITION            = "MOVE_POSITION"
)

type Command struct {
	kind     string
	intArgs  []int
	runeArgs []rune
}

func getInputData() []Command {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	commands := make([]Command, 0)

	swapPositionRe, _ := regexp.Compile(`swap position (\d+) with position (\d+)`)
	swapLetterRe, _ := regexp.Compile(`swap letter (\w+) with letter (\w+)`)
	rotateLeftRe, _ := regexp.Compile(`rotate left (\d+) step(s)?`)
	rotateRightRe, _ := regexp.Compile(`rotate right (\d+) step(s)?`)
	rotatePositionRe, _ := regexp.Compile(`rotate based on position of letter (\w+)`)
	reversePositionsRe, _ := regexp.Compile(`reverse positions (\d+) through (\d+)`)
	movePositionRe, _ := regexp.Compile(`move position (\d+) to position (\d+)`)

	for scanner.Scan() {
		line := scanner.Text()

		match := swapPositionRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			firstS, secondS := match[0][1], match[0][2]

			first, _ := strconv.Atoi(firstS)
			second, _ := strconv.Atoi(secondS)

			commands = append(commands, Command{
				kind:    SWAP_POSITION,
				intArgs: []int{first, second},
			})

			continue
		}

		match = swapLetterRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			firstS, secondS := match[0][1], match[0][2]

			commands = append(commands, Command{
				kind:     SWAP_LETTER,
				runeArgs: []rune{rune(firstS[0]), rune(secondS[0])},
			})

			continue
		}

		match = rotateLeftRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			firstS := match[0][1]

			first, _ := strconv.Atoi(firstS)

			commands = append(commands, Command{
				kind:    ROTATE_LEFT,
				intArgs: []int{first},
			})

			continue
		}

		match = rotateRightRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			firstS := match[0][1]

			first, _ := strconv.Atoi(firstS)

			commands = append(commands, Command{
				kind:    ROTATE_RIGHT,
				intArgs: []int{first},
			})

			continue
		}

		match = rotatePositionRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			firstS := match[0][1]

			commands = append(commands, Command{
				kind:     ROTATE_BASED_ON_POSITION,
				runeArgs: []rune{rune(firstS[0])},
			})

			continue
		}

		match = reversePositionsRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			firstS, secondS := match[0][1], match[0][2]

			first, _ := strconv.Atoi(firstS)
			second, _ := strconv.Atoi(secondS)

			commands = append(commands, Command{
				kind:    REVERSE_POSITIONS,
				intArgs: []int{first, second},
			})

			continue
		}

		match = movePositionRe.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			firstS, secondS := match[0][1], match[0][2]

			first, _ := strconv.Atoi(firstS)
			second, _ := strconv.Atoi(secondS)

			commands = append(commands, Command{
				kind:    MOVE_POSITION,
				intArgs: []int{first, second},
			})

			continue
		}

		panic(line)
	}

	return commands
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func substring(s string, from, to int) string {
	// fmt.Println(max(0, from), min(to, len(s)))
	return s[min(len(s), max(0, from)):max(0, min(to, len(s)))]
}

func swapPosition(s string, x, y int) string {
	if x > y {
		x, y = y, x
	}

	// fmt.Println(x, y)
	// fmt.Println(substring(s, 0, x-1), string(s[y]), substring(s, x+1, y-1), string(s[x]), substring(s, y+1, len(s)))
	return substring(s, 0, x) + string(s[y]) + substring(s, x+1, y) + string(s[x]) + substring(s, y+1, len(s))
}

func swapLetter(s string, x, y rune) string {
	var idx1 int
	var idx2 int

	for idx, c := range s {
		if c == x {
			idx1 = idx
		}
		if c == y {
			idx2 = idx
		}
	}

	return swapPosition(s, idx1, idx2)
}

func reversePositions(s string, x, y int) string {
	if x > y {
		x, y = y, x
	}

	newS := substring(s, 0, x)

	for idx := y; idx >= x; idx-- {
		newS += string(s[idx])
	}

	newS += substring(s, y+1, len(s))

	return newS
}

func rotateLeft(s string, positions int) string {
	pos := positions % len(s)
	return s[pos:] + s[:pos]
}

func rotateRight(s string, positions int) string {
	pos := len(s) - positions%len(s)
	return s[pos:] + s[:pos]
}

func movePosition(s string, x, y int) string {
	removed := substring(s, 0, x) + substring(s, x+1, len(s))

	return substring(removed, 0, y) + string(s[x]) + substring(removed, y, len(removed))
}

func rotateByLetterRight(s string, x rune) string {
	var xidx int

	for idx, c := range s {
		if c == x {
			xidx = idx
		}
	}

	rotations := 1 + xidx
	if xidx >= 4 {
		rotations++
	}

	return rotateRight(s, rotations)
}

func rotateByLetterLeft(s string, x rune) string {
	var xidx int

	for idx, c := range s {
		if c == x {
			xidx = idx
		}
	}

	rotations := 1 + xidx
	if xidx >= 4 {
		rotations++
	}

	return rotateLeft(s, rotations)
}

func scramble(s string, commands []Command) string {
	for _, cmd := range commands {
		switch cmd.kind {
		case SWAP_POSITION:
			s = swapPosition(s, cmd.intArgs[0], cmd.intArgs[1])
		case SWAP_LETTER:
			s = swapLetter(s, cmd.runeArgs[0], cmd.runeArgs[1])
		case ROTATE_LEFT:
			s = rotateLeft(s, cmd.intArgs[0])
		case ROTATE_RIGHT:
			s = rotateRight(s, cmd.intArgs[0])
		case ROTATE_BASED_ON_POSITION:
			s = rotateByLetterRight(s, cmd.runeArgs[0])
		case REVERSE_POSITIONS:
			s = reversePositions(s, cmd.intArgs[0], cmd.intArgs[1])
		case MOVE_POSITION:
			s = movePosition(s, cmd.intArgs[0], cmd.intArgs[1])
		}
	}

	return s
}

func unscramble(s string, commands []Command) string {
	for i := len(commands) - 1; i >= 0; i-- {
		cmd := commands[i]

		switch cmd.kind {
		case SWAP_POSITION:
			s = swapPosition(s, cmd.intArgs[1], cmd.intArgs[0])
		case SWAP_LETTER:
			s = swapLetter(s, cmd.runeArgs[0], cmd.runeArgs[1])
		case ROTATE_LEFT:
			s = rotateRight(s, cmd.intArgs[0])
		case ROTATE_RIGHT:
			s = rotateLeft(s, cmd.intArgs[0])
		case ROTATE_BASED_ON_POSITION:
			s = rotateByLetterLeft(s, cmd.runeArgs[0])
		case REVERSE_POSITIONS:
			s = reversePositions(s, cmd.intArgs[1], cmd.intArgs[0])
		case MOVE_POSITION:
			s = movePosition(s, cmd.intArgs[1], cmd.intArgs[0])
		}
	}

	return s
}

func main() {
	commands := getInputData()

	fmt.Println("Solution 1 is", scramble("abcdefgh", commands))
	fmt.Println("Solution 2 is", unscramble("fbgdceah", commands))
}
