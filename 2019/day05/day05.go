package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/aoc/intcode"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2019, 5)
	aoc.Test(run, "sample.txt", "999", "999")
	aoc.Test(run, "input.txt", "7692125", "14340395")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	program := f8l.Map(strings.Split(aoc.InputString(), ","), util.UnsafeAtoi)
	*p1 = strconv.Itoa(last(intcode.ExecIO(program, []int{1})))
	*p2 = strconv.Itoa(last(intcode.ExecIO(program, []int{5})))
}

func last[T any](v []T) T {
	return v[len(v)-1]
}
