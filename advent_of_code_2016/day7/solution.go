package main

import (
	"bufio"
	"fmt"
	"os"
)

type IP struct {
	inside  []string
	outside []string
}

func parseLine(line string) IP {
	ip := IP{}
	current := ""
	for idx, c := range line {
		if idx == len(line)-1 {
			ip.outside = append(ip.outside, current+string(c))
		} else if c == '[' {
			ip.outside = append(ip.outside, current)
			current = ""
		} else if c == ']' {
			ip.inside = append(ip.inside, current)
			current = ""
		} else {
			current += string(c)
		}
	}

	return ip
}

func getInputData() []IP {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := []IP{}

	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, parseLine(line))
	}

	return field
}

func hasAbba(str string) bool {
	for i := 0; i < len(str)-3; i++ {
		if str[i] == str[i+3] && str[i+1] == str[i+2] && str[i] != str[i+1] {
			return true
		}
	}
	return false
}

func supportsTLS(ip IP) bool {
	hasAbbaOutside := false

	for _, sequence := range ip.outside {
		if hasAbba(sequence) {
			hasAbbaOutside = true
			break
		}
	}

	hasAbbaInside := false
	for _, sequence := range ip.inside {
		if hasAbba(sequence) {
			hasAbbaInside = true
			break
		}
	}

	return hasAbbaOutside && !hasAbbaInside
}

func findABAs(sequences []string) []string {
	abas := make([]string, 0)
	for _, sequence := range sequences {
		for i := 0; i < len(sequence)-2; i++ {
			if sequence[i] == sequence[i+2] && sequence[i] != sequence[i+1] {
				abas = append(abas, sequence[i:i+3])
			}
		}
	}
	return abas
}

func supportsSSL(ip IP) bool {
	abas := findABAs(ip.outside)
	babs := findABAs(ip.inside)

	for _, aba := range abas {
		expectedBab := string(aba[1]) + string(aba[0]) + string(aba[1])
		for _, bab := range babs {
			if bab == expectedBab {
				return true
			}
		}
	}

	return false
}

func countIps(ips []IP, validIp func(IP) bool) int {
	count := 0
	for _, ip := range ips {
		if validIp(ip) {
			count++
		}
	}
	return count
}

func main() {
	ips := getInputData()
	fmt.Println("Solution 1 is", countIps(ips, supportsTLS))
	fmt.Println("Solution 2 is", countIps(ips, supportsSSL))
}
