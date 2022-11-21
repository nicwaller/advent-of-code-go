package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"strconv"
)

func main() {
	aoc.Select(2015, 1)
	aoc.Test(run, "sample.txt", "", "")
	aoc.Test(run, "input.txt", "280", "1797")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	a := analyze.CountDistinct(iter.StringIterator(aoc.InputString(), 1).List())
	*p1 = strconv.Itoa(a["("] - a[")"])

	level := 0
	for pos, x := range aoc.InputString() {
		switch x {
		case '(':
			level++
		case ')':
			level--
		}
		if level < 0 {
			*p2 = strconv.Itoa(pos + 1)
			break
		}
	}
}
