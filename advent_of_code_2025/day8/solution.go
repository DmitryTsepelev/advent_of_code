package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// An Item is something we manage in a priority queue.
type Item struct {
	p1       int
	p2       int
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
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

// // update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *Item, p1, p2 int, priority int) {
// 	item.p1 = p1
// 	item.p2 = p2
// 	item.priority = priority
// 	heap.Fix(pq, item.index)
// }

// =====

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func newUnionFind(numOfElements int) *UnionFind {
	// makeSet
	parent := make([]int, numOfElements)
	rank := make([]int, numOfElements)
	for i := 0; i < numOfElements; i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
		count:  numOfElements,
	}
}

// Time: O(logn) | Space: O(1)
func (uf *UnionFind) find(node int) int {
	for node != uf.parent[node] {
		// path compression
		uf.parent[node] = uf.parent[uf.parent[node]]
		node = uf.parent[node]
	}
	return node
}

// Time: O(1) | Space: O(1)
func (uf *UnionFind) union(node1, node2 int) {
	root1 := uf.find(node1)
	root2 := uf.find(node2)

	// already in the same set
	if root1 == root2 {
		return
	}

	if uf.rank[root1] > uf.rank[root2] {
		uf.parent[root2] = root1
	} else if uf.rank[root1] < uf.rank[root2] {
		uf.parent[root1] = root2
	} else {
		uf.parent[root2] = root1
		uf.rank[root2] += 1
	}
}

// func (uf *UnionFind) reset(node int) {
// 	uf.parent[node] = node
// 	uf.rank[node] = 0
// }

// func (uf *UnionFind) connected(node1 int, node2 int) bool {
// 	return uf.find(node1) == uf.find(node2)
// }

//----

func getInputData() []Point {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := []Point{}
	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), ",")

		x, _ := strconv.Atoi(cmp[0])
		y, _ := strconv.Atoi(cmp[1])
		z, _ := strconv.Atoi(cmp[2])

		points = append(points, Point{x, y, z})
	}

	return points
}

type Point struct {
	x, y, z int
}

func distance(p1, p2 Point) int {
	return int(math.Sqrt(float64((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y) + (p2.z-p1.z)*(p2.z-p1.z))))
}

func solve1(points []Point) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			heap.Push(pq, &Item{
				p1:       i,
				p2:       j,
				priority: distance(points[i], points[j]),
			})
		}
	}

	uf := newUnionFind(len(points))
	for i := 0; i < 1000; i++ {
		top := heap.Pop(pq).(*Item)
		uf.union(top.p1, top.p2)
	}

	sets := map[int]int{}

	for _, e := range uf.parent {
		sets[uf.find(e)]++
	}

	counts := []int{}
	for _, count := range sets {
		counts = append(counts, count)
	}

	sort.Ints(counts)
	mult := 1
	for i := len(counts) - 1; i >= len(counts)-3; i-- {
		mult *= counts[i]
	}

	return mult
}

func solve2(points []Point) int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			heap.Push(pq, &Item{
				p1:       i,
				p2:       j,
				priority: distance(points[i], points[j]),
			})
		}
	}

	uf := newUnionFind(len(points))
	var p1, p2 int
	for len(*pq) > 0 {
		top := heap.Pop(pq).(*Item)
		p1, p2 = top.p1, top.p2
		uf.union(top.p1, top.p2)

		sets := map[int]bool{}

		for _, e := range uf.parent {
			sets[uf.find(e)] = true
		}

		if len(sets) == 1 {
			return points[p1].x * points[p2].x
		}
	}

	panic("unreachable")
}

func main() {
	points := getInputData()

	fmt.Println("Part 1 solution is", solve1(points))
	fmt.Println("Part 2 solution is", solve2(points))
}
