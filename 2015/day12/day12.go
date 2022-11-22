package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2015, 12)
	aoc.TestLiteral(run, "[1,2,3]", "6", "")
	aoc.TestLiteral(run, "{\"a\":2,\"b\":4}", "6", "")
	aoc.TestLiteral(run, "[[[3]]]", "3", "")
	aoc.TestLiteral(run, "{\"a\":{\"b\":4},\"c\":-1}", "3", "")
	aoc.Test(run, "input.txt", "111754", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	sum := f8l.Sum(util.NumberFields(aoc.InputString()))
	*p1 = strconv.Itoa(sum)
	*p2 = strconv.Itoa(0)
}
