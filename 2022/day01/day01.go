package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"sort"
	"strconv"
)

func main() {
	aoc.Select(2022, 1)
	aoc.Test(run, "sample.txt", "24000", "45000")
	aoc.Test(run, "input.txt", "72017", "212520")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	calCounts := iter.Map[[]string, int](aoc.ParagraphsIterator(), func(para []string) int {
		return f8l.Reduce(f8l.Map(para, util.UnsafeAtoi), 0, util.IntSum)
	}).List()
	sort.Sort(sort.Reverse(sort.IntSlice(calCounts)))
	*p1 = strconv.Itoa(calCounts[0])
	*p2 = strconv.Itoa(f8l.Sum(calCounts[0:3]))
}
