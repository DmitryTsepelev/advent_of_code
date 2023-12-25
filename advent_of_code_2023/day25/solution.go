package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

type Graph struct {
	links     map[string]([]string)
	nodeNames []string
}

func (graph *Graph) AddLink(node1, node2 string) {
	if _, found := graph.links[node1]; !found {
		graph.links[node1] = make([]string, 0)
	}
	graph.links[node1] = append(graph.links[node1], node2)

	if _, found := graph.links[node2]; !found {
		graph.links[node2] = make([]string, 0)
	}
	graph.links[node2] = append(graph.links[node2], node1)
}

func (graph *Graph) IsAdjacent(node1, node2 string) bool {
	if nodes, found := graph.links[node1]; found {
		for _, node := range nodes {
			if node == node2 {
				return true
			}
		}
	}

	if nodes, found := graph.links[node2]; found {
		for _, node := range nodes {
			if node == node1 {
				return true
			}
		}
	}

	return false
}

func (graph *Graph) ComponentsWithout(disconnectedPairs []Pair) []int {
	components := make([]int, 0)
	visited := make(map[string]bool, len(graph.nodeNames))

	for _, node := range graph.nodeNames {
		if visited[node] {
			continue
		}
		visited[node] = true

		currentComponent := 0
		queue := []string{node}

		for len(queue) > 0 {
			nextQueueMap := make(map[string]bool)

			for _, current := range queue {
				currentComponent++

				for _, next := range graph.links[current] {
					isDisconnected := false
					for _, pair := range disconnectedPairs {
						if pair.from == current && pair.to == next || pair.to == current && pair.from == next {
							isDisconnected = true
							break
						}
					}

					if isDisconnected || visited[next] {
						continue
					}

					nextQueueMap[next] = true
					visited[next] = true
				}
			}

			queue = make([]string, 0)
			for next := range nextQueueMap {
				queue = append(queue, next)
			}
		}

		components = append(components, currentComponent)
		currentComponent = 0
	}

	return components
}

func getInputData() Graph {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	graph := Graph{links: make(map[string][]string), nodeNames: make([]string, 0)}

	nodeNamesMap := make(map[string]bool)

	for scanner.Scan() {
		line := scanner.Text()

		cmp := strings.Split(line, " ")
		from := cmp[0][:len(cmp[0])-1]
		nodeNamesMap[from] = true

		for _, to := range cmp[1:] {
			nodeNamesMap[to] = true
			graph.AddLink(from, to)
		}
	}

	for nodeName := range nodeNamesMap {
		graph.nodeNames = append(graph.nodeNames, nodeName)
	}

	sort.Strings(graph.nodeNames)

	return graph
}

type Pair struct {
	from string
	to   string
}

func (graph *Graph) dfs(destination string, currentPath []string, visited map[string]bool) *[]string {
	lastVisited := currentPath[len(currentPath)-1]

	if destination == lastVisited {
		return &currentPath
	}

	for _, point := range graph.links[lastVisited] {
		if visited[point] {
			continue
		}

		visited[point] = true

		path := graph.dfs(destination, append(currentPath, point), visited)
		if path != nil {
			return path
		}

		visited[point] = false
	}

	return nil
}

func (graph *Graph) GetPath(start, end string) *[]string {
	// return *graph.dfs(to, []string{from}, make(map[string]bool))

	prev := map[string]string{start: start}
	nodes := []string{start}
	seen := map[string]bool{start: true}
	for len(nodes) > 0 {
		new_nodes := []string{}
		for _, node := range nodes {
			for _, neighbour := range graph.links[node] {
				if _, found := seen[neighbour]; found {
					continue
				}
				seen[neighbour] = true
				prev[neighbour] = node
				new_nodes = append(new_nodes, neighbour)
			}
		}
		nodes = new_nodes
	}

	if _, found := prev[end]; !found {
		return nil
	}

	path := []string{}
	node := end
	for node != start {
		path = append(path, node)
		node = prev[node]
	}
	path = append(path, start)

	return &path
}

func main() {
	graph := getInputData()

	edgeCounts := make(map[Pair]int, 0)
	for i := 0; i < 10000; i++ {

		from := graph.nodeNames[rand.Int31n(int32(len(graph.nodeNames)))]
		to := graph.nodeNames[rand.Int31n(int32(len(graph.nodeNames)))]
		if from == to {
			continue
		}

		path := graph.GetPath(from, to)
		if path == nil {
			continue
		}
		for j := 0; j < len(*path)-1; j++ {
			edge := []string{(*path)[j], (*path)[j+1]}
			sort.Strings(edge)

			pair := Pair{from: edge[0], to: edge[1]}
			if _, found := edgeCounts[pair]; !found {
				edgeCounts[pair] = 1
			} else {
				edgeCounts[pair]++
			}
		}
	}

	edges := make([]Pair, 0)
	for edge := range edgeCounts {
		edges = append(edges, edge)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edgeCounts[edges[i]] > edgeCounts[edges[j]]
	})

	cmp := graph.ComponentsWithout(edges[:3])

	fmt.Println("Solution is", cmp[0]*cmp[1])
}
