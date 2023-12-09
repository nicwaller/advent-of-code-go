package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2023, 9)
	aoc.Test(run, "sample.txt", "114", "2")
	aoc.Test(run, "input.txt", "1884768153", "1031")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	var sumNext, sumPrev int
	for _, line := range aoc.InputLines() {
		z := util.NumberFields(line)
		sumNext += nextVal(z)
		sumPrev += prevVal(z)
	}
	*p1 = strconv.Itoa(sumNext)
	*p2 = strconv.Itoa(sumPrev)
}

func isAllZeroes(seq []int) bool {
	for i := range seq {
		if seq[i] != 0 {
			return false
		}
	}
	return true
}

func nextVal(seq []int) int {
	if isAllZeroes(seq) {
		return 0
	}
	ss := diffSeq(seq)
	ss = append(ss, nextVal(ss))
	y := seq[len(seq)-1] + ss[len(ss)-1]
	return y
}

func prevVal(seq []int) int {
	if isAllZeroes(seq) {
		return 0
	}
	ss := diffSeq(seq)
	ss = append([]int{prevVal(ss)}, ss...)
	y := seq[0] - ss[0]
	return y
}

func diffSeq(seq []int) []int {
	r := make([]int, len(seq)-1)
	i := 0
	for ite := iter.SlidingWindow[int](2, iter.ListIterator(seq)); ite.Next(); i++ {
		r[i] = ite.Value()[1] - ite.Value()[0]
	}
	return r
}
