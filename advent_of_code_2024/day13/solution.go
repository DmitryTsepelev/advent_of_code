package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	buttonA Point
	buttonB Point
	prize   Point
}

func parseButton(s string) (int, int) {
	cmp := strings.Split(s, ", ")
	x, _ := strconv.Atoi(cmp[0][2:])
	y, _ := strconv.Atoi(cmp[1][2:])
	return x, y
}

func parsePrize(s string) (int, int) {
	cmp := strings.Split(s, ", ")
	x, _ := strconv.Atoi(cmp[0][2:])
	y, _ := strconv.Atoi(cmp[1][2:])
	return x, y
}

func getInputData() []Machine {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	machines := []Machine{}

	for {
		scanner.Scan()
		if len(scanner.Text()) == 0 {
			break
		}
		ax, ay := parseButton(scanner.Text()[10:])

		scanner.Scan()
		bx, by := parseButton(scanner.Text()[10:])

		scanner.Scan()
		px, py := parsePrize(scanner.Text()[7:])

		scanner.Scan()

		machine := Machine{Point{ax, ay}, Point{bx, by}, Point{px, py}}

		machines = append(machines, machine)
	}

	return machines
}

type Point struct {
	x int
	y int
}

func findCost(buttonA, buttonB, prize Point) int {
	nom, den := prize.y*buttonB.x-prize.x*buttonB.y, buttonA.y*buttonB.x-buttonA.x*buttonB.y
	a := nom / den
	bx2 := prize.x - a*buttonA.x

	if nom%den == 0 && bx2%buttonB.x == 0 {
		return 3*a + (bx2 / buttonB.x)
	}

	return 0
}

func main() {
	machines := getInputData()

	var cost int
	for _, m := range machines {
		cost += findCost(m.buttonA, m.buttonB, m.prize)
	}
	fmt.Println("Part 1 solution is", cost)

	var cost2 int
	for _, m := range machines {
		multPrize := Point{m.prize.x + 10000000000000, m.prize.y + 10000000000000}
		cost2 += findCost(m.buttonA, m.buttonB, multPrize)
	}
	fmt.Println("Part 2 solution is", cost2)
}
