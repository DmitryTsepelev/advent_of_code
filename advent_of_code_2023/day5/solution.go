package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Mapping struct {
	sourceStart int64
	destStart   int64
	rangeWidth  int64
}

type Almanach struct {
	seeds                 []int64
	seedRanges            [][]int64
	seedToSoil            []Mapping
	soilToFertilizer      []Mapping
	fertilizerToWater     []Mapping
	waterToLight          []Mapping
	lightToTemperature    []Mapping
	temperatureToHumidity []Mapping
	humidityToLocation    []Mapping
}

func stringToIntList(s string) []int64 {
	sTrimmed := strings.Trim(s, " ")

	cmp := strings.Split(sTrimmed, " ")
	result := make([]int64, 0)
	for _, sNum := range cmp {
		trimmedNum := strings.Trim(sNum, " ")
		if len(trimmedNum) == 0 {
			continue
		}
		num, _ := strconv.ParseInt(trimmedNum, 10, 64)
		result = append(result, num)
	}
	return result
}

func parseMapping(scanner *bufio.Scanner) []Mapping {
	result := make([]Mapping, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		ints := stringToIntList(line)
		mapping := Mapping{sourceStart: ints[1], destStart: ints[0], rangeWidth: ints[2]}
		result = append(result, mapping)
	}
	return result
}

func getInputData() Almanach {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	almanach := Almanach{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, "seeds") {
			almanach.seeds = stringToIntList(line[7:])

			seedRanges := make([][]int64, 0)
			for i := 0; i < len(almanach.seeds)/2; i++ {
				seedRanges = append(seedRanges, []int64{almanach.seeds[i*2], almanach.seeds[i*2+1]})
			}
			almanach.seedRanges = seedRanges
		} else if line == "seed-to-soil map:" {
			almanach.seedToSoil = parseMapping(scanner)
		} else if line == "soil-to-fertilizer map:" {
			almanach.soilToFertilizer = parseMapping(scanner)
		} else if line == "fertilizer-to-water map:" {
			almanach.fertilizerToWater = parseMapping(scanner)
		} else if line == "water-to-light map:" {
			almanach.waterToLight = parseMapping(scanner)
		} else if line == "light-to-temperature map:" {
			almanach.lightToTemperature = parseMapping(scanner)
		} else if line == "temperature-to-humidity map:" {
			almanach.temperatureToHumidity = parseMapping(scanner)
		} else if line == "humidity-to-location map:" {
			almanach.humidityToLocation = parseMapping(scanner)
		}
	}

	return almanach
}

func lookup(mapping []Mapping, id int64) int64 {
	for _, relation := range mapping {
		sourceEnd := relation.sourceStart + relation.rangeWidth

		if id >= relation.sourceStart && id <= sourceEnd {
			return relation.destStart + id - relation.sourceStart
		}
	}

	return id
}

func solve1(almanach Almanach) int64 {
	minLocation := int64(math.MaxInt64)

	for _, seed := range almanach.seeds {
		soil := lookup(almanach.seedToSoil, seed)
		fertilizer := lookup(almanach.soilToFertilizer, soil)
		water := lookup(almanach.fertilizerToWater, fertilizer)
		light := lookup(almanach.waterToLight, water)
		temperature := lookup(almanach.lightToTemperature, light)
		humidity := lookup(almanach.temperatureToHumidity, temperature)
		location := lookup(almanach.humidityToLocation, humidity)

		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

// ---

func min(x, y int64) int64 {
	if x > y {
		return y
	}

	return x
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}

	return y
}

func applyTransformation(ranges [][]int64, mappings []Mapping) [][]int64 {
	sort.Slice(mappings, func(i, j int) bool {
		return mappings[i].sourceStart < mappings[j].sourceStart
	})

	nextRanges := make([][]int64, 0)

	for _, rng := range ranges {
		l, r := rng[0], rng[1]

		for idx, mapping := range mappings {
			// everything between mappings or before last mapping
			if l < mapping.sourceStart && l <= r {
				nextRanges = append(nextRanges, []int64{l, min(mapping.sourceStart-1, r)})
				l = mapping.sourceStart
			}

			// overlap with mapping
			sourceEnd := mapping.sourceStart + mapping.rangeWidth - 1
			if l <= r && l <= sourceEnd && mapping.sourceStart <= r {
				overlapL := max(mapping.sourceStart, l)
				overlapR := min(mapping.sourceStart+mapping.rangeWidth-1, r)
				shift := mapping.destStart - mapping.sourceStart
				nextRanges = append(nextRanges, []int64{overlapL + shift, overlapR + shift})
				l += overlapR - overlapL + 1
			}

			if idx == len(mappings)-1 && l <= r {
				nextRanges = append(nextRanges, []int64{l, r})
			}
		}
	}

	return nextRanges
}

func solve2(almanach Almanach) int64 {
	transforms := [][]Mapping{
		almanach.seedToSoil,
		almanach.soilToFertilizer,
		almanach.fertilizerToWater,
		almanach.waterToLight,
		almanach.lightToTemperature,
		almanach.temperatureToHumidity,
		almanach.humidityToLocation,
	}

	current := make([][]int64, 0)

	for i := 0; i < len(almanach.seeds)/2; i++ {
		start := almanach.seeds[i*2]
		width := almanach.seeds[i*2+1]
		current = append(current, []int64{start, start + width})
	}

	for _, transform := range transforms {
		current = applyTransformation(current, transform)
	}

	min := int64(math.MaxInt64)
	for _, locationRange := range current {
		if locationRange[0] < min {
			min = locationRange[0]
		}

		if locationRange[1] < min {
			min = locationRange[1]
		}
	}

	return min
}

func main() {
	almanach := getInputData()
	fmt.Println("Part 1 solution is", solve1(almanach))
	fmt.Println("Part 2 solution is", solve2(almanach))
}
