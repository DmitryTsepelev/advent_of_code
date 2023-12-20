package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Disk struct {
	startPos  int
	positions int
}

func getInputData() []Disk {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	re, _ := regexp.Compile(`Disc #\d+ has (\d+) positions; at time=0, it is at position (\d+).`)

	scanner := bufio.NewScanner(file)

	disks := []Disk{}

	for scanner.Scan() {
		line := scanner.Text()

		match := re.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			positionsS, startS := match[0][1], match[0][2]

			positions, _ := strconv.Atoi(positionsS)
			start, _ := strconv.Atoi(startS)

			disks = append(disks, Disk{
				startPos:  start,
				positions: positions,
			})

			continue
		}

		panic(line)
	}

	return disks
}

func solve(disks []Disk) int {
	time := 0
	for {
		bad := false

		for i := 1; i <= len(disks); i++ {
			disk := disks[i-1]
			if (time+i+disk.startPos)%disk.positions != 0 {
				bad = true
				break
			}
		}

		if !bad {
			break
		}

		time++
	}
	return time
}

func main() {
	disks := getInputData()
	fmt.Println("Solution 1 is", solve(disks))

	disks = append(disks, Disk{
		positions: 11,
		startPos:  0,
	})
	fmt.Println("Solution 2 is", solve(disks))
}
