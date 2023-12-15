package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	cmp := strings.Split(scanner.Text(), ",")

	return cmp
}

func calcHash(line string) int {
	currentValue := 0

	for _, r := range line {
		currentValue += int(r)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}

func sumHashes(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += calcHash(line)
	}
	return sum
}

const (
	REMOVE         = '-'
	Add_OR_REPLACE = '='
)

type Operation struct {
	label       string
	boxIdx      int
	kind        rune
	focalLength int
}

func parseOperation(line string) Operation {
	operation := Operation{}

	for idx, r := range line {
		if r == '-' || r == '=' {
			operation.boxIdx = calcHash(operation.label)
			operation.kind = r
			if operation.kind == Add_OR_REPLACE {
				focalLength, _ := strconv.Atoi(line[idx+1:])
				operation.focalLength = focalLength
			}
			break
		}

		operation.label += string(r)
	}

	return operation
}

type Node struct {
	label       string
	focalLength int
	next        *Node
}

type Boxes = [](*Node)

func addLens(boxes Boxes, boxIdx int, label string, focalLength int) {
	current := boxes[boxIdx]

	lens := &Node{label: label, focalLength: focalLength}
	if current == nil {
		boxes[boxIdx] = lens
		return
	}

	for {
		if current.label == label {
			// replacing existing lens
			current.focalLength = focalLength
			break
		}

		if current.next == nil {
			// adding to the end
			current.next = lens
			break
		}

		current = current.next
	}
}

func removeLens(boxes Boxes, boxIdx int, label string) {
	root := boxes[boxIdx]
	if root == nil {
		return
	}

	if root.label == label {
		// remove head of the list
		boxes[boxIdx] = root.next
		return
	}

	prev := root
	current := root.next
	for {
		if current == nil {
			break
		}

		if current.label == label {
			prev.next = current.next
			break
		}

		prev = current
		current = current.next
	}
}

func sumFocusingPower(lines []string) int {
	boxes := make(Boxes, 256)

	for _, line := range lines {
		operation := parseOperation(line)

		switch operation.kind {
		case Add_OR_REPLACE:
			addLens(boxes, operation.boxIdx, operation.label, operation.focalLength)
		case REMOVE:
			removeLens(boxes, operation.boxIdx, operation.label)
		}
	}

	sum := 0
	for boxNumber, box := range boxes {
		current := box
		slotNumber := 1
		for current != nil {
			sum += (boxNumber + 1) * slotNumber * current.focalLength
			current = current.next
			slotNumber++
		}
	}

	return sum
}

func main() {
	lines := getInputData()

	fmt.Println("Solution 1 is", sumHashes(lines))
	fmt.Println("Solution 2 is", sumFocusingPower(lines))
}
