package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"strconv"
)

func main() {
	aoc.Day(11)
	aoc.Test(run, "sample.txt", "1656", "195")
	aoc.Test(run, "input.txt", "1640", "312")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridNumbers()

	flashes := 0
	flashgrid := grid.TransformGrid[int](g, func(val int) bool {
		return false
	})
	for step := 1; ; step++ {
		// First, the energy level of each octopus increases by 1.
		g.MapAll(util.IntIncr)

		flashesInStep := 0
		flashgrid.Fill(false)
		cascading := true
		for cascading {
			// Then, any octopus with an energy level greater than 9 flashes.
			cellsCanFlash := g.Cells().Filter(func(c grid.Cell) bool {
				return g.Get(c) > 9 && flashgrid.Get(c) == false
			}).List()
			for _, cell := range cellsCanFlash {
				// This increases the energy level of all adjacent octopuses by 1,
				// including octopuses that are diagonally adjacent.
				neighbours := g.NeighboursSurround(cell, true)
				g.MapIter(util.IntIncr, iter.ListIterator[grid.Cell](neighbours))
				// An octopus can only flash at most once per step.
				flashgrid.Set(cell, true)
			}
			flashesInStep += len(cellsCanFlash)
			cascading = len(cellsCanFlash) > 0
		}
		// Finally, any octopus that flashed during this step has its energy level set to 0,
		// as it used all of its energy to flash.
		for tiredOcto := flashgrid.Filter(util.Identity[bool]); tiredOcto.Next(); {
			g.Set(tiredOcto.Value(), 0)
		}
		flashes += flashesInStep
		if step == 100 {
			*p1 = strconv.Itoa(flashes)
		}
		if flashesInStep == g.RowCount()*g.ColumnCount() {
			*p2 = strconv.Itoa(step)
			break
		}
	}
}
