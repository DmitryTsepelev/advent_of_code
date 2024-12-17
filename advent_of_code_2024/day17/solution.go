package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() ([3]int, []byte) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	registers := [3]int{}
	program := []byte{}

	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		if rowIdx < 3 {
			val, _ := strconv.Atoi(strings.Split(scanner.Text(), ": ")[1])
			registers[rowIdx] = val
		}

		if rowIdx == 4 {
			for _, snum := range strings.Split(strings.Split(scanner.Text(), ": ")[1], ",") {
				val, _ := strconv.Atoi(snum)
				program = append(program, byte(val))
			}
		}
	}

	return registers, program
}

const REG_A = 0
const REG_B = 1
const REG_C = 2

func pow2(n int) int {
	result := 1

	for i := 0; i < n; i++ {
		result *= 2
	}

	return result
}

type VM struct {
	registers   [3]int
	output      []byte
	iptr        int
	skipIptrInc bool
}

const OPCODE_ADV = 0
const OPCODE_BXL = 1
const OPCODE_BST = 2
const OPCODE_JNZ = 3
const OPCODE_BXC = 4
const OPCODE_OUT = 5
const OPCODE_BDV = 6
const OPCODE_CDV = 7

func (this *VM) execute(program []byte) {
	router := map[int](func(byte)){
		OPCODE_ADV: this.adv,
		OPCODE_BXL: this.bxl,
		OPCODE_BST: this.bst,
		OPCODE_JNZ: this.jnz,
		OPCODE_BXC: this.bxc,
		OPCODE_OUT: this.out,
		OPCODE_BDV: this.bdv,
		OPCODE_CDV: this.cdv,
	}

	for this.iptr < len(program) {
		opcode, operand := program[this.iptr], program[this.iptr+1]

		router[int(opcode)](operand)

		if this.skipIptrInc == false {
			this.iptr += 2
		}

		this.skipIptrInc = false
	}
}

func (this *VM) readCombo(operand byte) int {
	if operand <= 3 {
		return int(operand)
	}

	return this.registers[operand-4]
}

func (this *VM) adv(operand byte) {
	combo := this.readCombo(operand)
	this.registers[REG_A] /= pow2(combo)
}

func (this *VM) bxl(operand byte) {
	this.registers[REG_B] ^= int(operand)
}

func (this *VM) bst(operand byte) {
	this.registers[REG_B] = this.readCombo(operand) % 8
}

func (this *VM) jnz(operand byte) {
	if this.registers[REG_A] == 0 {
		return
	}

	this.skipIptrInc = true
	this.iptr = int(operand)
}

func (this *VM) bxc(operand byte) {
	this.registers[REG_B] ^= this.registers[REG_C]
}

func (this *VM) out(operand byte) {
	this.output = append(this.output, byte(this.readCombo(operand)%8))
}

func (this *VM) bdv(operand byte) {
	combo := this.readCombo(operand)
	this.registers[REG_B] = this.registers[REG_A] / pow2(combo)
}

func (this *VM) cdv(operand byte) {
	combo := this.readCombo(operand)
	this.registers[REG_C] = this.registers[REG_A] / pow2(combo)
}

func byteSliceToString(s []byte) string {
	output := strconv.Itoa(int(s[0]))
	for i := 1; i < len(s); i++ {
		output += "," + strconv.Itoa(int(s[i]))
	}

	return output
}

var NONE = -1

func sameSlice(s1 []byte, s2 []byte) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, s1n := range s1 {
		if s1n != s2[i] {
			return false
		}
	}

	return true
}

func findBestInput(program []byte, cursor int, mult int) int {
	for candidate := 0; candidate < 8; candidate++ {
		vm := VM{registers: [3]int{mult*8 + candidate, 0, 0}, output: []byte{}}
		vm.execute(program)

		if sameSlice(vm.output, program[cursor:]) {
			if cursor == 0 {
				return mult*8 + candidate
			}

			if result := findBestInput(program, cursor-1, mult*8+candidate); result != NONE {
				return result
			}
		}
	}

	return NONE
}

func executeProgram(vm *VM, program []byte) string {
	vm.execute(program)
	return byteSliceToString(vm.output)
}

func main() {
	registers, program := getInputData()

	vm := VM{registers: registers, output: []byte{}}
	fmt.Println("Part 1 solution is", executeProgram(&vm, program))

	fmt.Println("Part 2 solution is", findBestInput(program, len(program)-1, 0))
}
