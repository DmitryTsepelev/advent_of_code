package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MODULE_KIND_BROADCAST  = "BROADCAST"
	MODULE_KIND_FLIP_FLOP  = "FLIP_FLOP"
	MODULE_KIND_CONJUCTION = "CONJUCTION"
)

type Wires = map[string]([]string)
type ModuleKinds = map[string]string

const BUTTON = "button"
const BROADCASTER = "broadcaster"

func getInputData() (Wires, ModuleKinds) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wires := Wires{}
	moduleKinds := ModuleKinds{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, BROADCASTER) {
			destinations := strings.Split(line[15:], ", ")
			wires[BROADCASTER] = destinations
			moduleKinds[BROADCASTER] = MODULE_KIND_BROADCAST
		} else {
			idx := 0

			moduleKind := MODULE_KIND_FLIP_FLOP
			if line[idx] == '&' {
				moduleKind = MODULE_KIND_CONJUCTION
			}

			idx++ // skip kind
			source := ""
			for line[idx] != ' ' {
				source += string(line[idx])
				idx++
			}

			moduleKinds[source] = moduleKind

			idx += 4 // skip ->

			destinations := strings.Split(line[idx:], ", ")
			wires[source] = destinations
		}
	}

	return wires, moduleKinds
}

const (
	SIGNAL_PULSE_LOW  = "LOW"
	SIGNAL_PULSE_HIGH = "HIGH"
)

type Signal struct {
	kind      string
	senderId  string
	handlerId string
}

type FlipFlopMemo = map[string]bool

func buildFlipFlopMemo(wires Wires, moduleKinds ModuleKinds) FlipFlopMemo {
	memo := make(FlipFlopMemo)

	for moduleId, kind := range moduleKinds {
		if kind == MODULE_KIND_FLIP_FLOP {
			memo[moduleId] = false
		}
	}

	return memo
}

type ConjuctionMemo = map[string](map[string]string)

func buildConjuctionMemo(wires Wires, moduleKinds ModuleKinds) ConjuctionMemo {
	memo := make(ConjuctionMemo)

	for moduleId, kind := range moduleKinds {
		if kind == MODULE_KIND_CONJUCTION {
			memo[moduleId] = make(map[string]string)
		}
	}

	for sourceId, destinations := range wires {
		for _, destinationId := range destinations {
			if mapping, ok := memo[destinationId]; ok {
				mapping[sourceId] = SIGNAL_PULSE_LOW
			}
		}
	}

	return memo
}

func simulate(wires Wires, moduleKinds ModuleKinds, cb func(int, Signal) bool) {
	flipFlopMemo := buildFlipFlopMemo(wires, moduleKinds)
	conjuctionMemo := buildConjuctionMemo(wires, moduleKinds)

	halt := false

	iteration := 0
	for {
		iteration++

		signalQueue := []Signal{
			{
				senderId:  BUTTON,
				handlerId: BROADCASTER,
				kind:      SIGNAL_PULSE_LOW,
			},
		}

		for len(signalQueue) > 0 {
			currentSignal := signalQueue[0]

			halt = cb(iteration, currentSignal)
			if halt {
				break
			}

			signalQueue = signalQueue[1:]

			destinations := wires[currentSignal.handlerId]

			var newKind string

			switch moduleKinds[currentSignal.handlerId] {
			case MODULE_KIND_BROADCAST:
				newKind = currentSignal.kind
			case MODULE_KIND_FLIP_FLOP:
				if currentSignal.kind == SIGNAL_PULSE_HIGH {
					// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
					continue
				}

				if flipFlopMemo[currentSignal.handlerId] {
					// was on
					newKind = SIGNAL_PULSE_LOW
					flipFlopMemo[currentSignal.handlerId] = false // turns on
				} else {
					// was off
					newKind = SIGNAL_PULSE_HIGH
					flipFlopMemo[currentSignal.handlerId] = true // turns on
				}
			case MODULE_KIND_CONJUCTION:
				conjuctionMemo[currentSignal.handlerId][currentSignal.senderId] = currentSignal.kind

				allHigh := true
				for _, rememberedValue := range conjuctionMemo[currentSignal.handlerId] {
					if rememberedValue == SIGNAL_PULSE_LOW {
						allHigh = false
						break
					}
				}

				newKind = SIGNAL_PULSE_HIGH
				if allHigh {
					newKind = SIGNAL_PULSE_LOW
				}
			}

			for _, destination := range destinations {
				signalQueue = append(signalQueue, Signal{
					handlerId: destination,
					senderId:  currentSignal.handlerId,
					kind:      newKind,
				})
			}
		}

		if halt {
			break
		}
	}
}

func solve2(wires Wires, moduleKinds ModuleKinds) int {
	conjuctionMemo := buildConjuctionMemo(wires, moduleKinds)

	memo := make(map[string]int, 0)
	for id := range conjuctionMemo["rm"] {
		memo[id] = 0
	}

	simulate(wires, moduleKinds, func(iteration int, currentSignal Signal) bool {
		valuesFound := true
		for key, value := range memo {
			if currentSignal.senderId == key && currentSignal.kind == SIGNAL_PULSE_HIGH && value == 0 {
				memo[key] = iteration
			}

			if memo[key] == 0 {
				valuesFound = false
			}
		}

		return valuesFound
	})

	result := 1
	for _, mult := range memo {
		result *= mult
	}

	return result
}

func solve1(wires Wires, moduleKinds ModuleKinds) int {
	lowPulses, highPulses := 0, 0

	simulate(wires, moduleKinds, func(iteration int, currentSignal Signal) bool {
		if iteration > 1000 {
			return true
		}
		switch currentSignal.kind {
		case SIGNAL_PULSE_LOW:
			lowPulses++
		case SIGNAL_PULSE_HIGH:
			highPulses++
		}

		return false
	})

	return lowPulses * highPulses
}

func main() {
	wires, moduleKinds := getInputData()

	fmt.Println("Solution 1 is", solve1(wires, moduleKinds))
	fmt.Println("Solution 2 is", solve2(wires, moduleKinds))
}
