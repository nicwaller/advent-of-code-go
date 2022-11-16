package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2020, 1)
	aoc.Test(run, "sample.txt", "514579", "241861950")
	aoc.Test(run, "input.txt", "1010884", "253928438")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLinesInt()
	pair := iter.ProductV(lines, lines).Filter(func(pair []int) bool {
		return pair[0]+pair[1] == 2020
	}).TakeFirst()
	*p1 = strconv.Itoa(util.IntProductV(pair...))

	triplet := iter.ProductV(lines, lines, lines).Filter(func(pair []int) bool {
		return pair[0]+pair[1]+pair[2] == 2020
	}).TakeFirst()
	*p2 = strconv.Itoa(util.IntProductV(triplet...))
}
