package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD_OPCODE           = 1
	MULT_OPCODE          = 2
	READ_INPUT_OPCODE    = 3
	WRITE_OUTPUT_OPCODE  = 4
	JUMP_IF_TRUE_OPCODE  = 5
	JUMP_IF_FALSE_OPCODE = 6
	LESS_THAN_OPCODE     = 7
	EQUALS_OPCODE        = 8
	EXIT_OPCODE          = 99
)

type Mode = bool

func modeFor(fullOpcode int64, argIdx int) Mode {
	modes := fullOpcode / 100
	sModes := strconv.FormatInt(modes, 10)

	modeIdx := len(sModes) - argIdx - 1
	if modeIdx < 0 {
		return MODE_POSITION
	}

	return sModes[modeIdx] == '1'
}

type VM struct {
	instructionIdx int64
	program        []int64
	input          chan int64
	output         chan int64
}

func CreateVm(program []int64) VM {
	return VM{program: program, input: make(chan int64, 1), output: make(chan int64)}
}

const (
	MODE_POSITION  = false
	MODE_IMMEDIATE = true
)

func (this *VM) close() {
	var end bool
	for !end {
		select {
		case <-this.input:
		default:
			end = true
		}
	}
	close(this.input)
}

func (this *VM) Read() (int64, bool) {
	out, ok := <-this.output
	return out, ok
}

func (this *VM) readValue(parameter int64, mode Mode) int64 {
	switch mode {
	case MODE_POSITION:
		{
			position := parameter
			return this.program[position]
		}
	case MODE_IMMEDIATE:
		return parameter
	default:
		panic(mode)
	}
}

func (this *VM) Run() {
	var err error
	for err == nil {
		err = this.Next()
	}
}

func (this *VM) Next() error {
	fullOpcode := this.program[this.instructionIdx]

	opcode := fullOpcode % 100

	switch opcode {
	case ADD_OPCODE:
		{
			param1 := this.readValue(this.program[this.instructionIdx+1], modeFor(fullOpcode, 0))
			param2 := this.readValue(this.program[this.instructionIdx+2], modeFor(fullOpcode, 1))
			this.program[this.program[this.instructionIdx+3]] = param1 + param2

			this.instructionIdx += 4
		}
	case MULT_OPCODE:
		{
			param1 := this.readValue(this.program[this.instructionIdx+1], modeFor(fullOpcode, 0))
			param2 := this.readValue(this.program[this.instructionIdx+2], modeFor(fullOpcode, 1))
			this.program[this.program[this.instructionIdx+3]] = param1 * param2

			this.instructionIdx += 4
		}
	case READ_INPUT_OPCODE:
		{
			address := this.program[this.instructionIdx+1]
			this.program[address] = <-this.input

			this.instructionIdx += 2
		}
	case WRITE_OUTPUT_OPCODE:
		{
			value := this.readValue(this.program[this.instructionIdx+1], modeFor(fullOpcode, 0))
			this.output <- value

			this.instructionIdx += 2
		}
	case JUMP_IF_TRUE_OPCODE:
		{
			param1 := this.readValue(this.program[this.instructionIdx+1], modeFor(fullOpcode, 0))
			param2 := this.readValue(this.program[this.instructionIdx+2], modeFor(fullOpcode, 1))

			if param1 != 0 {
				this.instructionIdx = param2
			} else {
				this.instructionIdx += 3
			}
		}
	case JUMP_IF_FALSE_OPCODE:
		{
			param1 := this.readValue(this.program[this.instructionIdx+1], modeFor(fullOpcode, 0))
			param2 := this.readValue(this.program[this.instructionIdx+2], modeFor(fullOpcode, 1))

			if param1 == 0 {
				this.instructionIdx = param2
			} else {
				this.instructionIdx += 3
			}
		}
	case LESS_THAN_OPCODE:
		{
			param1 := this.readValue(this.program[this.instructionIdx+1], modeFor(fullOpcode, 0))
			param2 := this.readValue(this.program[this.instructionIdx+2], modeFor(fullOpcode, 1))

			if param1 < param2 {
				this.program[this.program[this.instructionIdx+3]] = 1
			} else {
				this.program[this.program[this.instructionIdx+3]] = 0
			}

			this.instructionIdx += 4
		}
	case EQUALS_OPCODE:
		{
			param1 := this.readValue(this.program[this.instructionIdx+1], modeFor(fullOpcode, 0))
			param2 := this.readValue(this.program[this.instructionIdx+2], modeFor(fullOpcode, 1))

			if param1 == param2 {
				this.program[this.program[this.instructionIdx+3]] = 1
			} else {
				this.program[this.program[this.instructionIdx+3]] = 0
			}

			this.instructionIdx += 4
		}
	case EXIT_OPCODE:
		{
			close(this.output)
			return fmt.Errorf("halt")
		}
	default:
		{
			close(this.output)
			return fmt.Errorf("error")
		}
	}

	return nil
}

