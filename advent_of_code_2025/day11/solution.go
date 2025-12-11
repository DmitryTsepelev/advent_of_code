package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph = map[string]([]string)

func getInputData() Graph {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := Graph{}
	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), ": ")

		from := cmp[0]
		to := strings.Split(cmp[1], " ")

		graph[from] = to
	}

	return graph
}

func countPaths(source, destination string, graph Graph) int {
	allNodes := make(map[string]bool)
	indegree := make(map[string]int)

	for source, destinations := range graph {
		allNodes[source] = true
		for _, dest := range destinations {
			indegree[dest]++
			allNodes[dest] = true
		}
	}

	// // Perform topological sort using Kahn's algorithm
	q := make([]string, 0)
	for node := range allNodes {
		if indegree[node] == 0 {
			q = append(q, node)
		}
	}

	topoOrder := make([]string, 0)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		topoOrder = append(topoOrder, node)

		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	// Array to store number of ways to reach each node
	ways := make(map[string]int)
	ways[source] = 1

	// Traverse in topological order
	for _, node := range topoOrder {
		for _, neighbor := range graph[node] {
			ways[neighbor] += ways[node]
		}
	}

	return ways[destination]
}

func main() {
	graph := getInputData()

	fmt.Println("Part 1 solution is", countPaths("you", "out", graph))
	result := countPaths("svr", "dac", graph)*countPaths("dac", "fft", graph)*countPaths("fft", "out", graph) +
		countPaths("svr", "fft", graph)*countPaths("fft", "dac", graph)*countPaths("dac", "out", graph)
	fmt.Println("Part 2 solution is", result)
}
