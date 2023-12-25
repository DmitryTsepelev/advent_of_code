package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

const (
	DOWN  = 'D'
	UP    = 'U'
	LEFT  = 'L'
	RIGHT = 'R'
)

var directionToDelta = map[byte][]int{
	DOWN:  {1, 0},
	UP:    {-1, 0},
	LEFT:  {0, -1},
	RIGHT: {0, 1},
}

var idxToDirection = map[int]byte{
	0: UP,
	1: DOWN,
	2: LEFT,
	3: RIGHT,
}

type Point struct {
	row int
	col int
}

func directionsFrom(passcode string, current Point) map[byte]Point {
	hash := GetMD5Hash(passcode)
	directions := map[byte]Point{}

	for idx, direction := range idxToDirection {
		if hash[idx] >= 'b' && hash[idx] <= 'f' {
			delta := directionToDelta[direction]
			point := Point{row: current.row + delta[0], col: current.col + delta[1]}

			if point.row >= 0 && point.row <= 3 && point.col >= 0 && point.col <= 3 {
				directions[direction] = point
			}
		}
	}

	return directions
}

func findPath(passcode string) string {
	queue := map[string]Point{passcode: {}}

	for {
		nextQueue := map[string]Point{}

		for path, point := range queue {
			if point.row == 3 && point.col == 3 {
				return path[len(passcode):]
			}

			for direction, nextPoint := range directionsFrom(path, point) {
				nextQueue[path+string(direction)] = nextPoint
			}
		}

		queue = nextQueue
	}
}

func findLongestPath(passcode string) int {
	queue := map[string]Point{passcode: {}}

	best := 0

	for len(queue) > 0 {
		nextQueue := map[string]Point{}

		for path, point := range queue {
			clearPath := path[len(passcode):]

			if point.row == 3 && point.col == 3 {
				if len(clearPath) > best {
					best = len(clearPath)
				}
			} else {
				for direction, nextPoint := range directionsFrom(path, point) {
					nextQueue[path+string(direction)] = nextPoint
				}
			}
		}

		queue = nextQueue
	}

	return best
}

func main() {
	passcode := "rrrbmfta"
	fmt.Println("Solution 1 is", findPath(passcode))
	fmt.Println("Solution 2 is", findLongestPath(passcode))
}
