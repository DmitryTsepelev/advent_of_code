package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type InstructionCpy struct {
	registerTo string
	value      string
}

type InstructionInc struct {
	register string
}

type InstructionDec struct {
	register string
}

type InstructionJnz struct {
	source string
	value  int
}

type Program = []interface{}

func getInputData() Program {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	cpyRe, _ := regexp.Compile(`cpy (\w|-?\d+) (\w)`)
	incRe, _ := regexp.Compile(`inc (\w)`)
	decRe, _ := regexp.Compile(`dec (\w)`)
	jnzRe, _ := regexp.Compile(`jnz (\w) (-?\d+)`)

	scanner := bufio.NewScanner(file)
	program := make(Program, 0)

	for scanner.Scan() {
		line := scanner.Text()

		cpy := cpyRe.FindAllStringSubmatch(line, -1)
		if len(cpy) > 0 {
			valueS, register := cpy[0][1], cpy[0][2]

			program = append(program, InstructionCpy{
				value:      valueS,
				registerTo: register,
			})

			continue
		}

		inc := incRe.FindAllStringSubmatch(line, -1)
		if len(inc) > 0 {
			register := inc[0][1]

			program = append(program, InstructionInc{
				register: register,
			})

			continue
		}

		dec := decRe.FindAllStringSubmatch(line, -1)
		if len(dec) > 0 {
			register := dec[0][1]

			program = append(program, InstructionDec{
				register: register,
			})

			continue
		}

		jnz := jnzRe.FindAllStringSubmatch(line, -1)
		if len(jnz) > 0 {
			source, valueS := jnz[0][1], jnz[0][2]

			value, _ := strconv.Atoi(valueS)

			program = append(program, InstructionJnz{
				source: source,
				value:  value,
			})

			continue
		}

		panic(line)
	}

	return program
}

type Registers = map[string]int

func readValue(registers Registers, value string) int {
	if value == "a" || value == "b" || value == "c" || value == "d" {
		return registers[value]
	}

	intValue, _ := strconv.Atoi(value)
	return intValue
}

func execute(program Program, registers Registers) int {
	instPtr := 0

	for instPtr < len(program) {
		instr := program[instPtr]

		switch typedInstr := instr.(type) {
		case InstructionCpy:
			registers[typedInstr.registerTo] = readValue(registers, typedInstr.value)
			instPtr++
		case InstructionInc:
			registers[typedInstr.register]++
			instPtr++
		case InstructionDec:
			registers[typedInstr.register]--
			instPtr++
		case InstructionJnz:
			value := readValue(registers, typedInstr.source)
			if value == 0 {
				instPtr++
			} else {
				instPtr += typedInstr.value
			}
		default:
			panic(typedInstr)
		}
	}

	return registers["a"]
}

func main() {
	program := getInputData()

	fmt.Println(
		"Solution 1 is",
		execute(
			program,
			Registers{
				"a": 0,
				"b": 0,
				"c": 0,
				"d": 0,
			},
		),
	)

	fmt.Println(
		"Solution 2 is",
		execute(
			program,
			Registers{
				"a": 0,
				"b": 0,
				"c": 1,
				"d": 0,
			},
		),
	)
}
