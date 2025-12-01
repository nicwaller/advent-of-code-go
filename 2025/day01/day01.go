package main

import (
	"strconv"

	"advent-of-code/lib/aoc"
)

func main() {
	aoc.Select(2025, 1)
	aoc.Test(run, "sample.txt", "3", "6")
	aoc.Test(run, "input.txt", "989", "5941")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	settleOnZero := 0
	ticksThroughZero := 0

	m := map[uint8]int{
		'L': -1,
		'R': 1,
	}

	dialPosition := 50
	for _, line := range aoc.InputLines() {
		direction := m[line[0]]
		magnitude, _ := strconv.Atoi(line[1:])
		delta := direction * magnitude
		target := dialPosition + delta
		for ; dialPosition != target; dialPosition += direction {
			if dialPosition%100 == 0 {
				ticksThroughZero++
			}
		}
		if dialPosition%100 == 0 {
			settleOnZero++
		}
	}

	*p1 = strconv.Itoa(settleOnZero)
	*p2 = strconv.Itoa(ticksThroughZero)
}
