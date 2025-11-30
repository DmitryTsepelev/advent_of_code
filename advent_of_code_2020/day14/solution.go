package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// =====

type Program struct {
	mask   [36]byte
	memory map[uint]([]byte)
}

func (this *Program) setMask(mask []byte) {
	for idx := 0; idx < 36; idx++ {
		val := mask[idx]
		this.mask[idx] = val
	}
}

func (this *Program) setValue(address uint, value uint) {
	binValue := fmt.Sprintf("%036s", strconv.FormatInt(int64(value), 2))

	result := make([]byte, 36)
	for idx := 0; idx < 36; idx++ {
		if this.mask[idx] == 'X' {
			result[idx] = binValue[idx]
		} else {
			result[idx] = this.mask[idx]
		}
	}

	this.memory[address] = result
}

func (this *Program) setValueV2(address uint, value uint) {
	binAddress := fmt.Sprintf("%036s", strconv.FormatInt(int64(address), 2))

	addresses := [][]byte{
		make([]byte, 36),
	}

	for idx := 0; idx < 36; idx++ {
		switch this.mask[idx] {
		case 'X':
			{
				newAddressMasks := make([][]byte, 0)
				for _, address := range addresses {
					withZero := make([]byte, 36)
					copy(withZero, address)
					withZero[idx] = '0'

					withOne := make([]byte, 36)
					copy(withOne, address)
					withOne[idx] = '1'

					newAddressMasks = append(newAddressMasks, withZero, withOne)
				}
				addresses = newAddressMasks
			}
		case '0':
			{
				for _, address := range addresses {
					address[idx] = binAddress[idx]
				}
			}
		case '1':
			{
				for _, address := range addresses {
					address[idx] = '1'
				}
			}
		}
	}

	binValue := []byte(fmt.Sprintf("%036s", strconv.FormatInt(int64(value), 2)))
	for _, address := range addresses {
		decimalAddress, _ := strconv.ParseInt(string(address), 2, 64)
		this.memory[uint(decimalAddress)] = binValue
	}
}

func (this *Program) sumMemory() int64 {
	var sum int64
	for _, value := range this.memory {
		decimalValue, _ := strconv.ParseInt(string(value), 2, 64)
		sum += decimalValue
	}
	return sum
}

// ====

func executeProgram(lines []string, useV2 bool) int64 {
	program := Program{mask: [36]byte{}, memory: make(map[uint]([]byte))}

	maskRegex := *regexp.MustCompile(`mask = (\w*)`)
	setRegex := *regexp.MustCompile(`mem\[(\w*)\] = (\w*)`)
	for _, line := range lines {

		if res := maskRegex.FindAllStringSubmatch(line, -1); len(res) > 0 {
			mask := res[0][1]
			program.setMask([]byte(mask))
		}

		if res := setRegex.FindAllStringSubmatch(line, -1); len(res) > 0 {
			address, _ := strconv.Atoi(res[0][1])
			value, _ := strconv.Atoi(res[0][2])
			if useV2 {
				program.setValueV2(uint(address), uint(value))
			} else {
				program.setValue(uint(address), uint(value))
			}
		}
	}

	return program.sumMemory()
}

func main() {
	lines := getInputData()

	fmt.Println("Part 1 solution is", executeProgram(lines, false))
	fmt.Println("Part 2 solution is", executeProgram(lines, true))
}
