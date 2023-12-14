package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Unit struct {
	Type  string
	HP    int
	Moved bool
	Point *Point
}

type Point struct {
	Row  int
	Cell int
}

type Field [][]string

func (field *Field) Show() {
	for _, row := range *field {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func (field *Field) Reachable(point *Point, target *Point) bool {
	value := (*field)[point.Row][point.Cell]
	return value == "." || *point == *target
}

func (field *Field) Neighbors(point *Point, target *Point) *[]*Point {
	candidates := []*Point{
		&Point{Row: point.Row - 1, Cell: point.Cell},
		&Point{Row: point.Row, Cell: point.Cell - 1},
		&Point{Row: point.Row, Cell: point.Cell + 1},
		&Point{Row: point.Row + 1, Cell: point.Cell},
	}

	results := []*Point{}
	for _, candidate := range candidates {
		if field.Reachable(candidate, target) {
			results = append(results, candidate)
		}
	}

	return &results
}

func (field *Field) FindPath(start *Point, end *Point) *[]*Point {
	// fmt.Println("FindPath", start, end)
	queue := []*Point{start}

	visited := map[Point]bool{}
	visited[*start] = true

	cameFrom := map[Point]*Point{}
	cameFrom[*start] = nil

	for {
		current := queue[0]
		if current == end {
			break
		}

		queue = queue[1:]
		for _, next := range *field.Neighbors(current, end) {
			_, isVisited := visited[*next]
			if !isVisited {
				queue = append(queue, next)
				visited[*next] = true
				cameFrom[*next] = current
			}
		}

		if len(queue) == 0 {
			break
		}
	}

	path := []*Point{}
	current, found := cameFrom[*end]
	// fmt.Println(current, found)
	if !found {
		return nil
	}

	for {
		if current == start {
			break
		}
		path = append([]*Point{current}, path...)
		current = cameFrom[*current]
	}

	return &path
}

func (field *Field) Move(units *[]*Unit) {
	goblins := []*Unit{}
	elves := []*Unit{}

	for _, unit := range *units {
		unit.Moved = false
		if unit.Type == "G" {
			goblins = append(goblins, unit)
		} else if unit.Type == "E" {
			elves = append(elves, unit)
		}
	}

	for rowNumber, row := range *field {
		for cellNumber, cell := range row {
			point := &Point{Row: rowNumber, Cell: cellNumber}

			var currentUnit *Unit
			for _, unit := range *units {
				if *unit.Point == *point {
					currentUnit = unit
					break
				}
			}

			if currentUnit == nil || currentUnit.Moved {
				continue
			}

			var enemies *[]*Unit
			if cell == "G" {
				enemies = &elves
			} else if cell == "E" {
				enemies = &goblins
			}

			if enemies == nil {
				continue
			}

			if len(*enemies) == 0 {
				// combat ended!!!
			}

			shortestPath := 1000
			var plan *Point
			for _, enemy := range *enemies {
				path := field.FindPath(point, enemy.Point)
				if path != nil {
					if len(*path) < shortestPath {
						shortestPath = len(*path)
						if len(*path) == 0 {
							// attack
						} else {
							plan = (*path)[0]
						}
					}
				}
			}

			if plan != nil {
				(*field)[plan.Row][plan.Cell] = cell
				(*field)[rowNumber][cellNumber] = "."

				currentUnit.Moved = true
				currentUnit.Point = plan
			}
		}
	}
}

func getData() (*Field, *[]*Unit) {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	field := Field{}
	units := []*Unit{}

	rowNumber := 0
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")

		for cellNumber, char := range chars {
			if char == "G" || char == "E" {
				point := Point{Row: rowNumber, Cell: cellNumber}
				units = append(units, &Unit{HP: 200, Type: char, Point: &point})
			}
		}

		field = append(field, chars)
		rowNumber++
	}

	return &field, &units
}

func main() {
	field, units := getData()

	field.Show()
	fmt.Println()

	field.Move(units)
	field.Show()
	fmt.Println()

	field.Move(units)
	field.Show()
	fmt.Println()

	// field.Move(units)
	// field.Show()
	// fmt.Println()
}
