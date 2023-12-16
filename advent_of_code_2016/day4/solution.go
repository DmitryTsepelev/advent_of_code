package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getInputData() []string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func parseRoom(room string) (string, string, int, []rune) {
	idx := 0
	name := ""
	fullName := ""
	for {
		if room[idx] >= '0' && room[idx] <= '9' {
			break
		}

		fullName += string(room[idx])

		if room[idx] != '-' {
			name += string(room[idx])
		}

		idx++
	}

	sectorS := ""
	for room[idx] != '[' {
		sectorS += string(room[idx])
		idx++
	}

	sectorId, _ := strconv.Atoi(sectorS)

	hash := []rune(room[idx+1 : len(room)-1])
	sort.Slice(hash, func(i, j int) bool {
		return hash[i] < hash[j]
	})

	return name, fullName, sectorId, hash
}

func nMostFrequentChars(s string, n int) []rune {
	counts := make(map[rune]int, 0)
	for _, r := range s {
		if _, ok := counts[r]; ok {
			counts[r]++
		} else {
			counts[r] = 1
		}
	}

	biggestCount := 0
	countGroups := make(map[int]([]rune), 0)
	for r, count := range counts {
		if count > biggestCount {
			biggestCount = count
		}
		if list, ok := countGroups[count]; ok {
			countGroups[count] = append(list, r)
		} else {
			countGroups[count] = []rune{r}
		}
	}

	result := make([]rune, 0)

	for count := biggestCount; count >= 1; count-- {
		groupRunes := countGroups[count]
		sort.Slice(groupRunes, func(i, j int) bool {
			return groupRunes[i] < groupRunes[j]
		})
		result = append(result, groupRunes...)
	}

	common := result[:5]
	sort.Slice(common, func(i, j int) bool {
		return common[i] < common[j]
	})

	return common
}

func sumSectors(lines []string) int {
	sumSectors := 0
	for _, line := range lines {
		number, _, sectorId, hash := parseRoom(line)
		common := nMostFrequentChars(number, 5)
		if string(common) == string(hash) {
			sumSectors += sectorId
		}
	}
	return sumSectors
}

func decryptName(name string, shift int) string {
	result := ""

	for _, c := range name {
		if c == '-' {
			result += " "
			continue
		}

		result += string((int(c)-'a'+shift)%26 + 'a')
	}

	return result[:len(result)-1]
}

func findStorage(lines []string) int {
	for _, line := range lines {
		_, encryptedName, sectorId, _ := parseRoom(line)
		decryptedName := decryptName(encryptedName, sectorId)

		if decryptedName == "northpole object storage" {
			return sectorId
		}
	}

	return 0
}

func main() {
	lines := getInputData()
	fmt.Println("Solution 1 is", sumSectors(lines))
	fmt.Println("Solution 2 is", findStorage(lines))
}
