package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2017, 1)
	aoc.TestLiteral(run, "1122", "3", "")
	aoc.TestLiteral(run, "1111", "4", "")
	aoc.TestLiteral(run, "1234", "0", "")
	aoc.TestLiteral(run, "91212129", "9", "")
	aoc.TestLiteral(run, "1212", "", "6")
	aoc.TestLiteral(run, "1221", "", "0")
	aoc.TestLiteral(run, "123425", "", "4")
	aoc.TestLiteral(run, "123123", "", "12")
	aoc.TestLiteral(run, "12131415", "", "4")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	input := aoc.InputString()

	repeat := iter.Repeat(iter.StringIterator(input, 1)).Take(len(input) + 1)
	pairs := iter.SlidingWindow(2, repeat)
	p1sum := iter.Map(pairs, func(p []string) int {
		if p[0] == p[1] {
			return util.UnsafeAtoi(p[0])
		} else {
			return 0
		}
	}).Reduce(util.IntSum, 0)
	*p1 = strconv.Itoa(p1sum)

	half := len(input) / 2
	repeat2 := iter.Repeat(iter.StringIterator(input, 1)).Take(len(input) + half)
	slide := iter.SlidingWindow(half+1, repeat2)
	p2sum := iter.Map(slide, func(p []string) int {
		if p[0] == p[len(p)-1] {
			return util.UnsafeAtoi(p[0])
		} else {
			return 0
		}
	}).Reduce(util.IntSum, 0)
	*p2 = strconv.Itoa(p2sum)
}
