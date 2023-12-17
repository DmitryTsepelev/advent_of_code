package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getInputData() [][]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	grid := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		gridRow := make([]int, len(line))
		for idx, c := range line {
			loss, _ := strconv.Atoi(string(c))
			gridRow[idx] = loss
		}

		grid = append(grid, gridRow)
	}

	return grid
}

const (
	INIT  = "INIT"
	RIGHT = "RIGHT"
	LEFT  = "LEFT"
	UP    = "UP"
	DOWN  = "DOWN"
)

type Node struct {
	row       int
	col       int
	distance  int
	direction string
}

type PQ = []Node

func swapPQ(pq *PQ, i int, j int) {
	tmp := (*pq)[i]

	(*pq)[i] = (*pq)[j]
	(*pq)[j] = tmp
}

func insertPQ(pq *PQ, el Node) {
	*pq = append(*pq, el)

	idx := len(*pq) - 1

	for {
		if idx == 0 {
			break
		}

		parentIdx := (idx - 1) / 2

		if (*pq)[parentIdx].distance > (*pq)[idx].distance {
			swapPQ(pq, idx, parentIdx)
		} else {
			break
		}

		idx = parentIdx
	}
}

func heapifyPQ(pq *PQ, n int, i int) {
	largest := i // Initialize largest as root
	l := 2*i + 1 // left = 2*i + 1
	r := 2*i + 2 // right = 2*i + 2

	// If left child is larger than root
	if l < n && (*pq)[l].distance < (*pq)[largest].distance {
		largest = l
	}

	// If right child is larger than largest so far
	if r < n && (*pq)[r].distance < (*pq)[largest].distance {
		largest = r
	}

	// If largest is not root
	if largest != i {
		swapPQ(pq, i, largest)

		// Recursively heapify the affected sub-tree
		heapifyPQ(pq, n, largest)
	}
}

func removePQ(pq *PQ) Node {
	v := (*pq)[0]

	swapPQ(pq, 0, len(*pq)-1)
	*pq = (*pq)[:len(*pq)-1]

	heapifyPQ(pq, len(*pq), 0)

	return v
}

// ----

var deltas = map[string]([]int){
	RIGHT: {0, 1},
	LEFT:  {0, -1},
	UP:    {-1, 0},
	DOWN:  {1, 0},
}

var validDirections = map[string]([]string){
	INIT:  {RIGHT, LEFT, UP, DOWN},
	RIGHT: {UP, DOWN},
	LEFT:  {UP, DOWN},
	UP:    {RIGHT, LEFT},
	DOWN:  {RIGHT, LEFT},
}

type Key struct {
	row       int
	col       int
	direction string
}

func dijkstra(graph [][]int, minMoves int, maxMoves int) int {
	distances := make(map[Key]int, len(graph))
	pq := PQ{}

	v := Node{row: 0, col: 0, distance: 0, direction: INIT}
	insertPQ(&pq, v)

	for len(pq) > 0 {
		current := removePQ(&pq)

		for _, validDirection := range validDirections[current.direction] {
			delta := deltas[validDirection]
			prevDistance := current.distance
			for shift := 1; shift <= maxMoves; shift++ {
				nextRow := current.row + shift*delta[0]
				nextCol := current.col + shift*delta[1]

				if nextRow < 0 || nextRow > len(graph)-1 || nextCol < 0 || nextCol > len(graph[0])-1 {
					continue
				}

				newNode := Node{
					row:       nextRow,
					col:       nextCol,
					direction: validDirection,
					distance:  prevDistance + graph[nextRow][nextCol],
				}

				prevDistance = newNode.distance
				if shift < minMoves {
					continue
				}

				key := Key{row: nextRow, col: nextCol, direction: validDirection}
				if _, ok := distances[key]; !ok {
					distances[key] = math.MaxInt32
				}
				if newNode.distance < distances[key] {
					distances[key] = newNode.distance
					insertPQ(&pq, newNode)
				}
			}
		}
	}

	minDistance := math.MaxInt32

	for key, distance := range distances {
		if key.row == len(graph)-1 && key.col == len(graph[0])-1 && minDistance > distance {
			minDistance = distance
		}
	}

	return minDistance
}

func main() {
	grid := getInputData()
	fmt.Println("Solution 1 is", dijkstra(grid, 1, 3))
	fmt.Println("Solution 2 is", dijkstra(grid, 4, 10))
}
