package main

import (
	"strconv"

	"golang.org/x/exp/slices"

	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	. "advent-of-code/lib/sequent/v1"
	. "advent-of-code/lib/sequent/v1/sugar"
	"advent-of-code/lib/util"
)

func main() {
	aoc.Select(2024, 2)
	aoc.Test(run, "sample.txt", "2", "4")
	aoc.Test(run, "input.txt", "246", "318")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	safeReports := Map(ListIterator(aoc.InputLines()), util.UIntFields).
		Filter(isSafe).
		Count()
	*p1 = strconv.Itoa(safeReports)

	tolerableReports := Map(ListIterator(aoc.InputLines()), util.UIntFields).
		Filter(isTolerable).
		Count()
	*p2 = strconv.Itoa(tolerableReports)
}

// Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.
func isTolerable(level []int) bool {
	if isSafe(level) {
		return true
	}

	for i := 1; i <= len(level); i++ {
		left := slices.Clone(level[:i-1])
		right := slices.Clone(level[i:])
		lvl2 := append(left, right...)
		if isSafe(lvl2) {
			return true
		}
	}

	return false
}

func isSafe(level []int) bool {
	// A report only counts as safe if both of the following are true:
	// - Any two adjacent levels differ by at least one and at most three.
	pairs := SlidingWindow(S[int]{}.List(level), 2).List()
	deltas := f8l.Map[[]int, int](pairs, func(p []int) int {
		return p[0] - p[1]
	})

	isSafeDelta := func(delta int) bool {
		abs := util.IntAbs(delta)
		return abs >= 1 && abs <= 3
	}

	for _, delta := range deltas {
		if !isSafeDelta(delta) {
			return false
		}
	}

	// - The levels are either all increasing or all decreasing.
	a := analyze.Analyze(deltas)
	if a.Min*a.Max < 0 {
		return false
	}

	return true
}
