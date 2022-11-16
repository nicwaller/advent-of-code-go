package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2020, 2)
	aoc.Test(run, "sample.txt", "2", "1")
	aoc.Test(run, "input.txt", "469", "267")
	aoc.Run(run)
}

func isValid1(line string) bool {
	f := strings.Fields(line)
	min, max := util.Pair(f8l.Map[string, int](strings.Split(f[0], "-"), util.UnsafeAtoi))
	letter := f[1][0:1]
	password := f[2]
	a := analyze.Analyze(iter.StringIterator(password, 1).List())
	return a.Frequency[letter] >= min && a.Frequency[letter] <= max
}

func isValid2(line string) bool {
	f := strings.Fields(line)
	i1, i2 := util.Pair(f8l.Map[string, int](strings.Split(f[0], "-"), util.UnsafeAtoi))
	letter := f[1][0]
	password := f[2]
	return (password[i1-1] == letter) != (password[i2-1] == letter)
}

func run(p1 *string, p2 *string) {
	*p1 = strconv.Itoa(aoc.InputLinesIterator().Filter(isValid1).Count())
	*p2 = strconv.Itoa(aoc.InputLinesIterator().Filter(isValid2).Count())
}
