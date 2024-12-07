package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	target int
	nums   []int
}

func getInputData() []Equation {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	equations := []Equation{}

	for scanner.Scan() {
		cmp := strings.Split(scanner.Text(), ": ")

		target, _ := strconv.Atoi(cmp[0])

		eq := Equation{target: target, nums: []int{}}

		cmp = strings.Split(cmp[1], " ")

		for _, s := range cmp {
			num, _ := strconv.Atoi(s)
			eq.nums = append(eq.nums, num)
		}

		equations = append(equations, eq)
	}

	return equations
}

type OperatorFunc = func(int, int) int

func findOperations(equation *Equation, i int, partial int, ops []func(a, b int) int) bool {
	if partial > equation.target {
		return false
	} else if i == len(equation.nums) {
		return partial == equation.target
	} else {
		for _, op := range ops {
			if findOperations(equation, i+1, op(partial, equation.nums[i]), ops) {
				return true
			}
		}
		return false
	}
}

func canProduceSum(target int, nums []int, operators [](OperatorFunc)) bool {
	sums := map[int]bool{nums[0]: true}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		nextSums := map[int]bool{}

		for sum := range sums {
			for _, op := range operators {
				candidate := op(sum, num)
				if candidate <= target {
					nextSums[candidate] = true
				}
			}
		}

		sums = nextSums
	}

	if _, ok := sums[target]; ok {
		return true
	}

	return false
}

func add(x, y int) int  { return x + y }
func mult(x, y int) int { return x * y }
func concat(x, y int) int {
	num, _ := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(y))

	return num
}

func solve(equations []Equation, operators []OperatorFunc) int {
	var sum int

	for _, eq := range equations {
		if canProduceSum(eq.target, eq.nums, operators) != findOperations(&eq, 1, eq.nums[0], operators) {
			fmt.Println(eq)
		}
		if canProduceSum(eq.target, eq.nums, operators) {
			sum += eq.target
		}
	}

	return sum
}

func main() {
	equations := getInputData()

	fmt.Println("Part 1 solution is", solve(equations, []OperatorFunc{add, mult}))
	fmt.Println("Part 2 solution is", solve(equations, []OperatorFunc{add, mult, concat}))
}
