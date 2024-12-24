package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const OPERATION_AND = 0
const OPERATION_OR = 1
const OPERATION_XOR = 2

var textToOp = map[string]int{
	"AND": OPERATION_AND,
	"OR":  OPERATION_OR,
	"XOR": OPERATION_XOR,
}

type Gate struct {
	in1       string
	in2       string
	outName   string
	operation int
}

type Inputs = map[string]bool
type Gates = map[string]Gate

func getInputData() (Inputs, Gates) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := Inputs{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}

		cmp := strings.Split(scanner.Text(), ": ")
		inputs[cmp[0]] = cmp[1] == "1"
	}

	gates := Gates{}
	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), " -> ")

		outName := cmp[1]

		cmp = strings.Split(cmp[0], " ")

		gate := Gate{
			in1:       cmp[0],
			in2:       cmp[2],
			operation: textToOp[cmp[1]],
			outName:   outName,
		}

		gates[gate.outName] = gate
	}

	return inputs, gates
}

func evalGate(gateKey string, gates Gates, inputs Inputs) bool {
	gate := gates[gateKey]

	if _, ok := inputs[gateKey]; !ok {
		if _, ok := inputs[gate.in1]; !ok {
			evalGate(gate.in1, gates, inputs)
		}
		if _, ok := inputs[gate.in2]; !ok {
			evalGate(gate.in2, gates, inputs)
		}

		switch gate.operation {
		case OPERATION_AND:
			inputs[gateKey] = inputs[gate.in1] && inputs[gate.in2]
		case OPERATION_OR:
			inputs[gateKey] = inputs[gate.in1] || inputs[gate.in2]
		case OPERATION_XOR:
			inputs[gateKey] = inputs[gate.in1] && !inputs[gate.in2] || !inputs[gate.in1] && inputs[gate.in2]
		}
	}

	return inputs[gateKey]
}

func part1(inputs Inputs, gates Gates) int64 {
	result := ""
	for i := 0; ; i++ {
		gateOut := "z" + strconv.Itoa(i)
		if i < 10 {
			gateOut = "z0" + strconv.Itoa(i)
		}

		if _, ok := gates[gateOut]; !ok {
			break
		}

		if evalGate(gateOut, gates, inputs) {
			result = "1" + result
		} else {
			result = "0" + result
		}
	}

	decResult, _ := strconv.ParseInt(result, 2, 64)

	return decResult
}

func main() {
	inputs, gates := getInputData()

	fmt.Println("Part 1 solution is", part1(inputs, gates))
}
