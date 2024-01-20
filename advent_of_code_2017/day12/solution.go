package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph = map[string](map[string]bool)

func getInputData() Graph {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := Graph{}

	for scanner.Scan() {
		line := scanner.Text()

		cmp := strings.Split(line, " <-> ")
		from := cmp[0]

		destinations := strings.Split(cmp[1], ", ")
		for _, to := range destinations {
			if _, ok := graph[from]; !ok {
				graph[from] = make(map[string]bool)
			}
			graph[from][to] = true

			if _, ok := graph[to]; !ok {
				graph[to] = make(map[string]bool)
			}
			graph[to][from] = true
		}
	}

	return graph
}

func visit(graph Graph, visited map[string]bool, group string) {
	queue := []string{group}

	for len(queue) > 0 {
		nextQueue := []string{}

		for _, node := range queue {
			for next := range graph[node] {
				if visited[next] {
					continue
				}
				visited[next] = true
				nextQueue = append(nextQueue, next)
			}
		}

		queue = nextQueue
	}
}

func findGroup(graph Graph, group string) int {
	visited := map[string]bool{}
	visit(graph, visited, group)
	return len(visited)
}

func countGroups(graph Graph) int {
	visited := map[string]bool{}

	groups := 0

	for len(visited) != len(graph) {
		for node := range graph {
			if visited[node] {
				continue
			}

			visit(graph, visited, node)
			groups++
		}
	}

	return groups
}

func main() {
	graph := getInputData()

	fmt.Println(findGroup(graph, "0"))

	fmt.Println(countGroups(graph))
}
