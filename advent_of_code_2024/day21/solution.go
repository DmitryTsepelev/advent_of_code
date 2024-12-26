package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	codes := []string{}

	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	return codes
}

const (
	Horizontal = iota
	Vertical
)

type Pos struct {
	row, col int
}

type Memo struct {
	seq   string
	depth byte
}

var dp = map[Memo]int{}  // Memoization map
var nkpStart = Pos{3, 2} // Numerical keypad start Pos
var dkpStart = Pos{0, 2} // Directional keypad start Pos

// Char to Pos on numeric keypad
func ctopn(c byte) Pos {
	if c == '0' {
		return Pos{3, 1}
	}
	if c == 'A' {
		return nkpStart
	}
	row := 2 - ((c - '0' - 1) / 3)
	col := (c - '0' - 1) % 3
	return Pos{int(row), int(col)}
}

// Char to Pos on directional keypad
func ctopd(d byte) Pos {
	switch d {
	case '^':
		return Pos{0, 1}
	case '<':
		return Pos{1, 0}
	case 'v':
		return Pos{1, 1}
	case '>':
		return Pos{1, 2}
	default:
		return Pos{0, 2}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pathWriter(off, dir int) []byte {
	var path []byte
	var c byte
	if dir == Horizontal {
		if off < 0 {
			c = '>'
		} else {
			c = '<'
		}
	} else {
		if off < 0 {
			c = 'v'
		} else {
			c = '^'
		}
	}
	for i := 0; i < abs(off); i++ {
		path = append(path, c)
	}
	return path
}

func shortestSeq(src, dst Pos, isNumPad bool) string {
	var path []byte

	dr := src.row - dst.row
	dc := src.col - dst.col

	movesV := pathWriter(dr, Vertical)
	movesH := pathWriter(dc, Horizontal)

	var onGap bool
	if isNumPad {
		onGap = (src.row == 3 && dst.col == 0) || (src.col == 0 && dst.row == 3)
	} else {
		onGap = (src.col == 0 && dst.row == 0) || (src.row == 0 && dst.col == 0)
	}

	goingLeft := dst.col < src.col

	if goingLeft != onGap {
		movesV, movesH = movesH, movesV
	}

	path = append(append([]byte{}, movesV...), movesH...)
	path = append(path, 'A')
	return string(path)
}

func dfs(memo Memo) int {
	if v, ok := dp[memo]; ok {
		return v
	}
	if memo.depth == 0 {
		return len(memo.seq)
	}

	var res int
	var path []string
	prev := dkpStart
	for _, c := range memo.seq {
		curr := ctopd(byte(c))
		path = append(path, shortestSeq(prev, curr, false))
		prev = curr
	}

	for _, p := range path {
		res += dfs(Memo{string(p), memo.depth - 1})
	}
	dp[memo] = res
	return res
}

func codeToNum(code string) int {
	digits := ""

	for _, r := range code {
		if r >= '0' && r <= '9' {
			digits += string(r)
		}
	}

	num, _ := strconv.Atoi(digits)
	return num
}

func solve(codes []string, depth byte) int {
	var res int

	var codeInt int
	for _, code := range codes {
		var path []string
		codeInt = codeToNum(string(code[:len(code)-1]))

		prev := nkpStart
		for _, c := range code {
			curr := ctopn(byte(c))
			path = append(path, shortestSeq(prev, curr, true))
			prev = curr
		}

		var pathLen int
		for _, code := range path {
			pathLen += dfs(Memo{string(code), depth})
		}
		res += pathLen * codeInt
	}

	return res
}

func main() {
	codes := getInputData()

	fmt.Println("Part 1 solution is", solve(codes, 2))
	fmt.Println("Part 2 solution is", solve(codes, 25))
}
