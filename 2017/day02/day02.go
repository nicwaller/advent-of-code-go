package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2017, 2)
	aoc.Test(run, "sample.txt", "18", "")
	aoc.Test(run, "sample2.txt", "", "9")
	//aoc.Test(run, "input.txt", "21845", "191")
	aoc.Run(run)
}

func maxdiff(l []int) int {
	a := analyze.Analyze(l)
	return a.Max - a.Min
}

func evenDivResult(l []int) int {
	pairs := iter.CombinationsN(l, 2)
	for pairs.Next() {
		a := analyze.Analyze(pairs.Value())
		if a.Max%a.Min == 0 {
			return a.Max / a.Min
		}
	}
	// the first sample input cannot be solved :(
	return 0
}

func run(p1 *string, p2 *string) {
	checksum := 0
	checksum2 := 0
	for _, line := range aoc.InputLines() {
		nf := util.NumberFields(line)
		checksum += maxdiff(nf)
		checksum2 += evenDivResult(nf)
	}
	*p1 = strconv.Itoa(checksum)
	*p2 = strconv.Itoa(checksum2)
}
