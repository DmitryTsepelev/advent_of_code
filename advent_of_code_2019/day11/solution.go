package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD_OPCODE                  = 1
	MULT_OPCODE                 = 2
	READ_INPUT_OPCODE           = 3
	WRITE_OUTPUT_OPCODE         = 4
	JUMP_IF_TRUE_OPCODE         = 5
	JUMP_IF_FALSE_OPCODE        = 6
	LESS_THAN_OPCODE            = 7
	EQUALS_OPCODE               = 8
	ADJUST_RELATIVE_BASE_OPCODE = 9
	EXIT_OPCODE                 = 99
)

type Mode = int

func modeFor(fullOpcode int64, argIdx int) Mode {
	modes := fullOpcode / 100
	sModes := strconv.FormatInt(modes, 10)

	modeIdx := len(sModes) - argIdx - 1
	if modeIdx < 0 {
		return MODE_POSITION
	}

	switch sModes[modeIdx] {
	case '0':
		return MODE_POSITION
	case '1':
		return MODE_IMMEDIATE
	case '2':
		return MODE_RELATIVE
	}

	panic("error")
}

type Program = map[int64]int64

type VM struct {
	instructionIdx int64
	relativeBase   int64
	program        Program
	input          chan int64
	output         chan int64
}

func CreateVm(program Program) VM {
	return VM{program: program, relativeBase: 0, input: make(chan int64, 1), output: make(chan int64)}
}

const (
	MODE_POSITION  = 0
	MODE_IMMEDIATE = 1
	MODE_RELATIVE  = 2
)

func (this *VM) Close() {
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

func (this *VM) readArgument(argumentPosition int) int64 {
	fullOpcode := this.program[this.instructionIdx]
	opcode := fullOpcode % 100
	parameter := this.program[this.instructionIdx+int64(argumentPosition)]

	var alwaysImmediateMode bool
	switch argumentPosition {
	case 1:
		alwaysImmediateMode = opcode == READ_INPUT_OPCODE
	case 2:
	case 3:
		alwaysImmediateMode = opcode == ADD_OPCODE || opcode == MULT_OPCODE || opcode == LESS_THAN_OPCODE || opcode == EQUALS_OPCODE
	}

	var mode Mode
	if alwaysImmediateMode {
		mode = MODE_IMMEDIATE
	} else {
		mode = modeFor(fullOpcode, argumentPosition-1)
	}

	switch mode {
	case MODE_POSITION:
		{
			return this.program[parameter]
		}
	case MODE_IMMEDIATE:
		return parameter
	case MODE_RELATIVE:
		{
			return this.program[this.relativeBase+parameter]
		}
	default:
		panic(mode)
	}
}

func (this *VM) getAddress(argumentPosition int) int64 {
	fullOpcode := this.program[this.instructionIdx]
	parameter := this.program[this.instructionIdx+int64(argumentPosition)]

	mode := modeFor(fullOpcode, argumentPosition-1)

	switch mode {
	case MODE_POSITION:
		{
			return parameter
		}
	case MODE_IMMEDIATE:
		panic("writing to immediate address")
	case MODE_RELATIVE:
		{
			return this.relativeBase + parameter
		}
	default:
		panic(mode)
	}
}

func (this *VM) WriteInput(value int64) {
	this.input <- value
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
			param1 := this.readArgument(1)
			param2 := this.readArgument(2)
			this.program[this.getAddress(3)] = param1 + param2

			this.instructionIdx += 4
		}
	case MULT_OPCODE:
		{
			param1 := this.readArgument(1)
			param2 := this.readArgument(2)
			this.program[this.getAddress(3)] = param1 * param2

			this.instructionIdx += 4
		}
	case READ_INPUT_OPCODE:
		{
			address := this.getAddress(1)
			this.program[address] = <-this.input

			this.instructionIdx += 2
		}
	case WRITE_OUTPUT_OPCODE:
		{
			value := this.readArgument(1)
			this.output <- value

			this.instructionIdx += 2
		}
	case JUMP_IF_TRUE_OPCODE:
		{
			param1 := this.readArgument(1)
			param2 := this.readArgument(2)

			if param1 != 0 {
				this.instructionIdx = param2
			} else {
				this.instructionIdx += 3
			}
		}
	case JUMP_IF_FALSE_OPCODE:
		{
			param1 := this.readArgument(1)
			param2 := this.readArgument(2)

			if param1 == 0 {
				this.instructionIdx = param2
			} else {
				this.instructionIdx += 3
			}
		}
	case LESS_THAN_OPCODE:
		{
			param1 := this.readArgument(1)
			param2 := this.readArgument(2)

			if param1 < param2 {
				this.program[this.getAddress(3)] = 1
			} else {
				this.program[this.getAddress(3)] = 0
			}

			this.instructionIdx += 4
		}
	case EQUALS_OPCODE:
		{
			param1 := this.readArgument(1)
			param2 := this.readArgument(2)

			if param1 == param2 {
				this.program[this.getAddress(3)] = 1
			} else {
				this.program[this.getAddress(3)] = 0
			}

			this.instructionIdx += 4
		}
	case ADJUST_RELATIVE_BASE_OPCODE:
		{
			this.relativeBase += this.readArgument(1)
			this.instructionIdx += 2
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

func getInputData() Program {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	cmp := strings.Split(input, ",")
	program := make(Program)
	for idx, c := range cmp {
		num, _ := strconv.Atoi(c)
		program[int64(idx)] = int64(num)
	}
	return program
}

var directionDeltas = [][2]int{
	{-1, 0}, // up
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
}

const (
	DIR_UP = iota
	DIR_RIGHT
	DIR_DOWN
	DIR_LEFT
)

func paintPanels(field *map[[2]int]int64) {
	currentPosition := [2]int{0, 0}
	currentDirection := DIR_UP

	program := getInputData()

	vm := CreateVm(program)
	defer vm.Close()
	go vm.Run()

	for {
		currentColor := (*field)[currentPosition]
		vm.WriteInput(currentColor)

		colorToPaint, ok := <-vm.output
		if !ok {
			break
		}
		turnTo := <-vm.output

		(*field)[currentPosition] = colorToPaint

		if turnTo == 1 {
			currentDirection++
			if currentDirection > DIR_LEFT {
				currentDirection = DIR_UP
			}
		} else {
			currentDirection--
			if currentDirection < DIR_UP {
				currentDirection = DIR_LEFT
			}
		}

		delta := directionDeltas[currentDirection]
		currentPosition[0] += delta[0]
		currentPosition[1] += delta[1]
	}
}

func countPaintedPanels() int {
	field := make(map[[2]int]int64)
	paintPanels(&field)
	return len(field)
}

func drawPaintedPanels() {
	field := make(map[[2]int]int64)
	field[[2]int{0, 0}] = 1
	paintPanels(&field)

	for rowIdx := 0; rowIdx < 6; rowIdx++ {
		for colIdx := 0; colIdx < 40; colIdx++ {
			if color, ok := field[[2]int{rowIdx, colIdx}]; ok && color == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Part 2 solution is")
	drawPaintedPanels()
}
