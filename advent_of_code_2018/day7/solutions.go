package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Row map[string]int
type Field map[string]Row

func getSteps() Field {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	restrictions := [][]string{}
	stepNames := []string{}

	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), " must be finished before step ")
		from := strings.Replace(cmp[0], "Step ", "", 1)
		to := strings.Replace(cmp[1], " can begin.", "", 1)

		restrictions = append(restrictions, []string{from, to})

		if !contains(stepNames, from) {
			stepNames = append(stepNames, from)
		}

		if !contains(stepNames, to) {
			stepNames = append(stepNames, to)
		}
	}

	field := make(Field)

	for _, i := range stepNames {
		row := make(Row)

		for _, j := range stepNames {
			row[j] = 0
		}

		field[i] = row
	}

	for _, restriction := range restrictions {
		from := restriction[0]
		to := restriction[1]

		field[from][to] = 1
	}

	return field
}

func getCandidates(field Field) []string {
	candidates := []string{}

	for candidate := range field {
		isCandidate := true

		for key, directions := range field {
			if candidate != key && directions[candidate] == 1 {
				isCandidate = false
			}
		}

		if isCandidate {
			candidates = append(candidates, candidate)
		}
	}
	return candidates
}

func solveTask1() string {
	field := getSteps()

	result := ""
	for {
		candidates := getCandidates(field)
		if len(candidates) == 0 {
			break
		}

		sort.Strings(candidates)
		nextStep := candidates[0]

		delete(field, nextStep)

		result += nextStep
	}

	return result
}

func secondsToComplete(step string) int {
	stepCmp := int(byte(step[0]) - 64)
	added := 60
	return stepCmp + added
}

type Worker struct {
	Job       string
	Remaining int
}

func solveTask2() int {
	field := getSteps()

	workers := []*Worker{
		&Worker{},
		&Worker{},
		&Worker{},
		&Worker{},
		&Worker{},
	}

	seconds := 0
	for {
		for _, worker := range workers {
			if worker.Job != "" {
				worker.Remaining--

				if worker.Remaining == 0 {
					delete(field, worker.Job)
					worker.Job = ""
				}
			}
		}

		candidates := getCandidates(field)
		noCandidates := len(candidates) == 0

		availableCandidates := []string{}
		for _, candidate := range candidates {
			isAvailable := true

			for _, worker := range workers {
				if worker.Job != "" && worker.Job == candidate {
					isAvailable = false
				}
			}

			if isAvailable {
				availableCandidates = append(availableCandidates, candidate)
			}
		}

		sort.Strings(availableCandidates)
		currentCandidateIdx := 0

		allWorkersFree := true
		for _, worker := range workers {
			if worker.Job == "" {
				if len(availableCandidates) > currentCandidateIdx {
					job := availableCandidates[currentCandidateIdx]
					worker.Job = job
					worker.Remaining = secondsToComplete(job)
					allWorkersFree = false
					currentCandidateIdx++
				}
			}
		}

		if noCandidates && allWorkersFree {
			return seconds
		}

		seconds++
	}
}

func main() {
	fmt.Println("Task 1 solution is", solveTask1())
	fmt.Println("Task 2 solution is", solveTask2())
}