func getInputData() []int64 {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	cmp := strings.Split(input, ",")
	program := make([]int64, len(cmp))
	for idx, c := range cmp {
		num, _ := strconv.Atoi(c)
		program[idx] = int64(num)
	}
	return program
}

func createCombinations(amplifierCount int, startFrom int) [][]int {
	combinations := [][]int{{}}

	for i := 0; i < amplifierCount; i++ {
		nextCombinations := [][]int{}

		for _, combination := range combinations {
			for phase := startFrom; phase <= startFrom+4; phase++ {
				invalid := false
				for _, prevPhase := range combination {
					if prevPhase == phase {
						invalid = true
					}
				}

				if invalid {
					continue
				}

				newCombination := make([]int, len(combination))
				copy(newCombination, combination)
				newCombination = append(newCombination, phase)
				nextCombinations = append(nextCombinations, newCombination)
			}
		}

		combinations = nextCombinations
	}

	return combinations
}

func findBestOutput(amplifierCount int) int64 {
	program := getInputData()
	combinations := createCombinations(5, 0)

	var best int64

	for _, combination := range combinations {
		var prevInput int64
		for ampIdx := 0; ampIdx < amplifierCount; ampIdx++ {
			programCopy := make([]int64, len(program))
			copy(programCopy, program)

			vm := CreateVm(programCopy)
			go vm.Run()
			vm.input <- int64(combination[ampIdx])
			vm.input <- int64(prevInput)

			prevInput = <-vm.output
		}

		best = max(best, prevInput)
	}

	return best
}

func provideInput(amp, prevAmp *VM, initial int64) {
	if prevAmp != nil {
		out, ok := prevAmp.Read()
		if !ok {
			return
		}
		amp.input <- out
	} else {
		amp.input <- initial
	}
}

func findBestOutputWithFeedback(amplifierCount int) int64 {
	program := getInputData()
	combinations := createCombinations(5, 5)

	var best int64

	for _, combination := range combinations {
		amplifiers := make([]*VM, amplifierCount)
		for ampIdx := 0; ampIdx < amplifierCount; ampIdx++ {
			ampProgram := make([]int64, len(program))
			copy(ampProgram, program)
			vm := CreateVm(ampProgram)

			amplifiers[ampIdx] = &vm
			amplifiers[ampIdx].input <- int64(combination[ampIdx])

			defer amplifiers[ampIdx].close()
			go amplifiers[ampIdx].Run()
		}

		var out, outTemp int64
		ok := true
		for ok {
			var prevAmp *VM
			for _, curAmp := range amplifiers {
				go provideInput(curAmp, prevAmp, out)
				prevAmp = curAmp
			}

			outTemp, ok = amplifiers[amplifierCount-1].Read()
			if ok {
				out = outTemp
			}
		}

		if out > best {
			best = out
		}
	}

	return best
}

func main() {
	fmt.Println("Part 1 solution is", findBestOutput(5))
	fmt.Println("Part 2 solution is", findBestOutputWithFeedback(5))
}
