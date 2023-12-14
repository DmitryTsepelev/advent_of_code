package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Guard struct {
	ID              int
	MostSleptMinute int
	Sleeps          map[string]*[60]int
	Counts          [60]int
	MinutesSlept    int
}

func loadGuards() map[int]*Guard {
	data := *getInputData()
	sort.Strings(data)

	guards := map[int]*Guard{}

	var guard *Guard
	startedAt := 0

	for _, line := range data {
		cmp := strings.Split(line, "] ")
		message := cmp[1]
		datetime := strings.Split(strings.Replace(cmp[0], "[", "", 1), " ")
		date := datetime[0]
		minute, _ := strconv.Atoi(
			strings.Split(datetime[1], ":")[1],
		)

		if strings.Contains(message, "begins shift") {
			id, _ := strconv.Atoi(
				strings.Replace(
					strings.Replace(message, "Guard #", "", 1),
					" begins shift",
					"",
					1,
				),
			)

			guard = guards[id]
			if guard == nil {
				guards[id] = &Guard{
					ID:     id,
					Sleeps: make(map[string]*[60]int),
				}
				guard = guards[id]
			}
		} else if strings.Contains(message, "falls asleep") {
			startedAt = minute
		} else if strings.Contains(message, "wakes up") {
			sleep := guard.Sleeps[date]
			if sleep == nil {
				guard.Sleeps[date] = &[60]int{}
			}

			for i := startedAt; i < minute; i++ {
				guard.Sleeps[date][i] = 1
			}
		}
	}

	for _, guard := range guards {
		for _, minutes := range guard.Sleeps {
			for i := range guard.Counts {
				if minutes[i] > 0 {
					guard.Counts[i] = guard.Counts[i] + 1
					guard.MinutesSlept = guard.MinutesSlept + 1
				}
			}
		}

		mostSleptMinuteCount := 0

		for i, sleepCount := range guard.Counts {
			if sleepCount > mostSleptMinuteCount {
				mostSleptMinuteCount = sleepCount
				guard.MostSleptMinute = i
			}
		}
	}

	return guards
}

func getInputData() *[]string {
	data := []string{}

	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return &data
}

func solveTask1() int {
	guards := loadGuards()

	var mostSleepingGuard *Guard

	for _, guard := range guards {
		if mostSleepingGuard == nil || guard.MinutesSlept > mostSleepingGuard.MinutesSlept {
			mostSleepingGuard = guard
		}
	}

	return mostSleepingGuard.ID * mostSleepingGuard.MostSleptMinute
}

func solveTask2() int {
	guards := loadGuards()

	var mostSleepingGuard *Guard

	for _, guard := range guards {
		if mostSleepingGuard == nil || guard.Counts[guard.MostSleptMinute] > mostSleepingGuard.Counts[mostSleepingGuard.MostSleptMinute] {
			mostSleepingGuard = guard
		}
	}

	return mostSleepingGuard.ID * mostSleepingGuard.MostSleptMinute
}

func main() {
	fmt.Println("Task 1 solution is", solveTask1())
	fmt.Println("Task 2 solution is", solveTask2())
}
