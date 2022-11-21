package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/util"
	"sort"
	"strconv"
)

func main() {
	aoc.Select(2015, 2)
	aoc.Test(run, "sample.txt", "58", "34")
	aoc.Test(run, "input.txt", "1588178", "")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	totalPaper := 0
	totalRibbon := 0
	for _, v := range aoc.InputLines() {
		dims := util.NumberFields(v)
		totalPaper += surfaceArea(dims)
		totalRibbon += ribbon(dims)
	}
	*p1 = strconv.Itoa(totalPaper)
	*p2 = strconv.Itoa(totalRibbon)
}

func surfaceArea(dim []int) int {
	l := dim[0]
	w := dim[1]
	h := dim[2]
	smallSide := util.IntMin(util.IntMin(l*w, w*h), h*l)
	return 2*l*w + 2*w*h + 2*h*l + smallSide
}

func ribbon(dim []int) int {
	sort.Ints(dim)
	wrap := dim[0]*2 + dim[1]*2
	bow := dim[0] * dim[1] * dim[2]
	return wrap + bow
}
