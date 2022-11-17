package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"sort"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2020, 5)
	aoc.Test(run, "sample.txt", "357", "")
	aoc.Test(run, "input.txt", "987", "603")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	seats := iter.Map[string, int](aoc.InputLinesIterator(), seatId).List()
	*p1 = strconv.Itoa(f8l.Reduce(seats, 0, util.IntMax))

	if len(seats) < 2 {
		return
	}

	sort.Ints(seats)
	mySeat := iter.
		SlidingWindow(2, iter.ListIterator(seats)).
		Filter(func(x []int) bool {
			return util.IntAbs(x[0]-x[1]) > 1
		}).TakeFirst()[0] + 1
	*p2 = strconv.Itoa(mySeat)
}

func seatId(code string) int {
	bit := map[string]string{
		"F": "0",
		"B": "1",
		"L": "0",
		"R": "1",
	}
	bits := iter.StringIterator(code, 1).Map(func(char string) string {
		return bit[char]
	}).List()
	return int(util.Must(strconv.ParseInt(strings.Join(bits, ""), 2, 16)))
}
