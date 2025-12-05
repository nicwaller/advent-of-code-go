package main

import (
	"sort"
	"strconv"

	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/assert"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2024, 1)
	aoc.Test(run, "sample.txt", "11", "31")
	aoc.Test(run, "input.txt", "1830467", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	listA := make([]int, 0)
	listB := make([]int, 0)
	for _, line := range aoc.InputLines() {
		a, b := util.Pair(util.NumberFields(line))
		listA = append(listA, a)
		listB = append(listB, b)
	}
	assert.Equal(len(listA), len(listB))

	sort.Ints(listA)
	sort.Ints(listB)

	delta := 0
	for i := 0; i < len(listA); i++ {
		delta += util.IntAbs(listA[i] - listB[i])
	}
	*p1 = strconv.Itoa(delta)

	freqs := analyze.Analyze(listB).Frequency
	simScore := 0
	for _, v := range listA {
		simScore += freqs[v] * v
	}
	*p2 = strconv.Itoa(simScore)
}
