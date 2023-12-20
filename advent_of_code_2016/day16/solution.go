package main

import "fmt"

func step(a []bool) []bool {
	b := []bool{}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] {
			b = append(b, false)
		} else {
			b = append(b, true)
		}
	}

	return append(append(a, false), b...)
}

func checksum(data []bool, size int) string {
	sum := data[:size]

	for {
		newSum := []bool{}
		for i := 0; i < len(sum)/2; i++ {
			x, y := sum[i*2], sum[i*2+1]

			newSum = append(newSum, x == y)
		}

		if len(newSum)%2 == 1 {
			sSum := ""
			for _, c := range newSum {
				if c {
					sSum += "1"
				} else {
					sSum += "0"
				}
			}
			return sSum
		}

		sum = newSum
	}
}

func fillDisk(data string, size int) string {
	bData := []bool{}

	for _, c := range data {
		if c == '1' {
			bData = append(bData, true)
		} else {
			bData = append(bData, false)
		}
	}

	for len(bData) < size {
		bData = step(bData)
	}

	return checksum(bData, size)
}

func main() {
	input := "10111011111001111"
	fmt.Println("Solution 1 is", fillDisk(input, 272))
	fmt.Println("Solution 2 is", fillDisk(input, 35651584))
}
