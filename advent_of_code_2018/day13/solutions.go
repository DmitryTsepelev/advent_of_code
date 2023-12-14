package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func move(field *[][]string, carts *[]*Cart) bool {
	for _, cart := range *carts {
		cart.Moved = false
	}

	notCrashed := 0
	var last *Cart
	for _, cart := range *carts {
		if !cart.Crashed {
			notCrashed++
			last = cart
		}
	}
	lastTick := notCrashed == 1

	for rowNumber, row := range *field {
		for cellNumber, _ := range row {
			for _, cart := range *carts {
				if cart.Row == rowNumber && cart.Cell == cellNumber && !cart.Moved && !cart.Crashed {
					cart.Moved = true

					if cart.Dir == ">" {
						cart.Cell++
						rightCell := (*field)[rowNumber][cellNumber+1]

						if rightCell == "\\" {
							cart.Dir = "v"
						} else if rightCell == "/" {
							cart.Dir = "^"
						} else if rightCell == "+" {
							if cart.NextTurn == "l" {
								cart.NextTurn = "s"
								cart.Dir = "^"
							} else if cart.NextTurn == "s" {
								cart.NextTurn = "r"
							} else if cart.NextTurn == "r" {
								cart.NextTurn = "l"
								cart.Dir = "v"
							}
						}
					} else if cart.Dir == "<" {
						cart.Cell--
						leftCell := (*field)[rowNumber][cellNumber-1]

						if leftCell == "\\" {
							cart.Dir = "^"
						} else if leftCell == "/" {
							cart.Dir = "v"
						} else if leftCell == "+" {

							if cart.NextTurn == "l" {
								cart.NextTurn = "s"
								cart.Dir = "v"
							} else if cart.NextTurn == "s" {
								cart.NextTurn = "r"
							} else if cart.NextTurn == "r" {
								cart.NextTurn = "l"
								cart.Dir = "^"
							}
						}
					} else if cart.Dir == "v" {
						cart.Row++
						bottomCell := (*field)[rowNumber+1][cellNumber]

						if bottomCell == "\\" {
							cart.Dir = ">"
						} else if bottomCell == "/" {
							cart.Dir = "<"
						} else if bottomCell == "+" {

							if cart.NextTurn == "l" {
								cart.NextTurn = "s"
								cart.Dir = ">"
							} else if cart.NextTurn == "s" {
								cart.NextTurn = "r"
							} else if cart.NextTurn == "r" {
								cart.NextTurn = "l"
								cart.Dir = "<"
							}
						}
					} else if cart.Dir == "^" {
						cart.Row--
						topCell := (*field)[rowNumber-1][cellNumber]

						if topCell == "\\" {
							cart.Dir = "<"
						} else if topCell == "/" {
							cart.Dir = ">"
						} else if topCell == "+" {

							if cart.NextTurn == "l" {
								cart.NextTurn = "s"
								cart.Dir = "<"
							} else if cart.NextTurn == "s" {
								cart.NextTurn = "r"
							} else if cart.NextTurn == "r" {
								cart.NextTurn = "l"
								cart.Dir = ">"
							}
						}
					}

					for _, anotherCart := range *carts {
						if !anotherCart.Crashed && cart != anotherCart && cart.Row == anotherCart.Row && cart.Cell == anotherCart.Cell {
							cart.Crashed = true
							anotherCart.Crashed = true
							fmt.Println(fmt.Sprintf("Crash %d,%d", cart.Cell, cart.Row))
						}
					}
				}
			}

			if lastTick {
				fmt.Println(fmt.Sprintf("Last cart %d,%d", last.Cell, last.Row))
				return true
			}
		}
	}

	return false
}

type Cart struct {
	Row      int
	Cell     int
	Dir      string
	Moved    bool
	Crashed  bool
	NextTurn string
}

func getData() (*[][]string, *[]*Cart) {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	carts := []*Cart{}
	field := [][]string{}
	rowNumber := 0
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")

		row := []string{}
		for cellNumber, cell := range chars {
			if cell == "^" || cell == "v" {
				carts = append(carts, &Cart{Row: rowNumber, Cell: cellNumber, Dir: cell, NextTurn: "l"})
				row = append(row, "|")
			} else if cell == ">" || cell == "<" {
				carts = append(carts, &Cart{Row: rowNumber, Cell: cellNumber, Dir: cell, NextTurn: "l"})
				row = append(row, "-")
			} else {
				row = append(row, cell)
			}
		}

		field = append(field, row)

		rowNumber++
	}

	return &field, &carts
}

func main() {
	field, carts := getData()

	for {
		completed := move(field, carts)

		if completed {
			break
		}
	}
}
