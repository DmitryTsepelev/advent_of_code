package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Sample struct {
	Before      *[]int
	Instruction *[]int
	After       *[]int
}

type Operation interface {
	Title() string
	Execute(registers *[]int, instruction *[]int)
}

func Check(op Operation, sample *Sample) bool {
	registers := make([]int, 4)
	copy(registers, *sample.Before)
	op.Execute(&registers, sample.Instruction)
	return reflect.DeepEqual(*sample.After, registers)
}

// -------------------------------

type Addr struct{}

func (op Addr) Title() string {
	return "addr"
}

func (_ Addr) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := (*registers)[instructionData[2]]

	(*registers)[instructionData[3]] = arg1 + arg2
}

// -------------------------------

type Addi struct{}

func (_ Addi) Title() string {
	return "addi"
}

func (_ Addi) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := instructionData[2]

	(*registers)[instructionData[3]] = arg1 + arg2
}

// -------------------------------

type Murl struct{}

func (_ Murl) Title() string {
	return "murl"
}

func (_ Murl) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := (*registers)[instructionData[2]]

	(*registers)[instructionData[3]] = arg1 * arg2
}

// -------------------------------

type Muli struct{}

func (_ Muli) Title() string {
	return "muli"
}

func (_ Muli) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := instructionData[2]

	(*registers)[instructionData[3]] = arg1 * arg2
}

// -------------------------------

type Banr struct{}

func (_ Banr) Title() string {
	return "banr"
}

func (_ Banr) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := (*registers)[instructionData[2]]

	(*registers)[instructionData[3]] = arg1 & arg2
}

// -------------------------------

type Bani struct{}

func (_ Bani) Title() string {
	return "bani"
}

func (_ Bani) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := instructionData[2]

	(*registers)[instructionData[3]] = arg1 & arg2
}

// -------------------------------

type Borr struct{}

func (_ Borr) Title() string {
	return "borr"
}

func (_ Borr) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := (*registers)[instructionData[2]]

	(*registers)[instructionData[3]] = arg1 | arg2
}

// -------------------------------

type Bori struct{}

func (_ Bori) Title() string {
	return "bori"
}

func (_ Bori) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := instructionData[2]

	(*registers)[instructionData[3]] = arg1 | arg2
}

// -------------------------------

type Setr struct{}

func (_ Setr) Title() string {
	return "setr"
}

func (_ Setr) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	(*registers)[instructionData[3]] = (*registers)[instructionData[1]]
}

// -------------------------------

type Seti struct{}

func (_ Seti) Title() string {
	return "seti"
}

func (_ Seti) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	(*registers)[instructionData[3]] = instructionData[1]
}

// -------------------------------

type Gtir struct{}

func (_ Gtir) Title() string {
	return "gtir"
}

func (_ Gtir) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := instructionData[1]
	arg2 := (*registers)[instructionData[2]]

	if arg1 > arg2 {
		(*registers)[instructionData[3]] = 1
	} else {
		(*registers)[instructionData[3]] = 0
	}
}

// -------------------------------

type Gtri struct{}

func (_ Gtri) Title() string {
	return "gtri"
}

func (_ Gtri) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := instructionData[2]

	if arg1 > arg2 {
		(*registers)[instructionData[3]] = 1
	} else {
		(*registers)[instructionData[3]] = 0
	}
}

// -------------------------------

type Gtrr struct{}

func (_ Gtrr) Title() string {
	return "gtrr"
}

func (_ Gtrr) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := (*registers)[instructionData[2]]

	if arg1 > arg2 {
		(*registers)[instructionData[3]] = 1
	} else {
		(*registers)[instructionData[3]] = 0
	}
}

// -------------------------------

type Eqir struct{}

func (_ Eqir) Title() string {
	return "eqir"
}

func (_ Eqir) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := instructionData[1]
	arg2 := (*registers)[instructionData[2]]

	if arg1 == arg2 {
		(*registers)[instructionData[3]] = 1
	} else {
		(*registers)[instructionData[3]] = 0
	}
}

// -------------------------------

type Eqri struct{}

func (_ Eqri) Title() string {
	return "eqri"
}

