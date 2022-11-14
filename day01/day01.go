package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Day(1)
	aoc.Test(run, "sample.txt", "7", "5")
	aoc.Test(run, "input.txt", "1167", "1130")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(part1())
	*p2 = strconv.Itoa(part2())
}

func part1() int {
	lines := aoc.InputLinesIterator()
	nums := iter.Map(lines, util.UnsafeAtoi)
	pairs := iter.SlidingWindow(2, nums)
	return pairs.Filter(func(p []int) bool { return p[1] > p[0] }).Count()
}

func part2() int {
	lines := aoc.InputLinesIterator()
	nums := iter.Map(lines, util.UnsafeAtoi)
	sliding := iter.SlidingWindow(3, nums)
	sums := iter.Map(sliding, func(arr []int) int {
		return f8l.Sum(arr)
	})
	pairs := iter.SlidingWindow(2, sums)
	return pairs.Filter(func(p []int) bool { return p[1] > p[0] }).Count()
}
