package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cube struct {
	x, y, z, w  int
	isHypercube bool
}

func getInitialState(isHypercube bool) map[Cube]bool {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := make(map[Cube]bool)

	for y := 0; scanner.Scan(); y++ {
		for x, val := range scanner.Text() {
			if val == '#' {
				state[Cube{x, y, 0, 0, isHypercube}] = true
			}
		}
	}

	return state
}

func iterateNeighbours(cube Cube, cb func(Cube)) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if cube.isHypercube {
					for dw := -1; dw <= 1; dw++ {
						if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
							continue
						}

						candidateCube := Cube{
							cube.x + dx,
							cube.y + dy,
							cube.z + dz,
							cube.w + dw,
							cube.isHypercube,
						}

						cb(candidateCube)
					}
				} else {
					if dx == 0 && dy == 0 && dz == 0 {
						continue
					}

					candidateCube := Cube{
						cube.x + dx,
						cube.y + dy,
						cube.z + dz,
						cube.w,
						cube.isHypercube,
					}

					cb(candidateCube)
				}
			}
		}
	}
}

func solve(isHypercube bool) int {
	state := getInitialState(isHypercube)

	cycles := 6
	for cycle := 0; cycle < cycles; cycle++ {
		candidateCubes := make(map[Cube]bool)
		for cube := range state {
			candidateCubes[cube] = true

			iterateNeighbours(cube, func(neighbour Cube) {
				candidateCubes[neighbour] = true
			})
		}

		nextState := make(map[Cube]bool)
		for cube := range candidateCubes {
			activeNeighbours := 0
			iterateNeighbours(cube, func(neighbour Cube) {
				if _, ok := state[neighbour]; ok {
					activeNeighbours++
				}
			})

			if _, isActive := state[cube]; activeNeighbours == 3 || isActive && activeNeighbours == 2 {
				nextState[cube] = true
			}
		}

		state = nextState
	}

	return len(state)
}

func main() {
	fmt.Println("Part 1 solution is", solve(false))
	fmt.Println("Part 2 solution is", solve(true))
}
