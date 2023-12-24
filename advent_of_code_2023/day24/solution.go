package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Hailstone struct {
	x  int64
	y  int64
	z  int64
	vx int64
	vy int64
	vz int64

	line Line
}

type Line struct {
	a float64
	b float64
	c float64
}

func getInputData() []Hailstone {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	re, _ := regexp.Compile(`(-?\d+),\s*(-?\d+),\s*(-?\d+)\s*@\s*(-?\d+),\s*(-?\d+),\s*(-?\d+)`)

	scanner := bufio.NewScanner(file)

	hailstones := []Hailstone{}

	for scanner.Scan() {
		line := scanner.Text()

		match := re.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			xS, yS, zS, xvS, yvS, zvS := match[0][1], match[0][2], match[0][3], match[0][4], match[0][5], match[0][6]

			x, _ := strconv.ParseInt(xS, 10, 64)
			y, _ := strconv.ParseInt(yS, 10, 64)
			z, _ := strconv.ParseInt(zS, 10, 64)

			vx, _ := strconv.ParseInt(xvS, 10, 64)
			vy, _ := strconv.ParseInt(yvS, 10, 64)
			vz, _ := strconv.ParseInt(zvS, 10, 64)

			m := float64(vy) / float64(vx)     // y2-y1 / x2-x1
			c := float64(y) - (m * float64(x)) // y = mx + c => c = y - mx
			line := Line{m, -1, c}

			hailstones = append(hailstones, Hailstone{x: x, y: y, z: z, vx: vx, vy: vy, vz: vz, line: line})

			continue
		}

		panic(line)
	}

	return hailstones
}

func findIntersection(s1, s2 Hailstone) *[]float64 {
	a1, b1, c1, a2, b2, c2 := s1.line.a, s1.line.b, s1.line.c, s2.line.a, s2.line.b, s2.line.c
	a1b2_a2b1 := (a1 * b2) - (a2 * b1)
	if a1b2_a2b1 == 0 { // parallel or same
		return nil
	}
	ix, iy := (b1*c2-b2*c1)/a1b2_a2b1, (c1*a2-c2*a1)/a1b2_a2b1
	return &[]float64{ix, iy}
}

func countIntersections(hailstones []Hailstone, testAreaFrom, testAreaTo float64) int {
	count := 0

	for lIdx := 0; lIdx < len(hailstones); lIdx++ {
		for rIdx := lIdx + 1; rIdx < len(hailstones); rIdx++ {
			s1 := hailstones[lIdx]
			s2 := hailstones[rIdx]

			intersection := findIntersection(s1, s2)

			if intersection == nil {
				continue
			}
			// fmt.Println(intersection)

			ix, iy := (*intersection)[0], (*intersection)[1]

			t1, t2 := (ix-float64(s1.x))/float64(s1.vx), (ix-float64(s2.x))/float64(s2.vx)

			if t1 > 0 && t2 > 0 && ix >= testAreaFrom && ix <= testAreaTo && iy >= testAreaFrom && iy <= testAreaTo {
				// fmt.Println("inside")
				count++
			}
			// fmt.Println()
		}
	}

	return count
}

func main() {
	hailstones := getInputData()
	fmt.Println(hailstones)
	fmt.Println(countIntersections(hailstones, 200000000000000, 400000000000000))

}
