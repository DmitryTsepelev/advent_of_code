package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	expressions := make([]string, 0)

	for y := 0; scanner.Scan(); y++ {
		expressions = append(expressions, scanner.Text())
	}

	return expressions
}

// Part 1

func evalLeftToRight(src string) int {
	src = strings.ReplaceAll(src, " ", "")

	stack := []any{}

	processNum := func(expr int) {
		arg1 := expr

		if len(stack) == 0 || stack[len(stack)-1] == '(' {
			stack = append(stack, arg1)
		} else {
			operand := stack[len(stack)-1]
			arg2 := stack[len(stack)-2].(int)

			stack = stack[:len(stack)-2]
			if operand == '*' {
				stack = append(stack, arg1*arg2)
			} else {
				stack = append(stack, arg1+arg2)
			}
		}
	}

	for _, expr := range src {
		if expr == ')' {
			parenResult := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			processNum(parenResult.(int))
		} else if expr >= '0' && expr <= '9' {
			processNum(int(expr - '0'))
		} else {
			stack = append(stack, expr)
		}
	}

	return stack[0].(int)
}

// Part 2

func isNumber(token rune) bool {
	return token >= '0' && token <= '9'
}

func top(stack []rune) rune {
	return stack[len(stack)-1]
}

func pop(stack []int) (int, []int) {
	if len(stack) == 0 {
		panic("pop failed")
	}

	return stack[len(stack)-1], stack[:len(stack)-1]
}

func evalRPN(tokens []rune) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case '+':
			var l, r int

			l, stack = pop(stack)
			r, stack = pop(stack)

			stack = append(stack, l+r)
		case '*':
			var l, r int

			l, stack = pop(stack)
			r, stack = pop(stack)

			stack = append(stack, l*r)
		default:
			stack = append(stack, int(token-'0'))
		}
	}

	return stack[0]
}

func shuntingYard(src string) []rune {
	output := make([]rune, 0)
	operatorStack := make([]rune, 0)

	for _, r := range src {
		if r >= '0' && r <= '9' {
			output = append(output, r)
		} else {
			if r == '*' {
				for len(operatorStack) > 0 && top(operatorStack) != '(' {
					output = append(output, top(operatorStack))
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
				operatorStack = append(operatorStack, r)
			} else if r == '+' {
				operatorStack = append(operatorStack, r)
			} else if r == '(' {
				operatorStack = append(operatorStack, r)
			} else if r == ')' {
				popped := top(operatorStack)
				operatorStack = operatorStack[:len(operatorStack)-1]
				for popped != '(' {
					output = append(output, popped)
					popped = top(operatorStack)
					operatorStack = operatorStack[:len(operatorStack)-1]
				}
			}
		}
	}
	for len(operatorStack) > 0 {
		output = append(output, top(operatorStack))
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return output
}

func evalAdvanced(src string) int {
	src = strings.ReplaceAll(src, " ", "")
	rpl := shuntingYard(src)
	return evalRPN(rpl)
}

// Solver

func solve(expressions []string, eval func(string) int) int {
	sum := 0
	for _, expression := range expressions {
		sum += eval(expression)
	}
	return sum
}

func main() {
	expressions := getInputData()

	fmt.Println("Part 1 solution is", solve(expressions, evalLeftToRight))

	fmt.Println("Part 2 solution is", solve(expressions, evalAdvanced))
}
