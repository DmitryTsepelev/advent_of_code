package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Priority queue

type Item struct {
	rowIdx int
	colIdx int
	dir    int
	score  int
	path   *[]Point
	index  int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// --------------

type Point struct {
	rowIdx int
	colIdx int
}

type Field struct {
	startPoint Point
	endPoint   Point
	wallMap    map[Point]bool
}

func getInputData() Field {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := Field{wallMap: make(map[Point]bool)}

	for rowIdx := 0; scanner.Scan(); rowIdx++ {
		line := scanner.Text()

		for colIdx, val := range line {
			point := Point{rowIdx, colIdx}

			switch val {
			case '#':
				field.wallMap[point] = true
			case 'S':
				field.startPoint = point
			case 'E':
				field.endPoint = point
			}
		}
	}

	return field
}

var UP = Point{1, 0}
var DOWN = Point{-1, 0}
var RIGHT = Point{0, 1}
var LEFT = Point{0, -1}

var directions = []Point{DOWN, RIGHT, UP, LEFT}

func getScore(field Field) (int, [][]Point) {
	var resultScore *int

	path := []Point{{field.startPoint.rowIdx, field.startPoint.colIdx}}
	queue := PriorityQueue{
		&Item{rowIdx: field.startPoint.rowIdx, colIdx: field.startPoint.colIdx, dir: 1, score: 0, path: &path},
	}
	heap.Init(&queue)

	paths := [][]Point{}

	visited := make(map[[3]int]int)

	for len(queue) > 0 {
		item := heap.Pop(&queue).(*Item)

		key := [3]int{item.rowIdx, item.colIdx, item.dir}
		if prevScore, ok := visited[key]; ok && prevScore < item.score {
			continue
		}
		visited[key] = item.score

		if item.rowIdx == field.endPoint.rowIdx && item.colIdx == field.endPoint.colIdx {
			if resultScore == nil {
				resultScore = &item.score
			}

			paths = append(paths, *item.path)
		}

		nextRowIdx := item.rowIdx + directions[item.dir].rowIdx
		nextColIdx := item.colIdx + directions[item.dir].colIdx

		if _, ok := field.wallMap[Point{nextRowIdx, nextColIdx}]; !ok {
			path := append([]Point{}, *item.path...)
			path = append(path, Point{nextRowIdx, nextColIdx})

			heap.Push(&queue, &Item{rowIdx: nextRowIdx, colIdx: nextColIdx, dir: item.dir, score: item.score + 1, path: &path})
		}

		path := append([]Point{}, *item.path...)
		heap.Push(&queue, &Item{rowIdx: item.rowIdx, colIdx: item.colIdx, dir: (item.dir + 1) % 4, score: item.score + 1000, path: &path})

		path = append([]Point{}, *item.path...)
		heap.Push(&queue, &Item{rowIdx: item.rowIdx, colIdx: item.colIdx, dir: (item.dir + 3) % 4, score: item.score + 1000, path: &path})
	}

	return *resultScore, paths
}

func main() {
	field := getInputData()

	score, paths := getScore(field)
	fmt.Println("Part 1 solution is", score)

	uniqPoints := make(map[Point]bool)

	for _, path := range paths {
		for _, point := range path {
			uniqPoints[point] = true
		}
	}

	fmt.Println("Part 2 solution is", len(uniqPoints))
}
