package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"strconv"
)

func main() {
	aoc.Select(2022, 6)
	aoc.Test(run, "sample.txt", "7", "")
	aoc.Test(run, "input.txt", "", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	// find four characters all different
	// how many characters before marker?

	ite := iter.SlidingWindow(4, iter.StringIterator(aoc.InputString(), 1))
	i := 4
	for ite.Next() {
		//abc := strings.Join(ite.Value(), "")
		if analyze.Analyze(ite.Value()).Distinct == 4 {
			*p1 = strconv.Itoa(i)
			break
		}
		i++
	}

	ite2 := iter.SlidingWindow(14, iter.StringIterator(aoc.InputString(), 1))
	i2 := 14
	for ite2.Next() {
		//abc := strings.Join(ite.Value(), "")
		if analyze.Analyze(ite2.Value()).Distinct == 14 {
			*p2 = strconv.Itoa(i2)
			break
		}
		i2++
	}

}
