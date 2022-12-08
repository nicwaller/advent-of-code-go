package main

import (
	"advent-of-code/lib/analyze"
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Select(2022, 8)
	aoc.Test(run, "sample.txt", "21", "8")
	aoc.Test(run, "input.txt", "1816", "383520")
	aoc.Out()
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridNumbers()

	lt := func(n int) func(int) bool { return func(v int) bool { return v < n } }
	gte := func(n int) func(int) bool { return func(v int) bool { return v >= n } }

	visible := 0
	viewDist := make([]int, 4)
	scenics := analyze.Box[int]{}
	for cellIterator := g.Cells(); cellIterator.Next(); {
		treeCell := cellIterator.Value()
		treeHeight := g.Get(treeCell)

		util.Clear(viewDist)
		visibleFromAnyEdge := false
		for k, ray := range g.NeighbourRays(treeCell) {
			vals := iter.Map[grid.Cell, int](ray, g.Get).List()
			iter.ListIterator(vals).Counting(&viewDist[k]).TakeWhile(lt(treeHeight)).Go()
			if len(f8l.Filter(vals, gte(treeHeight))) == 0 {
				visibleFromAnyEdge = true
			}
		}
		scenics.Put(util.IntProductV(viewDist...))
		if visibleFromAnyEdge {
			visible++
		}
	}

	*p1 = strconv.Itoa(visible)
	*p2 = strconv.Itoa(scenics.Max)
}
