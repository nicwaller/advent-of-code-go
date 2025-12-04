package main

import (
	"strconv"

	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
)

func main() {
	aoc.Select(2025, 4)
	aoc.Test(run, "sample.txt", "13", "43")
	aoc.Test(run, "input.txt", "1395", "8451")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridRunes()

	isPaper := func(c grid.Cell) bool {
		return g.Get(c) == "@"
	}

	isReachable := func(c grid.Cell) bool {
		neighbours := iter.ListIterator(g.NeighboursSurround(c, false))
		return neighbours.Filter(isPaper).Count() < 4
	}

	*p1 = strconv.Itoa(g.Cells().Filter(isPaper).Filter(isReachable).Count())

	removedTotal := 0
	for previous := -1; previous != removedTotal; {
		previous = removedTotal
		g.Cells().Filter(isPaper).Filter(isReachable).Counting(&removedTotal).Each(func(cell grid.Cell) {
			g.Set(cell, ".")
		})
	}

	*p2 = strconv.Itoa(removedTotal)
}
