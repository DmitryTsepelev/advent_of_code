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

func solveTask() int {
	ipInit, program := getProgram()

	zeroRegisterValue := 0
	bestZeroRegisterValue := 0
	smallestCommandCount := 10000

	for j := 832312; j < 832313; j++ {
		instructionPointer := 0

		ipRegister, _ := strconv.Atoi(strings.Replace(ipInit, "#ip ", "", 1))
		fmt.Println(zeroRegisterValue)
		registers := []int{zeroRegisterValue, 0, 0, 0, 0, 0}

		halt := true
		commandLimit := smallestCommandCount
		commandCount := 0
		for ; commandCount < commandLimit; commandCount++ {
			fmt.Println(instructionPointer)
			if instructionPointer >= len(*program) {
				halt = false
				fmt.Println("wat")
				break
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
			} else if strings.Contains(opcode, "bani") {
				registers[instruction[C]] = registers[instruction[A]] & instruction[B]
			} else if strings.Contains(opcode, "eqri") {
				if registers[instruction[A]] == instruction[B] {
					registers[instruction[C]] = 1
				} else {
					registers[instruction[C]] = 0
				}
			} else if strings.Contains(opcode, "bori") {
				registers[instruction[C]] = registers[instruction[A]] | instruction[B]
			} else if strings.Contains(opcode, "gtir") {
				if instruction[A] > registers[instruction[B]] {
					registers[instruction[C]] = 1
				} else {
					registers[instruction[C]] = 0
				}
			} else {
				panic(opcode)
			}

			instructionPointer = registers[ipRegister]
			instructionPointer++
		}

		fmt.Println("halt", halt, "commandCount", commandCount)
		if halt && commandCount < smallestCommandCount {
			smallestCommandCount = commandCount
			bestZeroRegisterValue = zeroRegisterValue
		}

		zeroRegisterValue++
	}

	return bestZeroRegisterValue
}

func main() {
	fmt.Println("Task 1 solution is", solveTask())
}
