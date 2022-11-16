package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
	"math"
	"sort"
	"strconv"
)

func main() {
	aoc.Select(2021, 9)
	aoc.Test(run, "sample.txt", "15", "1134")
	aoc.Test(run, "input.txt", "516", "1023660")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	g := aoc.InputGridNumbers()
	riskSum := 0
	minima := make([]grid.Cell, 0)
	for iterCell := g.Cells(); iterCell.Next(); {
		cell := iterCell.Value()
		neighbours := g.NeighboursAdjacent(cell, false)
		iterNeighbours := iter.ListIterator(neighbours)
		iterVals := iter.Transform[grid.Cell, int](iterNeighbours, func(c grid.Cell) int {
			return g.Get(c)
		})
		lowestNeighbour := iterVals.Reduce(util.IntMin, math.MaxInt32)
		cellValue := g.Get(cell)
		if cellValue < lowestNeighbour {
			minima = append(minima, cell)
			riskLevel := 1 + cellValue
			riskSum += riskLevel
		}
	}
	*p1 = strconv.Itoa(riskSum)

	basinSizes := make([]int, 0)
	for iterCell := iter.ListIterator(minima); iterCell.Next(); {
		cell := iterCell.Value()
		isBoundary := func(v int) bool { return v != 9 }
		changed := g.FloodFill(cell, isBoundary, 9)
		basinSizes = append(basinSizes, changed)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	basinMultiplyResult := basinSizes[0] * basinSizes[1] * basinSizes[2]

	*p2 = strconv.Itoa(basinMultiplyResult)
}
