package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getData() (int, int) {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	cmp := strings.Split(scanner.Text(), " players; last marble is worth ")

	players, _ := strconv.Atoi(cmp[0])
	points, _ := strconv.Atoi(strings.Replace(cmp[1], " points", "", 1))

	return players, points
}

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func insertAfter(node *Node, marble int) *Node {
	newNode := Node{Value: marble, Prev: node, Next: node.Next}

	newNode.Next = node.Next
	newNode.Prev = node
	node.Next.Prev = &newNode
	node.Next = &newNode

	return &newNode
}

func remove(node *Node) *Node {
	prev := node.Prev
	prev.Next = node.Next
	node.Next.Prev = prev

	return node.Next
}

func solveTask(players int, lastMarble int) int {
	current := &Node{Value: 0}
	current.Next = current
	current.Prev = current

	scores := []int{}
	for i := 0; i < players; i++ {
		scores = append(scores, 0)
	}

	currentPlayer := 0
	for marble := 1; marble <= lastMarble; marble++ {
		isScoringMarble := int(math.Mod(float64(marble), float64(23))) == 0

		if isScoringMarble {
			for i := 0; i < 7; i++ {
				current = current.Prev
			}

			scores[currentPlayer] += marble + current.Value
			current = remove(current)
		} else {
			for i := 0; i < 1; i++ {
				current = current.Next
			}

			current = insertAfter(current, marble)
		}

		currentPlayer++
		if currentPlayer == players {
			currentPlayer = 0
		}
	}

	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	players, points := getData()

	fmt.Println("Task 1 solution is", solveTask(players, points))
	fmt.Println("Task 2 solution is", solveTask(players, points*100))
}
