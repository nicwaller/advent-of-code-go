package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"sort"
	"strconv"
)

func main() {
	aoc.Select(2016, 3)
	//aoc.Test(run, "input.txt", "982", "1826")
	aoc.Run(run)
}

func isTriangle(edgeLength []int) bool {
	sort.Ints(edgeLength)
	return edgeLength[0]+edgeLength[1] > edgeLength[2]
}

func run(p1 *string, p2 *string) {
	count := 0
	for _, line := range aoc.InputLines() {
		if isTriangle(util.NumberFields(line)) {
			count++
		}
	}
	*p1 = strconv.Itoa(count)

	count = 0
	blocks := iter.Chunk(3, aoc.InputLinesIterator())
	for blocks.Next() {
		v := blocks.Value()
		nf := util.NumberFields(v[0] + v[1] + v[2])
		for offset := 0; offset < 3; offset++ {
			t := []int{nf[offset+0], nf[offset+3], nf[offset+6]}
			if isTriangle(t) {
				count++
			}
		}
	}
	*p2 = strconv.Itoa(count)
}
