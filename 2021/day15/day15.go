package main

import (
	"advent-of-code/lib/aoc"
	"advent-of-code/lib/grid"
	"github.com/yourbasic/graph"
	"strconv"
)

func main() {
	aoc.Select(2021, 15)
	aoc.Test(run, "sample.txt", "40", "315")
	aoc.Test(run, "input.txt", "363", "2835")
	aoc.Run(run)
}

func run(p1 *string, p2 *string) {
	smallGrid := grid.FromStringAsInt(aoc.InputString())

	smallGraph := GridToGraph(smallGrid)
	_, distance := graph.ShortestPath(smallGraph, 0, smallGrid.RowCount()*smallGrid.ColumnCount()-1)
	part1 := int(distance)
	*p1 = strconv.Itoa(part1)

	bigGrid := expandGrid(smallGrid)
	bigGraph := GridToGraph(bigGrid)
	_, distance = graph.ShortestPath(bigGraph, 0, bigGrid.RowCount()*bigGrid.ColumnCount()-1)
	part2 := int(distance)
	*p2 = strconv.Itoa(part2)
}

func expandGrid(smallGrid grid.Grid[int]) grid.Grid[int] {
	bigGrid := grid.NewGrid[int](smallGrid.RowCount()*5, smallGrid.ColumnCount()*5)
	for dstCellIter := bigGrid.Cells(); dstCellIter.Next(); {
		dstCell := dstCellIter.Value()
		srcCell := grid.Cell{
			dstCell[0] % smallGrid.Width(),
			dstCell[1] % smallGrid.Height(),
		}
		riskOffset := dstCell[0]/smallGrid.Width() + dstCell[1]/smallGrid.Height()
		srcVal := smallGrid.Get(srcCell)
		// this line here is a bit of straight fuckery
		// y u do this to me
		// > "However, risk levels above 9 wrap back around to 1"
		// SERIOUSLY?!
		dstVal := (srcVal+riskOffset)%10 + ((srcVal + riskOffset) / 10)
		bigGrid.Set(dstCell, dstVal)
	}
	return bigGrid
}

func GridToGraph(grd grid.Grid[int]) *graph.Mutable {
	graf := graph.New(grd.RowCount() * grd.ColumnCount())
	for c := grd.Cells(); c.Next(); {
		dst := grd.OffsetFromCell(c.Value())
		for _, n := range grd.NeighboursAdjacent(c.Value(), false) {
			src := grd.OffsetFromCell(n)
			graf.AddCost(src, dst, int64(grd.Get(c.Value())))
		}
	}
	return graf
}