func (_ Eqri) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := instructionData[2]

	if arg1 == arg2 {
		(*registers)[instructionData[3]] = 1
	} else {
		(*registers)[instructionData[3]] = 0
	}
}

// -------------------------------

type Eqrr struct{}

func (_ Eqrr) Title() string {
	return "eqrr"
}

func (_ Eqrr) Execute(registers *[]int, instruction *[]int) {
	instructionData := *instruction

	arg1 := (*registers)[instructionData[1]]
	arg2 := (*registers)[instructionData[2]]

	if arg1 == arg2 {
		(*registers)[instructionData[3]] = 1
	} else {
		(*registers)[instructionData[3]] = 0
	}
}

// -------------------------------

func getData() (*[]*Sample, *[]*[]int) {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	samples := []*Sample{}

	var sample *Sample
	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			if sample == nil {
				break
			}
			samples = append(samples, sample)
			sample = nil
		} else if sample == nil {
			before := []int{}
			input = strings.Replace(input, "Before: [", "", 1)
			input = strings.Replace(input, "]", "", 1)
			for _, c := range strings.Split(input, ", ") {
				intValue, _ := strconv.Atoi(c)
				before = append(before, intValue)
			}
			sample = &Sample{Before: &before}
		} else if sample.Instruction == nil {
			instruction := []int{}
			for _, c := range strings.Split(input, " ") {
				intValue, _ := strconv.Atoi(c)
				instruction = append(instruction, intValue)
			}
			sample.Instruction = &instruction
		} else {
			after := []int{}
			input = strings.Replace(input, "After:  [", "", 1)
			input = strings.Replace(input, "]", "", 1)
			for _, c := range strings.Split(input, ", ") {
				intValue, _ := strconv.Atoi(c)
				after = append(after, intValue)
			}
			sample.After = &after
		}
	}

	program := []*[]int{}

	scanner.Scan()
	scanner.Text()
	for scanner.Scan() {
		input := scanner.Text()
		row := []int{}
		for _, c := range strings.Split(input, " ") {
			intValue, _ := strconv.Atoi(c)
			row = append(row, intValue)
		}
		program = append(program, &row)
	}

	return &samples, &program
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	samples, program := getData()
	// samples, _ := getData()

	operations := []Operation{
		Addi{},
		Addr{},
		Murl{},
		Muli{},
		Banr{},
		Bani{},
		Borr{},
		Bori{},
		Setr{},
		Seti{},
		Gtir{},
		Gtri{},
		Gtrr{},
		Eqir{},
		Eqri{},
		Eqrr{},
	}

	behaveLikeThree := 0

	for _, sample := range *samples {
		matchingCount := 0

		for _, operation := range operations {
			if Check(operation, sample) {
				matchingCount++
			}
		}

		if matchingCount >= 3 {
			behaveLikeThree++
		}
	}

	fmt.Println("Task 1 solution is", behaveLikeThree)

	possible := [16][]string{}

	for _, sample := range *samples {
		for _, operation := range operations {
			if Check(operation, sample) {
				opcode := (*sample.Instruction)[0]
				if !contains(possible[opcode], operation.Title()) {
					possible[opcode] = append(possible[opcode], operation.Title())
				}
			}
		}
	}

	uniqueOps := map[int]string{}

	for {
		for op, candidates := range possible {
			if len(candidates) == 1 {
				uniqueOps[op] = candidates[0]
			}
		}

		for _, symOpcode := range uniqueOps {
			for op, candidates := range possible {
				newCandidates := []string{}

				for _, possibleOpcode := range candidates {
					if possibleOpcode != symOpcode {
						newCandidates = append(newCandidates, possibleOpcode)
					}
				}

				possible[op] = newCandidates
			}
		}

		if len(uniqueOps) == len(possible) {
			break
		}
	}

	registers := make([]int, 4)
	for _, instruction := range *program {
		opcode := uniqueOps[(*instruction)[0]]

		var operation *Operation
		for _, op := range operations {
			if op.Title() == opcode {
				operation = &op
				break
			}
		}

		(*operation).Execute(&registers, instruction)
	}

	fmt.Println("Task 2 solution is", registers[0])
}
