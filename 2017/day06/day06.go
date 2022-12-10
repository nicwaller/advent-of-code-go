package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2017, 6)
	aoc.Test(run, "sample.txt", "5", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func rebalance(banks []int) {
	mostBlocks := analyze.Analyze(banks).Max
	idx := slices.Index(banks, mostBlocks)
	banks[idx] -= mostBlocks
	for i := 0; i < mostBlocks; i++ {
		banks[(idx+1+i)%len(banks)]++
	}
}

func run(p1 *string, p2 *string) {
	banks := util.NumberFields(aoc.InputString())
	seen := set.New[string]()
	for cycle := 1; ; cycle++ {
		rebalance(banks)
		state := strings.Join(f8l.Map(banks, strconv.Itoa), ",")
		if seen.Contains(state) {
			*p1 = strconv.Itoa(cycle)
			seen = set.New[string]()
			seen.Add(state)
			break
		} else {
			seen.Add(state)
		}
	}

	for cycle := 1; ; cycle++ {
		rebalance(banks)
		state := strings.Join(f8l.Map(banks, strconv.Itoa), ",")
		if seen.Contains(state) {
			*p2 = strconv.Itoa(cycle)
			break
		} else {
			seen.Add(state)
		}
	}
}
