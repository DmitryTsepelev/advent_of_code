package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

type Machine struct {
	targetLights  []bool
	targetJoltage []int
	buttons       []([]int)
}

func getInputData() []Machine {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	machines := []Machine{}
	for scanner.Scan() {
		machine := Machine{}

		line := scanner.Text()

		cmp := strings.Split(line, " ")

		machine.buttons = [][]int{}

		for _, c := range cmp {
			if c[0] == '[' {
				machine.targetLights = []bool{}
				for i := 1; i < len(c)-1; i++ {
					machine.targetLights = append(machine.targetLights, c[i] == '#')
				}
			}

			if c[0] == '(' {
				buttons := []int{}
				sButtons := strings.Split(c[1:len(c)-1], ",")
				for _, b := range sButtons {
					bi, _ := strconv.Atoi(b)
					buttons = append(buttons, bi)
				}
				machine.buttons = append(machine.buttons, buttons)
			}

			if c[0] == '{' {
				machine.targetJoltage = []int{}
				sJoltage := strings.Split(c[1:len(c)-1], ",")
				for _, sj := range sJoltage {
					j, _ := strconv.Atoi(sj)
					machine.targetJoltage = append(machine.targetJoltage, j)
				}
			}
		}

		machines = append(machines, machine)
	}

	return machines
}

func fingerprintLights(nextLights []bool) string {
	fp := []byte{}
	for _, on := range nextLights {
		if on {
			fp = append(fp, '#')
		} else {
			fp = append(fp, '.')
		}
	}
	return string(fp)
}

func configureLights(machine Machine) int {
	visitedStates := map[string]bool{}

	queue := [][]bool{
		make([]bool, len(machine.targetLights)),
	}

	presses := 0
	for {
		presses++
		nextQueue := [][]bool{}

		for _, currentLights := range queue {
			for _, buttons := range machine.buttons {
				// copy state
				nextLights := make([]bool, len(currentLights))
				copy(nextLights, currentLights)

				// press buttons
				for _, button := range buttons {
					nextLights[button] = !nextLights[button]
				}

				targetFound := true
				for idx, isOn := range machine.targetLights {
					if nextLights[idx] != isOn {
						targetFound = false
						break
					}
				}

				if targetFound {
					return presses
				}

				fp := fingerprintLights(nextLights)
				if _, ok := visitedStates[fp]; !ok {
					visitedStates[fp] = true
					nextQueue = append(nextQueue, nextLights)
				}
			}
		}

		queue = nextQueue
	}
}

// ===

func getFinalJoltage(joltages string) []int {
	strJoltages := strings.Split(joltages[1:len(joltages)-1], ",")
	finalJoltage := make([]int, len(strJoltages))
	for i := 0; i < len(strJoltages); i++ {
		finalJoltage[i], _ = strconv.Atoi(strJoltages[i])
	}
	return finalJoltage
}

func configureJoltage(machine Machine) int {
	// finalJoltage := getFinalJoltage(joltages)
	// intButtons := getButtons(buttons)
	const maxClicks = 1000 // adjust as needed
	numJoltages := len(machine.targetJoltage)
	lp := golp.NewLP(0, len(machine.buttons))
	lp.SetVerboseLevel(golp.NEUTRAL)
	objectiveCoeffs := make([]float64, len(machine.buttons))
	for i := 0; i < len(machine.buttons); i++ {
		objectiveCoeffs[i] = 1.0
		lp.SetInt(i, true)
		lp.SetBounds(i, 0.0, float64(maxClicks))
	}
	lp.SetObjFn(objectiveCoeffs)
	for i := 0; i < numJoltages; i++ {
		var entries []golp.Entry
		for j, btn := range machine.buttons {
			if slices.Contains(btn, i) {
				entries = append(entries, golp.Entry{Col: j, Val: 1.0})
			}
		}
		targetValue := float64(machine.targetJoltage[i])
		if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
			panic(err)
		}
	}
	status := lp.Solve()
	if status != golp.OPTIMAL {
		return 0
	}
	solution := lp.Variables()
	clicks := 0
	for _, val := range solution {
		clicks += int(val + 0.5)
	}
	return clicks
}

func main() {
	machines := getInputData()

	totalPresses := 0
	for _, machine := range machines {
		totalPresses += configureLights(machine)
	}
	fmt.Println("Part 1 solution is", totalPresses)

	totalPresses = 0
	for _, machine := range machines {
		totalPresses += configureJoltage(machine)
	}
	fmt.Println("Part 2 solution is", totalPresses)
}
