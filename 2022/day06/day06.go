package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"strconv"
)

func main() {
	aoc.Select(2022, 6)
	aoc.Test(run, "sample.txt", "7", "19")
	aoc.Test(run, "input.txt", "1896", "3452")
	aoc.Run(run)
	aoc.Out()
}

func find(s string, n int) int {
	return n + iter.
		SlidingWindow(n, iter.StringIterator(s, 1)).
		TakeWhile(func(w []string) bool { return set.FromSlice(w).Size() != n }).
		Count()
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(find(aoc.InputString(), 4))
	*p2 = strconv.Itoa(find(aoc.InputString(), 14))
}
