package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	return VM{program: program, relativeBase: 0, input: make(chan int64, 0), output: make(chan int64, 0)}
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
	fmt.Println("WriteInput", value)
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
			fmt.Println("got input", this.program[address])

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

// ===

func countBlocks() int {
	program := getInputData()

	vm := CreateVm(program)
	defer vm.Close()
	go vm.Run()

	var blocks int
	for {
		_, ok := <-vm.output
		if !ok {
			break
		}
		<-vm.output
		tileId := <-vm.output

		if tileId == 2 {
			blocks++
		}
	}

	return blocks
}

type ObjectType = int
type Point = [2]int

type GameState struct {
	score     int
	walls     map[Point]bool
	blocks    map[Point]bool
	prevBall  Point
	ball      Point
	paddle    Point
	vm        VM
	nextInput int

	width, height int
}

func (this *GameState) prepareInitialState() int {
	this.walls = make(map[Point]bool)
	this.blocks = make(map[Point]bool)

	program := getInputData()

	this.vm = CreateVm(program)
	defer this.vm.Close()
	go this.vm.Run()

	var blocks int
	for {
		x, ok := <-this.vm.output
		if !ok {
			break
		}
		y := <-this.vm.output
		tileId := <-this.vm.output

		point := Point{int(x), int(y)}
		this.width = max(this.width, int(x))
		this.height = max(this.height, int(y))

		switch tileId {
		case KIND_WALL:
			this.walls[point] = true
		case KIND_BLOCK:
			this.blocks[point] = true
		case KIND_PADDLE:
			this.paddle = point
		case KIND_BALL:
			this.ball = point
		}
	}

	return blocks
}

const (
	KIND_EMPTY = iota
	KIND_WALL
	KIND_BLOCK
	KIND_PADDLE
	KIND_BALL
)

func objectToDrawable(objectType ObjectType) string {
	// switch objectType {
	// case KIND_EMPTY:
	// 	return " "
	// case KIND_WALL:
	// 	return "▉"
	// case KIND_BLOCK:
	// 	return "▉"
	// case KIND_PADDLE:
	// 	return "▂"
	// case KIND_BALL:
	// 	return "◉"
	// }

	switch objectType {
	case KIND_EMPTY:
		return "⬛️"
	case KIND_WALL:
		return "🟥"
	case KIND_BLOCK:
		return "👾"
	case KIND_PADDLE:
		return "↔️"
	case KIND_BALL:
		return "🥎"
	}

	panic("unreachable")
}

func (this *GameState) drawScreen() {
	fmt.Print("\033[H\033[2J")

	for y := 0; y <= this.height; y++ {
		for x := 0; x <= this.width; x++ {
			objectType := KIND_EMPTY

			point := Point{x, y}
			if point == this.ball {
				objectType = KIND_BALL
			} else if point == this.paddle {
				objectType = KIND_PADDLE
			} else if _, ok := this.walls[point]; ok {
				objectType = KIND_WALL
			} else if _, ok := this.blocks[point]; ok {
				objectType = KIND_BLOCK
			}

			fmt.Print(objectToDrawable(objectType))
		}
		fmt.Println()
	}
}

const (
	INPUT_LEFT    = -1
	INPUT_NEUTRAL = 0
	INPUT_RIGHT   = 1
)

func (this *GameState) nextBallPosition() Point {
	if this.prevBall[1] == this.ball[1] {
		return this.prevBall
	}

	return Point{
		this.ball[0] + (this.ball[0] - this.prevBall[0]),
		this.ball[1] + (this.ball[1] - this.prevBall[1]),
	}
}

// func (this *GameState) updateNextInput() {
// 	hdir := this.ball[0] - this.paddle[0]
// 	this.nextInput = hdir
// }

func (this *GameState) chooseNextInput() int {
	if this.ball == this.prevBall {
		return INPUT_NEUTRAL
	}
	// goesRight := this.ball[0]-this.prevBall[0] > 0
	goesDown := this.ball[1]-this.prevBall[1] > 0
	// fmt.Println("goesRight", goesRight, "goesDown", goesDown)

	if goesDown {
		targetX := this.ball[0] + (this.ball[0]-this.prevBall[0])*(this.paddle[1]-this.ball[1])

		if targetX > this.paddle[0] {
			return INPUT_RIGHT
		} else if targetX > this.paddle[0] {
			return INPUT_LEFT
		}

		return INPUT_NEUTRAL
	} else {
		return this.ball[0] - this.prevBall[0]
	}

	// if this.paddle[0]-this.ball[0] > 0 {
	// 	return INPUT_LEFT
	// }
	// if this.paddle[0]-this.ball[0] < 0 {
	// 	return INPUT_RIGHT
	// }

	// return INPUT_NEUTRAL
}

func (this *GameState) playGame() int {
	program := getInputData()

	// Memory address 0 represents the number of quarters that have been inserted; set it to 2 to play for free
	program[0] = 2

	this.vm = CreateVm(program)
	defer this.vm.Close()
	go this.vm.Run()

	output := make([]int, 0)

	var blocks int
	for {
		select {
		case this.vm.input <- int64(this.chooseNextInput()):
			{
			}
		case outputVal := <-this.vm.output:
			{
				output = append(output, int(outputVal))

				if len(output) == 3 {
					x, y, tileId := output[0], output[1], output[2]

					skipRender := false

					if x == -1 && y == 0 {
						this.score = tileId
					} else {
						point := Point{int(x), int(y)}
						switch tileId {
						case KIND_PADDLE:
							{
								this.paddle = point
							}
						case KIND_BALL:
							{
								if this.prevBall != point {
									this.prevBall = this.ball
									this.ball = point

									// remove touched box
									nextBall := this.nextBallPosition()
									if _, ok := this.blocks[nextBall]; ok {
										delete(this.blocks, nextBall)
									}
								}
							}
						default:
							{
								skipRender = true
							}
						}
					}

					if skipRender == false {
						time.Sleep(250 * time.Millisecond)
						this.drawScreen()
					}

					output = make([]int, 0)
				}
			}
		}
	}

	return blocks
}

func main() {
	// fmt.Println("Part 1 solution is", countBlocks())
	// fmt.Println("Part 2 solution is", playGame())

	game := GameState{}
	game.prepareInitialState()
	game.playGame()
	fmt.Println(game.score)
}
