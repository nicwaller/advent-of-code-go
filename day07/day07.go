package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"math"
	"strconv"
)

func main() {
	aoc.Day(7)
	aoc.Test(run, "sample.txt", "37", "168")
	aoc.Test(run, "input.txt", "345035", "97038163")
	aoc.Run(run)
}

var moveCosts = make([]int, 2000)

func init() {
	for i := 1; i < len(moveCosts); i++ {
		moveCosts[i] = moveCosts[i-1] + i
	}
}

func run(p1 *string, p2 *string) {
	input := util.NumberFields(aoc.InputString())
	aMin, aMax := analyze.MinMax(input)

	// PERF: this could be MUCH more efficient with gradient descent
	best1 := iter.Range(aMin, aMax).Map(func(hPos int) int {
		return iter.ListIterator(input).
			Map(func(crabX int) int { return util.IntAbs(crabX - hPos) }).
			Reduce(util.IntSum, 0)
	}).Reduce(util.IntMin, math.MaxInt32)
	*p1 = strconv.Itoa(best1)

	best2 := iter.Range(aMin, aMax).Map(func(hPos int) int {
		return iter.ListIterator(input).
			Map(func(crabX int) int { return util.IntAbs(crabX - hPos) }).
			Map(func(dist int) int { return moveCosts[dist] }).
			Reduce(util.IntSum, 0)
	}).Reduce(util.IntMin, math.MaxInt32)
	*p2 = strconv.Itoa(best2)
}
