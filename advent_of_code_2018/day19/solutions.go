package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getProgram() (string, *[]string) {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	ipInit := scanner.Text()

	program := []string{}
	for scanner.Scan() {
		program = append(program, scanner.Text())
	}

	return ipInit, &program
}

const A = 0
const B = 1
const C = 2

func solveTask(zeroRegisterValue int) int {
	ipInit, program := getProgram()

	instructionPointer := 0

	ipRegister, _ := strconv.Atoi(strings.Replace(ipInit, "#ip ", "", 1))

	registers := []int{zeroRegisterValue, 0, 0, 0, 0, 0}

	for {
		if instructionPointer >= len(*program) {
			break
		}

		if instructionPointer == 1 {
			total := 0
			for x := 1; x < registers[5]+1; x++ {
				if registers[5]%x == 0 {
					total += x
				}
			}
			return total
		}

		registers[ipRegister] = instructionPointer

		line := (*program)[instructionPointer]

		opcode := line[0:5]
		instrStrings := strings.Split(strings.Replace(line, opcode, "", 1), " ")

		instruction := []int{}
		for _, instrString := range instrStrings {
			register, _ := strconv.Atoi(instrString)
			instruction = append(instruction, register)
		}

		if strings.Contains(opcode, "addi") {
			registers[instruction[C]] = registers[instruction[A]] + instruction[B]
		} else if strings.Contains(opcode, "addr") {
			registers[instruction[C]] = registers[instruction[A]] + registers[instruction[B]]
		} else if strings.Contains(opcode, "seti") {
			registers[instruction[C]] = instruction[A]
		} else if strings.Contains(opcode, "setr") {
			registers[instruction[C]] = registers[instruction[A]]
		} else if strings.Contains(opcode, "mulr") {
			registers[instruction[C]] = registers[instruction[A]] * registers[instruction[B]]
		} else if strings.Contains(opcode, "muli") {
			registers[instruction[C]] = registers[instruction[A]] * instruction[B]
		} else if strings.Contains(opcode, "eqrr") {
			if registers[instruction[A]] == registers[instruction[B]] {
				registers[instruction[C]] = 1
			} else {
				registers[instruction[C]] = 0
			}
		} else if strings.Contains(opcode, "gtrr") {
			if registers[instruction[A]] > registers[instruction[B]] {
				registers[instruction[C]] = 1
			} else {
				registers[instruction[C]] = 0
			}
		}

		instructionPointer = registers[ipRegister]
		instructionPointer++
	}

	return registers[0]
}

func main() {
	fmt.Println("Task 1 solution is", solveTask(0))
	fmt.Println("Task 2 solution is", solveTask(1))
}
