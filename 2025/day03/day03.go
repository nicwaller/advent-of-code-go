package main

import (
	"strconv"

	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2025, 3)
	aoc.Test(run, "sample.txt", "357", "17321")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	sumJoltage := 0
	for _, bank := range aoc.InputLines() {
		sumJoltage += maxJoltage(bank)
	}
	*p1 = strconv.Itoa(sumJoltage)
}

func maxJoltage(bank string) int {
	batteries := f8l.Map(iter.StringIterator(bank, 1).List(), util.UnsafeAtoi)
	bestJoltage := 0
	for i := 0; i < len(batteries); i++ {
		for j := i + 1; j < len(batteries); j++ {
			joltage := batteries[i]*10 + batteries[j]
			if joltage > bestJoltage {
				bestJoltage = joltage
			}
		}
	}
	return bestJoltage
}
