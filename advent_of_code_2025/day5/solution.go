package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInputData() ([][2]int, []int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	freshRanges := [][2]int{}
	for scanner.Scan() && len(scanner.Text()) > 0 {
		cmp := strings.Split(scanner.Text(), "-")
		l, _ := strconv.Atoi(cmp[0])
		r, _ := strconv.Atoi(cmp[1])

		freshRanges = append(freshRanges, [2]int{l, r})
	}

	available := []int{}
	for scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())
		available = append(available, id)
	}

	return freshRanges, available
}

func merge(intervals [][2]int) [][2]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][2]int, 0)
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}

func main() {
	freshRanges, available := getInputData()

	count := 0
	for _, id := range available {
		for _, rng := range freshRanges {
			if id >= rng[0] && id <= rng[1] {
				count++
				break
			}
		}
	}
	fmt.Println("Part 1 solution is", count)

	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0]
	})

	mergedRanges := merge(freshRanges)

	var total int64
	for _, rng := range mergedRanges {
		total += int64(rng[1] - rng[0] + 1)
	}
	fmt.Println("Part 2 solution is", total)
}
