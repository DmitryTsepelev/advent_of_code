package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	RECT = 0
	ROW  = 1
	COL  = 2
)

type Command struct {
	kind int
	arg1 int
	arg2 int
}

func parseLine(line string) Command {
	if strings.HasPrefix(line, "rect ") {
		idx := 5

		arg1s := ""
		for line[idx] != 'x' {
			arg1s += string(line[idx])
			idx++
		}
		arg1, _ := strconv.Atoi(arg1s)

		idx++

		arg2s := ""
		for idx < len(line) {
			arg2s += string(line[idx])
			idx++
		}
		arg2, _ := strconv.Atoi(arg2s)

		return Command{kind: RECT, arg1: arg1, arg2: arg2}
	}

	if strings.HasPrefix(line, "rotate row") {
		idx := 13

		arg1s := ""
		for line[idx] != ' ' {
			arg1s += string(line[idx])
			idx++
		}
		arg1, _ := strconv.Atoi(arg1s)

		idx += 4

		arg2s := ""
		for idx < len(line) {
			arg2s += string(line[idx])
			idx++
		}
		arg2, _ := strconv.Atoi(arg2s)

		return Command{kind: ROW, arg1: arg1, arg2: arg2}
	}

	if strings.HasPrefix(line, "rotate column") {
		idx := 16

		arg1s := ""
		for line[idx] != ' ' {
			arg1s += string(line[idx])
			idx++
		}
		arg1, _ := strconv.Atoi(arg1s)

		idx += 4

		arg2s := ""
		for idx < len(line) {
			arg2s += string(line[idx])
			idx++
		}
		arg2, _ := strconv.Atoi(arg2s)

		return Command{kind: COL, arg1: arg1, arg2: arg2}
	}

	panic("unreachable")
}

func getInputData() []Command {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := []Command{}

	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, parseLine(line))
	}

	return field
}

type Screen = [][]bool

func initScreen(rows, columns int) Screen {
	screen := make(Screen, rows)

	for row := 0; row < rows; row++ {
		screen[row] = make([]bool, columns)
	}

	return screen
}

func rect(screen Screen, width, height int) {
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			screen[row][col] = true
		}
	}
}

func rotateColumn(screen Screen, col int, shift int) {
	newCol := make([]bool, len(screen))

	for row := 0; row < len(screen); row++ {
		if !screen[row][col] {
			continue
		}
		shifted := (row + shift) % len(screen)
		newCol[shifted] = true
	}

	for row := 0; row < len(screen); row++ {
		screen[row][col] = newCol[row]
	}
}

func rotateRow(screen Screen, row int, shift int) {
	newRow := make([]bool, len(screen[0]))

	for col := 0; col < len(screen[0]); col++ {
		if !screen[row][col] {
			continue
		}
		shifted := (col + shift) % len(screen[0])
		newRow[shifted] = true
	}

	for col := 0; col < len(screen[0]); col++ {
		screen[row][col] = newRow[col]
	}
}

func showScreen(screen Screen) {
	for row := 0; row < len(screen); row++ {
		for col := 0; col < len(screen[0]); col++ {
			if screen[row][col] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	screen := initScreen(6, 50)
	commands := getInputData()

	for _, command := range commands {
		switch command.kind {
		case RECT:
			rect(screen, command.arg1, command.arg2)
		case ROW:
			rotateRow(screen, command.arg1, command.arg2)
		case COL:
			rotateColumn(screen, command.arg1, command.arg2)
		}
	}

	count := 0
	for row := 0; row < len(screen); row++ {
		for col := 0; col < len(screen[0]); col++ {
			if screen[row][col] {
				count++
			}
		}
	}

	fmt.Println("Solution 1 is", count)

	fmt.Println("Solution 2 is")
	showScreen(screen)
}
