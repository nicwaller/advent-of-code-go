package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"strconv"
	"strings"
)

func main() {
	aoc.Day(20)
	aoc.Test(run, "sample.txt", "35", "")
	aoc.Test(run, "input.txt", "5359", "")
	aoc.Run(run)
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	lines := aoc.InputLinesIterator()
	enhanceAlgoStr := lines.MustTakeArray(2)[0]
	if len(enhanceAlgoStr) != 512 {
		panic(len(enhanceAlgoStr))
	}
	g := grid.FromString(strings.Join(lines.Map(func(s string) string { return s }).List(), "\n"))
	g.Grow(8, ".")
	g.Print()
	enhance := func(g grid.Grid[string]) grid.Grid[string] {
		// nice way to make a new grid? copy!
		newG := g.Copy()
		// FIXME: enlarge!
		// MapAll() should also provide cell reference, not just value
		// Fill() should take a function or a value?
		for cellIter := newG.Cells(); cellIter.Next(); {
			cell := cellIter.Value()
			nCells := iter.ListIterator(newG.NeighboursSurround(cell, true))
			// Neighbours() should also get values? this is cumbersome.
			nVals := iter.Transform(nCells, func(c grid.Cell) string {
				return g.Get(c)
			}).List()
			nStr := strings.Join(nVals, "")
			nStr = strings.ReplaceAll(nStr, "#", "1")
			nStr = strings.ReplaceAll(nStr, ".", "0")
			x, _ := strconv.ParseInt(nStr, 2, 16)
			newG.Set(cell, enhanceAlgoStr[x:x+1])
		}
		return newG
	}
	for i := 0; i < 2; i++ {
		g = enhance(g)
		g.Print()
	}
	// FIXME: hack to subtract -4 to deal with the corners
	*p1 = strconv.Itoa(g.Filter(func(s string) bool { return s == "#" }).Count() - 4)
	//*p2 = strconv.Itoa()
}
