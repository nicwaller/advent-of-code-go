package main

import (
	"strconv"

	"advent-of-code/lib/aoc"
)

func main() {
	aoc.Select(2025, 1)
	aoc.Test(run, "sample.txt", "3", "")
	aoc.Test(run, "input.txt", "989", "0")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	dialPosition := 50

	c := 0
	for _, line := range aoc.InputLines() {
		dirPart := line[0]
		numPart := line[1:]
		m := 1
		switch dirPart {
		case 'L':
			m = -1
		case 'R':
			m = 1
		default:
			panic(dirPart)
		}
		n, _ := strconv.Atoi(numPart)

		dialDelta := m * n
		dialPosition += dialDelta
		dialPosition += 100
		dialPosition %= 100

		//fmt.Printf("delta=%d, pos=%d\n", dialDelta, dialPosition)

		if dialPosition%100 == 0 {
			c++
		}
	}

	*p1 = strconv.Itoa(c)
	*p2 = ""
}
