package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2022, 4)
	aoc.Test(run, "sample.txt", "2", "4")
	aoc.Test(run, "input.txt", "513", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	fullyContainedPairs := 0
	partialOverlaps := 0
	for _, line := range aoc.InputLines() {
		pair := strings.Split(line, ",")
		e1 := f8l.Map(strings.Split(pair[0], "-"), util.UnsafeAtoi)
		e2 := f8l.Map(strings.Split(pair[1], "-"), util.UnsafeAtoi)
		if e1[0] >= e2[0] && e1[1] <= e2[1] {
			fullyContainedPairs++
		} else if e2[0] >= e1[0] && e2[1] <= e1[1] {
			fullyContainedPairs++
		}
		if e1[0] >= e2[0] && e1[0] <= e2[1] {
			partialOverlaps++
		} else if e2[0] >= e1[0] && e2[0] <= e1[1] {
			partialOverlaps++
		}
	}
	*p1 = strconv.Itoa(fullyContainedPairs)
	*p2 = strconv.Itoa(partialOverlaps)
}
