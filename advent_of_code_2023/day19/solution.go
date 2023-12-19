package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	isGt         bool
	propertyName string
	value        int
}

type Rule struct {
	condition *Condition
	target    string
}

type Workflow = []Rule

type Workflows = map[string]Workflow

func parseWorkflow(line string) (string, Workflow) {
	cmp := strings.Split(line, "{")
	name := cmp[0]

	cmp = strings.Split(cmp[1][:len(cmp[1])-1], ",")

	workflow := Workflow{}

	for _, desc := range cmp {
		if strings.Index(desc, ":") == -1 {
			workflow = append(workflow, Rule{target: desc})
		} else {
			ruleCmp := strings.Split(desc, ":")

			value, _ := strconv.Atoi(ruleCmp[0][2:])
			condition := Condition{
				propertyName: string(ruleCmp[0][0]),
				isGt:         ruleCmp[0][1] == '>',
				value:        value,
			}

			workflow = append(workflow, Rule{target: ruleCmp[1], condition: &condition})
		}
	}

	return name, workflow
}

type Part = map[string]int

func parsePart(line string) Part {
	cmp := strings.Split(line[1:len(line)-1], ",")

	part := Part{}

	for _, desc := range cmp {
		propCmp := strings.Split(desc, "=")
		value, _ := strconv.Atoi(propCmp[1])
		part[propCmp[0]] = value
	}

	return part
}

func getInputData() (Workflows, []Part) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	parsingWorkflows := true
	workflows := Workflows{}
	parts := make([]Part, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			parsingWorkflows = false
			continue
		}

		if parsingWorkflows {
			name, workflow := parseWorkflow(line)
			workflows[name] = workflow
		} else {
			parts = append(parts, parsePart(line))
		}
	}

	return workflows, parts
}

// ------------------------------------------------------------------

func route(workflow Workflow, part Part) string {
	for _, rule := range workflow {
		if rule.condition == nil {
			return rule.target
		}

		propertyValue := part[rule.condition.propertyName]
		if rule.condition.isGt {
			if propertyValue > rule.condition.value {
				return rule.target
			}
		} else {
			if propertyValue < rule.condition.value {
				return rule.target
			}
		}
	}

	panic("impossible")
}

func partRating(part Part) int {
	rating := 0
	for _, v := range part {
		rating += v
	}
	return rating
}

func sumAcceptedRating(workflows Workflows, parts []Part) int {
	rating := 0
	for _, part := range parts {
		workflowName := "in"
		for workflowName != "R" && workflowName != "A" {
			workflow := workflows[workflowName]
			workflowName = route(workflow, part)
		}

		if workflowName == "A" {
			rating += partRating(part)
		}
	}
	return rating
}

// ------------------------------------------------------------------

type Ranges = map[string]([]int)

func copyRanges(ranges Ranges) Ranges {
	return Ranges{
		"x": []int{ranges["x"][0], ranges["x"][1]},
		"m": []int{ranges["m"][0], ranges["m"][1]},
		"a": []int{ranges["a"][0], ranges["a"][1]},
		"s": []int{ranges["s"][0], ranges["s"][1]},
	}
}

func splitRange(ranges Ranges, rule Rule) (Ranges, Ranges) {
	leftRanges := copyRanges(ranges)
	rightRanges := copyRanges(ranges)

	pName := rule.condition.propertyName
	if rule.condition.isGt {
		leftRanges[pName][0] = rule.condition.value + 1
		rightRanges[pName][1] = rule.condition.value
	} else {
		leftRanges[pName][1] = rule.condition.value - 1
		rightRanges[pName][0] = rule.condition.value
	}

	return leftRanges, rightRanges
}

func combinations(ranges Ranges) int {
	result := 1

	for _, rng := range ranges {
		result *= rng[1] - rng[0] + 1
	}

	return result
}

func traverse(workflowName string, workflows Workflows, ranges Ranges) int {
	if workflowName == "R" {
		return 0
	}

	if workflowName == "A" {
		return combinations(ranges)
	}

	workflow := workflows[workflowName]
	sum := 0
	currentRanges := ranges
	for _, rule := range workflow {
		if rule.condition == nil {
			sum += traverse(rule.target, workflows, currentRanges)
		} else {
			left, right := splitRange(currentRanges, rule)
			sum += traverse(rule.target, workflows, left)
			currentRanges = right
		}
	}

	return sum
}

func countAcceptedCombinations(workflows Workflows) int {
	ranges := Ranges{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}
	return traverse("in", workflows, ranges)
}

// ------------------------------------------------------------------

func main() {
	workflows, parts := getInputData()

	fmt.Println("Solution 1 is", sumAcceptedRating(workflows, parts))
	fmt.Println("Solution 2 is", countAcceptedCombinations(workflows))
}
