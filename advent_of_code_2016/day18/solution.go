package main

import "fmt"

func countSafe(prevRow string, rows int) int {
	safe := 0

	for _, c := range prevRow {
		if c == '.' {
			safe++
		}
	}

	for i := 1; i < rows; i++ {
		nextRow := ""

		for j := 0; j < len(prevRow); j++ {
			leftIsTrap := false
			if j-1 >= 0 && prevRow[j-1] == '^' {
				leftIsTrap = true
			}

			centerIsTrap := false
			if prevRow[j] == '^' {
				centerIsTrap = true
			}

			rightIsTrap := false
			if j+1 < len(prevRow) && prevRow[j+1] == '^' {
				rightIsTrap = true
			}

			isTrap := leftIsTrap && centerIsTrap && !rightIsTrap ||
				!leftIsTrap && centerIsTrap && rightIsTrap ||
				leftIsTrap && !centerIsTrap && !rightIsTrap ||
				!leftIsTrap && !centerIsTrap && rightIsTrap

			if isTrap {
				nextRow += "^"
			} else {
				safe += 1
				nextRow += "."
			}
		}

		prevRow = nextRow
	}

	return safe
}

func main() {
	input := "^^^^......^...^..^....^^^.^^^.^.^^^^^^..^...^^...^^^.^^....^..^^^.^.^^...^.^...^^.^^^.^^^^.^^.^..^.^"

	fmt.Println("Solution 1 is", countSafe(input, 40))
	fmt.Println("Solution 2 is", countSafe(input, 400000))
}
