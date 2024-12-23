package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Graph = map[string](map[string]bool)

func getInputData() Graph {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := Graph{}

	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), "-")

		if _, ok := graph[cmp[0]]; !ok {
			graph[cmp[0]] = make(map[string]bool)
		}
		graph[cmp[0]][cmp[1]] = true

		if _, ok := graph[cmp[1]]; !ok {
			graph[cmp[1]] = make(map[string]bool)
		}
		graph[cmp[1]][cmp[0]] = true
	}

	return graph
}

func part1(graph Graph) int {
	res := make(map[[3]string]bool, 0)
	for c1, computer1connections := range graph {
		if c1[0] != 't' {
			continue
		}

		for c2 := range computer1connections {
			if c1 == c2 {
				continue
			}
			for c3 := range graph[c2] {
				if c2 == c3 {
					continue
				}
				if _, ok := graph[c3][c1]; ok {
					unsortedKey := []string{c1, c2, c3}
					sort.Strings(unsortedKey)

					key := [3]string{unsortedKey[0], unsortedKey[1], unsortedKey[2]}

					res[key] = true
				}
			}
		}
	}
	return len(res)
}

func part2(graph Graph) string {
	bestSet := []string{}
	for initialComputer := range graph {
		queue := []string{initialComputer}
		subgraph := map[string]bool{}
		visited := map[string]bool{}

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			for next := range graph[current] {
				if _, ok := visited[current]; !ok {
					queue = append(queue, next)
					visited[next] = true
				}
			}

			if len(subgraph) == 0 {
				subgraph[current] = true
			} else {
				canAdd := true
				for prev := range subgraph {
					if _, ok := graph[prev][current]; !ok {
						canAdd = false
						break
					}
				}

				if canAdd {
					subgraph[current] = true
				}
			}
		}

		candidate := []string{}
		for k := range subgraph {
			candidate = append(candidate, k)
		}
		sort.Strings(candidate)

		if len(candidate) > len(bestSet) {
			bestSet = candidate
		}
	}

	return strings.Join(bestSet, ",")
}

func main() {
	graph := getInputData()
	fmt.Println("Part 1 solution is", part1(graph))
	fmt.Println("Part 2 solution is", part2(graph))
}
