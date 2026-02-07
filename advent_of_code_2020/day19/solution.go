package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputData(overrideRules bool) (map[int]Rule, []string) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int]Rule, 0)
	messages := make([]string, 0)

	parsingMessages := false

	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			parsingMessages = true
		} else if parsingMessages {
			messages = append(messages, text)
		} else {
			cmp := strings.Split(text, ": ")

			id, _ := strconv.Atoi(cmp[0])

			if overrideRules {
				switch id {
				case 8:
					cmp[1] = "42 | 42 8"
				case 11:
					cmp[1] = "42 31 | 42 11 31"
				}
			}

			if cmp[1][0] == '"' {
				rules[id] = CharRule{cmp[1][1]}
			} else {
				cmp = strings.Split(cmp[1], " | ")
				sequenceRules := make([]SequenceRule, 0)

				for _, data := range cmp {
					cmp = strings.Split(data, " ")
					ids := make([]int, 0)
					for _, sid := range cmp {
						id, _ := strconv.Atoi(sid)
						ids = append(ids, id)
					}

					sequenceRules = append(sequenceRules, SequenceRule{ids})
				}

				if len(sequenceRules) == 1 {
					rules[id] = sequenceRules[0]
				} else {
					rules[id] = OrRule{sequenceRules}
				}
			}
		}
	}

	return rules, messages
}

type Rule interface {
	Match(string, int, map[int]Rule) (bool, int)
}

type CharRule struct {
	Value byte
}

func (this CharRule) Match(message string, messageIdx int, rules map[int]Rule) (bool, int) {
	return this.Value == message[messageIdx], messageIdx + 1
}

type SequenceRule struct {
	RuleIds []int
}

func (this SequenceRule) Match(message string, messageIdx int, rules map[int]Rule) (bool, int) {
	hasMatch := true

	for _, ruleId := range this.RuleIds {
		hasMatch, messageIdx = rules[ruleId].Match(message, messageIdx, rules)
		if hasMatch == false {
			return hasMatch, messageIdx
		}
	}

	return hasMatch, messageIdx
}

type OrRule struct {
	Rules []SequenceRule
}

func (this OrRule) Match(message string, messageIdx int, rules map[int]Rule) (bool, int) {
	for _, rule := range this.Rules {
		hasMatch, endIdx := rule.Match(message, messageIdx, rules)
		if hasMatch {
			return hasMatch, endIdx
		}
	}

	return false, messageIdx
}

func matchRule(message string, rules map[int]Rule) bool {
	hasMatch, endIdx := rules[0].Match(message, 0, rules)

	return hasMatch && endIdx == len(message)
}

func countMatches(overrideRules bool) int {
	rules, messages := getInputData(overrideRules)

	count := 0
	for _, message := range messages {
		if matchRule(message, rules) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("Part 1 solution is", countMatches(false))
	fmt.Println("Part 2 solution is", countMatches(true)) // not workin
}
