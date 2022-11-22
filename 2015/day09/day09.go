package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grapheasy"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/set"
	"advent-of-code/lib/util"
	"strconv"
	"strings"
)

func main() {
	aoc.Select(2015, 9)
	aoc.Test(run, "sample.txt", "605", "982")
	aoc.Test(run, "input.txt", "117", "909")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLines()
	g := grapheasy.New[bool](len(lines))
	locs := set.New[string]()
	for _, line := range lines {
		f := strings.Fields(line)
		g.AddBothCost(f[0], f[2], util.Must(strconv.Atoi(f[4])))
		locs.Extend(f[0], f[2])
	}
	a := analyze.Box[int]{}
	for p := iter.Permute[string](locs.Items()); p.Next(); {
		permutation := p.Value()
		var sumCost int
		for sliding := iter.SlidingWindow[string](2, iter.ListIterator(permutation)); sliding.Next(); {
			pair := sliding.Value()
			sumCost += g.Cost(pair[0], pair[1])
		}
		a.Put(sumCost)
	}
	*p1 = strconv.Itoa(a.Min)
	*p2 = strconv.Itoa(a.Max)
}
