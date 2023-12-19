package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type State = map[int]([]int)

const (
	BOT    = "bot"
	OUTPUT = "output"
)

type Instruction struct {
	lowKind  string
	lowId    int
	highKind string
	highId   int
}

type Transitions = map[int](Instruction)

func addToState(state State, id, value int) {
	if values, ok := state[id]; ok {
		state[id] = append(values, value)
	} else {
		state[id] = []int{value}
	}
}

func getInputData() (State, Transitions) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	initRe, _ := regexp.Compile(`value (\d+) goes to bot (\d+)`)
	giveRe, _ := regexp.Compile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)

	scanner := bufio.NewScanner(file)
	state := make(State)
	transitions := make(Transitions)

	for scanner.Scan() {
		line := scanner.Text()

		init := initRe.FindAllStringSubmatch(line, -1)

		if len(init) > 0 {
			valueS, botIdS := init[0][1], init[0][2]

			value, _ := strconv.Atoi(valueS)
			botId, _ := strconv.Atoi(botIdS)

			addToState(state, botId, value)

			continue
		}

		give := giveRe.FindAllStringSubmatch(line, -1)

		if len(give) > 0 {
			botIdS, lowKind, lowIdS, highKind, highIdS := give[0][1], give[0][2], give[0][3], give[0][4], give[0][5]

			botId, _ := strconv.Atoi(botIdS)
			lowId, _ := strconv.Atoi(lowIdS)
			highId, _ := strconv.Atoi(highIdS)

			transitions[botId] = Instruction{
				lowKind: lowKind, lowId: lowId,
				highKind: highKind, highId: highId,
			}

			continue
		}

		panic(line)
	}

	return state, transitions
}

func main() {
	state, transitions := getInputData()
	outputs := make(State)

	for {
		stateIsEmpty := true
		for _, chips := range state {
			if len(chips) > 0 {
				stateIsEmpty = false
			}
		}
		if stateIsEmpty {
			break
		}

		for botId, chips := range state {
			if len(chips) < 2 {
				continue
			}

			lowChip, highChip := chips[0], chips[1]
			if lowChip > highChip {
				lowChip, highChip = chips[1], chips[0]
			}

			if lowChip == 17 && highChip == 61 {
				fmt.Println("Solution 1 is", botId)
			}

			transition := transitions[botId]
			if transition.lowKind == BOT {
				addToState(state, transition.lowId, lowChip)
			} else {
				addToState(outputs, transition.lowId, lowChip)
			}

			if transition.highKind == BOT {
				addToState(state, transition.highId, highChip)
			} else {
				addToState(outputs, transition.highId, highChip)
			}

			state[botId] = []int{}
		}
	}

	mult := 1
	for i := 0; i <= 2; i++ {
		chips := outputs[i]
		for _, chip := range chips {
			mult *= chip
		}
	}

	fmt.Println("Solution 2 is", mult)
}
