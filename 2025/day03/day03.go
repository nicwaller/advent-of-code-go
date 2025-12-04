package main

import (
	"slices"
	"strconv"
	"strings"

	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2025, 3)
	aoc.Test(run, "sample.txt", "357", "3121910778619")
	aoc.Test(run, "input.txt", "17321", "171989894144198")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	partOneSum := 0
	partTwoSum := 0
	for _, line := range aoc.InputLines() {
		bank := f8l.Map(iter.StringIterator(line, 1).List(), util.UnsafeAtoi)
		partOneSum += intify(maxJoltageLong(bank, 2))
		partTwoSum += intify(maxJoltageLong(bank, 12))
	}
	*p1 = strconv.Itoa(partOneSum)
	*p2 = strconv.Itoa(partTwoSum)
}

func maxJoltageLong(bank []int, desiredLength int) []int {
	if desiredLength == 0 {
		return []int{}
	}
	threshold := len(bank) - desiredLength + 1
	best := analyze.Analyze(bank[:threshold]).Max
	rest := maxJoltageLong(bank[slices.Index(bank, best)+1:], desiredLength-1)
	return append([]int{best}, rest...)
}

func intify(bank []int) int {
	return util.UnsafeAtoi(strings.Join(f8l.Map(bank, strconv.Itoa), ""))
}
