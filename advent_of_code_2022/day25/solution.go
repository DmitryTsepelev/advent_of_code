package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	snafuNums := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		snafuNums = append(snafuNums, line)
	}

	return snafuNums
}

var snafuSyms = []rune{'=', '-', '0', '1', '2'}

func decToSnafuHelper(num, mult, register int, pad string) (string, bool) {
	if register < 0 {
		if num == 0 {
			return "", true
		}

		return "", false
	}
	snafuNum := ""

	for _, sym := range snafuSyms {
		snafuSuffix, valid := decToSnafuHelper(num-snafuSymToDec(sym, mult), mult/5, register-1, pad+"  ")
		if valid {
			return snafuNum + string(sym) + snafuSuffix, true
		}
	}

	return "", false
}

var trans_rev = map[int]string{
	-2: "=",
	-1: "-",
	0:  "0",
	1:  "1",
	2:  "2",
}

func decToSnafu(num int) string {
	if num == 0 {
		return ""
	}
	m1 := num % 5
	m1 = ((m1 + 2) % 5) - 2
	d0 := trans_rev[m1]
	rest := decToSnafu((num - m1) / 5)
	rest += d0
	return rest
}

func snafuSymToDec(sym rune, mult int) int {
	switch sym {
	case '1':
		return mult
	case '2':
		return mult * 2
	case '-':
		return -mult
	case '=':
		return -mult * 2
	}

	return 0
}

func snafuToDec(snafuNum string) int {
	decNum := 0
	mult := 1

	for i := len(snafuNum) - 1; i >= 0; i-- {
		decNum += snafuSymToDec(rune(snafuNum[i]), mult)
		mult *= 5
	}

	return decNum
}

func main() {
	snafuNums := getInputData()

	var sum int
	for _, snafuNum := range snafuNums {
		sum += snafuToDec(snafuNum)
	}
	fmt.Println(decToSnafu(sum))
}
