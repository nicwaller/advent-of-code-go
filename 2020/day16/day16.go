package main

import (
	"advent-of-code/lib/aoc"
	"strconv"
)

func main() {
	aoc.Select(2020, 16)
	aoc.Test(run, "sample.txt", "", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(0)
	*p2 = strconv.Itoa(0)
}
