package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData() *[]int {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	data := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		stringValues := strings.Split(text, " ")

		for _, stringValue := range stringValues {
			value, _ := strconv.Atoi(stringValue)
			data = append(data, value)
		}

	}

	return &data
}

type Node struct {
	Children *[]Node
	Metadata *[]int
}

func buildNode(data *[]int, startIdx int) (*Node, int) {
	childrenCount := (*data)[startIdx]
	metaCount := (*data)[startIdx+1]

	children := []Node{}

	startIdx += 2

	for i := 0; i < childrenCount; i++ {
		child, newIndex := buildNode(data, startIdx)
		startIdx = newIndex
		children = append(children, *child)
	}

	metadata := []int{}
	for i := 0; i < metaCount; i++ {
		meta := (*data)[startIdx]
		metadata = append(metadata, meta)
		startIdx++
	}

	node := Node{
		Children: &children,
		Metadata: &metadata,
	}

	return &node, startIdx
}

func countMeta(node *Node) int {
	sum := 0

	for _, meta := range *node.Metadata {
		sum += meta
	}

	for _, child := range *node.Children {
		sum += countMeta(&child)
	}

	return sum
}

func solveTask1() int {
	data := getData()
	node, _ := buildNode(data, 0)
	return countMeta(node)
}

func countValue(node *Node) int {
	sum := 0

	if len(*node.Children) == 0 {
		for _, meta := range *node.Metadata {
			sum += meta
		}

		return sum
	}

	for _, meta := range *node.Metadata {
		children := *node.Children

		if meta > 0 && meta <= len(children) {
			child := children[meta-1]
			sum += countValue(&child)
		}
	}

	return sum
}

func solveTask2() int {
	data := getData()
	node, _ := buildNode(data, 0)
	return countValue(node)
}

func main() {
	fmt.Println("Task 1 solution is", solveTask1())
	fmt.Println("Task 2 solution is", solveTask2())
}
