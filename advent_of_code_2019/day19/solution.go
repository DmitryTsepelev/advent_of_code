package main

import (
	"bufio"
	"fmt"
	"math"
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
	waitForInput   chan interface{}
}

func CreateVm(program Program) VM {
	return VM{program: program, relativeBase: 0, input: make(chan int64, 0), output: make(chan int64, 0), waitForInput: make(chan interface{}, 0)}
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
			this.waitForInput <- 1
			address := this.getAddress(1)
			this.program[address] = <-this.input
			// fmt.Println("got input", this.program[address])

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

// ---

type Point struct {
	x, y int
}

func draw(field map[Point]int) {
	fmt.Print("\033[H\033[2J")
	minX, maxX, minY, maxY := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt

	for point := range field {
		minX = min(minX, point.x)
		minY = min(minY, point.y)
		maxX = max(maxX, point.x)
		maxY = max(maxY, point.y)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			point := Point{x, y}

			if val, found := field[point]; found {
				sVal := "."
				if val == 1 {
					sVal = "#"
				}
				fmt.Print(sVal)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func checkPointCovered(point Point) bool {
	program := getInputData()
	vm := CreateVm(program)
	defer vm.Close()
	go vm.Run()

	<-vm.waitForInput
	vm.input <- int64(point.x)

	<-vm.waitForInput
	vm.input <- int64(point.y)

	result := <-vm.output

	return int(result) == 1
}

func part1() int {
	count := 0
	field := map[Point]int{}

	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			point := Point{x, y}
			if checkPointCovered(point) {
				count++
				field[point] = 1
			} else {
				field[point] = 0
			}
		}
	}

	draw(field)

	return count
}

func getLineBoundaries(y int) (int, int) {
	l, r := -1, -1

	for x := 0; ; x++ {
		program := getInputData()
		vm := CreateVm(program)
		defer vm.Close()
		go vm.Run()

		<-vm.waitForInput
		vm.input <- int64(x)

		<-vm.waitForInput
		vm.input <- int64(y)

		result := <-vm.output

		if result == 0 {
			if l >= 0 {
				return l, r
			}
		} else {
			if l == -1 {
				l = x
			} else {
				r = x
			}
		}
	}
}

func findMaxAreaCorner(width int) Point {
	leftEdgeX := 0
	for y := 0; ; y++ {
		for checkPointCovered(Point{leftEdgeX, y}) == false {
			leftEdgeX++
		}
		for x := leftEdgeX; checkPointCovered(Point{x, y}); x++ {
			if checkPointCovered(Point{x + width - 1, y}) &&
				checkPointCovered(Point{x, y + width - 1}) &&
				checkPointCovered(Point{x + width - 1, y + width - 1}) {
				return Point{x, y}
			}
		}
	}
}

func part2() int {
	width := 100
	point := findMaxAreaCorner(width)
	return point.x*10000 + point.y
}

func main() {
	fmt.Println("Part 1 solution is", part1())
	fmt.Println("Part 2 solution is", part2())
}
